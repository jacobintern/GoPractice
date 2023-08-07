package service

type OrderHeader struct {
	OHId     int         `gorm:"type:INT;autoIncrement:true;primaryKey"`
	UserName string      `gorm:"type:VARCHAR(100)"`
	Phone    string      `gorm:"type:VARCHAR(10)"`
	Address  int         `gorm:"type:INT"`
	Bodys    []OrderBody `gorm:"foreignKey:OHId;references:OHId"`
}

type OrderBody struct {
	OBId     int         `gorm:"type:INT;autoIncrement:true;primaryKey"`
	PId      int         `gorm:"type:INT"`
	Quantity int         `gorm:"type:INT"`
	OHId     int         `gorm:"type:INT"`
	Header   OrderHeader `gorm:"foreignKey:OHId"`
}

func OrderList() *[]OrderHeader {
	orders := &[]OrderHeader{}

	db := Context()

	db.Table("OrderHeader").Find(&orders)

	return orders
}

func Create(model *OrderHeader) bool {

	return true
}
