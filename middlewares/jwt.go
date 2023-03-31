package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	// "github.com/jeypc/go-jwt-mux/config"
	"github.com/ditobayu/task-5-vix-btpns-Dito-Bayu-Adhitya/helpers"
)


var JWT_KEY = []byte("ashdjqy9283409bsdklkg8hda01")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				helpers.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}
		// mengambil token value
		tokenString := c.Value

		claims := &JWTClaim{}
		// parsing token jwt
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				// token invalid
				response := map[string]string{"message": "Unauthorized"}
				helpers.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				// token expired
				response := map[string]string{"message": "Unauthorized, Token expired!"}
				helpers.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "Unauthorized"}
				helpers.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			helpers.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)
	})
}