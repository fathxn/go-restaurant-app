package model

type OrderStatus string

type Order struct {
	ID            string `gorm:"primaryKey"`
	Status        OrderStatus
	ProductOrders []ProductOrder
}

type ProductOrderStatus string

type ProductOrder struct {
	ID         string `gorm:"primaryKey"`
	OrderID    string
	OrderCode  string
	Quantity   int
	TotalPrice int64
	Status     ProductOrderStatus
}

type OrderMenuProductRequest struct {
	OrderCode string
	Quantity  int
}

type OrderMenuRequest struct {
	OrderProducts []OrderMenuProductRequest
}

type GetOrderInfoRequest struct {
	OrderID string
}
