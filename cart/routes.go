package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
	"net/http"
)

func TestFunc(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var promos []Promo
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

	var item = GetItemFromDB(r.FormValue("item_id"))
	var cart = GetCartFromDB(mux.Vars(r)["cart_id"])
	//cart.Items = append(cart.Items, GetItemFromDB(itemId)

	_, err := db.
		NewInsert().
		Model(
			&CartItem{
				ItemID:    item.ID,
				CartID:    cart.ID,
				Price:     &item.Price,
				OrigPrice: item.Price,
			}).
		Exec(ctx)
	if err != nil {
		w.WriteHeader(http.StatusCreated)
		http.Error(w, err.Error(), 422)
		//panic(err)
	} else {
		w.WriteHeader(http.StatusCreated)
		if err != nil {
			panic(err)
		}
	}
}

func ApplyPromoToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart = GetCartFromDB(mux.Vars(r)["cart_id"])
	cart.Promocode = r.FormValue("promocode")
	_, err := db.NewUpdate().Model(cart).WherePK().Exec(ctx)
	if err != nil {
		panic(err)
	}
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
