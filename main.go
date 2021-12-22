package main

import (
	"fmt"
	"github.com/jerevick83/handler"
	"github.com/jerevick83/utils"
	"net/http"
)

func init() {
	handler.RenderTemplate()
}

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	fmt.Println("Server running on port: ", utils.PortName)
	http.ListenAndServe(utils.PortName, nil)
}
