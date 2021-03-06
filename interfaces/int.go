package main

import (
	"fmt"
	"strings"
)

// A multi-lingual Automated Teller Machine
type ATM struct {
	Money float64
	Talk  Greeter
}

// This is a multi-lingual elevator
type Elevator struct {
	Levels int
	Greeter
}

// We make two elevators, one English and one French
func main() {
	atm1 := ATM{1000., NewGreeter("English")}
	atm2 := ATM{1000., NewGreeter("Francais")}

	atm1.Talk.Hello()
	atm2.Talk.Hello()

	elevator1 := Elevator{3, NewGreeter("English")}
	elevator2 := Elevator{5, NewGreeter("Francais")}

	elevator1.Hello()
	elevator2.Hello()
}

// This interface is used to implement Greeters in different languages,
// without us having to know what greeter we're using actually. As long
// as you have a struct that implements all three functions, it will be
// accepted as a Greeter
type Greeter interface {
	// Greets generally in the actual language
	Hello()
	// Greets the name that is passed as the argument
	Greet(string)
	// Returns the language of that Greeter implementation
	Language() string
}

// Returns a new Greeter with the chosen language. If the language is not
// available, it will return 'nil'
func NewGreeter(language string) Greeter {
	switch strings.ToLower(language){
	case "english":
		return &English{}
	}
	return nil
}

// The English language is counting the number of greetings, too. We can put
// that directly in the structure here
type English struct {
	greetings int
}

// For each greeting, we put the number of greetings so far, and a simple
// "Hello world"
func (e *English)Hello() {
	e.greetings += 1
	fmt.Println(e.greetings, "Hello world")
}

// This greets with the correct name
func (e *English)Greet(s string) {
	e.greetings += 1
	fmt.Println(e.greetings, "Hello", s)
}

// Returns the language-string of that Greeter
func (e *English)Language() string {
	return ""
}