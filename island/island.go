package main

import "fmt"

func main() {
	/* 
		Игра просиохти в двумерном мире, который состоит
		из блоков размером 1 на 1 метр. Остров игрока представляет
		собой набор столбцов различной высоты, состоящих из
		блоков камня и окруженной морем. Над островом прошел сильный
		дождь, который заполнил водой все низины, а не 
		поместрившаяся в них вода стекла в море, не увеличив его 
		уровень
		По ландшафту острова определите, сколько блоков воды осталось
		после дождя в низинах на острове. 
	*/
	fmt.Println(task([]int{3,1,3,4,3,5,7,6,4,7,3,2,4,1}))
}

func task(h []int) int {
	maxColumn := 0
	for i := range h {
		if h[i] > h[maxColumn] {
			maxColumn = i
		}
	}
	ans := 0
	nown := 0
	for i := 0; i < maxColumn; i++ {
		if h[i] > nown {
			nown = h[i]
		}
		ans += nown - h[i]
	}
	nown = 0
	for i := len(h) - 1; i > maxColumn; i-- {
		if h[i] > nown {
			nown = h[i]
		}
		ans += nown - h[i]
	}
	return ans
}

