package main

func main() {
	// Дана строка (в кодировке UTF-8)
	// Найти самый часто встречающийся в ней символ.
	// Если несколько символов встречаются одинаково
	// часто, то можно вывести любой.
	// solve1("affssdfasdfff")
}

func solve1(str string) (string, int) {
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
	return result, max
}

func solve2(str string) (string, int) {
	anscnt := 0
	var ans string
	dict := make(map[string]int)
	for _, now := range str {
		dict[string(now)] += 1
	}
	for key := range dict {
		if dict[key] > anscnt {
			anscnt = dict[key]
			ans = key
		}
	}
	return ans, anscnt
}


