package rpg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Identite struct {
	Name string
	Age  int
}

func Unmarshal() {
	empJsonData := `{"Name":"George Smith","Age":30,}`
	empBytes := []byte(empJsonData)
	var emp Identite
	json.Unmarshal(empBytes, &emp)
	fmt.Println(emp.Name)
	fmt.Println(emp.Age)
}

func Marshal(emp Identite) []byte {
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(empData))
	return empData
}

func main() {
	file, err := os.Create("Save.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File created successfully")
	defer file.Close()
	emp := Identite{Name: "George Smith", Age: 30}
	marshaled_data := Marshal(emp)
	read, err := os.OpenFile("Save.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer read.Close() // on ferme automatiquement à la fin de notre programme

	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(string(marshaled_data))
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile("Save.txt") // lire le fichier
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
	fmt.Println(string(marshaled_data)) // conversion de byte en string

}
