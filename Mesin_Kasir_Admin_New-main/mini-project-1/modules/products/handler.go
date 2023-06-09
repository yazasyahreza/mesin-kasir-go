package products

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := handler.Usecase.UcGetAllProducts()
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "Data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	_, err = json.Marshal(products)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "Data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response := &ProductsResponse{
		Message: "Data found",
		Data:    products,
	}

	json.NewEncoder(w).Encode(response)
}

func (handler Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	product, err := handler.Usecase.UcGetProductById(r.Context())
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "Data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response := &ProductResponse{
		Message: "Data found",
		Data:    product,
	}

	json.NewEncoder(w).Encode(response)
}

func (handler Handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "Failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	err = handler.Usecase.UcAddProduct(&product)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "Data was not successfully added"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response := &ResponseAddAndEditData{
		Message: "Data added successfully",
		Data:    product,
	}

	json.NewEncoder(w).Encode(response)
}

func (handler Handler) EditProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "Failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	if err := handler.Usecase.UcEditProduct(r.Context(), &product); err != nil {
		if err != nil {
			if err == ErrPoductHasBeenRemoved {
				json.NewEncoder(w).Encode(map[string]string{
					"Message": "This product has been removed",
				})
				return
			} else if err == ErrProductIdNotFound {
				json.NewEncoder(w).Encode(map[string]string{
					"Message": "Product id no found",
				})
				return
			}
		}
	}

	response := &ResponseAddAndEditData{
		Message: "Data changed successfully",
		Data:    product,
	}

	json.NewEncoder(w).Encode(response)
}

func (handler Handler) SoftDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestBody RequesBodyStatus
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "Failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	product, err := handler.Usecase.SoftDelete(r.Context(), requestBody.Status)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"Message": err.Error(),
		})
		return
	}

	response := &ProductResponse{
		Message: "Success",
		Data:    product,
	}

	json.NewEncoder(w).Encode(response)
}
