package noerr

import (

)

func noerr(err error) (v interface{}, e error) {
    if err != nil {
        return v, e
    }
    return nil, nil
}

func NoErr(err error) (v interface{}, e error) {
    return noerr(err)
}
