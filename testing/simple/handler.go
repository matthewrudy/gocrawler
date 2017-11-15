package simple

import (
	"fmt"
	"net/http"
)

// a handler to serve the simple testing example
func Handler(templatePath string) *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(templatePath))

	// always error
	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		fmt.Fprintln(w, "500 error!!!")
	})

	firstTime := true

	// fail the first time
	mux.HandleFunc("/eventually", func(w http.ResponseWriter, r *http.Request) {
		if firstTime {
			firstTime = false

			w.WriteHeader(503)
			fmt.Fprintln(w, "503 error!!! Try again")
			return
		}
		http.ServeFile(w, r, templatePath+"/eventually.html")
	})

	// send other files directly
	mux.Handle("/", fs)

	return mux
}
