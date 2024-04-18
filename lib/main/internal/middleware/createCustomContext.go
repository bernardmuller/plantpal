package middleware

import (
	"context"
	"domain-app/internal/utils"
	"domain-app/internal/views"
	"fmt"
	"net/http"
)

func CreateCustomContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create custom context
		renderer := views.NewTemplate()

		customContext := utils.CustomContext{
			Context:  context.Background(),
			Data:     nil,
			Renderer: renderer,
		}

		// Pass the custom context to the request context
		ctx := context.WithValue(r.Context(), "customContext", customContext)

		fmt.Println("Custom context created")
		// Call the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
