package noerr

import (
    "testing"
    "errors"
)

func TestNoErr(t *testing.T){
    e := errors.New("Wrong here.") 

    _, got := NoErr(nil)
    _, want := NoErr(e) 

    if got != want {
        t.Errorf("got %v != want %v", got, want)
    }
}

func TestOked(t *testing.T) {
    ok := true
    no := false

    got := oked(ok)
    want := oked(!no)

    if got != want {
        t.Errorf("got %v != want %v", got, want)
    }
}
