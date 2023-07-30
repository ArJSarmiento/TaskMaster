package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"crud_ql/graph/model"
	"crud_ql/repository"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(db *repository.DB, cognitoClient *CognitoClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				next.ServeHTTP(w, r)
				return
			}

			// The Authorization header is in form "Bearer <token>"
			splitToken := strings.Split(token, "Bearer ")
			if len(splitToken) != 2 {
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}

			pubKeyURL := "https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json"
			formattedURL := fmt.Sprintf(pubKeyURL, os.Getenv("AWS_DEFAULT_REGION"), cognitoClient.UserPoolId)

			keySet, err := jwk.Fetch(r.Context(), formattedURL)
			if err != nil {
				log.Printf("Could not fetch public keys from Cognito: %s", err)
			}

			auth_token, err := jwt.Parse(
				[]byte(splitToken[1]),
				jwt.WithKeySet(keySet),
				jwt.WithValidate(true),
			)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			sub := auth_token.Subject()
			log.Print(sub)
			if sub == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			user, err := db.GetUserBySub(sub)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
