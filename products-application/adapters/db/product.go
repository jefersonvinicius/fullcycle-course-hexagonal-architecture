package db

import (
	"database/sql"
	"log"

	"github.com/jefersonvinicius/fullcycle-course-hexagonal-architecture/products-application/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (p *ProductDB) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDB) Save(product application.ProductInterface) (application.ProductInterface, error) {

	if productAlreadyExists(product.GetID(), p) {
		stmt, err := p.db.Prepare("update products set name = ?, price = ?, status = ? where id = ?")
		if err != nil {
			return nil, err
		}

		_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
		if err != nil {
			return nil, err
		}

		return product, nil

	} else {
		stmt, err := p.db.Prepare("insert into products (id, name, price, status) values (?, ?, ?, ?)")
		if err != nil {
			return nil, err
		}

		_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
		if err != nil {
			return nil, err
		}

		return product, nil
	}
}

func productAlreadyExists(id string, productDB *ProductDB) bool {
	stmt, err := productDB.db.Prepare("select count(*) as count from products where id = ?")
	checkError(err)

	var count int
	err = stmt.QueryRow(id).Scan(&count)
	checkError(err)

	return count == 1
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
