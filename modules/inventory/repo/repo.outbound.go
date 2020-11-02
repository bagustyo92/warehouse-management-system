package repo

import "github.com/bagustyo92/wms/modules/inventory/model"

func (pr *ProductRepo) GetOutbound(outboundID uint) (*model.Outbound, error) {
	outbound := &model.Outbound{}
	outbound.ID = outboundID

	if err := pr.gdb.Model(&outbound).Find(&outbound).Error; err != nil {
		return nil, err
	}

	return outbound, nil
}

func (pr *ProductRepo) GetOutbounds(query *model.Query) (interface{}, error) {
	outbounds := []model.Outbound{}

	if err := pr.gdb.Find(&outbounds).Error; err != nil {
		return nil, err
	}

	return outbounds, nil
}
