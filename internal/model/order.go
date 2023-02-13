package model

type OrderStatus string

type Order struct {
	ID				string `gorm:"primarykey"`
	Status			OrderStatus
	ProductOrders	[]OrderStatus
	ReferenceID	string
}

type ProductOrderStatus	string

type ProductOrders struct {
	ID			string	`gorm:"primarykey"`
	OrderID		string
	OrderCode	string
	Quantity	int
	TotalPrice	int64
	Status		ProductOrderStatus
}

type OrderMenuProductRequest struct {
	OrderCode	string
	Quantity	int
}

type OrderMenuRequest struct {
	OrderProducts []OrderMenuProductRequest
	ReferenceID	string
}

type GetOrderInfoRequest struct {
	OrderID		string
}