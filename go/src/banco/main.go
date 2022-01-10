package main

import (
	"fmt"

	"banco/contas"
)

func main() {
	contaSilvia := contas.ContaCorrente{
		Titular: "Silvia",
		Saldo:   300,
	}

	contaGustavo := contas.ContaCorrente{
		Titular: "Gustavo",
		Saldo:   100,
	}

	status := contaSilvia.Transferir(200, &contaGustavo)

	fmt.Println(status)
	fmt.Println(contaSilvia)
	fmt.Println(contaGustavo)
}
