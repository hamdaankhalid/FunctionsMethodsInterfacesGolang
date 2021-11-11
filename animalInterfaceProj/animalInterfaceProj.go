package main

import (
	"os"
	"fmt"
	"bufio"
	"errors"
	"strings"
)

// interface for animal
type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (cow Cow) Eat(){
	fmt.Println("grass")
}

func (cow Cow) Move(){
	fmt.Println("walk")
}

func (cow Cow) Speak(){
	fmt.Println("moo")
}

type Bird struct {
	name string
}

func (bird Bird) Eat(){
	fmt.Println("worms")
}

func (bird Bird) Move(){
	fmt.Println("fly")
}

func (bird Bird) Speak(){
	fmt.Println("peep")
}

type Snake struct {
	name string
}

func (snake Snake) Eat(){
	fmt.Println("mice")
}

func (snake Snake) Move(){
	fmt.Println("slither")
}

func (snake Snake) Speak(){
	fmt.Println("hsss")
}

func createAnimal(name string, animalType string) (Animal, error){
	var animal Animal
	var err error
	switch animalType {
	case "cow":
		animal = Cow{name}
	case "bird":
		animal = Bird{name}
	case "snake":
		animal = Snake{name}
	default:
		err = errors.New("Animal must be of type cow, bird, snake")
	}
	return animal, err
}

func delegateAnimal( name string, animalMap map[string]Animal ) (Animal, error){
	if val, ok := animalMap[name]; ok{
		return val, nil
	}else{
		return nil, errors.New("Animal name given has not been created.")
	}
}

func main(){
	animalNameMap := make(map[string]Animal)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome, you can either create an animal or query an animal you have previously created.") 
	fmt.Println("To create an animal enter: newanimal nameOfAnimal typeOfAnimal \n To query an animal enter: query nameOfAnimal method")
	fmt.Println("Animal Types Available: cow, bird, snake. \n Methods Available: eat, move, speak ")
	
	for {
		fmt.Print("> ")
		userInput, err := reader.ReadString('\n')
		if(err!=nil){
			fmt.Println("Error occurred", err)
			continue
		}

		userInput = strings.Replace(userInput, "\n", "", -1)
		words := strings.Fields(userInput)

		if (words[0]=="quit"){
			fmt.Println("Bye")
			break
		}else if(words[0]=="newanimal" && len(words)==3){
			animal, err := createAnimal(words[1], words[2])
			if(err!=nil){
				fmt.Println("Invalid input")
			}
			animalNameMap[words[1]] = animal
		} else if(words[0]=="query" && len(words)==3){
			animal, err := delegateAnimal(words[1], animalNameMap)
			if (err!=nil){
				fmt.Println(err)
				continue
			}
			switch words[2] {
			case "speak":
				animal.Speak()
			case "move":
				animal.Move()
			case "eat":
				animal.Eat()
			default:
				fmt.Println("Method has to be of type speak, move, eat.")
			}

		} else {
			fmt.Println("Please read the instructions")
		}
	}
}