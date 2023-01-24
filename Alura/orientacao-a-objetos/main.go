package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	guilherme := ContaCorrente{
		titular:       "teste",
		numeroAgencia: 10,
		numeroConta:   12,
		saldo:         1231,
	}

	bruna := ContaCorrente{
		"teste",
		10,
		12,
		1231,
	}

	rogerio := ContaCorrente{"Salve", 10, 10, 10}

	fmt.Println(guilherme)
	fmt.Println(bruna)
	fmt.Println(rogerio)
}
