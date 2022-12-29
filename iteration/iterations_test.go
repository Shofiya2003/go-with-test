package iterations

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	got := Repeat("a", 7)
	want := "aaaaaaa"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleRepeat() {
	res := Repeat("s", 4)
	fmt.Println(res)
	// Output: ssss

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
