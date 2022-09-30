package rpg

import (
	"fmt"
	"math/rand"
)

// Trouver une photo pour quand on gagne

func Fight(joueur Player, enemy PNJ) bool {
	// Le joueur choisi ces équipements
	if len(joueur.Inventaire[0]) == 0 {
		fmt.Println("Vous n'avez pas d'arme")
	} else {
		arme := ""
		fmt.Println("Choisissez votre arme, écrivez le nom de l'arme que vous souhaitez : ")
		fmt.Print(joueur.Inventaire[0])
		fmt.Scan(&arme)
		joueur.Attaque += joueur.Inventaire[0][arme]
	}
	if len(joueur.Inventaire[1]) == 0 {
		fmt.Println("Vous n'avez pas d'armure")
	} else {
		armure := ""
		fmt.Println("Choisissez votre armure, écrivez le nom de l'armure que vous souhaitez : ")
		fmt.Print(joueur.Inventaire[1])
		fmt.Scan(&armure)
		joueur.LP += joueur.Inventaire[1][armure]
	}

	// Le combat commence
	fmt.Printf("Vous avez %v PV", joueur.LP)
	fmt.Printf("\n")
	fmt.Printf("%v à %v PV", enemy.Name, enemy.LP)
	fmt.Printf("\n")
	for (joueur.LP > 0) && (enemy.LP > 0) {
		fmt.Printf("Écrivez un  nombre entre 0 et 1 : ")
		var pileface int
		fmt.Scan(&pileface)
		nbr := rand.Intn(2)
		nbr2 := 0
		if nbr == 0 {
			nbr2 = 1
		} else if nbr == 1 {
			nbr2 = 0
		}
		if pileface == nbr {
			enemy.LP -= joueur.Attaque
			fmt.Printf("\n")
			fmt.Printf("Vous inflgez %v points de dégats", joueur.Attaque)
			fmt.Printf("\n")
			fmt.Printf("Vous avez %v PV", joueur.LP)
			fmt.Printf("\n")
			fmt.Printf("Il reste à %v %v PV", enemy.Name, enemy.LP)
			fmt.Printf("\n")
			fmt.Printf("\n")
		} else if pileface == nbr2 {
			joueur.LP -= enemy.Attaque
			fmt.Printf("\n")
			fmt.Printf("Vous avez reçu %v points de dégâts", enemy.Attaque)
			fmt.Printf("\n")
			fmt.Printf("Il vous reste %v PV", joueur.LP)
			fmt.Printf("\n")
			fmt.Printf("Il reste à %v %v PV", enemy.Name, enemy.LP)
			fmt.Printf("\n")
			fmt.Printf("\n")
		} else {
			fmt.Printf("Merci d'écrire les choix affichés !")
		}
	}
	if enemy.LP <= 0 {
		fmt.Printf("Bien joué %v, vous avez vaincu %v", joueur.Name, enemy.Name)
		fmt.Printf("\n")
		return true
		//fmt.Printf(enemi.Ficha)
	} else {
		fmt.Println("You are a piece of shit")
		fmt.Printf("\n")
		fmt.Println("⣿⣿⣿⣿⣿⣿⣿⣿⣻⡯⠝⠛⠊⠉⠉⠉⠉⠓⠛⠫⢽⣻⢿⣿⣿⣿⣿⣿⣿⣿")
		fmt.Println("⣿⣿⣿⣿⣿⣻⠗⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠻⣟⣿⣿⣿⣿⣿")
		fmt.Println("⣿⣿⣿⣿⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⡟⣿⣿⣿")
		fmt.Println("⣿⣿⡽⠁⠀⠀⠀⠀⠀⠀⠀⣀⣀⣎⠀⠀⠈⠁⠒⠠⢀⠀⠀⠀⠀⠀⠘⢯⢿⣿")
		fmt.Println("⡿⡸⠁⠀⠀⠀⠰⠶⣶⣤⣌⡀⠀⠀⠈⠁⠒⠤⡀⠀⠀⠑⢔⠒⠢⢄⠀⠈⢎⢿")
		fmt.Println("⣷⠃⠀⠀⠀⠀⠀⠀⠀⠉⠉⠀⠉⠁⠐⠠⢀⠀⠈⠐⠀⠀⠀⠑⠀⠀⠢⠀⠘⣿")
		fmt.Println("⡾⠀⠀⠀⠄⠀⠀⠀⠀⠠⠊⡲⠤⠄⡀⠀⠀⠑⠀⠀⠀⠀⠀⠀⠀⠀⠀⢃⠀⢿")
		fmt.Println("⡇⠀⢀⠊⠀⠀⠀⠀⣰⣁⠀⠦⣀⠀⠈⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠀⢸")
		fmt.Println("⢧⡰⠃⠀⠀⠀⢀⣌⠀⠉⠙⠳⠦⣭⣔⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⡾")
		fmt.Println("⣿⠁⠀⠀⠀⢠⠾⣿⣿⣷⣶⣤⣀⡀⠈⠙⠢⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢱⢿")
		fmt.Println("⣇⣇⠀⠀⡰⠃⠀⠹⣿⣿⣿⣿⣿⣿⣷⣶⣤⣈⡑⢠⣀⠀⠀⠀⠀⠀⠀⠀⢸⢸")
		fmt.Println("⣿⣿⣧⡼⠁⠀⠀⠀⠈⠛⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⡆⠀⠀⠀⠀⠀⠀⢸⢺")
		fmt.Println("⣿⣿⣿⣽⣦⡀⠀⠀⠀⠀⠀⠀⠉⠙⠛⠛⠛⠋⠉⢸⡇⠱⠀⠀⠀⠀⠀⠀⣈⣼")
		fmt.Println("⣿⣿⣿⣿⣿⣽⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡇⠀⣀⠀⠀⠀⠀⡠⣻⣿")
		fmt.Println("⣿⣿⣿⣿⣿⣿⣿⣷⣭⣒⣦⣤⣄⣀⣀⣀⣀⣠⢤⣼⣳⣯⣿⣷⣿⣤⣯⣾⣿⣿")
		fmt.Printf("\n")
		fmt.Println("Il vous reste la moitié de vos PV")
		fmt.Printf("\n")
		return false
		//fmt.Printf(enemi.Photo)
	}
}
