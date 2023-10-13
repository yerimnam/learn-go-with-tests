package main

import "testing"

func TestHello(t *testing.T) {
	// todo -> 각 나라별 테스트 케이스 짜기
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Eldie", "Spanish")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
