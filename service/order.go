package service

type OrderHeader struct {
	OHId     int `gorm:"primaryKey"`
	UserName string
	Phone    string
	Address  int
}

func OrderList() OrderHeader {
	res := OrderHeader{}

	return res
}
