package main

import (
	"database/sql"

	dbAdapter "github.com/jefersonvinicius/fullcycle-course-hexagonal-architecture/products-application/adapters/db"
	"github.com/jefersonvinicius/fullcycle-course-hexagonal-architecture/products-application/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDb := dbAdapter.NewProductDB(db)

	productService := application.NewProductService(productDb)
	product, _ := productService.Create("Control", 10)
	productService.Enable(product)
}
