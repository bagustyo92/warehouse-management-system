package repo

import "github.com/bagustyo92/wms/modules/inventory/model"

func (pr *ProductRepo) GetInbound(inboundID uint) (*model.Inbound, error) {
	inbound := &model.Inbound{}
	inbound.ID = inboundID

	if err := pr.gdb.Model(&inbound).Find(&inbound).Error; err != nil {
		return nil, err
	}

	return inbound, nil
}

func (pr *ProductRepo) GetInbounds(query *model.Query) (interface{}, error) {
	inbounds := []model.Inbound{}

	if err := pr.gdb.Find(&inbounds).Error; err != nil {
		return nil, err
	}

	return inbounds, nil
}
