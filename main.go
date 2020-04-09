package main

import (
	"fmt"
	"net/http"

	"github.com/csornyei/simple-go-web-service/controllers"
)

func main() {
	controllers.RegisterControllers()
	fmt.Println("Starting server on port 3000...")
	http.ListenAndServe(":3000", nil)
}
