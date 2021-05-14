// noerr support the simple way to get rid of obvious judge like
// if err != nil;
package noerr

// noerr return two value, one of is interface{} type v and 
// another is error type e, indicates error value.
func noerr(err error) (v interface{}, e error) {
    if err != nil {
        return v, e
    }
    return nil, nil
}

// NoErr call the function noerr to return the value, which is private
// and unable to calling obviously.
func NoErr(err error) (v interface{}, e error) {
    return noerr(err)
}

// oked checks if ok is true or false
func oked(ok bool) bool {
    if ok {
        return true
    }
    return false
}
