package uuidv4

import (
	"strings"
	"testing"
	"fmt"
)

func TestUUIDBase(t *testing.T) {
	var id UUIDV4 = RandV4()
	fmt.Printf("uuid HexV4: %v\n", id.HexV4())
	fmt.Printf("uuid RawV4: %v\n", id.RawV4())
	id1, err := FromStrV4("1870747d-b26c-4507-9518-1ca62bc66e5d")
	if err != nil {
		t.Errorf("UUIDV4 base error %v", err)

	}
	id2 := MustFromStrV4("1870747db26c450795181ca62bc66e5d")
	fmt.Printf("uuid FromStrV4: %v\n", id1)
	fmt.Printf("uuid MustFromStrV4: %v\n", id2)
	if  id1 != id2{
		t.Errorf("uuid FromStrV4 %v | not MustFromStrV4 %v\n", id1, id2)
	}
}

// TestDuplicates check the first 1024 generated UUIDV4 for duplicates
func TestDuplicates(t *testing.T) {
	var id UUIDV4
	var hex string
	dupSMap := make(map[string]bool)
	for i := 0; i < 1024; i++ {
		id = RandV4()
		hex = id.HexV4()
		fmt.Printf("TestDuplicates id: %v, hex: %v\n", id, hex)
		if dupSMap[hex] {
			t.Errorf("Duplicates after %d iterations", i+1)
			t.FailNow()
		}
		dupSMap[hex] = true
	}
}

// TestFromStrSanity test FromStr(id.Hex()) == id
func TestFromStrSanity(t *testing.T) {
	var id, id2 UUIDV4
	for i := 0; i < 18; i++ {
		id = RandV4()
		id2 = MustFromStrV4(id.HexV4())
		fmt.Printf("TestFromStrSanity id: %v, id2: %v\n", id, id2)
		if id2 != id {
			t.Errorf("Sanity check fail for UUIDV4 string %s\n\tid:  %v\n\tid2: %v", id.HexV4(), id, id2)
			t.FailNow()
		}
	}
}

// TestHex does a simple test to make sure Hex string returns the elements in the right position
func TestHex(t *testing.T) {
	x := [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s := strings.ToLower(UUIDV4(x).HexV4())
	fmt.Printf("TestHex hex: %v, uuidStr: %v\n", x, s)
	if s != "00010203-0405-0607-0809-0a0b0c0d0e0f" {
		t.Errorf("Hex fail:\n\tBinary: %v,\n\tBad hex: %s", x, s)
		t.FailNow()
	}
}
