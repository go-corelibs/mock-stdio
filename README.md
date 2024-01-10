[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/mock-stdio)
[![codecov](https://codecov.io/gh/go-corelibs/mock-stdio/graph/badge.svg?token=NzJRec8e5Q)](https://codecov.io/gh/go-corelibs/mock-stdio)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/mock-stdio)](https://goreportcard.com/report/github.com/go-corelibs/mock-stdio)

# mock-stdio - utilities for mocking os.Stdin, os.Stdout and os.Stderr

# Installation

``` shell
> go get github.com/go-corelibs/mock-stdio@latest
```

# Examples

## NewStdin

``` go
func Test(t *testing.T) {
    inio := stdio.NewStdin("testing!\n")
    err := inio.Capture()
    So(err, ShouldEqual, nil)
    var data []byte
    data, err = io.ReadAll(os.Stdin)
    So(err, ShouldEqual, nil)
    So(data, ShouldEqual, []byte("testing!\n"))
}
```

## NewStdout

``` go
func Test(t *testing.T) {
    outio := stdio.NewStdout()
    err := outio.Capture()
    So(err, ShouldEqual, nil)
    _, err = os.Stdout.WriteString("testing!\n")
    So(err, ShouldEqual, nil)
    So(outio.Data(), ShouldEqual, []byte("testing!\n"))
}
```

## NewStderr

``` go
func Test(t *testing.T) {
    errio := stdio.NewStderr()
    err := errio.Capture()
    So(err, ShouldEqual, nil)
    _, err = os.Stderr.WriteString("testing!\n")
    So(err, ShouldEqual, nil)
    So(errio.Data(), ShouldEqual, []byte("testing!\n"))
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2024 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
