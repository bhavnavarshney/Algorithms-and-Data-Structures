package strings

import "fmt"

func TestKMP() {
	fmt.Println("KMP : Substring found at positions:", kmp("AAAAABAAABA", "AAAA")) // {0, 1, 2, 3}

}
