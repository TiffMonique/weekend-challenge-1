package main

import (
	"fmt"
	"os"

	"moniquedev.lat/challenge/evenodd"
	"moniquedev.lat/challenge/loops"
	"moniquedev.lat/challenge/physics"
)

func main() {
	args := os.Args
	functionName := args[1]
	switch functionName {
	case "evenodd":
		var num int
		fmt.Println("Qué número quieres evaluar?")
		fmt.Scan(&num)
		fmt.Println(evenodd.Evenodd(num))
	case "ohm":
		var v float32
		var r float32
		var i float32
		fmt.Println("Dame los valores para v,r e i")
		fmt.Scan(&v, &r, &i)
		fmt.Println(physics.Ohm(v, r, i))
	case "foobar":
		var num int
		fmt.Println("Cual sera el limite de foobar?")
		fmt.Scan(&num)
		fmt.Println(loops.Foobar(num))
	}
}
