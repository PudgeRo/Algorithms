package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(quadraticEquation(-5, 4, 1))
}

func quadraticEquation(a, b, c float64) string {
	if a == 0 {
		if b != 0 {
			return fmt.Sprintf("%.2f", -c / b)
		}
		if b == 0 && c == 0 {
			return "Infinity number of solutions"
		}
		if b == 0 && c != 0 {
			return "There are no solutions"
		}
	} else {
		d := b * b - 4 * a * c
		if d < 0 {
			return "There are no solutions"
		} else if d == 0 {
			x := -b / (2 * a)
			return fmt.Sprintf("%.2f", x)
		} else if d > 0 {
			x1 := (-b + math.Sqrt(d)) / (2 * a)
			x2 := (-b - math.Sqrt(d)) / (2 * a)
			if x1 > x2 {
				return fmt.Sprintf("%.2f, %.2f", x2, x1)
			} else {
				return fmt.Sprintf("%.2f, %.2f", x1, x2)
			}
		}		 
	}
	return "Oops"
}