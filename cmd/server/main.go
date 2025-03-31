package main

import (
	"github.com/sondoannam/go-ecommerce-backend-api/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run(":8000")
}
