package rpg

import (
	"fmt"
	"math/rand"
)

// Il faut encore rajouter les prints pour que le joueur voit l'avencement du combat.
// Il faut tester le code.
// Faire respawn le joueur avec la moitier de ces LP si il perd.

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
		fmt.Printf("You are a piece of shit")
		return false
		//fmt.Printf(enemi.Photo)
	}
}









func distributeur(tableau []Item, joueur *Player) []Item {
	validation := ""
	tab := []Item{}
	fmt.Printf("Do you want to spend 10 credits to get an Item : yes/no ")
	fmt.Printf("Write your answer : ")
	fmt.Scan(&validation)
	nbr := rand.Intn(5)
	switch validation {
	case "yes":
		joueur.Inventaire = append(joueur.Inventaire, tableau[nbr])
		joueur.Credits -= 10
		fmt.Printf("You win %v. You can see it in your inventory", tableau[nbr].Name)
	}
	for i := 0; i < len(tab); i++ {
		if i != nbr {
			tab = append(tab, tableau[i])
		}
	}
	tableau = tab
	return tab
}










func Staires(nbr int) int {
	switch nbr {
	case 0:
		validation := ""
		fmt.Printf("Do you want to go upstraire : yes/no")
		fmt.Printf("Write your anwser : ")
		fmt.Scan(&validation)
		if validation == "yes" {
			nbr++
			return nbr
		} else if validation == "no" {
			break
		}
	case 1:
		validation := ""
		fmt.Printf("Where do you want to go : up/down/stay")
		fmt.Printf("Write your anwser : ")
		fmt.Scan(&validation)
		if validation == "up" {
			nbr++
			return nbr
		} else if validation == "down" {
			nbr--
			return nbr
		} else if validation == "stay" {
			break
		}
	case 2:
		validation := ""
		fmt.Printf("Where do you want to go : up/down/stay")
		fmt.Printf("Write your anwser : ")
		fmt.Scan(&validation)
		if validation == "up" {
			nbr++
			return nbr
		} else if validation == "down" {
			nbr--
			return nbr
		} else if validation == "stay" {
			break
		}
	case 3:
		validation := ""
		fmt.Printf("Do you want to go downstraire : yes/no")
		fmt.Printf("Write your anwser : ")
		fmt.Scan(&validation)
		if validation == "yes" {
			nbr--
			return nbr
		} else if validation == "no" {
			break
		}
	}
	return nbr
}
