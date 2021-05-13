package noerr

import (
    "testing"
)

func TestNoErr(t *testing.T){
    var e error
    
    _, got := NoErr(nil)
    _, want := NoErr(e) 

    if got != want {
        t.Errorf("got %v != want %v", got, want)
    }
}
