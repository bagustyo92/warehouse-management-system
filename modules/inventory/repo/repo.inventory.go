package repo

import (
	"fmt"

	"github.com/bagustyo92/wms/modules/inventory/model"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
)

func (pr *ProductRepo) InsertProduct(product *model.Product) error {
	return pr.gdb.Create(&product).Error
}

func (pr *ProductRepo) GetProduct(productID uint) (*model.Product, error) {
	product := model.Product{}
	stock := []model.Stock{}
	inbound := []model.Inbound{}
	outbound := []model.Outbound{}

	product.ID = productID
	if err := pr.gdb.Preloads(&stock).Preloads(&inbound).Preloads(&outbound).Find(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (pr *ProductRepo) GetProducts(query *model.Query) (interface{}, error) {
	products := []model.Product{}

	var finalQuery *gorm.DB
	fil := make(map[string]interface{})

	if query.Search != nil {
		finalQuery = pr.gdb.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *query.Search))
	} else {
		if query.SKU != nil {
			fil["sku"] = query.SKU
		}

		finalQuery = pr.gdb.Where(fil)
	}

	// pagination execution
	paginator := pagination.Paging(&pagination.Param{
		DB:      finalQuery.Find(&products),
		Page:    query.PageNumber,
		Limit:   query.PageLimit,
		OrderBy: []string{"id desc"},
		ShowSQL: false,
	}, &products)

	return paginator, nil
}

func (pr *ProductRepo) DeleteProduct(productID uint) error {
	return pr.gdb.Where("id = ?", productID).Delete(&model.Product{}).Error
}

func (pr *ProductRepo) UpdateProduct(product *model.Product) error {
	return pr.gdb.Model(product).Update(product).Error
}
