package controller

import (
	"fastFood/1/src/utils"
	"fmt"
	"time"
)

func OrderController() {
	for {
		var choice int 

		utils.Clear()
		fmt.Println("O que deseja adicionar ?")
		fmt.Scan(&choice)

		fmt.Println(choice)

		time.Sleep(10 * time.Second)
	}
}