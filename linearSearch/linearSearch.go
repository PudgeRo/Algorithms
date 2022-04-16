package main

import (
	"math"
	"strings"
	"unicode/utf8"
)

func main() {
	
}

func firstEnter(seq []int, x int) int {
	// Дана последовательность чисел длиной N
	// Найти левое вхождение положительного числа Х
	// в нее или вывести -1б если число Х не встречалось
	ans := -1
	for i := range seq {
		if ans == -1 && seq[i] == x {
			ans = i
		}
	}
	return ans
}

func lastEnter(seq []int, x int) int {
	// Дана последовательность чисел длиной N
	// Найти правое вхождение положительного числа Х
	// в нее или вывести -1б если число Х не встречалось
	ans := -1
	for i := range seq {
		if seq[i] == x {
			ans = i
		}
	}
	return ans
}

func maxFromSeq(seq []int, x int) int {
	// Дана последовательность чисел длиной N (N > 0)
	// Найти максимальное число в последовательности
	ans := 0
	for i := range seq {
		if seq[i] > seq[ans] {
			ans = i
		}
	}
	return seq[ans]
}

func twoMaxSeq(seq []int) (int, int) {
	// Дана последовательность чисел длиной N (N > 1)
	// Найти максимальное число в последовательности 
	// и второе по величине число (такое, которое будет
	// максимальным, если вычеркунть из последовательности
	// одно максимальное число).
	max1 := math.Max(float64(seq[0]), float64(seq[1]))
	max2 := math.Min(float64(seq[0]), float64(seq[1]))
	for i := range seq {
		if float64(seq[i]) > max1 {
			max2 = max1
			max1 = float64(seq[i])
		} else if float64(seq[i]) > max2 {
			max2 = float64(seq[i])
		}
	}
	return int(max1), int(max2)
}

func findMinEven(seq []int) int {
	// Дана последовательность чисел длиной N
	// Найти минимальное четное число в последовательности
	// или вывести -1, если такеого не существует
	ans := -1
	flag := false
	for i := range seq {
		if seq[i] % 2 == 0 && (!flag || seq[i] < ans) {
			ans = seq[i]
			flag = true
		}
	}
	return ans
}

func shortWords(seq []string) string {
	minLen := utf8.RuneCountInString(seq[0])
	for _, word := range seq {
		if utf8.RuneCountInString(word) < minLen {
			minLen = utf8.RuneCountInString(word)
		}
	}
	ans := []string{}
	for _, word := range seq {
		if utf8.RuneCountInString(word) == minLen {
			ans = append(ans, word)
		}
	}
	return strings.Join(ans, " ")
}