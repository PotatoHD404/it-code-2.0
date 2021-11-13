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
	//Order("order.id ASC").
	id, _ := shortid.Generate()
	cardItem := &CartItem{
		CartItemID: id,
		ItemID:     item.ID,
		CartID:     cart.ID,
		Price:      &item.Price,
		OrigPrice:  item.Price,
	}

	cart.Items = append(cart.Items, cardItem)
	_, err := db.
		NewInsert().
		Model(cardItem).
		Exec(ctx)
	if err != nil {
		w.WriteHeader(http.StatusCreated)
		http.Error(w, err.Error(), 422)
		//panic(err)
	}
	cart.ResetCart()
	cart.ApplyPromocode()
	_, err = db.NewUpdate().Model(cart).WherePK().Exec(ctx)

	if err != nil {
		w.WriteHeader(http.StatusCreated)
		http.Error(w, err.Error(), 422)
		//panic(err)
	}
	//values := db.NewValues(&cart.Items)
	//_, err = db.NewUpdate().
	//	With("_data", values).
	//	Model((*CartItem)(nil)).
	//	TableExpr("_data").
	//	Where("card_item.id = _data.id").
	//	Exec(ctx)
	_, err = db.NewUpdate().Model(&cart.Items).Column("price", "orig_price", "cart_item_id").Bulk().Exec(ctx)
	if err != nil {
		//w.WriteHeader(http.StatusCreated)
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

	id, _ := shortid.Generate()
	var cart = Cart{CartID: id}
	_, err := db.NewInsert().Model(&cart).Exec(ctx)
	if err != nil {
		panic(err)
	}
	type Response struct {
		ShortID string `json:"cart_id"`
	}
	if err != nil {
		//w.WriteHeader(http.StatusCreated)
		http.Error(w, err.Error(), 422)
		//panic(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(Response{ShortID: id})
		if err != nil {
			panic(err)
		}
	}
}
