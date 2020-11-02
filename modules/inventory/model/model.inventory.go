package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name      string
	SKU       string
	Expirable bool
	Stocks    []Stock    `gorm:"foreignKey:productID"`
	Inbounds  []Inbound  `gorm:"foreignKey:productID"`
	Outbounds []Outbound `gorm:"foreignKey:productID"`
}

type Stock struct {
	gorm.Model
	InboundDate     time.Time
	ExpiredDate     time.Time
	ProductID       uint
	InboundQuantity int64
	CurrentStock    int64
}

type Inbound struct {
	gorm.Model
	InboundDate    time.Time
	ExpiredDate    time.Time
	ProductID      uint
	Quantity       int64
	ItemPrice      int64
	TotalPrice     int64
	PurchaseNumber string
}

type Outbound struct {
	gorm.Model
	OutboundDate    time.Time
	ProductID       uint
	Quantity        int64
	ItemPrice       int64
	TotalPrice      int64
	UseCase         string
	ReferenceNumber string
}

type Query struct {
	Search *string
	SKU    *string

	PageNumber int
	PageLimit  int
}
