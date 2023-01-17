package integer

import (
	"fmt"
	"testing"
)

func TestInteger(t *testing.T) {
	got := Add(2, 3)
	expected := 5
	if got != expected {
		t.Errorf("got '%d', expected '%d'", got, expected)
	}
}

func ExampleAdd() {
	sum := Add(4, 7)
	fmt.Println(sum)
	// Output: 11
}
