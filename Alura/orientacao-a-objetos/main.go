package main

import (
	"fmt"
	"orientacao-a-objetos/clientes"
	"orientacao-a-objetos/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	guilherme := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Guilherme",
			CPF:       "436.736.123-18",
			Profissao: "Dev",
		},
		NumeroAgencia: 10,
		NumeroConta:   12,
	}

	PagarBoleto(&guilherme, 60)

	// bruna := contas.ContaCorrente{
	// 	"teste",
	// 	10,
	// 	12,
	// 	1231,
	// }

	// rogerio := contas.ContaCorrente{"Salve", 10, 10, 10}

	fmt.Println(guilherme)
	// fmt.Println(bruna)
	// fmt.Println(rogerio)
}
