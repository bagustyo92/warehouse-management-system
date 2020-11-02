package repo

import "github.com/bagustyo92/wms/modules/inventory/model"

func (ps *ProductRepo) GetStock(stockID uint) (*model.Stock, error) {
	stock := model.Stock{}
	stock.ID = stockID

	if err := ps.gdb.Model(&stock).Find(&stock).Error; err != nil {
		return nil, err
	}

	return &stock, nil
}

func (ps *ProductRepo) GetStocks(query *model.Query) (interface{}, error) {
	stocks := []model.Stock{}

	if err := ps.gdb.Find(&stocks).Error; err != nil {
		return nil, err
	}

	return stocks, nil
}
