package rpg

import (
	"fmt"
	"math/rand"
)

// Trouver une photo pour quand on gagne

func Fight(joueur Player, enemy PNJ) bool {
	// Le joueur choisi ces équipements
	if len(joueur.Inventaire[0]) == 0 {
		fmt.Println("You don't have weapon")
	} else {
		arme := ""
		fmt.Println("Choose your weapon, Write the name of the weapon that you want : ")
		fmt.Print(joueur.Inventaire[0])
		fmt.Scan(&arme)
		joueur.Attaque += joueur.Inventaire[0][arme]
	}
	if len(joueur.Inventaire[1]) == 0 {
		fmt.Println("You don't have armor")
	} else {
		armure := ""
		fmt.Println("Choose your armor, Write the name of the weapon that you want : ")
		fmt.Print(joueur.Inventaire[1])
		fmt.Scan(&armure)
		joueur.LP += joueur.Inventaire[1][armure]
	}

	// Le combat commence
	fmt.Printf("You have %v LP", joueur.LP)
	fmt.Printf("\n")
	fmt.Printf("%v have %v LP", enemy.Name, enemy.LP)
	fmt.Printf("\n")
	for (joueur.LP > 0) && (enemy.LP > 0) {
		fmt.Printf("Please write a number between 0 and 1 : ")
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
			fmt.Printf("You have inflicted %v damages points", joueur.Attaque)
			fmt.Printf("\n")
			fmt.Printf("You have %v LP", joueur.LP)
			fmt.Printf("\n")
			fmt.Printf("%v have %v LP left", enemy.Name, enemy.LP)
			fmt.Printf("\n")
			fmt.Printf("\n")
		} else if pileface == nbr2 {
			joueur.LP -= enemy.Attaque
			fmt.Printf("\n")
			fmt.Printf("You received %v damages points", enemy.Attaque)
			fmt.Printf("\n")
			fmt.Printf("You have %v LP left", joueur.LP)
			fmt.Printf("\n")
			fmt.Printf("%v have %v LP", enemy.Name, enemy.LP)
			fmt.Printf("\n")
			fmt.Printf("\n")
		} else {
			fmt.Printf("Wrong input. Please try again")
		}
	}
	if enemy.LP <= 0 {
		fmt.Printf("Well play %v, you won against %v", joueur.Name, enemy.Name)
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
		fmt.Println("You got now the half of your life point")
		fmt.Printf("\n")
		return false
		//fmt.Printf(enemi.Photo)
	}
}
