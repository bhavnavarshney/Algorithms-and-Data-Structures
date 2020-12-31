package strings

import "fmt"

func KMP(haystack string, needle string) int {
	Nstr, Npattern := len(haystack), len(needle)
	if Nstr == 0 && Npattern == 0 || Npattern == 0 {
		return 0
	}
	table := fillTable(needle)
	fmt.Println(table)
	str, pattern := 0, 0

	for str < Nstr {
		if haystack[str] == needle[pattern] {
			str++
			pattern++
		}
		if pattern == Npattern {
			return str - pattern
		} else if str < Nstr && haystack[str] != needle[pattern] {
			if pattern != 0 {
				pattern = table[pattern-1]
			} else {
				str++
			}
		}

	}
	return -1
}

func fillTable(needle string) []int {
	n := len(needle)
	table := make([]int, n)
	i := 0
	j := 1

	for j < n {
		if needle[i] == needle[j] {
			table[j] = i + 1
			i++
			j++
		} else {
			if i != 0 {
				i = table[i-1]
			} else {
				table[j] = 0
				j++
			}
		}
	}
	return table
}
