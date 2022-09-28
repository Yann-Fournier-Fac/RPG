package main

import (
	"fmt"
	"rpg"
)

// chercher la fonction clear() sur internet puis la faire
// Mettre des clear() un peu partout

func main() {
	for {
		str := ""
		fmt.Println("Please, write what you want to do")
		fmt.Println("newgame")
		fmt.Println("about")
		fmt.Println("save")
		fmt.Println("exit")
		fmt.Scan(&str)
		switch str {
		case "newgame":
			level := 0
			fmt.Printf("\n")
			name := ""
			fmt.Printf("Enter your name : ")
			fmt.Scan(&name)

			// Creation des enemies
			var player rpg.Player
			player.Name = name
			player.LP = 50
			player.Attaque = 50
			player.Credits = 50
			player.Inventaire = [3]map[string]int{}

			var Nesrine rpg.PNJ
			Nesrine.Name = "Nesrine"
			Nesrine.Enemies = true
			Nesrine.Attaque = 10
			Nesrine.LP = 100
			Nesrine.Credits = 110

			var Guillaume rpg.PNJ
			Guillaume.Name = "Guillaume"
			Guillaume.Enemies = true
			Guillaume.Attaque = 50
			Guillaume.LP = 150
			Guillaume.Credits = 200

			var Paul rpg.PNJ
			Paul.Name = "Paul"
			Paul.Enemies = true
			Paul.Attaque = 200
			Paul.LP = 500
			Paul.Credits = 700

			for {
				if level == 10 {
					return
				} else {
					if level == 0 {
						f0, nb0 := rpg.Floor0()
						level = rpg.Moove(f0, nb0, player, Nesrine, Guillaume, Paul)
					} else if level == 1 {
						f1, nb1 := rpg.Floor1()
						level = rpg.Moove(f1, nb1, player, Nesrine, Guillaume, Paul)
					} else if level == 2 {
						f2, nb2 := rpg.Floor2()
						level = rpg.Moove(f2, nb2, player, Nesrine, Guillaume, Paul)
					} else if level == 3 {
						f3, nb3 := rpg.Floor3()
						level = rpg.Moove(f3, nb3, player, Nesrine, Guillaume, Paul)
					}
				}
			}
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
