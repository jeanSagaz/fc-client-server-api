package main

import (
	"log"

	"github.com/jeanSagaz/client/internal/application/services"
)

func main() {
	response, err := services.GetQuotationTax()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	err = services.CreateFile(response)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
