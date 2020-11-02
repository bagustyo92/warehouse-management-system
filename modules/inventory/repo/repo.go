package repo

import (
	"github.com/bagustyo92/wms/modules/inventory/model"
	"github.com/jinzhu/gorm"
)

type ProductsInterface interface {
	// Inventory
	InsertProduct(product *model.Product) error
	GetProduct(productID uint) (*model.Product, error)
	GetProducts(query *model.Query) (interface{}, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(productID uint) error

	// Stocks
	GetStock(stockID uint) (*model.Stock, error)
	GetStocks(query *model.Query) (interface{}, error)

	// Inbound
	GetInbound(inboundID uint) (*model.Inbound, error)
	GetInbounds(query *model.Query) (interface{}, error)

	// Outbound
	GetOutbound(outboundID uint) (*model.Outbound, error)
	GetOutbounds(query *model.Query) (interface{}, error)
}

type ProductRepo struct {
	gdb *gorm.DB
	// rdb *redis.Conn
}

func NewProductsRepo(gdb *gorm.DB) ProductsInterface {
	return &ProductRepo{gdb}
}
