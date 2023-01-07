package main

import (
	"fmt"

	"github.com/thanhfphan/bitset/bigbitset"
)

var (
	friendIDMap = map[string]int{
		"alice": 0,
		"bob":   1,
		"long":  2,
		"gate":  3,
		"thanh": 4,
	}
)

func main() {
	alice := bigbitset.New(friendIDMap["alice"], friendIDMap["long"], friendIDMap["thanh"]) // 10101
	bob := bigbitset.New(friendIDMap["bob"], friendIDMap["thanh"])                          // 10010

	fmt.Printf("BitSet Alice: %s\n", alice.String())
	fmt.Printf("BitSet Bob: %s\n", bob.String())
	fmt.Printf("Is Alice friend with Bob: %t\n", alice.Contains(friendIDMap["bob"]))     // False
	fmt.Printf("Is Alice friend with Thanh: %t\n", alice.Contains(friendIDMap["thanh"])) // True
	fmt.Printf("Is Bob friend with Thanh: %t\n", bob.Contains(friendIDMap["thanh"]))     // True
	fmt.Printf("Is Bob friend with Alice: %t\n", bob.Contains(friendIDMap["alice"]))     // False

}
