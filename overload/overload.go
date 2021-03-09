package overload

import "fmt"

type Animal struct {
	ID 	string `default:"1"`
}
type Cat struct {
	Animal
	Name	string
}


func (animal *Animal) Run() {
	fmt.Println("Animal is running...")
}

func (cat *Cat) Run() {
	fmt.Printf("Cat %s is running...", cat.Name)
}