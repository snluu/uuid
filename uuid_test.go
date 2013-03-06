package uuid

import (
	"strings"
	"testing"
)

// TestDuplicates check the first 1024 generated UUID for duplicates
func TestDuplicates(t *testing.T) {
	var id UUID
	var hex string

	dups := make(map[string]bool)
	for i := 0; i < 1024; i++ {
		id = Rand()
		hex = id.Hex()
		if dups[hex] {
			t.Errorf("Duplicates after %d iterations", i+1)
			t.FailNow()
		}
		dups[hex] = true
	}
}

// TestFromStrSanity test FromStr(id.Hex()) == id
func TestFromStrSanity(t *testing.T) {
	var id, id2 UUID
	for i := 0; i < 18; i++ {
		id = Rand()
		id2 = MustFromStr(id.Hex())
		if id2 != id {
			t.Errorf("Sanity check fail for UUID string %s\n\tid:  %v\n\tid2: %v", id.Hex(), id, id2)
			t.FailNow()
		}
	}
}

// TestHex does a simple test to make sure Hex string returns the elements in the right position
func TestHex(t *testing.T) {
	x := [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	s := strings.ToLower(UUID(x).Hex())
	if s != "01020304-0506-0708-090a-0b0c0d0e0f10" {
		t.Errorf("Hex fail:\n\tBinary: %v,\n\tBad hex: %s", x, s)
		t.FailNow()
	}
}
