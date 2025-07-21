package solutions_2015_03

import "testing"

func TestPartOne(t *testing.T) {
	t.Run("Delivers presents to 2 houses", func(t *testing.T) {
		got := PartOne(">")
		want := 2
		assertCorrectMessage(t, got, want)
	})
	t.Run("Delivers presents to 4 houses", func(t *testing.T) {
		got := PartOne("^>v<")
		want := 4
		assertCorrectMessage(t, got, want)
	})
	t.Run("Delivers presents to 2 houses", func(t *testing.T) {
		got := PartOne("^v^v^v^v^v")
		want := 2
		assertCorrectMessage(t, got, want)
	})
}

func TestPartTwo(t *testing.T) {
	t.Run("Delivers presents to 3 houses", func(t *testing.T) {
		got := PartTwo("^v")
		want := 3
		assertCorrectMessage(t, got, want)
	})
	t.Run("Delivers presents to 3 houses", func(t *testing.T) {
		got := PartTwo("^>v<")
		want := 3
		assertCorrectMessage(t, got, want)
	})
	t.Run("Delivers presents to 2 houses", func(t *testing.T) {
		got := PartTwo("^v^v^v^v^v")
		want := 11
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got \"%v\" want \"%v\"", got, want)
	}
}
