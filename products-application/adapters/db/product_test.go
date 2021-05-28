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

func TestProductDb_Save(t *testing.T) {
	setUp()

	defer Db.Close()

	product := application.NewProduct()
	product.Name = "Control"
	product.Price = 20

	productDb := db.NewProductDB(Db)
	result, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "Control", result.GetName())
	require.Equal(t, 20.0, result.GetPrice())
	require.Equal(t, 1, getAmountProductOnDatabase())

}

func getAmountProductOnDatabase() int {
	stmt, err := Db.Prepare("select count(*) as count from products")
	if err != nil {
		log.Fatal(err.Error())
	}
	result, err := stmt.Query()
	if err != nil {
		log.Fatal(err.Error())
	}
	var count int
	result.Scan(&count)
	return count
}
