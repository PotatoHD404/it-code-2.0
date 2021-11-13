package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"log"
	"net/http"
)

var db *bun.DB
var ctx context.Context

func main() {
	ctx = context.Background()
	var router = newRouter()
	db = newDB()
	log.Fatal(http.ListenAndServe(":8080", router))
}
