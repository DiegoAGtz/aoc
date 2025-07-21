package solutions_2015_02

import "testing"

func TestPartOne(t *testing.T) {
	t.Run("lader dimension", func(t *testing.T) {
		got := partOne([]string{"2x3x4"})
		want := 58
		assertCorrectMessage(t, got, want)
	})
	t.Run("two equal sides", func(t *testing.T) {
		got := partOne([]string{"1x1x10"})
		want := 43
		assertCorrectMessage(t, got, want)
	})
}

func TestPartTwo(t *testing.T) {
	t.Run("lader dimension", func(t *testing.T) {
		got := partTwo([]string{"2x3x4"})
		want := 34
		assertCorrectMessage(t, got, want)
	})
	t.Run("two equal sides", func(t *testing.T) {
		got := partTwo([]string{"1x1x10"})
		want := 14
		assertCorrectMessage(t, got, want)
	})
}

func TestArea(t *testing.T) {
	t.Run("quadratic area", func(t *testing.T) {
		got := area(4, 4)
		want := 16
		assertCorrectMessage(t, got, want)
	})
	t.Run("rectangle area", func(t *testing.T) {
		got := area(4, 6)
		want := 24
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty area", func(t *testing.T) {
		got := area(0, 6)
		want := 0
		assertCorrectMessage(t, got, want)
	})
	t.Run("negative area", func(t *testing.T) {
		got := area(-2, 6)
		want := -12
		assertCorrectMessage(t, got, want)
	})
}

func TestVolume(t *testing.T) {
	t.Run("lader volume", func (t *testing.T) {
		got := volume(2, 3, 4)
		want := 24
		assertCorrectMessage(t, got, want)
	})
	t.Run("two equal sides volume", func (t *testing.T) {
		got := volume(1, 1, 10)
		want := 10
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got \"%v\" want \"%v\"", got, want)
	}
}
