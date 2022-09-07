package main

import (
	"fastFood/1/src/controller"
	"fastFood/1/src/utils"
	"fastFood/1/src/view"
	"fmt"
)

func main(){
	for{
		view.MainMenu()
		choice := utils.GetChoice(view.MainMenu, view.MaxMainMenu)

		switch choice {
			case 1: controller.MenuController()
			case 2: controller.OrderController()
			case 3: fmt.Println("3")
			case 4: fmt.Println("4")
			case 5: fmt.Println("5")
			case 6: return
			default:fmt.Println("defuuu")
		}	
	}

}