package controller

import (
	"fastFood/1/src/utils"
	print "fastFood/1/src/view"
	"fmt"
)

func MenuController() {
	for {
		print.Order()
		choice := utils.GetChoice(print.Order, print.MaxMenu);

		switch choice {
		case 1:
			print.Foods()
			fmt.Scanln()
		case 2:
			print.Drinks()
			fmt.Scanln()
		case 3:
			print.Servings()
			fmt.Scanln()
		case 4:
			print.Additional()
			fmt.Scanln()
		case 5:
			return
		default:
			fmt.Println("defuuu")
		}
	}
}