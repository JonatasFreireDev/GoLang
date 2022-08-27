package print

import (
	"fastFood/1/src/utils"
	"fmt"
	"time"
)
const MaxOrder int = 5
const MaxFoods int = 5
const MaxDrinks int = 5
const MaxServings int = 5
const MaxAdditional int = 5

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

func Foods(){
	utils.Clear()
	fmt.Println("|------ 2 - Bebidas     ------|")
	time.Sleep(2 * time.Second)
}

func Drinks(){
	utils.Clear()
	fmt.Println("|------ 2 - Bebidas     ------|")
}

func Servings(){
	utils.Clear()
	fmt.Println("|------ 2 - Bebidas     ------|")
}

func Additional(){
	utils.Clear()
	fmt.Println("|------ 2 - Bebidas     ------|")
}
