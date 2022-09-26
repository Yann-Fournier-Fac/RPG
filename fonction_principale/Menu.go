package main

import (
	"fmt"
	"rpg"
)

func main() {
	for true {
		str := ""
		fmt.Println("Please, write what you want to do")
		fmt.Println("newgame")
		fmt.Println("about")
		fmt.Println("save")
		fmt.Println("exit")
		fmt.Scan(&str)
		switch str {
		case "newgame":
			fmt.Printf("\n")
			var player rpg.Player
			name := ""
			fmt.Printf("Enter your name :")
			fmt.Scan(&name)
			player.Name = name
			fmt.Printf("\n")	

		case "about":
			fmt.Println("readme")
		case "save":
			fmt.Println("start a save")
		case "exit":
			return
		default:
			fmt.Println("wrong input")
		}
	}
}
