package noerr

import (
    "testing"
    "errors"
)

var e = errors.New("Wrong here.")

func TestNoErr(t *testing.T){
    _, got := NoErr(nil)
    _, want := NoErr(e)

    if got != want {
        t.Errorf("got %v != want %v", got, want)
    }
}

func TestNoErrP(t *testing.T) {
    NoErrP(e)
}

func TestNoErrln(t *testing.T) {
    NoErrln(e)
}

func TestNoErrF(t *testing.T) {
    NoErrF(e)
}

func TestNoErrPc(t *testing.T) {
    NoErrPc(e)
}

func TestOked(t *testing.T) {
    ok := true
    no := false

    got := Oked(ok)
    want := Oked(!no)

    if got != want {
        t.Errorf("got %v != want %v", got, want)
    }
}
