package simple

import (
	"fmt"
	"net/http"
)

// a handler to serve the simple testing example
func Handler() *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("testing/simple"))
	mux.Handle("/", fs)
	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		fmt.Fprintln(w, "500 error!!!")
	})
	return mux
}
