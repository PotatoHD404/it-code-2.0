package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
	"net/http"
)

func AddItemToCart(w http.ResponseWriter, r *http.Request) {
	item, err := GetItemFromDB(r.FormValue("item_id"))
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	cart, err := GetCartFromDB(mux.Vars(r)["cart_id"])
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	id, _ := shortid.Generate()
	price := item.Price
	cardItem := &CartItem{
		CartItemID: id,
		ItemID:     item.ID,
		CartID:     cart.ID,
		Price:      &price,
		OrigPrice:  item.Price,
	}

	cart.Items = append(cart.Items, cardItem)
	_, err = db.
		NewInsert().
		Model(cardItem).
		Exec(ctx)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	if len(cart.Promos) > 0 {
		_, err := db.NewDelete().Model(&cart.Promos).WherePK().Exec(ctx)
		if err != nil {
			http.Error(w, err.Error(), 422)
			return
		}
	}
	err = cart.ResetCart()
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	err = cart.ApplyPromocode()
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	if len(cart.Promos) > 0 {
		_, err := db.NewInsert().Model(&cart.Promos).Exec(ctx)
		if err != nil {
			http.Error(w, err.Error(), 422)
			return
		}
	}
	_, err = db.NewUpdate().Model(cart).WherePK().Exec(ctx)

	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	_, err = db.NewUpdate().Model(&cart.Items).Column("price", "orig_price", "cart_item_id").Bulk().Exec(ctx)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func ApplyPromoToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cart, err := GetCartFromDB(mux.Vars(r)["cart_id"])
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	cart.Promocode = r.FormValue("promocode")
	if len(cart.Promos) > 0 {
		_, err := db.NewDelete().Model(&cart.Promos).WherePK().Exec(ctx)
		if err != nil {
			http.Error(w, err.Error(), 422)
			return
		}
	}
	err = cart.ResetCart()
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	err = cart.ApplyPromocode()
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	if len(cart.Promos) > 0 {
		_, err := db.NewInsert().Model(&cart.Promos).Exec(ctx)
		if err != nil {
			http.Error(w, err.Error(), 422)
			return
		}
	}
	_, err = db.NewUpdate().Model(cart).WherePK().Exec(ctx)

	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	_, err = db.NewUpdate().Model(&cart.Items).Column("price", "orig_price", "cart_item_id").Bulk().Exec(ctx)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

}
func ArrContainsCartPromo(s []*Promo, e *CartPromo) int {
	for num, a := range s {
		if a.ID == (*e).PromoID {
			return num
		}
	}
	return -1
}
func GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart, err = GetCartFromDB(mux.Vars(r)["cart_id"])
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	if cart == nil {
		http.Error(w, "Cart not found", 422)
		return
	}
	for _, item := range cart.Items {
		if item.Price == nil {
			item.Price = new(float)
		}
	}
	var promos []*Promo
	for _, promo := range cart.Promos {
		promos = append(promos, &Promo{ID: promo.PromoID})
	}
	promos, err = GetPromosFromDB(promos)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	for _, promo := range cart.Promos {
		num := ArrContainsCartPromo(promos, promo)
		promo.Title = promos[num].Title
	}
	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
}

func CreateCart(w http.ResponseWriter, _ *http.Request) {

	id, _ := shortid.Generate()
	var cart = Cart{CartID: id}
	_, err := db.NewInsert().Model(&cart).Exec(ctx)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	type Response struct {
		ShortID string `json:"cart_id"`
	}
	if err != nil {
		http.Error(w, err.Error(), 422)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(Response{ShortID: id})
		if err != nil {
			http.Error(w, err.Error(), 422)
			return
		}
	}
}
