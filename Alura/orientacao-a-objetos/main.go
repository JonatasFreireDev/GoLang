package main

import (
	"fmt"
	"orientacao-a-objetos/contas"
)

func main() {
	guilherme := contas.ContaCorrente{
		Titular:       "teste",
		NumeroAgencia: 10,
		NumeroConta:   12,
		Saldo:         1231,
	}

	bruna := contas.ContaCorrente{
		"teste",
		10,
		12,
		1231,
	}

	rogerio := contas.ContaCorrente{"Salve", 10, 10, 10}

	fmt.Println(guilherme)
	fmt.Println(bruna)
	fmt.Println(rogerio)
}
