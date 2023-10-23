package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Log the incoming request
		lg := log.New(os.Stdout, "", 0)
		logMsg := fmt.Sprintf("[%s] \033[34m[INFO]\033[0m  method=%s uri=%s ", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.RequestURI)

		// Output the log message using the standard logger
		lg.Output(2, logMsg)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}