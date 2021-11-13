package main

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"log"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", CreateCart).Methods("POST")
	r.HandleFunc("/{cart_id}/items", AddItemToCart).Methods("POST")
	r.HandleFunc("/{cart_id}/promocode", ApplyPromoToCart).Methods("POST")
	r.HandleFunc("/{cart_id}", GetCart).Methods("GET")
	return r
}

func newDB() *bun.DB {
	dsn := "itcode2021:itcode2021@/itcode"
	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, mysqldialect.New())
	db.RegisterModel((*PromoConditionItem)(nil), (*PromoExclusions)(nil), (*PromoGiftItem)(nil), (*PromoItemSelector)(nil), (*CartItem)(nil), (*CartPromo)(nil), (*Promo)(nil), (*Item)(nil))
	return db
}

type Item struct {
	bun.BaseModel `bun:"items"`
	ID            uint64  `json:"id"`
	Title         string  `bun:"title" json:"title"`
	Price         float64 `bun:"price" json:"price"`
}

var ActionType = [...]string{
	"percent_discount",
	"price_discount",
	"gift",
}

var ScopeType = [...]string{
	"order",
	"item",
}

type Promo struct {
	bun.BaseModel  `bun:"promos"`
	ID             uint64  `json:"id"`
	Promocode      string  `bun:"promocode" json:"promocode"`
	Priority       uint64  `bun:"priority" json:"priority"`
	Action         string  `bun:"action" json:"-"`
	Discount       float64 `bun:"discount" json:"discount"`
	Title          string  `bun:"title" json:"title"`
	Scope          string  `bun:"scope" json:"-"`
	ConditionItems []Item `bun:"m2m:promo_condition_item" json:"condition_items"`
	//SelectorItems  []Item  `bun:"m2m:promo_item_selector,join:Promo=Item" json:"selector_items"`
	//GiftItems      []Item  `bun:"m2m:promo_gift_items,join:Promo=Item" json:"gift_items"`
	//Exclusions     []Promo `bun:"m2m:promo_exclusions,join:Promo=ExPromo" json:"exclusions"`
}

type PromoConditionItem struct {
	bun.BaseModel `bun:"promo_condition_item"`
	ID            uint64
	//PromoID       uint64 `bun:"promo_id" json:"promo_id"`
	Promo         *Promo `bun:"promo_id,rel:belongs-to" json:"promo"`
	//ItemID        uint64 `bun:"item_id" json:"item_id"`
	Item          *Item  `bun:"item_id,rel:belongs-to" json:"item"`
}

type PromoExclusions struct {
	bun.BaseModel `bun:"promo_exclusions"`
	ID            uint64 `json:"id"`
	Promo         *Promo `bun:"promo_id,rel:belongs-to,join:promo_id=id" json:"promo"`
	ExPromo       *Promo `bun:"exclusion_promo_id,rel:belongs-to,join:exclusion_promo_id=id" json:"ex_promo"`
}

type PromoGiftItem struct {
	bun.BaseModel `bun:"promo_gift_items"`
	ID            uint64 `json:"id"`
	Promo         *Promo `bun:"promo_id,rel:belongs-to,join:promo_id=id" json:"promo"`
	Item          *Item  `bun:"item_id,rel:belongs-to,join:item_id=id" json:"item"`
}

type PromoItemSelector struct {
	bun.BaseModel `bun:"promo_item_selector"`
	//ID            uint64 `json:"id"`
	Promo *Promo `bun:"promo_id,rel:belongs-to,join:promo_id=id" json:"promo"`
	Item  *Item  `bun:"item_id,rel:belongs-to,join:item_id=id" json:"item"`
}

//type Profile struct {
//	ID     int64 `bun:",pk"`
//	Lang   string
//	Active bool
//	UserID int64
//}
//
//// User has many profiles.
//type User struct {
//	ID       int64 `bun:",pk"`
//	Name     string
//	Profiles []*Profile `bun:"rel:has-many,join:id=user_id"`
//}

type Cart struct {
	bun.BaseModel `bun:"orders"`
	ID            uint64 `json:"id"`
	CartId        string `bun:"cart_id" json:"cart_id"`
	Items         []Item `bun:"m2m:cart_items,join:Cart=Item" json:"items"`
	Promos        []Item `bun:"m2m:cart_promos,join:Cart=Promo" json:"promos"`
	Promocode     string `bun:"promocode" json:"promocode"`
}

type CartItem struct {
	bun.BaseModel `bun:"cart_items"`
	ID            uint64          `json:"id"`
	Price         sql.NullFloat64 `bun:"price" json:"price"`
	Cart          *Cart           `bun:"cart_id,rel:belongs-to,join:cart_id=id" json:"cart"`
	Item          *Item           `bun:"item_id,rel:belongs-to,join:item_id=id" json:"item"`
}

type CartPromo struct {
	bun.BaseModel `bun:"cart_promos"`
	ID            uint64          `json:"id"`
	Price         sql.NullFloat64 `bun:"price" json:"price"`
	Cart          *Cart           `bun:"cart_id,rel:belongs-to,join:cart_id=id" json:"cart"`
	Promo         *Promo          `bun:"promo_id,rel:belongs-to,join:promo_id=id" json:"item"`
}

func main() {
	var db = newDB()
	err := db.Close()
	if err != nil {
		panic(err)
	}
	var router = newRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func CreateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	//var promo []Promo
	var test []PromoConditionItem
	var db = newDB()
	//defer func(db *bun.DB) {
	//	err := db.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(db)
	//params := mux.Vars(r)
	////params["id"]
	if err := db.
		NewSelect().
		Model(&test).
		Relation("Promo").
		Relation("Item").
		Scan(ctx); err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(test); err != nil {
		panic(err)
	}
}

func AddItemToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//ctx := context.Background()
	//var user User
	//db := newDB()
	//defer func(db *bun.DB) {
	//	err := db.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(db)
	//db.RegisterModel((*User)(nil))
	//params := mux.Vars(r)
	////params["id"]
	//if err := db.
	//	NewSelect().
	//	Model(&user).
	//	Where("? = ?", bun.Ident("id"), params["id"]).
	//	Scan(ctx); err != nil {
	//	panic(err)
	//}
	//if err := json.NewEncoder(w).Encode(user); err != nil {
	//	panic(err)
	//}
}

func ApplyPromoToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//ctx := context.Background()
	//var user User
	//db := newDB()
	//defer func(db *bun.DB) {
	//	err := db.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(db)
	//db.RegisterModel((*User)(nil))
	//params := mux.Vars(r)
	////params["id"]
	//if err := db.
	//	NewSelect().
	//	Model(&user).
	//	Where("? = ?", bun.Ident("id"), params["id"]).
	//	Scan(ctx); err != nil {
	//	panic(err)
	//}
	//if err := json.NewEncoder(w).Encode(user); err != nil {
	//	panic(err)
	//}
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//ctx := context.Background()
	//var user User
	//db := newDB()
	//defer func(db *bun.DB) {
	//	err := db.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(db)
	//db.RegisterModel((*User)(nil))
	//params := mux.Vars(r)
	////params["id"]
	//if err := db.
	//	NewSelect().
	//	Model(&user).
	//	Where("? = ?", bun.Ident("id"), params["id"]).
	//	Scan(ctx); err != nil {
	//	panic(err)
	//}
	//if err := json.NewEncoder(w).Encode(user); err != nil {
	//	panic(err)
	//}
}

//
////https://github.com/uptrace/bun/blob/master/example/fixture/main.go
//

//func GetUsers(w http.ResponseWriter, _ *http.Request) {
//	ctx := context.Background()
//	var users []User
//	db := newDB()
//	defer func(db *bun.DB) {
//		err := db.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(db)
//	db.RegisterModel((*User)(nil))
//	if err := db.NewSelect().Model(&users).Scan(ctx); err != nil {
//		panic(err)
//	}
//	if err := json.NewEncoder(w).Encode(users); err != nil {
//		panic(err)
//	}
//}

// Init books var as a slice Book struct

//// GetBooks Get all books
//func GetBooks(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(books)
//}
//
//// GetBook Get single book
//func GetBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r) // Gets params
//	// Loop through books and find one with the id from the params
//	for _, item := range books {
//		if item.ID == params["id"] {
//			_ = json.NewEncoder(w).Encode(item)
//			return
//		}
//	}
//	_ = json.NewEncoder(w).Encode(&Book{})
//}
//
//// CreateBook Add new book
//func CreateBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var book Book
//	_ = json.NewDecoder(r.Body).Decode(&book)
//	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
//	books = append(books, book)
//	_ = json.NewEncoder(w).Encode(book)
//}
//
//// UpdateBook Update book
//func UpdateBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//	for index, item := range books {
//		if item.ID == params["id"] {
//			books = append(books[:index], books[index+1:]...)
//			var book Book
//			_ = json.NewDecoder(r.Body).Decode(&book)
//			book.ID = params["id"]
//			books = append(books, book)
//			_ = json.NewEncoder(w).Encode(book)
//			return
//		}
//	}
//}
//
//// DeleteBook Delete book
//func DeleteBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//	for index, item := range books {
//		if item.ID == params["id"] {
//			books = append(books[:index], books[index+1:]...)
//			break
//		}
//	}
//	_ = json.NewEncoder(w).Encode(books)
//}

//Request sample
//{
//"isbn":"4545454",
//"title":"Book Three",
//"author":{"firstname":"Harry", "lastname":"White"}
//}
