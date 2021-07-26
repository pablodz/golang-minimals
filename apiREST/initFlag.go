package main

import (
	"flag"
	"log"
)

var name string
var age int

func init() {
	flag.StringVar(&name, "name", "stranger", "your beautiful name")
	flag.IntVar(&age, "age", 0, "your age")
}

func main() {
	flag.Parse()
	log.Printf("Hello %s (%d years). Welcome!", name, age)
}
