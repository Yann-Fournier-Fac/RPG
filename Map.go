package rpg

import (
	"fmt"

	"github.com/fatih/color"
	//"strconv"
)

/*type Color struct{
	Red
	Yellow
	Pink
	Brown
	Blue
	White
	Green
}*/

type Case struct { //
	Walkable       bool
	Affichage      Affichage
	Enemies        bool
	Item           bool
	Watchable      bool
	Isplayer       bool
	Isstaire       bool
	Isdistributeur bool
}

type Affichage struct {
	text  string
	color *color.Color
}

func Wall() [12][12]Case {
	// Création des murs
	Wall := [12][12]Case{}
	var block Case
	for i := 0; i <= 11; i++ {
		for j := 0; j <= 11; j++ {
			if (i == 0) || (i == 11) {
				Wall[i][j] = block
				Wall[i][j].Walkable, Wall[i][j].Enemies, Wall[i][j].Item, Wall[i][j].Isplayer, Wall[i][j].Isstaire, Wall[i][j].Isdistributeur = false, false, false, false, false, false
				Wall[i][j].Watchable = true
				Wall[i][j].Affichage.color = color.New(color.BgWhite)
				Wall[i][j].Affichage.text = "   "
				//Wall[i][j].Affichage = " | "
			} else if (j == 0) || (j == 11) {
				Wall[i][j] = block
				Wall[i][j].Walkable, Wall[i][j].Enemies, Wall[i][j].Item, Wall[i][j].Isplayer, Wall[i][j].Isstaire, Wall[i][j].Isdistributeur = false, false, false, false, false, false
				Wall[i][j].Watchable = true
				Wall[i][j].Affichage.color = color.New(color.BgWhite)
				Wall[i][j].Affichage.text = "   "
				//Wall[i][j].Affichage = "---"
			}
		}
	}
	return Wall
}

func Floor0() ([10][10]Case, int) {
	// Création du RDC
	var block Case
	f0 := [10][10]Case{}
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			f0[i][j] = block
			f0[i][j].Enemies, f0[i][j].Item, f0[i][j].Isplayer, f0[i][j].Isstaire, f0[i][j].Isdistributeur = false, false, false, false, false
			f0[i][j].Walkable, f0[i][j].Watchable = true, true
			f0[i][j].Affichage.color = color.New(color.BgBlue)
			f0[i][j].Affichage.text = "   "
			//f0[i][j].Affichage = "   "

		}
	}
	f0[0][0].Affichage.color = color.New(color.BgGreen)
	f0[0][0].Affichage.text = "   "
	//f0[0][0].Affichage = " J "
	f0[0][0].Isplayer = true

	f0[7][7].Affichage.color = color.New(color.BgHiYellow)
	f0[7][7].Affichage.text = "   "
	f0[7][7].Enemies = true
	//f0[7][7].Affichage = " N "

	f0[5][9].Affichage.color = color.New(color.BgHiMagenta)
	f0[5][9].Affichage.text = "   "
	//f0[9][5].Affichage = " S "
	f0[5][9].Isstaire = true

	return f0, 0
}

func Floor1() ([10][10]Case, int) {
	// Création du RDC
	var block Case
	f1 := [10][10]Case{}
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			f1[i][j] = block
			f1[i][j].Enemies, f1[i][j].Item, f1[i][j].Isplayer, f1[i][j].Isstaire, f1[i][j].Isdistributeur = false, false, false, false, false
			f1[i][j].Walkable, f1[i][j].Watchable = true, true
			f1[i][j].Affichage.color = color.New(color.BgBlue)
			f1[i][j].Affichage.text = "   "
			//f1[i][j].Affichage = "   "

		}
	}
	f1[9][0].Affichage.color = color.New(color.BgHiYellow)
	f1[0][0].Affichage.text = "   "
	f1[9][0].Enemies = true
	//f1[9][0].Affichage = " G "

	f1[5][9].Affichage.color = color.New(color.BgHiMagenta)
	f1[5][9].Affichage.text = "   "
	//f1[9][5].Affichage = " S "
	f1[5][9].Isstaire = true

	return f1, 1
}

func Floor2() ([10][10]Case, int) {
	// Création du RDC
	var block Case
	f2 := [10][10]Case{}
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			f2[i][j] = block
			f2[i][j].Enemies, f2[i][j].Item, f2[i][j].Isplayer, f2[i][j].Isstaire, f2[i][j].Isdistributeur = false, false, false, false, false
			f2[i][j].Walkable, f2[i][j].Watchable = true, true
			f2[i][j].Affichage.color = color.New(color.BgBlue)
			f2[i][j].Affichage.text = "   "
			//f2[i][j].Affichage = "   "

		}
	}
	f2[5][0].Affichage.color = color.New(color.BgBlack)
	f2[5][0].Affichage.text = "   "
	f2[5][0].Isdistributeur = true

	f2[9][5].Affichage.color = color.New(color.BgBlack)
	f2[9][5].Affichage.text = "   "
	f2[9][5].Isdistributeur = true

	f2[0][4].Affichage.color = color.New(color.BgBlack)
	f2[0][4].Affichage.text = "   "
	f2[0][4].Isdistributeur = true

	f2[5][9].Affichage.color = color.New(color.BgHiMagenta)
	f2[5][9].Affichage.text = "   "
	//f1[9][5].Affichage = " S "
	f2[5][9].Isstaire = true

	return f2, 2
}

func Floor3() ([10][10]Case, int) {
	// Création du RDC
	var block Case
	f3 := [10][10]Case{}
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			f3[i][j] = block
			f3[i][j].Enemies, f3[i][j].Item, f3[i][j].Isplayer, f3[i][j].Isstaire, f3[i][j].Isdistributeur = false, false, false, false, false
			f3[i][j].Walkable, f3[i][j].Watchable = true, true
			f3[i][j].Affichage.color = color.New(color.BgBlue)
			f3[i][j].Affichage.text = "   "
			//f3[i][j].Affichage = "   "

		}
	}
	f3[8][5].Affichage.color = color.New(color.BgYellow)
	f3[8][5].Affichage.text = "   "
	f3[8][5].Enemies = true

	f3[5][9].Affichage.color = color.New(color.BgHiMagenta)
	f3[5][9].Affichage.text = "   "
	//f1[9][5].Affichage = " S "
	f3[5][9].Isstaire = true

	return f3, 3
}

func Map(floor [10][10]Case, p int) {
	fmt.Println("Floor", p)
	Wall := Wall()
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			if ((i == 0) || (i == 11)) || ((j == 0) || (j == 11)) {
				Wall[j][i].Affichage.color.Print(Wall[j][i].Affichage.text)
				//fmt.Printf(Wall[j][i].Affichage)
			} else {
				floor[j-1][i-1].Affichage.color.Print(floor[j-1][i-1].Affichage.text)
				//fmt.Printf(floor[j][i].Affichage)
			}
		}
		fmt.Printf("\n")
	}
}

// Voici une fonction a tester pour afficher toutes les maps
// vous pouvez la copier dans le fichier main.go du dossier test

/*func main() {
	f0, nb0 := rpg.Floor0()
	f1, nb1 := rpg.Floor1()
	f2, nb2 := rpg.Floor2()
	f3, nb3 := rpg.Floor3()
	rpg.Map(f0, nb0)
	rpg.Map(f1, nb1)
	rpg.Map(f2, nb2)
	rpg.Map(f3, nb3)
}*/
