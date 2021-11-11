package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"errors"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}


func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func main(){
	animalMap := make(map[string]Animal)
	animalMap["cow"] = Animal{"grass", "walk", "moo"}
	animalMap["bird"] = Animal{"worms", "fly", "peep"}
	animalMap["snake"] = Animal{"mice", "slither", "hsss"}
	// Repeat loop for user to enter animal name, and attribute
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter an Animal from the choices Cow, Bird, Snake, and then enter attribute eat, move, or speak. Enter X to quit:")
		fmt.Print(">")
		animalInput, err := reader.ReadString('\n')
		animalInput = strings.Replace(animalInput, "\n", "", -1)

		if(err!=nil){
			fmt.Println("Error occurred", err)
			continue
		}

		if animalInput == "X"{
			break
		}

		animal, err := delegateAnimal(animalInput, animalMap)

		if err!=nil{
			fmt.Println("Please choose animal from choices: cow, snake, bird")
			continue
		}
	
		fmt.Print(">")
		animalAttribute, err := reader.ReadString('\n')
		animalAttribute = strings.Replace(animalAttribute, "\n", "", -1)

		if(err!=nil){
			fmt.Printf("Error occurred %s", err)
		}

		if animalAttribute == "X"{
			break
		}

		delegateMethod(animalAttribute, &animal)
	}

}

func delegateAnimal( inputString string, animalMap map[string]Animal ) (Animal, error){
	if val, ok := animalMap[inputString]; ok{
		return val, nil
	}else{
		return Animal{}, errors.New("Invalid input")
	}
}

func delegateMethod( method string, animal *Animal){
	if method == "speak"{
		animal.Speak()
	}else if method == "move"{
		animal.Move()
	}else if method == "eat"{
		animal.Eat()
	}else{
		fmt.Println(method, "Invalid Method")
	}
}
