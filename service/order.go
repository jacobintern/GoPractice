package service

import (
	"fmt"
)

type OrderHeader struct {
	OrderHeaderID int         `gorm:"column:OHId;type:INT;autoIncrement:true;primaryKey"`
	UserName      string      `gorm:"type:VARCHAR(100)"`
	Phone         string      `gorm:"type:VARCHAR(10)"`
	Address       int         `gorm:"type:INT"`
	OrderBodys    []OrderBody `gorm:"foreignKey:OrderHeaderID"`
}

type OrderBody struct {
	OrderBodyID   int `gorm:"column:OBId;type:INT;autoIncrement:true;primaryKey"`
	ProductID     int `gorm:"column:PId;type:INT"`
	Quantity      int `gorm:"type:INT"`
	OrderHeaderID int `gorm:"column:OHId;type:INT"`
}

func (*OrderBody) TableName() string {
	return "OrderBody"
}

func (*OrderHeader) TableName() string {
	return "OrderHeader"
}

func GetOrderList() *[]OrderHeader {
	var orders []OrderHeader

	db := Context()

	err := db.Preload("OrderBodys").Find(&orders).Error

	if err != nil {
		fmt.Println(err)
	}

	return &orders
}

type IOrder interface {
	GetOrder()
}

func GetOrder(id int) *OrderHeader {
	var order OrderHeader

	db := Context()

	err := db.Preload("OrderBodys").First(&order, id).Error

	if err != nil {
		fmt.Println(err)
	}

	return &order
}

// func GetOrder(name string) *OrderHeader {
// 	var order OrderHeader

// 	db := Context()

// 	err := db.Preload("OrderBodys").First(&order, name).Error

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return &order
// }

func Create(model *OrderHeader) bool {

	return true
}
