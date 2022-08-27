package main

import (
	"fastFood/1/src/controller"
	"fastFood/1/src/print"
	"fastFood/1/src/utils"
	"fmt"
)

func main(){
	for{

		print.Menu()
		choice := utils.GetChoice(print.Menu, print.MaxMenu)

		switch choice {
			case 1: controller.OrderController()
			case 2: fmt.Println("2")
			case 3: fmt.Println("3")
			case 4: fmt.Println("4")
			case 5: fmt.Println("5")
			case 6: return
			default:fmt.Println("defuuu")
		}	
	}

}