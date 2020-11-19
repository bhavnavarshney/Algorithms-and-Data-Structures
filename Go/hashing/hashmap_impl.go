package hashing

import "fmt"

func TestHashing() {
	hashmap, err := createHashMap(5)
	if err != nil {
		fmt.Printf("Error in creating hashmap. Error:%v", err)
	}

	done := hashmap.putValue("test", "hashmap")
	if !done {
		fmt.Printf("Error adding value in hashmap. Error:%v", err)
	}
	fmt.Println(hashmap, done)
	// value, found := hashmap.getValue("test")
	// if !found {
	// 	fmt.Printf("Value not found in hashmap")
	// }
	// fmt.Println("Value found:", value.value)
}
