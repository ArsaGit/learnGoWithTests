package iteration

import "testing"

func TestMyToUpper(t *testing.T) {
	got := MyToUpper("qwerty")
	want := "QWERTY"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
