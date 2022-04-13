package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Входные данные: строка в виде последовательности чисел
	// без пробела
	fmt.Println(sumSequence("123321"))
	fmt.Println(maxSequence("12351233543241"))
}

func sumSequence(str string) int {
	var seq []int
	for _, i := range str {
		num, _ := strconv.Atoi(string(i))
		seq = append(seq, num)
	}
	if len(seq) == 0 {
		return 0
	} else {
		sum := seq[0]
		for i := 1; i < len(seq); i++ {
			sum += seq[i]
		}
		return sum
	}
}

func maxSequence(str string) string {
	var seq []int
	for _, i := range str {
		num, _ := strconv.Atoi(string(i))
		seq = append(seq, num)
	}
	if len(seq) == 0 {
		return "Empty sequence"
	} else {
		max := seq[0]
		for i := 1; i < len(seq); i++ {
			if seq[i] > max {
				max = seq[i]
			}
		}
		return strconv.Itoa(max)
	}
}