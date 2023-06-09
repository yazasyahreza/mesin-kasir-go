package products

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetAllProducts() ([]Product, error) {
	var products []Product
	result := repo.DB.Unscoped().Find(&products)

	return products, result.Error
}

func (repo Repository) GetProductById(id int) (*Product, error) {
	var product *Product
	result := repo.DB.First(&product, id)

	return product, result.Error
}

func (repo Repository) AddProduct(product *Product) error {
	result := repo.DB.Create(&product)

	return result.Error
}

func (repo Repository) EditProduct(id int, product *Product) error {

	result := repo.DB.Select("*").Where(id).Updates(&product)

	return result.Error
}

func (repo Repository) Updates(product *Product) error{
	result := repo.DB.Updates(&product)

	return result.Error
}

func (repo Repository) Save(product *Product) error {
	result := repo.DB.Save(&product)

	return result.Error
}
