package main

import (
	"log"
	Lib "github.com/jeffotoni/modhello/modlib/lib"
	//Lib "/modhello/modlib/lib"
)

func main() {

	log.Println("start")
	run(1, 2)
}

func run(a, b int) {

	log.Printf("Sum(%d,%d) = %d", a, b, Lib.Sum(a, b))
	//log.Println("run func")
}
