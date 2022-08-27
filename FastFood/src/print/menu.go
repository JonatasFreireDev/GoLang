package print

import (
	u "fastFood/1/src/utils"
	"fmt"
)

const MaxMenu int = 6

func Menu(){
	u.Clear()
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
}

