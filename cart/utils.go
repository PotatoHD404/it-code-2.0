package main

import (
	"context"
	"strconv"
)

func GetCartFromDB(cartId string) *Cart {
	var cart Cart
	ctx := context.Background()
	err := db.NewSelect().
		Model(&cart).
		Where("cart_id = ?", cartId).
		Relation("Items").
		Relation("Promos").
		Scan(ctx)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil
	} else if err != nil {
		panic(err)
	}
	if &cart != nil {
		if cart.Items == nil {
			cart.Items = []*CartItem{}
		}
		if cart.Promos == nil {
			cart.Promos = []*Promo{}
		}
	}
	var prev = Item{}
	for num, m := range cart.Items {
		if prev.ID != m.ItemID {
			m.Item = GetItemFromDB(strconv.Itoa(int(m.ItemID)))
			prev = *m.Item
		} else {
			m.Item = &prev
		}

		m.CartItemID = strconv.Itoa(num + 1)
		m.Title = m.Item.Title
		if m.Price != nil {
			m.Discount = m.OrigPrice - *m.Price
			cart.Sum += *m.Price
		} else {
			m.Discount = m.OrigPrice
		}
		cart.Discount += m.Discount
	}
	return &cart
}

func GetItemFromDB(itemID string) *Item {
	var item Item
	ctx := context.Background()
	err := db.NewSelect().
		Model(&item).
		Where("id = ?", itemID).
		Scan(ctx)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil
	} else if err != nil {
		panic(err)
	}
	return &item
}
