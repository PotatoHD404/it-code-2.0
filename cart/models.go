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
	ID             uint32   `bun:"id,pk" json:"promo_id"`
	Promocode      string   `bun:"promocode" json:"-"`
	Priority       uint32   `bun:"priority" json:"-"`
	Action         string   `bun:"action" json:"-"`
	Discount       float    `bun:"discount" json:"-"`
	Title          string   `bun:"title" json:"title"`
	Scope          string   `bun:"scope" json:"-"`
	MinOrderSum    *float   `bun:"min_order_sum" json:"-"`
	ConditionItems []*Item  `bun:"m2m:promo_condition_item,join:Promo=Item" json:"-"`
	SelectorItems  []*Item  `bun:"m2m:promo_item_selector,join:Promo=Item" json:"-"`
	GiftItems      []*Item  `bun:"m2m:promo_gift_items,join:Promo=Item" json:"-"`
	Exclusions     []*Promo `bun:"m2m:promo_exclusions,join:Promo=ExPromo" json:"-"`
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
	ID            uint32       `bun:"id,pk" json:"-"`
	CartID        string       `bun:"cart_id" json:"-"`
	Items         []*CartItem  `bun:"rel:has-many,join:id=cart_id" json:"items"`
	Promos        []*CartPromo `bun:"rel:has-many,join:id=cart_id" json:"promos"`
	Promocode     string       `bun:"promocode" json:"promocode"`
	Sum           float        `bun:"cart_sum" json:"cart_sum"`
	Discount      float        `bun:"cart_discount" json:"cart_discount"`
	OrigPrice     float        `bun:"-" json:"-"`
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
	CartID        uint32 `bun:"cart_id" json:"-"`
	Cart          *Cart  `bun:"rel:belongs-to,join:cart_id=id" json:"-"`
	PromoID       uint32 `bun:"promo_id" json:"promo_id"`
	Promo         *Promo `bun:"rel:belongs-to,join:promo_id=id" json:"-"`
	Title         string `bun:"-" json:"title"`
}
