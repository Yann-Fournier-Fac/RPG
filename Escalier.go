package rpg

import "fmt"

func Staires(level int) int {
	switch level {
	case 0:
		validation := ""
		fmt.Println("Do you want to go upstraire : yes/no")
		fmt.Printf("Write your anwser : ")
		fmt.Scan(&validation)
		if validation == "yes" {
			level++
			return level
		} else if validation == "no" {
			break
		}
	case 1:
		validation := ""
		fmt.Println("Where do you want to go : up/down/stay")
		fmt.Printf("Write your anwser : ")
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
		fmt.Println("Where do you want to go : up/down/stay")
		fmt.Printf("Write your anwser : ")
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
		fmt.Println("Do you want to go downstraire : yes/no")
		fmt.Printf("Write your anwser : ")
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
