package print

import (
	"fmt"
	"time"
)

func Menu() int{
	var choice int

	fmt.Print("\033[H\033[2J")
	fmt.Println("")
	fmt.Println("+---------------------Menu---------------------+")
	fmt.Println("|------ 1 - Visualizar o Cárdapio        ------|")
	fmt.Println("|------ 2 - Fazer novo pedido            ------|")
	fmt.Println("|------ 3 - Consultar Pedidos em preparo ------|")
	fmt.Println("|------ 4 - Consultar Pedidos Prontos    ------|")
	fmt.Println("|------ 5 - Cancelar um Pedido           ------|")
	fmt.Println("|------ 6 - Sair do Programa             ------|")
	fmt.Println("+----------------------------------------------+")
	fmt.Println("")

	_, e := fmt.Scan(&choice)

	if e != nil {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Digito Incorreto, Tente novamente")
		time.Sleep(5 * time.Second)
		Menu()
	}

	return choice
}