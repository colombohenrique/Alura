package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaGuilherme := ContaCorrente{
		titular:       "Guilherme",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.49,
	}

	contaBruna := ContaCorrente{
		"Bruna",
		222,
		111222,
		200,
	}

	fmt.Println(contaGuilherme)
	fmt.Println(contaBruna)
}
