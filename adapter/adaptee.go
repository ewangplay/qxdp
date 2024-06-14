package adapter

import "fmt"

type Adaptee struct{
	name string
}

func NewAdaptee(name string) *Adaptee {
	return &Adaptee{name: name}
}

func (a *Adaptee) Hello() {
	fmt.Println("Hello, " + a.name)
}
