package main

import (
	"fmt"
	"github.com/Jared-Prime/bootcamp/pkg/exercise"
)

func main() {
	p := exercise.Player{}
	p.Id = 42
	p.Name = "Jared"
	p.Location = "Chicago"
	fmt.Println(p.Greetings())
}
