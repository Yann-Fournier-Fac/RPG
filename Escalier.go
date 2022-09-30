package rpg

import "fmt"

func Staires(level int) int {
	switch level {
	case 0:
		validation := ""
		fmt.Println("Vouslez-vous monter d'un étage ? yes/no")
		fmt.Printf("Écrivez votre réponse : ")
		fmt.Scan(&validation)
		if validation == "yes" {
			level++
			return level
		} else if validation == "no" {
			break
		}
	case 1:
		validation := ""
		fmt.Println("Où voulez-vous aller ? up/down/stay")
		fmt.Printf("Écrivez votre réponse : ")
		fmt.Scan(&validation)
		if validation == "up" {
			level++
			return level
		} else if validation == "down" {
			level--
			return level
		} else if validation == "stay" {
			break
		}
	case 2:
		validation := ""
		fmt.Println("Où voulez-vous aller ? up/down/stay")
		fmt.Printf("Écrivez votre réponse : ")
		fmt.Scan(&validation)
		if validation == "up" {
			level++
			return level
		} else if validation == "down" {
			level--
			return level
		} else if validation == "stay" {
			break
		}
	case 3:
		validation := ""
		fmt.Println("Voulez-vous descendre d'un étage ? yes/no")
		fmt.Printf("Écrivez votre réponse : ")
		fmt.Scan(&validation)
		if validation == "yes" {
			level--
			return level
		} else if validation == "no" {
			break
		}
	}
	return level
}
