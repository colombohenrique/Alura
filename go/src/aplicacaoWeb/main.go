package main

import (
	"net/http"

	"github.com/Alura/go/src/aplicacaoWeb/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
