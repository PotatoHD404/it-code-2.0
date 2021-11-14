package main

import (
	"context"
	"github.com/teris-io/shortid"
)

func (cart *Cart) ResetCart() error {
	cart.Sum = 0.0
	var itemsToDelete []*CartItem
	for i := 0; i < len(cart.Items); i++ {
		if cart.Items[i].Price == nil {
			itemsToDelete = append(itemsToDelete, cart.Items[i])
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
		} else {
			cart.Items[i].Price = new(float)
			*cart.Items[i].Price = cart.Items[i].OrigPrice
			cart.Items[i].Discount = 0.0
			cart.Sum += cart.Items[i].OrigPrice
		}
	}
	if len(itemsToDelete) > 0 {
		_, err := db.NewDelete().Model(&itemsToDelete).WherePK().Exec(ctx)
		if err != nil {
			return err
		}
	}
	cart.Discount = 0.0
	cart.Promos = []*CartPromo{}
	return nil
}
func ArrContainsItem(s []*Item, e *CartItem) int {
	for num, a := range s {
		if a.ID == e.ItemID {
			return num
		}
	}
	return -1
}
func ArrContainsPromo(s []Promo, e *Promo) int {
	for num, a := range s {
		if a.ID == (*e).ID {
			return num
		}
	}
	return -1
}
func (m *CartItem) ApplyPromo(promo *Promo) {
	if promo.Action == "percent_discount" {
		*m.Price -= *m.Price * promo.Discount / 100.0
	} else if promo.Action == "price_discount" {
		*m.Price -= promo.Discount
	}
	if *m.Price < 0 {
		*m.Price = 0
	}
	m.Discount = m.OrigPrice - *m.Price
	m.Used = true
	m.Used = true
}

func (cart *Cart) FlushItems() {
	for _, item := range cart.Items {
		item.Selected = false
		item.Used = false
	}
}
func (cart *Cart) ApplyItemPromo(promo *Promo) {
	tempItems := promo.SelectorItems
	for _, item := range cart.Items {
		if !item.Used {
			if num := ArrContainsItem(tempItems, item); num != -1 && item.Price != nil {
				prevPrice := *item.Price
				item.ApplyPromo(promo)
				cart.Sum -= prevPrice - *item.Price
				tempItems = append(tempItems[:num], tempItems[num+1:]...)
			}
		}
		if len(tempItems) == 0 {
			break
		}
	}
	cart.Promos = append(cart.Promos, &CartPromo{CartID: cart.ID, PromoID: promo.ID})
	promo.AppliesCount++
}

func (cart *Cart) CheckConditions(promo *Promo) (bool, error) {
	var tempItems []*Item
	tempItems = promo.ConditionItems
	if promo.MinOrderSum == nil || cart.Sum >= *promo.MinOrderSum {
		for j := 0; j < len(cart.Items); j++ {
			item := cart.Items[j]
			if !item.Selected {
				if num := ArrContainsItem(tempItems, item); num != -1 {
					item.Selected = true
					tempItems = append(tempItems[:num], tempItems[num+1:]...)
				}
			}
			if len(tempItems) == 0 {
				return true, nil
			}
		}
	}
	return false, nil
}

func (promo *Promo) Exclude(promocodes *[]Promo, i *int) {
	if promo.AppliesCount > 0 {
		for _, exPromo := range promo.Exclusions {
			if num := ArrContainsPromo(*promocodes, exPromo); num != -1 {
				*promocodes = append((*promocodes)[:num], (*promocodes)[num+1:]...)
				if num <= *i {
					*i--
				}
			}
		}
	}
}

func (cart *Cart) ApplyOrderPromo(promo *Promo) error {
	prevPrice := cart.Sum
	if promo.Action == "percent_discount" {
		cart.Sum -= cart.Sum * promo.Discount / 100.0
	} else if promo.Action == "price_discount" {
		cart.Sum -= promo.Discount
	} else /*if promo.Action == "gift"*/ {
		var newItems []*CartItem
		for _, gift := range promo.GiftItems {
			id, _ := shortid.Generate()
			newItem := &CartItem{
				ItemID:     gift.ID,
				CartID:     cart.ID,
				Price:      nil,
				OrigPrice:  0,
				CartItemID: id,
			}
			newItems = append(newItems, newItem)
			cart.Items = append(cart.Items, newItem)
		}
		if newItems != nil {
			_, err := db.
				NewInsert().
				Model(&newItems).
				Exec(ctx)
			if err != nil {
				return err
			}
		}
	}
	if cart.Sum < 0 {
		cart.Sum = 0
	}
	promo.AppliesCount++
	cart.Discount += prevPrice - cart.Sum
	cart.Promos = append(cart.Promos, &CartPromo{CartID: cart.ID, PromoID: promo.ID})
	return nil
}

func (cart *Cart) ApplyPromocode() error {
	var promocodes []Promo
	err := db.NewSelect().
		Model(&promocodes).
		Relation("ConditionItems").
		Relation("SelectorItems").
		Relation("GiftItems").
		Relation("Exclusions").
		Where("promocode = ?", "").
		WhereOr("promocode = ?", cart.Promocode).
		Order("priority ASC").
		Scan(ctx)
	if err != nil {
		return err
	}

	cart.Sum = 0
	for _, m := range cart.Items {
		if m.Price != nil {
			m.Discount = m.OrigPrice - *m.Price
			cart.Sum += *m.Price
		}
	}
	cart.OrigPrice = cart.Sum
	for i := 0; i < len(promocodes); i++ {
		promo := promocodes[i]
		cond := true
		for cond {
			cond, err = cart.CheckConditions(&promo)
			if err != nil {
				return err
			}
			if cond {
				if promo.Scope == "order" {
					err := cart.ApplyOrderPromo(&promo)
					if err != nil {
						return err
					}
					break
				} else {
					if promo.Action == "gift" {
						err := cart.ApplyOrderPromo(&promo)
						if err != nil {
							return err
						}
					} else {
						cart.ApplyItemPromo(&promo)
					}
				}
			}
		}
		promo.Exclude(&promocodes, &i)
		cart.FlushItems()
	}
	return nil
}

func GetCartFromDB(cartId string) (*Cart, error) {
	var cart Cart
	ctx := context.Background()
	err := db.NewSelect().
		Model(&cart).
		Where("cart_id = ?", cartId).
		Relation("Items").
		Relation("Promos").
		Scan(ctx)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	if &cart != nil {
		if cart.Items == nil {
			cart.Items = []*CartItem{}
		}
		if cart.Promos == nil {
			cart.Promos = []*CartPromo{}
		}
	}
	var items []*Item
	for _, m := range cart.Items {
		items = append(items, &Item{ID: m.ItemID})
	}

	items, err = GetItemsFromDB(items)
	if err != nil {
		return nil, err
	}
	for _, m := range cart.Items {
		num := ArrContainsItem(items, m)

		m.Title = items[num].Title
		if m.Price != nil {
			m.Discount = m.OrigPrice - *m.Price
		} else {
			m.Discount = m.OrigPrice
		}

	}
	return &cart, nil
}

func GetItemFromDB(itemID string) (*Item, error) {
	var item Item
	ctx := context.Background()
	err := db.NewSelect().
		Model(&item).
		Where("id = ?", itemID).
		Scan(ctx)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &item, nil
}

func GetItemsFromDB(items []*Item) ([]*Item, error) {
	ctx := context.Background()
	if items == nil {
		return []*Item{}, nil
	}
	err := db.NewSelect().Model(&items).WherePK().Scan(ctx)

	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return items, nil
}

func GetPromosFromDB(items []*Promo) ([]*Promo, error) {
	ctx := context.Background()
	if items == nil {
		return []*Promo{}, nil
	}
	err := db.NewSelect().Model(&items).WherePK().Scan(ctx)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return items, nil
}
