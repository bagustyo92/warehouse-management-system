package service

import (
	"github.com/bagustyo92/wms/modules/inventory/model"
	"github.com/bagustyo92/wms/modules/inventory/repo"
)

type ProductInterface interface {
	InsertProduct(product *model.Product) error
	GetProduct(productID uint) (*model.Product, error)
	GetProducts(query *model.Query) (interface{}, error)
	DeleteProduct(productID uint) error
	UpdateProduct(productID *model.Product) error
}

type ProductService struct {
	ps repo.ProductsInterface
}

func NewProductService(productRepo repo.ProductsInterface) ProductInterface {
	return &ProductService{productRepo}
}
