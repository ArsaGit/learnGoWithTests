package iteration

import (
	"fmt"
	"testing"
)

func TestMyToUpper(t *testing.T) {
	got := MyToUpper("qwerty")
	want := "QWERTY"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func ExampleMyToUpper() {
	input := "qwerty"
	output := MyToUpper(input)
	fmt.Println(output)
	// Output: QWERTY
}