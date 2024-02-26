package main

import (
    "testing"
    "regexp"
)

func TestFly(t *testing.T) {
    string := "Fly!"
    want := regexp.MustCompile(string)
    msg := fly_away()
    if !want.MatchString(msg) {
        t.Fatalf(`fly_away() = %q, want match for %#q`, msg, want)
    }
}
