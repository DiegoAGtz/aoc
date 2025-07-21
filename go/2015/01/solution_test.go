package solutions_2015_01

import "testing"

func TestPartOne(t *testing.T) {
    t.Run("going to third floor", func(t *testing.T) {
        got := partOne("(((")
        want := 3
        assertCorrectMessage(t, got, want)
    })
    t.Run("going to basement", func(t *testing.T) {
        got := partOne("((())))")
        want := -1
        assertCorrectMessage(t, got, want)
    })
    t.Run("going to fifht floor", func(t *testing.T) {
        got := partOne("(()))()(())()()(())()()()())((((()))((()(((()")
        want := 5
        assertCorrectMessage(t, got, want)
    })
}

func TestPartTwo(t *testing.T) {
    t.Run("going to basement after one move", func(t *testing.T) {
        got := partTwo(")")
        want := 1
        assertCorrectMessage(t, got, want)
    })
    t.Run("going to basement after seven moves", func(t *testing.T) {
        got := partTwo("(())())()()()))()(((()))()))()")
        want := 7
        assertCorrectMessage(t, got, want)
    })
    t.Run("never pass through basement", func(t *testing.T) {
        got := partTwo("(((())(()())())")
        want := -1
        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want any) {
    t.Helper()
    if got != want {
        t.Errorf("got \"%v\" want \"%v\"", got, want)
    }
}
