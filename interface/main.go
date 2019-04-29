package main

import "fmt"

//if there is any other type inside of the program
//that has a function called getGreeting
//associated it that returns a string then
//that type is also an honorary kind of like
//promoted automatically to also being something of type bot

// so englishBot and spanishBot are also type bot
type bot interface {
	getGretting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGretting())
}

func (englishBot) getGretting() string {
	return "Hi There!"
}

func (spanishBot) getGretting() string {
	return "Hola!"
}
