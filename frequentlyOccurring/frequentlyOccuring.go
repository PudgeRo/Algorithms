package main

func main() {
	// Дана строка (в кодировке UTF-8)
	// Найти самый часто встречающийся в ней символ.
	// Если несколько символов встречаются одинаково
	// часто, то можно вывести любой.
	solve1("affssdfasdfff")
}

func solve1(str string) string {
	max := 0
	result := ""
	counter := 0
	for i, k := range str {
		counter = 0
		for _, j := range str {
			if string(k) == string(j) {
				counter++
			}
		}
		if counter > max {
			max = counter
			result = string(str[i])
		}
	}
	return result
}


