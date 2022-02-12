package sample

import "testing"

func TestSample(t *testing.T) {
    want := "This is a great sample!"
    if got := Sample(); got != want {
        t.Errorf("Sample() = %q, want %q", got, want)
    }
}
