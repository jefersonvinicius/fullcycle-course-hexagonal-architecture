package application_test

import (
	"testing"

	"github.com/jefersonvinicius/fullcycle-course-hexagonal-architecture/products-application/application"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Controle"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	if err != nil {
		t.Errorf("'err' must be nil; but got '%s'\n", err.Error())
	}

	product.Price = 0

	err = product.Enable()
	if err.Error() != "the price must be greater than zero to enable the product" {
		t.Errorf("'err' must be Error (The price must be greater than zero to enable the product); but got '%s'\n", err.Error())
	}
}
