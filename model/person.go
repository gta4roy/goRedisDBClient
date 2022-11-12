package model

type Person struct {
	Name   string  `json:"name"`
	Phone  string  `json:"phone"`
	Salary float32 `json:"salary"`
	Age    int     `json:"age"`
}

func NewPerson() *Person {
	p := Person{}
	p.Age = 0
	p.Name = ""
	p.Phone = ""
	p.Salary = 0.0
	return &p
}

func NewPersonWithArgs(name string, age int, phone string, salary float32) *Person {
	p := Person{}
	p.Age = age
	p.Name = name
	p.Phone = phone
	p.Salary = salary
	return &p
}
