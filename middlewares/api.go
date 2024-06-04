package middlewares

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func APIKEYMiddleware(next http.Handler) http.Handler {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("API-Key")
		expectedAPIKey := os.Getenv("API_KEY") // Change this to your actual API key

		// Check if the API key is provided
		if apiKey == "" {
			http.Error(w, "API key is missing", http.StatusUnauthorized)
			return
		}

		// Compare the API key
		if apiKey != expectedAPIKey {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}

		// If the API key is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
