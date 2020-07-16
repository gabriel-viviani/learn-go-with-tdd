package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Vivi", "")
		want := "Hello, Vivi"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when string is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Julia", "es")
		want := "Hola, Julia"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Je", "fr")
		want := "Bonjour, Je"
		assertCorrectMessage(t, got, want)
	})
}
