package main

import (
	"fmt"
	"os"

	"moniquedev.lat/challenge/bmi"
	"moniquedev.lat/challenge/evenodd"
	"moniquedev.lat/challenge/loops"
	"moniquedev.lat/challenge/mario"
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
	case "bmi":
		var weight, height float32

		fmt.Println("How much do you weigh? (don't lie)")
		fmt.Scanln(&weight)

		fmt.Println("How tall are you? (barefoot)")
		fmt.Scanln(&height)

		bmi := bmi.Bmi(weight, height)
		fmt.Printf("Right now your BMI is %.2f\n", bmi)

		if bmi < 18.5 {
			fmt.Println("You are underweight, add more potato to the broth")
		} else if bmi < 25 {
			fmt.Println("You have a normal weight, I have healthy envy of you")
		} else {
			fmt.Println("You are overweight, I know, the pandemic has affected us all")
		}
	case "mario":
		var height int
		for {
			fmt.Print("Pyramid height: ")
			_, err := fmt.Scanln(&height)
			if err != nil || height < 1 || height > 8 {
				fmt.Println("Invalid input. Please enter a number between 1 and 8.")
			} else {
				break
			}
		}
		mario.Mario(height)
	}
}
