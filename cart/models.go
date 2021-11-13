package main

import (
	"fmt"
	"github.com/uptrace/bun"
)

type float float32

func (n float) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", n)), nil
}

type Item struct {
	bun.BaseModel `bun:"items"`
	ID            uint32 `bun:"id,pk" json:"id"`
	Title         string `bun:"title" json:"title"`
	Price         float  `bun:"price" json:"price"`
}

type Promo struct {
	bun.BaseModel  `bun:"promos"`
	ID             uint32   `bun:"id,pk" json:"id"`
	Promocode      string   `bun:"promocode" json:"promocode"`
	Priority       uint32   `bun:"priority" json:"priority"`
	Action         string   `bun:"action" json:"action"`
	Discount       float    `bun:"discount" json:"discount"`
	Title          string   `bun:"title" json:"title"`
	Scope          string   `bun:"scope" json:"scope"`
	ConditionItems []*Item  `bun:"m2m:promo_condition_item,join:Promo=Item" json:"condition_items"`
	SelectorItems  []*Item  `bun:"m2m:promo_item_selector,join:Promo=Item" json:"selector_items"`
	GiftItems      []*Item  `bun:"m2m:promo_gift_items,join:Promo=Item" json:"gift_items"`
	Exclusions     []*Promo `bun:"m2m:promo_exclusions,join:Promo=ExPromo" json:"exclusions"`
	Applied        bool     `bun:"-" json:"-"`
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
	ID            uint32      `bun:"id,pk" json:"-"`
	CartID        string      `bun:"cart_id" json:"-"`
	Items         []*CartItem `bun:"rel:has-many,join:id=cart_id" json:"items"`
	Promos        []*Promo    `bun:"m2m:cart_promos,join:Cart=Promo" json:"promos"`
	Promocode     string      `bun:"promocode" json:"promocode"`
	Sum           float       `bun:"-" json:"cart_sum"`
	Discount      float       `bun:"-" json:"cart_discount"`
}

type CartItem struct {
	bun.BaseModel `bun:"cart_items"`
	ID            int    `bun:"id,pk" json:"-"`
	CartItemID    string `bun:"cart_item_id" json:"cart_item_id"`
	ItemID        uint32 `bun:"item_id" json:"item_id"`
	Title         string `bun:"-" json:"title"`
	Price         *float `bun:"price" json:"price"`
	OrigPrice     float  `bun:"orig_price" json:"original_price"`
	CartID        uint32 `bun:"cart_id" json:"-"`
	Cart          *Cart  `bun:"rel:belongs-to,join:cart_id=id" json:"-"`
	Item          *Item  `bun:"rel:belongs-to,join:item_id=id" json:"-"`
	Discount      float  `bun:"-" json:"discount"`
	Selected      bool   `bun:"-" json:"-"`
	Used          bool   `bun:"-" json:"-"`
}

type CartPromo struct {
	bun.BaseModel `bun:"cart_promos"`
	ID            uint32 `bun:"id,pk" json:"-"`
	Price         *float `bun:"price" json:"price"`
	CartID        uint32 `bun:"cart_id" json:"-"`
	Cart          *Cart  `bun:"rel:belongs-to,join:cart_id=id" json:"cart"`
	PromoID       uint32 `bun:"promo_id" json:"-"`
	Promo         *Promo `bun:"rel:belongs-to,join:promo_id=id" json:"item"`
}
