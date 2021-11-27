# noerr

> You will never see &#34;if err != nil&#34; in the go file again!

## Getting Started

### Install

```sh
go get https://github.com/i0Ek3/noerr
```

### Usage

```Go
package main

import "github.com/i0Ek3/noerr"

func main() {
    // show err then break
    noerr.Xerr(err, msg...)
    noerr.NoErr(err)
    noerr.NoErrP(err)
    noerr.NoErrF(err)
    noerr.NoErrPc(err)
    noerr.NoErrln(err)
}
```

## TODO

- [x] support ok checking
- [ ] refactor


## Contributing

Pull requests and Issue are also welcome.

## Show your support

Give a ⭐️ if this project helped you!

## License

MIT.
