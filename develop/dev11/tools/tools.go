package tools
import (
	"fmt"
	"log"
	"net/http"
	"time"
	"mbatimel/WB_Tech-L2/tree/main/develop/dev11/event"
)

func IsValidDate(event *event.Event) bool {
	value, err := time.Parse("2006-01-02", string(event.Date))
	if err != nil {
		log.Println(err)
		return false
	}
	event.Time = value
	return true
}

func JsonResult(respondText string) []byte {
	return []byte(fmt.Sprintf(`{"result": %s}`, respondText))
}

func JsonError(respondText string) []byte {
	return []byte(fmt.Sprintf(`{"error": "%s"}`, respondText))
}

func MakeJsonRespond(w http.ResponseWriter, code int, data []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(data)
	if err != nil {
		log.Println(err)
	}
}

func MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println(
			"method", r.Method,
			"path", r.URL.EscapedPath(),
			"duration", time.Since(start),
		)
		next(w, r)
	}
}