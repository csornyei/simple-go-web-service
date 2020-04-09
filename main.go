package main

import (
	"fmt"

	"github.com/csornyei/simple-go-web-service/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Matthew",
		LastName:  "Cerni",
	}

	fmt.Println(u)
}
