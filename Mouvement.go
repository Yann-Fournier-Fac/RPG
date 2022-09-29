package rpg

import (
	"fmt"

	"github.com/fatih/color"
)

// Revoir comment faire bien spawn le joueur apres un combat annulé sans effacer la case jaune

func Moove(floor [10][10]Case, nbr int, joueur Player, Nesrine PNJ, Guillaume PNJ, Paul PNJ) int {
	Map(floor, nbr)
	var i int = 0
	var j int = 0
	if nbr != 0 {
		i = 5
		j = 8
	}
	Dweapon := Weapon()
	Darmor := Armor()
	Dheal := Heal()
	longueur := len(floor)
	Result := false
	fmt.Printf("Here are the commands that you can use to play this game : z/s/q/d/i/s/exit ")
	fmt.Printf("\n")
	for {
		Mouvement := ""
		fmt.Scan(&Mouvement)
		switch Mouvement {
		case "z":
			if j > 0 {
				floor[i][j-1].Isplayer = true
				floor[i][j-1].Affichage.color = color.New(color.BgGreen)
				floor[i][j-1].Affichage.text = "   "
				floor[i][j].Isplayer = false
				floor[i][j].Affichage.color = color.New(color.BgBlue)
				floor[i][j].Affichage.text = "   "
				Map(floor, nbr)
				j -= 1
			}
		case "s":
			if j < longueur-1 {
				floor[i][j+1].Isplayer = true
				floor[i][j+1].Affichage.color = color.New(color.BgGreen)
				floor[i][j+1].Affichage.text = "   "
				floor[i][j].Isplayer = false
				floor[i][j].Affichage.color = color.New(color.BgBlue)
				floor[i][j].Affichage.text = "   "
				Map(floor, nbr)
				j += 1
			}
		case "d":
			if i < longueur-1 {
				floor[i+1][j].Isplayer = true
				floor[i+1][j].Affichage.color = color.New(color.BgGreen)
				floor[i+1][j].Affichage.text = "   "
				floor[i][j].Isplayer = false
				floor[i][j].Affichage.color = color.New(color.BgBlue)
				floor[i][j].Affichage.text = "   "
				Map(floor, nbr)
				i += 1
			}
		case "q":
			if i > 0 {
				floor[i-1][j].Isplayer = true
				floor[i-1][j].Affichage.color = color.New(color.BgGreen)
				floor[i-1][j].Affichage.text = "   "
				floor[i][j].Isplayer = false
				floor[i][j].Affichage.color = color.New(color.BgBlue)
				floor[i][j].Affichage.text = "   "
				Map(floor, nbr)
				i -= 1
			}
		case "i":
			fmt.Print(joueur.Inventaire)
		case "h":
			if len(joueur.Inventaire[2]) == 0 {
				fmt.Println("You don't have heal")
				break
			}
			validation := ""
			fmt.Println("Do you want to heal you ? yes/no")
			fmt.Printf("Write your answer : ")
			fmt.Scan(&validation)
			if validation == "yes" {
				fmt.Println(joueur.Inventaire[2])
				soin := ""
				fmt.Printf("Choose the heal that you want to use, Write the name : ")
				fmt.Scan(&soin)
				joueur.LP += joueur.Inventaire[2][soin]
			}
		case "exit":
			fmt.Println("Do you want to save : yes/no")
			fmt.Println("Write you anwser : ")
			save := ""
			fmt.Scan(&save)
			if save == "yes" {
				fmt.Printf("Your game has been save")
				// fonction save
				return 10
			} else {
				return 10
			}
		}

		if floor[i][j].Isdistributeur {
			if (i == 5) && (j == 0) {
				Dweapon = distributeur(Dweapon, &joueur, 0)
				floor[i][j+1].Isplayer = true
				floor[i][j+1].Affichage.color = color.New(color.BgGreen)
				floor[i][j+1].Affichage.text = "   "
				floor[i][j].Isplayer = false
				floor[i][j].Affichage.color = color.New(color.BgBlack)
				floor[i][j].Affichage.text = "   "
				j += 1
			} else if (i == 9) && (j == 5) {
				Darmor = distributeur(Darmor, &joueur, 1)
				floor[i-1][j].Isplayer = true
				floor[i-1][j].Affichage.color = color.New(color.BgGreen)
				floor[i-1][j].Affichage.text = "   "
				floor[i][j].Isplayer = false
				floor[i][j].Affichage.color = color.New(color.BgBlack)
				floor[i][j].Affichage.text = "   "
				i -= 1
			} else if (i == 0) && (j == 4) {
				Dheal = distributeur(Dheal, &joueur, 2)
				floor[i+1][j].Isplayer = true
				floor[i+1][j].Affichage.color = color.New(color.BgGreen)
				floor[i+1][j].Affichage.text = "   "
				floor[i][j].Isplayer = false
				floor[i][j].Affichage.color = color.New(color.BgBlack)
				floor[i][j].Affichage.text = "   "
				i += 1
			}
			floor[i][j].Affichage.color = color.New(color.BgBlack)
			floor[i][j].Affichage.text = "   "

		} else if floor[i][j].Enemies {
			validation := ""
			if nbr == 0 {
				fmt.Println("Do you want to fight against Nesrine : yes/no")
				fmt.Printf("Write your answer : ")
				fmt.Scan(&validation)
				if validation == "yes" {
					Result = Fight(joueur, Nesrine)
					if !Result {
						joueur.LP = joueur.LP / 2
					} else {
						joueur.Credits += Nesrine.Credits
					}
				}
				fmt.Printf("Here are the commands that you can use to play this game : z/s/q/d/i/s/exit ")
				fmt.Printf("\n")
				floor[i][j-1].Isplayer = true
				floor[i][j-1].Affichage.color = color.New(color.BgGreen)
				floor[i][j-1].Affichage.text = "   "
				floor[i][j].Isplayer = false
				floor[i][j].Affichage.color = color.New(color.BgYellow)
				floor[i][j].Affichage.text = "   "
				j -= 1
			} else if nbr == 1 {
				fmt.Println("Do you want to fight against Guillaume: yes/no")
				fmt.Printf("Write your answer : ")
				fmt.Scan(&validation)
				if validation == "yes" {
					Result = Fight(joueur, Guillaume)
					if !Result {
						joueur.LP = joueur.LP / 2

					} else {
						joueur.Credits += Guillaume.Credits

					}
				}
				fmt.Printf("Here are the commands that you can use to play this game : z/s/q/d/i/s/exit ")
				fmt.Printf("\n")
				floor[i][j+1].Isplayer = true
				floor[i][j+1].Affichage.color = color.New(color.BgGreen)
				floor[i][j+1].Affichage.text = "   "
				floor[i][j].Isplayer = false
				floor[i][j].Affichage.color = color.New(color.BgYellow)
				floor[i][j].Affichage.text = "   "
				j += 1
			} else if nbr == 3 {
				fmt.Println("Do you want to fight against Paul : yes/no")
				fmt.Printf("Write your answer : ")
				fmt.Scan(&validation)
				if validation == "yes" {
					Result = Fight(joueur, Paul)
					if !Result {
						joueur.LP = joueur.LP / 2
					} else {
						fmt.Printf("Bravo tu as sauvé Ynov")
						return 10
					}
				}
				fmt.Printf("Here are the commands that you can use to play this game : z/s/q/d/i/s/exit ")
				fmt.Printf("\n")
				floor[i-1][j].Isplayer = true
				floor[i-1][j].Affichage.color = color.New(color.BgGreen)
				floor[i-1][j].Affichage.text = "   "
				floor[i][j].Isplayer = false
				floor[i][j].Affichage.color = color.New(color.BgYellow)
				floor[i][j].Affichage.text = "   "
				i -= 1
			}
		} else if (i == 5) && (j == 9) {
			nbr = Staires(nbr)
			return nbr
		}
	}
}
