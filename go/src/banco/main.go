package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteBruno := clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.111.123.12",
		Profissao: "Desenvolvedor Go",
	}

	contaBruno := contas.ContaCorrente{
		Titular:       clienteBruno,
		NumeroAgencia: 123,
		NumeroConta:   123456,
		Saldo:         100,
	}

	fmt.Println(contaBruno)
}
