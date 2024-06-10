package pkg

import (
	"fmt"
)

type Flyer interface {
	Fly() string
}

type Eater interface {
	Eat() string
}

type Bird struct {
	Name string
}

func (b Bird) Fly() string {
	return fmt.Sprintf("%s is flying", b.Name)
}

type FruitEater struct {
	Name string
}

func (f FruitEater) Eat() string {
	return fmt.Sprintf("%s is eating fruits", f.Name)
}

type MeatEater struct {
	Name string
}

func (m MeatEater) Eat() string {
	return fmt.Sprintf("%s is eating meat", m.Name)
}

func RunStruct() {
	bird := Bird{Name: "Sparrow"}
	fruitEater := FruitEater{Name: "Monkey"}
	meatEater := MeatEater{Name: "Lion"}

	fmt.Println(bird.Fly())

	fmt.Println(fruitEater.Eat())

	fmt.Println(meatEater.Eat())
}
