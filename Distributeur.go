package rpg

import (
	"fmt"
	"math/rand"
)

func distributeur(tableau []Item, joueur *Player, nbr int) []Item {
	// nbr correspond aux differents distributeur
	validation := ""
	tab2 := len(tableau)
	tab := []Item{}
	fmt.Println("Do you want to spend 10 credits to get an Item : yes/no ")
	fmt.Printf("Write your answer : ")
	fmt.Scan(&validation)
	nb := rand.Intn(tab2)
	switch validation {
	case "yes":
		if nbr == 0 {
			joueur.Inventaire[nbr][tableau[nb].Name] = tableau[nb].Attaque
			joueur.Credits -= 50
			fmt.Printf("You win %v. You can see it in your inventory", tableau[nb].Name)
		} else if nbr == 1 {
			joueur.Inventaire[nbr][tableau[nb].Name] = tableau[nb].LP
			joueur.Credits -= 50
			fmt.Printf("You win %v. You can see it in your inventory", tableau[nb].Name)
		} else if nbr == 2 {
			joueur.Inventaire[nbr][tableau[nb].Name] = tableau[nb].LP
			joueur.Credits -= 50
			fmt.Printf("You win %v. You can see it in your inventory", tableau[nb].Name)
		}
	}
	for i := 0; i < len(tab); i++ {
		if i != nbr {
			tab = append(tab, tableau[i])
		}
	}
	return tab
}
