package products

import (
	"context"
	"time"
)

type Usecase struct {
	Repo Repository
}

func (uc Usecase) UcGetAllProducts() ([]Product, error) {
	products, err := uc.Repo.GetAllProducts()
	return products, err
}

func (uc Usecase) UcGetProductById(ctx context.Context) (*Product, error) {
	idPrms := ctx.Value("idPrms")
	id := idPrms.(int)
	product, err := uc.Repo.GetProductById(id)
	return product, err
}

func (uc Usecase) UcAddProduct(product *Product) error {
	err := uc.Repo.AddProduct(product)
	return err
}

func (uc Usecase) UcEditProduct(ctx context.Context, product *Product) error {
	idPrms := ctx.Value("idPrms")
	id := idPrms.(int)
	products, err := uc.Repo.GetProductById(id)
	if err != nil {
		return ErrProductIdNotFound
	}
	
	if products.DeletedAt != nil{
		return ErrPoductHasBeenRemoved
	}

	product.ID = id

	if err := uc.Repo.Updates(product); err != nil{
		return ErrChangedStatus
	}

	return err
}

func (uc Usecase) SoftDelete(ctx context.Context, status string) (*Product, error) {
	idPrms := ctx.Value("idPrms")
	id := idPrms.(int)
	product, err := uc.Repo.GetProductById(id)
	if err != nil {
		return nil, err
	}

	if status == "active" {
		if product.DeletedAt == nil {
			return nil, ErrProductNotDeleted
		} else if product.DeletedAt != nil {
			product.DeletedAt = nil
		}
	} else if status == "inactive" {
		if product.DeletedAt == nil {
			deleteAt := time.Now()
			product.DeletedAt = &deleteAt
		} else if product.DeletedAt != nil {
			return nil, ErrProductAlreadyDeleted
		}
	} else {
		return nil, ErrInvalidStatus
	}

	if err := uc.Repo.Save(product); err != nil {
		return nil, ErrChangedStatus
	}
	
	return product, nil
}
