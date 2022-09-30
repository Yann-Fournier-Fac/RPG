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
	fmt.Println("Voulez-vous dépenser 10 crédits pour obtenir un objet ? yes/no ")
	fmt.Printf("Écrivez votre réponse : ")
	fmt.Scan(&validation)
	nb := rand.Intn(tab2 - 1)
	switch validation {
	case "yes":
		if nbr == 0 {
			joueur.Inventaire[nbr][tableau[nb].Name] = tableau[nb].Attaque
			joueur.Credits -= 10
			fmt.Printf("Vous avez gagné %v. Il est désormais visible dans votre inventaire", tableau[nb].Name)
			fmt.Println("\n")
		} else if nbr == 1 {
			joueur.Inventaire[nbr][tableau[nb].Name] = tableau[nb].LP
			joueur.Credits -= 10
			fmt.Printf("Vous avez gagné %v. Il est désormais visible dans votre inventaire", tableau[nb].Name)
			fmt.Println("\n")
		} else if nbr == 2 {
			joueur.Inventaire[nbr][tableau[nb].Name] = tableau[nb].LP
			joueur.Credits -= 10
			fmt.Printf("Vous avez gagné %v. Il est désormais visible dans votre inventaire", tableau[nb].Name)
			fmt.Println("\n")
		}
	case "no":
		return tableau
	}
	for i := 0; i < len(tableau); i++ {
		if i != nb {
			tab = append(tab, tableau[i])
		}
	}
	fmt.Printf("Il vous reste : %v credits\n", joueur.Credits)
	return tab
}
