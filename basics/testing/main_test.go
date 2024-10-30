package testing

import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello World"
	got := sayHello()

	if got != want {
		t.Errorf("Wanted %s, got %s", want, got)
	}
}
