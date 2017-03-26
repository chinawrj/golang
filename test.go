package main

import "fmt"
import "log"
import "net/http"
import _ "net/http/pprof"

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

func test_array () {
    var a = []int {1, 2, 3, 4, 5};
    var b = a;
    log.Println("array origin", a);
    b[1]++;
    log.Println("array copy", a, b);
}

func test_struct () {
    p1 := new(person);
    p1.speak();

    p2 := person{};
    p2.speak();

    p3 := person{name: "person name"};
    p3.speak();

    p4 := person{"person name"};
    p4.speak();
}

func main() {
    test_struct();

    p1 := person{"GS"};
    c1 := cat{"Kitty "};
    animal_speak(p1);
    animal_speak(c1);
    test_array();
    log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
}
