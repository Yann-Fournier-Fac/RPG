package rpg

import (
	"fmt"

	"github.com/fatih/color"
)

// Revoir comment organiser le code (peut etre faire plusieurs fonctions Moove pour chacun des floor)

func Moove(floor [10][10]Case, nbr int, joueur Player, Nesrine PNJ, Guillaume PNJ, Paul PNJ) {
	Map(floor, nbr)
	var i int = 0
	var j int = 0
	longueur := len(floor)
	Dweapon := []Item{}
	Darmor := []Item{}
	Dheal := []Item{}
	Result := false
	fmt.Printf("Here are the commands that you can use to play this game : z/s/q/d/i/exit ")
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
		case "exit":
			fmt.Printf("Do you want to save : yes/no")
			fmt.Printf("Write you anwser : ")
			save := ""
			fmt.Scan(&save)
			if save == "yes" {
				fmt.Printf("Your game has been save")
				// fonction save
				return
			} else {
				return
			}
		}
		if floor[i][j].Isdistributeur {
			if (i == 5) && (j == 0) {
				Dweapon = distributeur(Dweapon, &joueur)
			} else if (i == 9) && (j == 5) {
				Darmor = distributeur(Darmor, &joueur)
			} else if (i == 0) && (j == 4) {
				Dheal = distributeur(Dheal, &joueur)
			}
		} else if floor[i][j].Enemies {
			if nbr == 0 {
				Result = Fight(joueur, Nesrine)
				if !Result {
					joueur.LP = joueur.LP / 2
				}
			} else if nbr == 1 {
				Result = Fight(joueur, Guillaume)
				if !Result {
					joueur.LP = joueur.LP / 2
				}
			} else if nbr == 3 {
				Result = Fight(joueur, Paul)
				if !Result {
					joueur.LP = joueur.LP / 2
				}
			}
		} else if (i == 5) && (j == 9) {
			nbr = Staires(nbr)
		}
	}
}
