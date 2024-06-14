package adapter

type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter(name string) *Adapter {
	return &Adapter{
		adaptee: NewAdaptee(name),
	}
}

func (a *Adapter) Greet() {
	a.adaptee.Hello()
}
