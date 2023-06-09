package admin

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) Login(w http.ResponseWriter, r *http.Request) {
	var userInput User
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}
	defer r.Body.Close()

	signedToken, err := handler.Usecase.Login(userInput.Username, userInput.Password)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"Message": "Success",
		"Data":    signedToken,
	})
}
