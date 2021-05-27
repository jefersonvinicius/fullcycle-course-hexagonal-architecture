package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/jefersonvinicius/fullcycle-course-hexagonal-architecture/products-application/adapters/db"
	"github.com/jefersonvinicius/fullcycle-course-hexagonal-architecture/products-application/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	sql := `
		CREATE TABLE products (
			id string,
			name string,
			price float,
			status string
		)
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	sql := `
		INSERT INTO products VALUES ("abc", "Product", 10, "disabled")
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()

	defer Db.Close()

	productDb := db.NewProductDB(Db)
	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, application.DISABLED, product.GetStatus())
}
