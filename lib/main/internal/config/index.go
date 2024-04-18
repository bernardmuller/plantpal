package config

import (
	"fmt"
	"net/http"
)

func CreateRouter() *http.ServeMux {
	mux := http.ServeMux{}

	//mux.Handle("/static/images/", http.StripPrefix("/static/images/", http.FileServer(http.Dir("images"))))
	//mux.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("css"))))
	fmt.Println(http.Dir("views"))
	//mux.Handle()

	return &mux
}
