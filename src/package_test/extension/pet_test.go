package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("pet speaking ")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("to", host)
}

type Dog struct {
	Pet
}

//func (d *Dog) Speak() {
//	fmt.Print("dog speaking ")
//}
//
//func (d *Dog) SpeakTo(host string) {
//	d.Speak()
//	fmt.Println("to", host)
//}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	fmt.Println()
}
