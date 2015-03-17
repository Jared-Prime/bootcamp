package main

import (
	"fmt"
	"github.com/Jared-Prime/bootcamp/pkg/exercise"
)

func main() {
	p := exercise.Player{
		&exercise.User{Id: 42, Name: "Jared", Location: "Chicago"},
		1234}
	fmt.Println(p.Greetings())
}
