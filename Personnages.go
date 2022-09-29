package rpg

// changer la facon de creer les Dweapon, Darmor, Dheal. faire des dictionnaires

type PNJ struct {
	Name    string
	Enemies bool
	Attaque int
	LP      int
	Credits int
	//Inventaire []Item
	// Photo   ...
	// Ficha   ...
}

type Player struct {
	Name       string
	Attaque    int
	LP         int
	Credits    int
	Inventaire [3]map[string]int
}

type Item struct {
	Name     string
	Isarmor  bool
	Isweapon bool
	Isheal   bool
	LP       int
	Attaque  int
}

// Creation des armes
func Weapon() []Item {

	var Epee Item
	Epee.Name = "epee"
	Epee.Isarmor = false
	Epee.Isweapon = true
	Epee.Isheal = false
	Epee.LP = 0
	Epee.Attaque = 55

	var Mouse Item
	Mouse.Name = "souris"
	Mouse.Isarmor = false
	Mouse.Isweapon = true
	Mouse.Isheal = false
	Mouse.LP = 0
	Mouse.Attaque = 35

	var Elec Item
	Elec.Name = "rallonge"
	Elec.Isarmor = false
	Elec.Isweapon = true
	Elec.Isheal = false
	Elec.LP = 0
	Elec.Attaque = 45

	var Chair Item
	Chair.Name = "chaise"
	Chair.Isarmor = false
	Chair.Isweapon = true
	Chair.Isheal = false
	Chair.LP = 0
	Chair.Attaque = 75

	var Ascii Item
	Ascii.Name = "ascii"
	Ascii.Isarmor = false
	Ascii.Isweapon = true
	Ascii.Isheal = false
	Ascii.LP = 0
	Ascii.Attaque = 2

	var Table Item
	Table.Name = "table"
	Table.Isarmor = false
	Table.Isweapon = true
	Table.Isheal = false
	Table.LP = 0
	Table.Attaque = 90

	var Money Item
	Money.Name = "piece"
	Money.Isarmor = false
	Money.Isweapon = true
	Money.Isheal = false
	Money.LP = 0
	Money.Attaque = 67

	Dweapon := []Item{Epee, Mouse, Elec, Chair, Ascii, Table, Money}
	/*Dweapon := map[string]int{
		Elec.Name:  Elec.Attaque,
		Mouse.Name: Mouse.Attaque,
		Elec.Name:  Elec.Attaque,
		Chair.Name: Chair.Attaque,
		Ascii.Name: Ascii.Attaque,
		Table.Name: Table.Attaque,
		Money.Name: Money.Attaque,
	}*/

	return Dweapon
}

// Creation des Armures
func Armor() []Item {

	var BDE Item
	BDE.Name = "bde"
	BDE.Isarmor = true
	BDE.Isweapon = false
	BDE.Isheal = false
	BDE.LP = 25
	BDE.Attaque = 0

	var Processor Item
	Processor.Name = "processeur"
	Processor.Isarmor = true
	Processor.Isweapon = false
	Processor.Isheal = false
	Processor.LP = 55
	Processor.Attaque = 0

	var Card Item
	Card.Name = "carteynov"
	Card.Isarmor = true
	Card.Isweapon = false
	Card.Isheal = false
	Card.LP = 45
	Card.Attaque = 0

	var Coat Item
	Coat.Name = "veste"
	Coat.Isarmor = true
	Coat.Isweapon = false
	Coat.Isheal = false
	Coat.LP = 100
	Coat.Attaque = 0

	var Navigo Item
	Navigo.Name = "cartenavigo"
	Navigo.Isarmor = true
	Navigo.Isweapon = false
	Navigo.Isheal = false
	Navigo.LP = 15
	Navigo.Attaque = 0

	var PC Item
	PC.Name = "pcportable"
	PC.Isarmor = true
	PC.Isweapon = false
	PC.Isheal = false
	PC.LP = 65
	PC.Attaque = 0

	var Bagpack Item
	Bagpack.Name = "sacados"
	Bagpack.Isarmor = true
	Bagpack.Isweapon = false
	Bagpack.Isheal = false
	Bagpack.LP = 75
	Bagpack.Attaque = 0

	var Kapla Item
	Kapla.Name = "kapla"
	Kapla.Isarmor = true
	Kapla.Isweapon = false
	Kapla.Isheal = false
	Kapla.LP = 10
	Kapla.Attaque = 0

	var Polar Item
	Polar.Name = "vestepolaire"
	Polar.Isarmor = true
	Polar.Isweapon = false
	Polar.Isheal = false
	Polar.LP = 115
	Polar.Attaque = 0

	Darmor := []Item{BDE, Processor, Card, Coat, Navigo, PC, Bagpack, Kapla, Polar}

	return Darmor
}

// Creation des Soins
func Heal() []Item {

	var Water Item
	Water.Name = "gourdeaeau"
	Water.Isarmor = false
	Water.Isweapon = false
	Water.Isheal = true
	Water.LP = 85
	Water.Attaque = 0

	var USB Item
	USB.Name = "cleusb"
	USB.Isarmor = false
	USB.Isweapon = false
	USB.Isheal = true
	USB.LP = 35
	USB.Attaque = 0

	var Food Item
	Food.Name = "repas"
	Food.Isarmor = false
	Food.Isweapon = false
	Food.Isheal = true
	Food.LP = 110
	Food.Attaque = 0

	var Snack Item
	Snack.Name = "snack"
	Snack.Isarmor = false
	Snack.Isweapon = false
	Snack.Isheal = true
	Snack.LP = 55
	Snack.Attaque = 0

	var Power Item
	Power.Name = "batterie"
	Power.Isarmor = false
	Power.Isweapon = false
	Power.Isheal = true
	Power.LP = 45
	Power.Attaque = 0

	Dheal := []Item{Water, USB, Food, Snack, Power}

	return Dheal
}
