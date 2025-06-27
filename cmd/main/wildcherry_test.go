package main

import "testing"

func Example_wildcherry() {
	goMain([]string{"wildcherry", "../../testdata/LICENSE"})
	// Output:
	//      21     169    1071 ../../testdata/LICENSE
}

func Example_url() {
	goMain([]string{"wildcherry", "https://raw.githubusercontent.com/tamada/wildcherry/refs/heads/main/LICENSE"})
	// Output:
	//      21     169    1071 https://raw.githubusercontent.com/tamada/wildcherry/refs/heads/main/LICENSE
}

func TestHello(t *testing.T) {
	got := hello()
	want := "Welcome to WildCherry!"
	if got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}
