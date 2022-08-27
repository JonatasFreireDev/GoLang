package controller

import (
	"fastFood/1/src/print"
	"fastFood/1/src/utils"
	"fmt"
)

func OrderController() {
	for {
		print.Order()
		choice := utils.GetChoice(print.Order, print.MaxOrder);

		switch choice {
		case 1:
			print.Foods()
		case 2:
			print.Drinks()
		case 3:
			print.Servings()
		case 4:
			print.Additional()
		case 5:
			return
		default:
			fmt.Println("defuuu")
		}
	}
}