package main

import (
	"fmt"
	"rpg"
)

// chercher la fonction clear() sur internet puis la faire
// Mettre des clear() un peu partout

func main() {
	// Histoire
	fmt.Println(" _____ _____ _____ _____ _____ _____ _____ _____ _____ _____ ")
	fmt.Println("/_____/_____/_____/_____/_____/_____/_____/_____/_____/_____/")
	fmt.Println("\n")
	fmt.Println("	__  __                    __    _ ____   ")
	fmt.Println("	\\ \\/ /___  ____ _   __   / /   (_) __/__ ")
	fmt.Println("	 \\  / __ \\/ __ \\ | / /  / /   / / /_/ _ \\")
	fmt.Println("	 / / / / / /_/ / |/ /  / /___/ / __/  __/")
	fmt.Println("	/_/_/ /_/\\____/|___/  /_____/_/_/  \\___/")
	fmt.Println("\n")
	fmt.Println(" _____ _____ _____ _____ _____ _____ _____ _____ _____ _____ ")
	fmt.Println("/_____/_____/_____/_____/_____/_____/_____/_____/_____/_____/")
	fmt.Println("\n")

	fmt.Println("Tout commence, le 01 Septembre 2022, dans le campus d’YNOV situé à Nanterre. ")
	fmt.Println("À peine arrivé dans le hall pour votre premier jour, vous ne pouvez vous empêcher d'écouter")
	fmt.Println("les discussions de couloirs, il semblerait qu'un pickpocket se soit infiltré sur le campus.")
	fmt.Println("Cela n'était pas réellement votre problème jusqu'à ce qu'à l'heure du déjeuner, vous réalisiez")
	fmt.Println("que votre carte graphique achetée ce matin même est manquante..")
	fmt.Println("Selon des sources sûrs, un certain Paul se serait lancé dans le minage de")
	fmt.Println("Crypto monnaies et heureux hasard.. Il a toujours voulu devenir pickpocket.")
	fmt.Println("Après avoir entamé les négociations et sur demande du staff d'YNOV, vous seul")
	fmt.Println("êtes en mesure de repousser l'invasion des mentors. ")
	fmt.Println("Il est grand temps de vous équiper et de partir récupérer les biens dérobés.")
	fmt.Printf("\n")

	for { // Boucle infinie (le jeu tourneras à l'infini)
		str := ""
		fmt.Println("Écrivez ce que vous désirez faire :") // Affichage du menu principales
		fmt.Println("1) newgame")
		fmt.Println("2) about")
		fmt.Println("3) save")
		fmt.Println("4) exit")
		fmt.Scan(&str) // demande au joueur de choisir ce qu'il veut faire

		armes := make(map[string]int) // creation des map qui composeront l'inventeire du joueur
		armures := make(map[string]int)
		soins := make(map[string]int)

		switch str {
		case "newgame":
			level := 0
			fmt.Printf("\n")

			// Le joueur entre son nom
			name := ""
			fmt.Printf("Entrez votre pseudo : ")
			fmt.Scan(&name)

			// Creation du player
			var player rpg.Player
			player.Name = name
			player.LP = 50
			player.Attaque = 50
			player.Credits = 50
			player.Inventaire = [3]map[string]int{armes, armures, soins}

			// Creation des enemis
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
			Paul.Attaque = 70
			Paul.LP = 300
			Paul.Credits = 700

			// boucle infinie qui fait tourner le jeu
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
			fmt.Println("Ce jeu a été développé par 4 étudiants de Paris Ynov Campus. (Yann FOURNIER, Thomas RUCKTOOA et Ouardia AOUDAR)")
			return
		case "save":
			fmt.Println("start a save")
		case "exit":
			return
		default:
			fmt.Println("Merci d'écrire les choix affichés !")
			fmt.Printf("\n")
		}
	}
}
