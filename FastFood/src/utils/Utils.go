package utils

import (
	"fmt"
	"time"
)

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func GetChoice(printFunc func(), maxOption int) int{
	var choice int

	_, e := fmt.Scan(&choice)

	if e != nil || choice > maxOption {
		Clear()
		fmt.Println("Digito Incorreto, Tente novamente")
		time.Sleep(4 * time.Second)
		printFunc()
	}

	return choice
}