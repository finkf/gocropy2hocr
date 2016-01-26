# gocropy2Hocr
Convert [ocropy](https://github.com/tmbdev/ocropy) hocr files into a
format that is more suitable for
[PoCoTo](https://github.com/cisocrgroup/PoCoTo).

# Installation and usage
* install `go`
* set `GOPATH` environment variable `export GOPATH=some/go/path`
* setup go environment `mdkir $GOPATH && cd $GOPATH`
* install package `go get github.com/finkf/gocropy2hocr`
* compile and install `go install github.com/finkf/gocropy2hocr`
* execute test run `bin/gocropy2hocr
  src/github.com/finkf/gocropy2hocr/data/ersch.html
  src/github.com/finkf/gocropy2hocr/data/book/0001`
