package uuidv4

import (
	"fmt"
	"strings"
	"testing"
)

// BenchmarkDuplicates check the first 1024 generated UUIDV4 for duplicates
func BenchmarkDuplicates(b *testing.B) {
	var id UUIDV4
	var hex string

	dupSMap := make(map[string]bool)
	for i := 0; i < 1024; i++ {
		id = RandV4()
		hex = id.HexV4()
		fmt.Printf("BenchmarkDuplicates id: %v, hex: %v\n", id, hex)
		if dupSMap[hex] {
			b.Errorf("Duplicates after %d iterations", i+1)
			b.FailNow()
		}
		dupSMap[hex] = true
	}
}

// BenchmarkFromStrSanity test FromStr(id.Hex()) == id
func BenchmarkFromStrSanity(b *testing.B) {
	var id, id2 UUIDV4
	for i := 0; i < 18; i++ {
		id = RandV4()
		id2 = MustFromStrV4(id.HexV4())
		fmt.Printf("BenchmarkFromStrSanity id: %v, id2: %v\n", id, id2)
		if id2 != id {
			b.Errorf("Sanity check fail for UUIDV4 string %s\n\tid:  %v\n\tid2: %v", id.HexV4(), id, id2)
			b.FailNow()
		}
	}
}

// BenchmarkHex does a simple test to make sure Hex string returns the elements in the right position
func BenchmarkHex(b *testing.B) {
	x := [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s := strings.ToLower(UUIDV4(x).HexV4())
	fmt.Printf("BenchmarkHex hex: %v, uuidStr: %v\n", x, s)
	if s != "00010203-0405-0607-0809-0a0b0c0d0e0f" {
		b.Errorf("Hex fail:\n\tBinary: %v,\n\tBad hex: %s", x, s)
		b.FailNow()
	}
}
