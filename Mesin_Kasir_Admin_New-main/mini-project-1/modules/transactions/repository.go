package transactions

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetAll() ([]Transaction, error){
	var transactions []Transaction
	result := repo.DB.Select("id", "timestamp", "total").Find(&transactions)

	return transactions, result.Error
}

func (repo Repository) GetById(id int) (*Transaction, error){
	var transaction *Transaction
	result := repo.DB.Preload("Admin").Preload("Items.Product").First(&transaction, id)

	return transaction, result.Error
}

func (repo Repository) Create(transaction *Transaction) error{
	result := repo.DB.Create(transaction)

	return result.Error
}