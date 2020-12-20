package strings

import "fmt"

func kmp(haystack, needle string) []int {
	subsPositions := []int{}
	patternTable := fillTable(needle)
	fmt.Println("FillTable: ", patternTable)
	return subsPositions
}

func fillTable(needle string) []int {
	lenNeedle := len(needle)
	table := make([]int, lenNeedle)
	i, j := 0, 1

	for j < lenNeedle {
		if needle[i] == needle[j] {
			table[j] = table[j-1] + 1
			i++
			j++
			continue
		}
		if i == 0 {
			table[j] = i
			j++
		} else {
			i = table[i-1]
		}
	}
	return table
}
