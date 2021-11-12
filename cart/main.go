package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"net/http"
)

var db *bun.DB

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	//r.HandleFunc("/api/books", CreateBook).Methods("POST")
	//r.HandleFunc("/api/books/{id}", UpdateBook).Methods("PUT")
	//r.HandleFunc("/api/books/{id}", DeleteBook).Methods("DELETE")
	return r
}

func newDB() *bun.DB {
	dsn := "mysql://itcode2021:itcode2021@(mysql:3306)/itcode"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}

// User struct
type User struct {
	bun.BaseModel `bun:"users"`
	ID            int64  `json:"id"`
	Name          string `bun:"name" json:"name"`
	Surname       string `bun:"surname" json:"surname"`
}

//MYSQL_ROOT_PASSWORD: root
//MYSQL_DATABASE: itcode
//MYSQL_USER: itcode2021
//MYSQL_PASSWORD: itcode2021

func main() {
	db = newDB()
	var router = newRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var user User
	db := newDB()
	defer func(db *bun.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	db.RegisterModel((*User)(nil))
	params := mux.Vars(r)
	//params["id"]
	if err := db.
		NewSelect().
		Model(&user).
		Where("? = ?", bun.Ident("id"), params["id"]).
		Scan(ctx); err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

//
////https://github.com/uptrace/bun/blob/master/example/fixture/main.go
//

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()
	var users []User
	db := newDB()
	defer func(db *bun.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	db.RegisterModel((*User)(nil))
	if err := db.NewSelect().Model(&users).Scan(ctx); err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

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
