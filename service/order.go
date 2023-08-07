package service

type OrderHeader struct {
	OHId     int `gorm:"primaryKey"`
	UserName string
	Phone    string
	Address  int
}

func OrderList() *[]OrderHeader {
	orders := &[]OrderHeader{}

	db := Context()

	db.Table("OrderHeader").Find(&orders)

	return orders
}
