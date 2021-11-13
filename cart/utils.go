package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"net/http"
	"strconv"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", CreateCart).Methods("POST")
	r.HandleFunc("/{cart_id}/promocode", ApplyPromoToCart).Methods("POST")
	r.HandleFunc("/test", TestFunc).Methods("GET")
	r.HandleFunc("/{cart_id}/items", AddItemToCart).Methods("POST")
	r.HandleFunc("/{cart_id}", GetCart).Methods("GET")
	return r
}

func (cart *Cart) ResetCart() {
	cart.Sum = 0.0
	for i := 0; i < len(cart.Items); i++ {
		if cart.Items[i].Price == nil {
			_, err := db.NewDelete().Model(cart.Items[i]).WherePK().Exec(ctx)
			if err != nil {
				panic(err)
			}

		} else {
			cart.Items[i].Price = &cart.Items[i].OrigPrice
			cart.Items[i].Discount = 0.0
			cart.Sum += cart.Items[i].OrigPrice
		}
	}
	cart.Discount = 0.0
	cart.Promos = []*Promo{}
	_, err := db.NewUpdate().Model(cart).WherePK().Exec(ctx)
	if err != nil {
		panic(err)
	}
}
func ArrContains(s []*Item, e *CartItem) int {
	for num, a := range s {
		if a.ID == e.ItemID {
			return num
		}
	}
	return -1
}
func (m *CartItem) applyPromo(promo Promo) {
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
}

func (cart *Cart) ApplyPromo(w http.ResponseWriter) {
	var promocodes []Promo
	err := db.NewSelect().
		Model(&promocodes).
		Relation("ConditionItems").
		Relation("SelectorItems").
		Relation("GiftItems").
		Relation("Exclusions").
		Where("promocode = ?", "").
		WhereOr("promocode = ?", cart.Promocode).
		Order("id ASC").
		Scan(ctx)
	var tempItems []*Item
	for _, promo := range promocodes {
		copy(promo.ConditionItems, tempItems)
		for _, item := range cart.Items {
			if num := ArrContains(promo.ConditionItems, item); num != -1 {
				item.Selected = true
				promo.ConditionItems = append(promo.ConditionItems[:num], promo.ConditionItems[num+1:]...)
			}
			if len(promo.ConditionItems) == 0 {
				copy(tempItems, promo.ConditionItems)
				copy(promo.SelectorItems, tempItems)
				for _, item = range cart.Items {
					if num := ArrContains(promo.SelectorItems, item); num != -1 && !item.Used {
						item.applyPromo(promo)
						promo.SelectorItems = append(promo.SelectorItems[:num], promo.SelectorItems[num+1:]...)
					}
					if len(promo.ConditionItems) == 0 {
						break
					}
				}
				copy(tempItems, promo.SelectorItems)
			}
		}
	}
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(w).Encode(promocodes)
	if err != nil {
		panic(err)
	}

	_, err = db.NewUpdate().Model(cart).WherePK().Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func newDB() *bun.DB {
	dsn := "itcode2021:itcode2021@/itcode"
	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, mysqldialect.New())
	db.RegisterModel(
		(*PromoExclusions)(nil),
		(*PromoConditionItem)(nil),
		(*PromoGiftItem)(nil),
		(*PromoItemSelector)(nil),
		(*CartItem)(nil),
		(*CartPromo)(nil))
	return db
}

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
	for _, m := range cart.Items {
		if prev.ID != m.ItemID {
			m.Item = GetItemFromDB(strconv.Itoa(int(m.ItemID)))
			prev = *m.Item
		} else {
			m.Item = &prev
		}
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
