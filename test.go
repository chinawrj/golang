package main

import "fmt"

type animal interface {
    speak()
}

type person struct {
    name string
}

func (p person) speak() {
    fmt.Println("name is", p.name)
}

type cat struct {
    pet string
}

func (c cat) speak() {
    fmt.Println("Pet name is", c.pet)
}

func animal_speak(a animal) {
    a.speak();
}

func main() {
    p1 := person{"GS"};
    c1 := cat{"Kitty "};
    animal_speak(p1);
    animal_speak(c1);
}
