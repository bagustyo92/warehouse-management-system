package service

import "github.com/bagustyo92/wms/modules/inventory/model"

func (ps ProductService) InsertProduct(product *model.Product) error {
	return ps.ps.InsertProduct(product)
}

func (ps ProductService) GetProduct(productID uint) (*model.Product, error) {
	return ps.ps.GetProduct(productID)
}

func (ps ProductService) GetProducts(query *model.Query) (interface{}, error) {
	return ps.ps.GetProducts(query)
}

func (ps ProductService) DeleteProduct(productID uint) error {
	return ps.ps.DeleteProduct(productID)
}

func (ps ProductService) UpdateProduct(productID *model.Product) error {
	return ps.ps.UpdateProduct(productID)
}
