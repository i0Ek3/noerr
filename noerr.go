// noerr support the simple way to get rid of obvious judge like
// if err != nil;
package noerr

import (
    "log"
)

// noerr return two value, one of is interface{} type v and 
// another one is error type e, it's returns indicate the 
// needs of error checking
func noerr(err error) (v interface{}, e error) {
    if err != nil {
        return nil, err
    }
    return v, e
}

// NoErr call the function noerr to return the value
func NoErr(err error) (v interface{}, e error) {
    return noerr(err)
}

func NoErrPc(err error) {
    if err != nil {
        panic(err)
    }
}

// NoErrP indicats the standard error showing then print
func NoErrP(err error) {
    if err != nil {
        log.Print(err)
    }
}

func NoErrln(err error) {
    if err != nil {
        log.Println(err)
    }
}

// NoErrF indicats the standard error showing then break
func NoErrF(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

// oked checks if ok is true or false
func oked(ok bool) bool {
    if ok {
        return true
    }
    return false
}

// Oked is the public function to calling to oked
func Oked(ok bool) bool {
    return oked(ok)
}
