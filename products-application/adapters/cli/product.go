package cli

import (
	"fmt"

	"github.com/jefersonvinicius/fullcycle-course-hexagonal-architecture/products-application/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}
		res, err := service.Enable(product)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product %s has been enabled.", res.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}
		res, err := service.Disable(product)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product %s has been disabled.", res.GetName())
	case "get":
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s", res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())
	default:
		result = fmt.Sprintf("Action '%s' doesn't exists", action)
	}

	return result, nil

}
