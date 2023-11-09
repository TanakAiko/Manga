package main

import (
	"fmt"
	hl "manga/internals/handlers"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", hl.RenderHandler("home.html"))
	http.HandleFunc("/about", hl.RenderHandler("about.html"))
	fmt.Println("-----------------> http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
