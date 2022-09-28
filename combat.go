package rpg

import (
	"fmt"
	"math/rand"
)

// demander au joueur de choisir quelle arme ou quelle armur équiper

func Fight(joueur Player, enemy PNJ) bool {
	fmt.Printf("You have %v LP", joueur.LP)
	fmt.Printf("\n")
	fmt.Printf("%v have %v LP", enemy.Name, enemy.LP)
	fmt.Printf("\n")
	for (joueur.LP != 0) && (enemy.LP != 0) {
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
		fmt.Println("nbr2 ", nbr2)
		fmt.Println("nbr ", nbr)
		fmt.Println("pileface ", pileface)
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
	if enemy.LP == 0 {
		fmt.Printf("Well play %v, you won against %v", joueur.Name, enemy.Name)
		return true
		//fmt.Printf(enemi.Ficha)
	} else {
		fmt.Println("You are a piece of shit")
		fmt.Println("You got the half of your life point")
		return false
		//fmt.Printf(enemi.Photo)
	}
}
