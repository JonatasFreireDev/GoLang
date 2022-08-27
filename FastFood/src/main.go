package main

import (
	"fastFood/1/src/print"
	"fmt"
	"time"
)

func main(){
	
	for{
		switch choice := print.Menu(); choice {
			case 1: fmt.Println("1")
			case 2: fmt.Println("2")
			case 3: fmt.Println("3")
			case 4: fmt.Println("4")
			case 5: fmt.Println("5")
			default:fmt.Println("defuuu")
		}

		time.Sleep(4 * time.Second)
		
	}

}