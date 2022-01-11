package main

import (
	"fmt"
	"github/Alura/go/src/banco/clientes"
	"github/Alura/go/src/banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

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
	}

	contaBruno.Depositar(100)
	PagarBoleto(&contaBruno, 55)

	fmt.Println(contaBruno.ObterSaldo())

	contaDenis := contas.ContaPoupanca{}
	contaDenis.Depositar(100)
	PagarBoleto(&contaDenis, 60)

	fmt.Println(contaDenis.ObterSaldo())
}
