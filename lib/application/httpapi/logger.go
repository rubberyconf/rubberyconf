package httpapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rubberyconf/rubberyconf/lib/core/logs"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		message := fmt.Sprintf("%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start))

		logs.GetLogs().WriteMessage(logs.INFO, message, r.Header)

	})
}
