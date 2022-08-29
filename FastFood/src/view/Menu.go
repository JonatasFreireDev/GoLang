package print

import (
	"fastFood/1/src/utils"
	"fmt"
)
const MaxMenu int = 5

func Menu(){
	utils.Clear()
	fmt.Println("")
	fmt.Println("+----------Cárdapio-----------+")
	fmt.Println("|------ 1 - Comidas     ------|")
	fmt.Println("|------ 2 - Bebidas     ------|")
	fmt.Println("|------ 3 - Porções     ------|")
	fmt.Println("|------ 4 - Adicionais  ------|")
	fmt.Println("|------ 5 - Voltar      ------|")
	fmt.Println("+-----------------------------+")
	fmt.Println("")
}

func Foods(){
	utils.Clear()
	fmt.Println("")
	fmt.Println("+--------------Comidas---------------+")
	fmt.Println("|---- 1 - X-Burguer....RS 10,00  ----|")
	fmt.Println("|---- 2 - X-Frango.....RS 10,00  ----|")
	fmt.Println("|---- 3 - X-Bacon......RS 10,00  ----|")
	fmt.Println("|---- 4 - X-Salada.....RS 10,00  ----|")
	fmt.Println("|---- 5 - X-Gordinho...RS 10,00  ----|")
	fmt.Println("+------------------------------------+")
	fmt.Println("+---------Press Enter to Exit--------+")
	fmt.Println("+------------------------------------+")
	fmt.Println("")
	fmt.Scanln()
}

func Drinks(){
	utils.Clear()
	fmt.Println("")
	fmt.Println("+---------------Bebidas----------------+")
	fmt.Println("|---- 1 - Coca-Cola......RS 10,00  ----|")
	fmt.Println("|---- 2 - Fanta Laranja..RS 10,00  ----|")
	fmt.Println("|---- 3 - Suco Laranja...RS 10,00  ----|")
	fmt.Println("|---- 4 - Agua...........RS 10,00  ----|")
	fmt.Println("|---- 5 - Agua C/Gas.....RS 10,00  ----|")
	fmt.Println("+--------------------------------------+")
	fmt.Println("+----------Press Enter to Exit---------+")
	fmt.Println("+--------------------------------------+")
	fmt.Println("")
	fmt.Scanln()
}

func Servings(){
	utils.Clear()
	fmt.Println("")
	fmt.Println("+-----------------Porçoes-------------------+")
	fmt.Println("|---- 1 - Frango..............RS 30,00  ----|")
	fmt.Println("|---- 2 - Batata..............RS 10,00  ----|")
	fmt.Println("|---- 3 - Batata com Cheddar..RS 10,00  ----|")
	fmt.Println("|---- 4 - Cebolas.............RS 10,00  ----|")
	fmt.Println("|---- 5 - Coxinha.............RS 10,00  ----|")
	fmt.Println("+-------------------------------------------+")
	fmt.Println("+-------------Press Enter to Exit-----------+")
	fmt.Println("+-------------------------------------------+")
	fmt.Println("")
	fmt.Scanln()
}

func Additional(){
	utils.Clear()
	fmt.Println("")
	fmt.Println("+--------------Comidas---------------+")
	fmt.Println("|---- 1 - Bacon........RS  5,00  ----|")
	fmt.Println("|---- 2 - Ovo..........RS  1,00  ----|")
	fmt.Println("|---- 3 - Cheddar......RS  5,00  ----|")
	fmt.Println("|---- 4 - Hamburguer...RS  6,00  ----|")
	fmt.Println("+------------------------------------+")
	fmt.Println("+---------Press Enter to Exit--------+")
	fmt.Println("+------------------------------------+")
	fmt.Println("")
	fmt.Scanln()
}
