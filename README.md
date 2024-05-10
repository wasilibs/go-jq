# go-jq

go-jq is a distribution of [jq][1], that can be built with Go. It does not actually reimplement any
functionality of jq in Go, instead compiling the original source to WebAssembly, and 
executing with the pure Go Wasm runtime [wazero][2]. This means that `go install` or `go run`
can be used to execute it, with no need to rely on separate package managers such as homebrew,
on any platform that Go supports.

Note that there is an excellent Go port of jq, [gojq][3]. It has some small differences from upstream,
so this version may be appropriate where those come into play. It also intends to be a case study
on bringing tools to Go without a significant rewrite.

## Installation

Precompiled binaries are available in the [releases](https://github.com/wasilibs/go-jq/releases).
Alternatively, install the plugin you want using `go install`.

```bash
$ go install github.com/wasilibs/go-jq/cmd/jq@latest
```

To avoid installation entirely, it can be convenient to use `go run`

```bash
$ echo '{"foo": 0}' | go run github.com/wasilibs/go-jq/cmd/jq@latest
```

Note that jq is generally used for very small operations, so the overhead of `go run` will often be
noticable.

Note that due to the sandboxing of the filesystem when using Wasm, currently only files that descend
from the current directory when executing the tool are accessible to it, i.e., `../sql/my.json` or
`/separate/root/my.json` will not be found.

[1]: https://github.com/jqlang/jq
[2]: https://wazero.io/
[3]: https://github.com/itchyny/gojq
