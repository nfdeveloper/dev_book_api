package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nfdeveloper/dev_book_api/src/config"
	"github.com/nfdeveloper/dev_book_api/src/router"
)

func main() {
	config.Carregar()
	fmt.Println(config.Porta)

	fmt.Println("Rodando API.")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
