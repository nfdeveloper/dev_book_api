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
	r := router.Gerar()

	fmt.Println("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
