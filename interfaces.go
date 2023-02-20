package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func interface() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//very custon logic for english greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	//very custon logic for english greeting
	return "Hola!"
}
