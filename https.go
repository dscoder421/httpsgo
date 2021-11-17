package httpsgo

import (
	"fmt"
	"net/http"
)

func RegisterHandler(url string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(url, handler)
}

func StartHttpServer(port string) {
	http.ListenAndServe(port, nil)
}

func StartHttpsServer(port string, certFile string, keyFile string) {
	http.ListenAndServeTLS(port, certFile, keyFile, nil)
}

func SetHeader(key string, value string, w http.ResponseWriter) {
	w.Header().Set(key, value)
}

func SendStatus(code int, rw http.ResponseWriter) {
	fmt.Fprint(rw, http.StatusText(code))
}

func Send(text string, rw http.ResponseWriter) {
	fmt.Fprint(rw, text)
}
