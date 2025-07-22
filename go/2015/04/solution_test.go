package solutions_2015_04

import "testing"

func TestPartOne(t *testing.T) {
    t.Run("hash for abcdef", func(t *testing.T) {
        got := PartOne("abcdef")
        want := 609043
        assertCorrectMessage(t, got, want)
    })
    t.Run("hash for pqrstuv", func(t *testing.T) {
        got := PartOne("pqrstuv")
        want := 1048970
        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want any) {
    t.Helper()
    if got != want {
        t.Errorf("got \"%v\" want \"%v\"", got, want)
    }
}
