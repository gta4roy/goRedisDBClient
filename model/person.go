package model

type Person struct {
	Name   string
	Phone  string
	Salary float32
	Age    int
}

func NewPerson() *Person {
	p := Person{}
	p.Age = 0
	p.Name = ""
	p.Phone = ""
	p.Salary = 0.0
	return &p
}
