package main

import (
	"log"

	base62 "../base62"
)

func main() {
	x := 100
	log.Println("usebase.go")
	base62String := base62.ToBase62(x)
	log.Println(base62String)
	normalNumber := base62.ToBase10(base62String)
	log.Println(normalNumber)
}
