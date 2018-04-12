package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		assertMessage(t, Hello("Weylin", ""), "Hello, Weylin")
	})
	t.Run("say Hello, World when empty string is supplied", func(t *testing.T) {
		assertMessage(t, Hello("", ""), "Hello, World")
	})
	t.Run("say Hello in Spanish", func(t *testing.T) {
		assertMessage(t, Hello("Weylin", "Spanish"), "Hola, Weylin")
	})
	t.Run("say Hello in French", func(t *testing.T) {
		assertMessage(t, Hello("Weylin", "French"), "Bonjour, Weylin")
	})
}
