package controller

import (
	"fastFood/1/src/utils"
	"fastFood/1/src/view"
	"fmt"
)

func OrderController() {
	for {
		view.Order()
		choice := utils.GetChoice(view.Order, 5);

		switch choice {
			case 1: view.Foods()
			case 2: view.Drinks()
			case 3: view.Servings()
			case 4: view.Additional()
			case 5: return
			default: fmt.Println("defuuu")
		}
	}
}