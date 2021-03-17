# nakedretrun

[![pkg.go.dev][gopkg-badge]][gopkg]

`nakedretrun` finds naked returns.

```go
package a

func f() (n int) {
        return // want "should not use naked return"
}
```

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gostaticanalysis/nakedretrun
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gostaticanalysis/nakedretrun?status.svg

