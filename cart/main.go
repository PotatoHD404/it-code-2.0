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
	"os"
)

func newDB() *bun.DB {

	dsn := os.Getenv("MYSQL_USER") +
		":" +
		os.Getenv("MYSQL_PASSWORD") +
		"@tcp(" + os.Getenv("MYSQL_HOST") +
		":" +
		os.Getenv("MYSQL_PORT") +
		")/" +
		os.Getenv("MYSQL_DB")
	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
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

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", CreateCart).Methods("POST")
	r.HandleFunc("/{cart_id}/promocode", ApplyPromoToCart).Methods("POST")
	r.HandleFunc("/{cart_id}/items", AddItemToCart).Methods("POST")
	r.HandleFunc("/{cart_id}", GetCart).Methods("GET")
	return r
}

var db *bun.DB
var ctx context.Context

func main() {
	ctx = context.Background()
	var router = newRouter()
	db = newDB()
	log.Fatal(http.ListenAndServe(":8080", router))
}
