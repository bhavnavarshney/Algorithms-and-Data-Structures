package hashing

import (
	"crypto/sha1"
	"errors"
	"hash"
	"strconv"
)

type Node struct {
	key   string
	value interface{}
}

type HashMap struct {
	count   int
	size    int
	buckets [][]Node
}

func getHash(s string) int {
	// hsha2 := sha256.Sum256([]byte(s))
	// fmt.Println(hsha2)
	// hashbytes := hsha2.Sum(nil)
	// var mySlice = []byte{244, 244, 244, 244, 244, 244, 244, 244}
	// data := binary.BigEndian.Uint64(mySlice)
	// hash, err := strconv.Atoi(fmt.Sprintf("%s", hsha2))
	// if err != nil {
	// 	fmt.Printf("Error in getHash function: %v", err)
	// 	return -1
	// }
	// return hash

	var hashVal hash.Hash
	hashVal = sha1.New()
	hashVal.Write([]byte(s))

	var bytes []byte

	bytes = hashVal.Sum(nil)
	byteToInt, _ := strconv.Atoi(string(bytes))
	return byteToInt
}

func (h *HashMap) getIndex(key string) int {
	hashInt := getHash(key)
	if hashInt != -1 {
		return getHash(key) % h.size
	}
	return -1
}

func (h *HashMap) getValue(key string) (*Node, bool) {
	index := h.getIndex(key)
	if index == -1 {
		return nil, false
	}
	chain := h.buckets[index]

	for _, node := range chain {
		if node.key == key {
			return &node, true
		}
	}
	return nil, false
}

func (h *HashMap) putValue(key string, value interface{}) bool {
	index := h.getIndex(key)
	if index == -1 {
		return false
	}
	chain := h.buckets[index]

	//found := false

	for i := range chain {
		node := &chain[i]
		if node.key == key {
			node.value = value
			return true
		}
	}

	if h.size == h.count {
		return false
	}

	node := Node{key: key, value: value}
	chain = append(chain, node)
	h.buckets[index] = chain
	h.count++
	return true
}

func createHashMap(expectedSize int) (*HashMap, error) {
	if expectedSize < 1 {
		return nil, errors.New("Size of the hashmap cannot be less than 1")
	}
	b := make([][]Node, expectedSize)
	for i := 0; i < expectedSize; i++ {
		b[i] = make([]Node, 0)
	}
	h := &HashMap{
		count:   0,
		size:    expectedSize,
		buckets: b,
	}
	return h, nil
}

func (h *HashMap) getCount() int {
	return h.count
}

func (h *HashMap) getsize() int {
	return h.size
}
