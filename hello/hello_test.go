package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hellow to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("say, 'Hello, World' when an empty string is suppleid", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
