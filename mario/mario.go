package mario

import "fmt"

func Mario(height int) {
	for i := 1; i <= height; i++ {
		for j := 1; j <= height-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
