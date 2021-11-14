package main

import (
	"fmt"
)

type apple struct {
	name  string
	score int
}

func (a *apple) say() string {
	return a.name
}

type banana struct {
	job   string
	price int
}

func (b *banana) speak() string {
	return b.job
}

type shop struct {
	a *apple
	b *banana
}

func (s *shop) sail() string {
	return s.a.say() + s.b.speak()
}

func Init1() {
	a := NewA()
	b := NewB()
	s := NewS(a, b)
	fmt.Println(s.sail())
}

func NewS(a *apple, b *banana) shop {
	s := shop{
		a: a,
		b: b,
	}
	return s
}

func NewB() *banana {
	b := &banana{
		job:   "/cccc",
		price: 0,
	}
	return b
}

func NewA() *apple {
	a := &apple{
		name:  "asdfasdf",
		score: 0,
	}
	return a
}

func main() {
	//Init1()
	s := InitializeShop()
	sail := s.sail()
	fmt.Println(sail)
}
