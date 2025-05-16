package models

import "time"

// Models (equivalent to tables)
type Customer struct {
	CustomerID string `gorm:"primaryKey;size:50"`
	Name       string
	Email      string
	Address    string
}

type Product struct {
	ProductID string `gorm:"primaryKey;size:50"`
	Name      string
	Category  string
}

type Order struct {
	OrderID      string `gorm:"primaryKey;size:50"`
	CustomerID   string `gorm:"size:50"`
	Region       string
	DateOfSale   time.Time
	PaymentType  string
	ShippingCost float64
}

type OrderDetails struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	OrderID      string `gorm:"size:50"`
	ProductID    string `gorm:"size:50"`
	QuantitySold int
	UnitPrice    float64
	Discount     float64
}

type OverAllData struct {
	Quantity float64 `json:"quantity" gorm:"Column:quantity"`
	Category string  `json:"Category,omitempty" gorm:"Column:category"`
	Region   string  `json:"Region,omitempty" gorm:"Column:region"`
}

// type Customer struct {
// 	CustomerID string `gorm:"primaryKey;size:50"`
// 	Name       string
// 	Email      string
// 	Address    string
// }

// type Order struct {
// 	OrderID      string   `gorm:"primaryKey;size:50"`
// 	CustomerID   string   `gorm:"size:50;not null"`
// 	Customer     Customer `gorm:"foreignKey:CustomerID;references:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// 	Region       string
// 	DateOfSale   time.Time
// 	PaymentType  string
// 	ShippingCost float64
// }

// type OrderItem struct {
// 	ID           uint    `gorm:"primaryKey;autoIncrement"`
// 	OrderID      string  `gorm:"size:50;not null"` // foreign key column in order_items table
// 	Order        Order   `gorm:"foreignKey:OrderID;references:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// 	ProductID    string  `gorm:"size:50;not null"`
// 	Product      Product `gorm:"foreignKey:ProductID;references:ProductID"`
// 	QuantitySold int
// 	UnitPrice    float64
// 	Discount     float64
// }

// type Product struct {
// 	ProductID string `gorm:"primaryKey;size:50"`
// 	Name      string
// 	Category  string
// }
