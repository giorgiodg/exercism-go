package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(string) string
}

// Next, implement a function SayHello that accepts the name of the visitor and anything that implements the Greeter interface as arguments
// and returns the desired greeting string. For example, imagine a German Greeter implementation for which LanguageName returns "German" and Greet returns "Hallo {name}!":

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

// Now your job is to make the robot work for people that scan Italian passports.
// For that, create a struct Italian and implement the two methods that are needed for the struct to fulfill the
// Greeter interface you set up in task 1. You can greet someone in Italian with "Ciao {name}!".

type Italian struct{}

func (Italian) LanguageName() string {
	return "Italian"
}

func (Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct{}

func (Portuguese) LanguageName() string {
	return "Portuguese"
}

func (Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
