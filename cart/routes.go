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
	//w.Header().Set("Content-Type", "application/json")

	var item = GetItemFromDB(r.FormValue("item_id"))
	var cart = GetCartFromDB(mux.Vars(r)["cart_id"])
	//cart.Items = append(cart.Items, GetItemFromDB(itemId)
	//Order("order.id ASC").
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
	_, err := db.
		NewInsert().
		Model(cardItem).
		Exec(ctx)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
		//panic(err)
	}
	if len(cart.Promos) > 0 {
		_, err := db.NewDelete().Model(&cart.Promos).WherePK().Exec(ctx)
		if err != nil {
			http.Error(w, err.Error(), 422)
			panic(err)
		}
	}
	cart.ResetCart()
	cart.ApplyPromocode(w)
	if len(cart.Promos) > 0 {
		_, err := db.NewInsert().Model(&cart.Promos).Exec(ctx)
		if err != nil {
			http.Error(w, err.Error(), 422)
			panic(err)
		}
	}
	_, err = db.NewUpdate().Model(cart).WherePK().Exec(ctx)

	if err != nil {
		http.Error(w, err.Error(), 422)
		return
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
		http.Error(w, err.Error(), 422)
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
	if len(cart.Promos) > 0 {
		_, err := db.NewDelete().Model(&cart.Promos).WherePK().Exec(ctx)
		if err != nil {
			http.Error(w, err.Error(), 422)
			panic(err)
		}
	}
	cart.ResetCart()
	cart.ApplyPromocode(w)
	if len(cart.Promos) > 0 {
		_, err := db.NewInsert().Model(&cart.Promos).Exec(ctx)
		if err != nil {
			http.Error(w, err.Error(), 422)
			panic(err)
		}
	}
	_, err := db.NewUpdate().Model(cart).WherePK().Exec(ctx)

	if err != nil {
		panic(err)
	}

	_, err = db.NewUpdate().Model(&cart.Items).Column("price", "orig_price", "cart_item_id").Bulk().Exec(ctx)
	if err != nil {
		panic(err)
	}

	//err = json.NewEncoder(w).Encode(cart)
	//if err != nil {
	//	panic(err)
	//}

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
	var cart = GetCartFromDB(mux.Vars(r)["cart_id"])
	for _, item := range cart.Items {
		if item.Price == nil {
			item.Price = new(float)
		}
	}
	var items []*Item
	for _, m := range cart.Items {
		items = append(items, &Item{ID: m.ItemID})
	}
	var promos []*Promo
	for _, promo := range cart.Promos {
		promos = append(promos, &Promo{ID: promo.PromoID})
	}
	promos = GetPromosFromDB(promos)
	for _, promo := range cart.Promos {
		num := ArrContainsCartPromo(promos, promo)
		promo.Title = promos[num].Title
	}
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
