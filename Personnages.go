package rpg

type PNJ struct {
	Name       string
	Enemies    bool
	Attaque    int
	LP         int
	Credits    int
	Inventaire []Item
	// Photo   ...
	// Ficha   ...
}

type Player struct {
	Name       string
	Attaque    int
	LP         int
	Credits    int
	Inventaire []Item
}

type Item struct {
	Name     string
	Isarmor  bool
	Isweapon bool
	Isheal   bool
	LP       bool
	Attaque  int
}
