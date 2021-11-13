package main

import (
	"database/sql"
	"github.com/uptrace/bun"
)

type Item struct {
	bun.BaseModel `bun:"items"`
	ID            uint32  `bun:"id,pk" json:"id"`
	Title         string  `bun:"title" json:"title"`
	Price         float32 `bun:"price" json:"price"`
}

type Promo struct {
	bun.BaseModel  `bun:"promos"`
	ID             uint32  `bun:"id,pk" json:"id"`
	Promocode      string  `bun:"promocode" json:"promocode"`
	Priority       uint32  `bun:"priority" json:"priority"`
	Action         string  `bun:"action" json:"action"`
	Discount       float32 `bun:"discount" json:"discount"`
	Title          string  `bun:"title" json:"title"`
	Scope          string  `bun:"scope" json:"scope"`
	ConditionItems []Item  `bun:"m2m:promo_condition_item,join:Promo=Item" json:"condition_items"`
	SelectorItems  []Item  `bun:"m2m:promo_item_selector,join:Promo=Item" json:"selector_items"`
	GiftItems      []Item  `bun:"m2m:promo_gift_items,join:Promo=Item" json:"gift_items"`
	Exclusions     []Promo `bun:"m2m:promo_exclusions,join:Promo=ExPromo" json:"exclusions"`
}

type PromoConditionItem struct {
	bun.BaseModel `bun:"promo_condition_item"`
	ID            uint32 `bun:"id,pk" json:"-"`
	PromoID       uint32 `bun:"promo_id" json:"-"`
	Promo         *Promo `bun:"rel:belongs-to,join:promo_id=id" json:"promo"`
	ItemID        uint32 `bun:"item_id" json:"-"`
	Item          *Item  `bun:"rel:belongs-to,join:item_id=id" json:"item"`
}

type PromoExclusions struct {
	bun.BaseModel `bun:"promo_exclusions"`
	ID            uint32 `bun:"id,pk" json:"-"`
	PromoID       uint32 `bun:"promo_id" json:"-"`
	Promo         *Promo `bun:"rel:belongs-to,join:promo_id=id" json:"promo"`
	ExPromoID     uint32 `bun:"exclusion_promo_id" json:"-"`
	ExPromo       *Promo `bun:"rel:belongs-to,join:exclusion_promo_id=id" json:"ex_promo"`
}

type PromoGiftItem struct {
	bun.BaseModel `bun:"promo_gift_items"`
	ID            uint32 `bun:"id,pk" json:"-"`
	PromoID       uint32 `bun:"promo_id" json:"-"`
	Promo         *Promo `bun:"rel:belongs-to,join:promo_id=id" json:"promo"`
	ItemID        uint32 `bun:"item_id" json:"-"`
	Item          *Item  `bun:"rel:belongs-to,join:item_id=id" json:"item"`
}

type PromoItemSelector struct {
	bun.BaseModel `bun:"promo_item_selector"`
	ID            uint32 `bun:"id,pk" json:"-"`
	PromoID       uint32 `bun:"promo_id" json:"-"`
	Promo         *Promo `bun:"rel:belongs-to,join:promo_id=id" json:"promo"`
	ItemID        uint32 `bun:"item_id" json:"-"`
	Item          *Item  `bun:"rel:belongs-to,join:item_id=id" json:"item"`
}

type Cart struct {
	bun.BaseModel `bun:"orders"`
	ID            uint32 `bun:"id,pk" json:"-"`
	CartID        string `bun:"cart_id" json:"cart_id"`
	Items         []Item `bun:"m2m:cart_items,join:Cart=Item" json:"items"`
	Promos        []Promo `bun:"m2m:cart_promos,join:Cart=Promo" json:"promos"`
	Promocode     string `bun:"promocode" json:"promocode"`
}

type CartItem struct {
	bun.BaseModel `bun:"cart_items"`
	ID            uint32          `bun:"id,pk" json:"-"`
	Price         sql.NullFloat64 `bun:"price" json:"price"`
	CartID        uint32          `bun:"cart_id" json:"-"`
	Cart          *Cart           `bun:"rel:belongs-to,join:cart_id=id" json:"cart"`
	ItemID        uint32          `bun:"item_id" json:"-"`
	Item          *Item           `bun:"rel:belongs-to,join:item_id=id" json:"item"`
}

type CartPromo struct {
	bun.BaseModel `bun:"cart_promos"`
	ID            uint32          `bun:"id,pk" json:"-"`
	Price         sql.NullFloat64 `bun:"price" json:"price"`
	CartID        uint32          `bun:"cart_id" json:"-"`
	Cart          *Cart           `bun:"rel:belongs-to,join:cart_id=id" json:"cart"`
	PromoID       uint32          `bun:"promo_id" json:"-"`
	Promo         *Promo          `bun:"rel:belongs-to,join:promo_id=id" json:"item"`
}
