package rpg

type Save struct {
	jouer  *Player // le joueur
	level  int     // save de l'étage
	weapon []Item  // état des distributeurs
	armor  []Item
	heal   []Item
	x      int //coordonnée => i
	y      int // coordonnée => j
}

func Sauveg(play Player, level int, weapon []Item, armor []Item, heal []Item, x, y int) Save {
	var Sauvegard Save
	Sauvegard.jouer = &play
	Sauvegard.level = level
	Sauvegard.weapon = weapon
	Sauvegard.armor = armor
	Sauvegard.heal = heal
	Sauvegard.x = x
	Sauvegard.y = y
	return Sauvegard
}
