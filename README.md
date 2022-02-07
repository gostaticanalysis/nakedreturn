# nakedreturn

[![pkg.go.dev][gopkg-badge]][gopkg]

`nakedreturn` finds naked returns.

```go
package a

func f() (n int) {
        return // want "should not use naked return"
}
```

`fixnakedreturn` command check and fix (remove) naked return.

```sh
$ go install github.com/gostaticanalysis/nakedreturn/cmd/fixnakedreturn
$ fixnakedreturn ./...
```
<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gostaticanalysis/nakedreturn
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gostaticanalysis/nakedreturn?status.svg

