package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	Name string
	City string
	Age  int
}

// Метод структуры Human
func (h *Human) Greeting() {
	fmt.Printf("Hello, my name is %s, I live in %s and I am %d years old\n", h.Name, h.City, h.Age)
}

// Структура Action, в которую встраивается Human, соответственно она имеет доступ ко всем полям и методам Human
type Action struct {
	Human
	Activity string
}

// Метод структуры Action имеет доступ как к своим полям, так и к полям встроенной в себя структуры Human
func (a *Action) DoActivity() {
	fmt.Printf("%s is %s\n", a.Name, a.Activity)
}

// Добавил просто чтобы показать, что методы встроенной структуры можно вызывать из методов Action
func (a *Action) GreetingWhileActivity() {
	a.Greeting()
	fmt.Printf("Right now I am %s\n", a.Activity)
}

func main() {
	action := Action{
		Human: Human{
			Name: "Gopher",
			City: "Moscow",
			Age:  20,
		},
		Activity: "Learning Go",
	}
	// Можем вызвать метод встоенной в Action структуры
	action.Greeting()

	// И понятное дело метод самой Action
	action.DoActivity()

	// Метод Action, который внутри себя вызывает метод встроенной структуры Human
	action.GreetingWhileActivity()
}
