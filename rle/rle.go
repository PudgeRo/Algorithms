package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(task1("AABCCCXXXXYZZ"))
	fmt.Println(task2("AABCCCXXXXYZZ"))
}

func task1(s string) string {
	/*
		Дана строка (возможно, пустая), состоящая из букв A-Z
		AABCCCXXXXYZZ
		Нужно написать функцию, которая на выходе даст строку
		вида:
		ABCXYZ
	*/
	lastSym := string(s[0])
	ans := []string{}
	for i := range s {
		if string(s[i]) != lastSym {
			ans = append(ans, lastSym)
			lastSym = string(s[i])
		}
	}
	ans = append(ans, lastSym)
	return strings.Join(ans, "")
}

func pack(s string, cnt int) string {
	if cnt > 1 {
		return s + fmt.Sprintf("%v", cnt)
	}
	return s
}

func task2(s string) string {
	/*
		Дана строка (возможно, пустая), состоящая из букв A-Z
		AABCCCXXXXYZZ
		Нужно написать функцию, которая на выходе даст строку
		вида:
		ABCXYZ
	*/
	lastSym := string(s[0])
	lastPos := 0
	ans := []string{}
	for i := range s {
		if string(s[i]) != lastSym {
			ans = append(ans, pack(lastSym, i - lastPos))
			lastPos = i
			lastSym = string(s[i])
		}
	}
	ans = append(ans, pack(string(s[lastPos]), len(s) - lastPos))
	return strings.Join(ans, "")
}

