package midware

import (
	"fmt"
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before handler is executed")
		w.Write([]byte("Adding response via middleware\n"))
		log.Println(r.URL.Path)
		next.ServeHTTP(w, r)
		fmt.Println("After handler is executed")
	})
}
