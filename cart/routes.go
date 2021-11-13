package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
	"net/http"
)

func TestFunc(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var promos []Promo
	ctx := context.Background()
	if err := db.NewSelect().
		Model(&promos).
		Relation("ConditionItems").
		Relation("SelectorItems").
		Relation("GiftItems").
		Relation("Exclusions").
		Scan(ctx); err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(promos); err != nil {
		panic(err)
	}
}

func AddItemToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func ApplyPromoToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func GetCartFromDB(cartId string) *Cart {
	var cart = new(Cart)
	ctx := context.Background()
	err := db.NewSelect().
		Model(cart).
		Where("cart_id = ?", cartId).
		Relation("Items").
		Relation("Promos").
		Scan(ctx)
	if err != nil && err.Error() != "sql: no rows in result set" {
		panic(err)
	}
	if cart.Items == nil {
		cart.Items = []Item{}
	}
	if cart.Promos == nil {
		cart.Promos = []Promo{}
	}
	return cart
}
func GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart = GetCartFromDB(mux.Vars(r)["cart_id"])
	err := json.NewEncoder(w).Encode(cart)
	if err != nil {
		panic(err)
	}

}

func CreateCart(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := shortid.Generate()
	var cart = Cart{CartID: id}
	ctx := context.Background()
	_, err := db.NewInsert().Model(&cart).Exec(ctx)
	if err != nil {
		panic(err)
	}
	type Response struct {
		ShortID string `json:"cart_id"`
	}
	err = json.NewEncoder(w).Encode(Response{ShortID: id})
	if err != nil {
		panic(err)
	}
}
