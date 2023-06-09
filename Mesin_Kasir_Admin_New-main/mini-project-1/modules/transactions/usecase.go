package transactions

import (
	"context"
	"fmt"
	"mini-project/modules/products"
	"time"
)

type Usecase struct {
	TransacationRepo Repository
	ProductRepo      products.Repository
}

func (uc Usecase) GetAll() ([]Transaction, error) {
	transactions, err := uc.TransacationRepo.GetAll()
	return transactions, err
}

func (uc Usecase) GetById(ctx context.Context) (*Transaction, error) {
	idPrms := ctx.Value("idPrms")
	id := idPrms.(int)
	transaction, err := uc.TransacationRepo.GetById(id)
	return transaction, err
}

func (uc Usecase) CreateTransaction(ctx context.Context, req *products.CreateTransactionRequest) (*Transaction, error) {

	id_admin := ctx.Value("id_admin")

	items := []products.TransactionItem{}
	totalPrice := 0

	for _, i := range req.Items {
		product, err := uc.ProductRepo.GetProductById(int(i.ProductID))
		if err != nil {
			return nil, ErrProductIdNotFound
		}

		if product.DeletedAt != nil {
			return nil, ErrPoductHasBeenRemoved
		}

		var stock int

		if i.Quantity > product.Stock {
			return nil, ErrStockNotEnough
		} else if i.Quantity <= product.Stock {
			stock = product.Stock - i.Quantity
		}

		product.Stock = stock

		price := int(i.Quantity) * product.Price

		ResponseAddTransaction := &products.ResponseAddTransaction{
			ID:   product.ID,
			Name: product.Name,
		}

		fmt.Println(ResponseAddTransaction)

		item := &products.TransactionItem{
			ProductID: uint(i.ProductID),
			Quantity:  i.Quantity,
			Price:     price,
		}

		items = append(items, *item)

		totalPrice += price

		err = uc.ProductRepo.EditProduct(int(i.ProductID), product)
		if err != nil {
			return nil, err
		}
	}

	transaction := &Transaction{
		AdminID:   id_admin.(float64),
		Timestamp: time.Now(),
		Total:     totalPrice,
		Items:     items,
	}

	err := uc.TransacationRepo.Create(transaction)
	if err != nil {
		return nil, err
	}

	newTransaction, err := uc.TransacationRepo.GetById(transaction.ID)
	if err != nil {
		return nil, err
	}

	return newTransaction, nil
}
