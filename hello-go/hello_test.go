package main

import "testing"

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper() //if this helper reports a failure, Go points to the calling test instead of this function
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Frank", "French")
		want := "Bonjour, Frank"
		assertCorrectMessage(t, got, want)
	})
}
