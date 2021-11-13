package main

import (
	"context"
	"database/sql"
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
	r.HandleFunc("/{cart_id}/promocode", ApplyPromoToCart).Methods("POST")
	r.HandleFunc("/test", TestFunc).Methods("GET")
	r.HandleFunc("/{cart_id}/items", AddItemToCart).Methods("POST")

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
	db.RegisterModel(
		(*PromoExclusions)(nil),
		(*PromoConditionItem)(nil),
		(*PromoGiftItem)(nil),
		(*PromoItemSelector)(nil),
		(*CartItem)(nil),
		(*CartPromo)(nil))
	return db
}

var db *bun.DB
var ctx context.Context

//var _ bun.AfterScanRowHook = (*CartItem)(nil)
//
//func (m *CartItem) AfterScanRow(_ context.Context) error {
//	m.Title = m.Item.Title
//	if m.Price != nil {
//		m.Discount = m.OrigPrice - *m.Price
//	} else {
//		m.Discount = m.OrigPrice
//	}
//	return nil
//}

func main() {
	ctx = context.Background()
	var router = newRouter()
	db = newDB()
	log.Fatal(http.ListenAndServe(":8080", router))
}



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
