package main

import (
	"fmt"
	"github.com/Jared-Prime/bootcamp/pkg/exercise"
)

func main() {
	p := exercise.NewPlayer(42, "Jared", "Chicago", 1234)
	fmt.Println(p.Greetings())
}
