package transactions

import (
	"errors"
	"mini-project/modules/products"
	"time"
)

type Transaction struct {
	ID        int                        `gorm:"primaryKey" json:"id"`
	Timestamp time.Time                  `json:"timestamp"`
	Total     int                        `gorm:"not null" json:"total"`
	AdminID   float64                    `json:"admin_id" gorm:"not null;ForeignKey:UserID"`
	Admin     User                       `json:"admin"`
	Items     []products.TransactionItem `json:"items"`
}

type User struct {
	ID   int
	Name string
}

var (
	ErrProductIdNotFound    = errors.New("Product id not found")
	ErrStockNotEnough       = errors.New("Stock not enough")
	ErrPoductHasBeenRemoved = errors.New("Product has been removed")
)
