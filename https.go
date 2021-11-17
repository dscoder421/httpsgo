package httpsgo

import (
	"net/http"
)

func RegisterHandler(url string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(url, handler)
}

func StartServer(port string) {
	http.ListenAndServe(port, nil)
}

func SetHeader(key string, value string, w http.ResponseWriter) {
	w.Header().Set(key, value)
}

func main() {
	RegisterHandler("/hello", func(rw http.ResponseWriter, r *http.Request) {
		// rw.Header().Set("", "")
		SetHeader("", "", rw)
	})
	StartServer(":3000")
}
