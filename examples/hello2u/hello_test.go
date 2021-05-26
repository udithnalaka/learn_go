package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("sayine hello to people", func(t *testing.T) {
		got := Hello("Udith", "English")
		want := "Hello, Udith"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying 'Hello, World' when an empty string passed", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("hello in Spanish", func(t *testing.T) {
		got := Hello("Nalaka", "Spanish")
		want := "Hola, Nalaka"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("hello in French", func(t *testing.T) {
		got := Hello("Udith", "French")
		want := "Bonjour, Udith"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("hello in Sinhala", func(t *testing.T) {
		got := Hello("Udith", "Sinhala")
		want := "Aayubowan, Udith"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
