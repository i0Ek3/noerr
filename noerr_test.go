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
    
    //NoErrf(e)
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
