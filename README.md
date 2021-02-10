# spinlock

[![GoDoc](https://img.shields.io/badge/api-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/tidwall/spinlock)


A [spinlock](https://en.wikipedia.org/wiki/Spinlock) implementation for Go.

It shares the same interface as [sync.Mutex](https://golang.org/pkg/sync/#Mutex), and is intended to be used to synchronize exceptionally short lived operations.

## Install

```
go get -u github.com/tidwall/spinlock
```

## Contact

Josh Baker [@tidwall](http://twitter.com/tidwall)

## License

spinlock source code is available under the MIT [License](/LICENSE).
