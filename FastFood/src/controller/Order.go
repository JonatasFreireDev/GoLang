package controller

import (
	"fastFood/1/src/utils"
	print "fastFood/1/src/view"
	"fmt"
)

func OrderController() {
	for {
		print.Order()
		choice := utils.GetChoice(print.Order, 5);

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