package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"mini-project/modules/admin"
	"net/http"
	"strconv"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func MiddlewareJWTAuthorization(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			messageErr, _ := json.Marshal(map[string]string{"message": "invalid token"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(messageErr)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.ParseWithClaims(tokenString, &admin.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return admin.JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(*admin.MyClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, "idPrms", id)
		ctx = context.WithValue(ctx, "id_admin", claims.ID)
		ctx = context.WithValue(ctx, "name", claims.Name)

		next(w, r.WithContext(ctx))
	})

}
