package products

import (
	"errors"
	"time"
)

type Product struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"varchar(300)" json:"name"`
	Price     int    `gorm:"int" json:"price"`
	Stock     int    `gorm:"int" json:"stock"`
	DeletedAt *time.Time
}

type TransactionItem struct {
	ID            int                     `gorm:"primaryKey" json:"id"`
	TransactionID uint                    `gorm:"not null;foreignKey:TransactionID"`
	ProductID     uint                    `gorm:"not null;foreignKey:ProductID"`
	Quantity      int                     `gorm:"not null" json:"quantity"`
	Price         int                     `gorm:"not null" json:"price"`
	Product       *Product `json:"product"`
}

type CreateItemRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CreateTransactionRequest struct {
	Items []CreateItemRequest
}

type ResponseAddTransaction struct{
	ID int
	Name string
}

type ProductsResponse struct {
	Message string
	Data    []Product
}

type ProductResponse struct {
	Message string
	Data    *Product
}

type ResponseAddAndEditData struct {
	Message string
	Data    Product
}

type RequesBodyStatus struct {
	Status string `json:"status"`
}

var (
	ErrProductAlreadyDeleted = errors.New("product already deleted")
	ErrProductNotDeleted     = errors.New("product is not deleted yet")
	ErrInvalidStatus         = errors.New("invalid status")
	ErrChangedStatus         = errors.New("status data cannot changed")
	ErrPoductHasBeenRemoved  = errors.New("product has been removed")
	ErrProductIdNotFound     = errors.New("data has been deleted")
)
