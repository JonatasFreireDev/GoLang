package view

import (
	"fastFood/1/src/utils"
	"fmt"
)

func Order(){
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
