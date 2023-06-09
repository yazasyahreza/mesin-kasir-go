package transactions

import (
	"encoding/json"
	"mini-project/modules/products"
	"net/http"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transactions, err := handler.Usecase.GetAll()
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	_, err = json.Marshal(transactions)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	result := []map[string]interface{}{}

	for _, data := range transactions {
		t := map[string]interface{}{
			"Id":        data.ID,
			"Timestamp": data.Timestamp,
			"Total":     data.Total,
		}
		result = append(result, t)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"Message": "data found",
		"Data":    result,
	})
}

func (handler Handler) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transaction, err := handler.Usecase.GetById(r.Context())
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	_, err = json.Marshal(transaction)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"Message": "data found",
		"Data":    transaction,
	})
}

func (handler Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request products.CreateTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "Failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	transaction, err := handler.Usecase.CreateTransaction(r.Context(), &request)
	if err != nil {
		if err == ErrProductIdNotFound {
			json.NewEncoder(w).Encode(map[string]string{
				"Message": "Product Id Not Found",
			})
			return
		} else if err == ErrPoductHasBeenRemoved {
			json.NewEncoder(w).Encode(map[string]string{
				"Message": "This Product Has Been Removed",
			})
			return
		} else if err == ErrStockNotEnough {
			json.NewEncoder(w).Encode(map[string]string{
				"Message": "Stock Not Enough",
			})
			return
		}
	}

	json.NewEncoder(w).Encode(transaction)
}
