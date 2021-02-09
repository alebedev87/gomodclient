# How to use Go module from GitHub

## Do not use the public proxy
If your git repository is private (i.e. corporate), no need to query the public proxy:
```bash
$ go env -w GOPRIVATE="mycorp.net/*"
```

## Build and run the example
```bash
$ go build main.go
go: finding module for package github.com/alebedev87/gomodhello
go: downloading github.com/alebedev87/gomodhello v0.0.3
go: found github.com/alebedev87/gomodhello in github.com/alebedev87/gomodhello v0.0.3
$ ./main
Hello, world.
Concurrency is not parallelism.
Hello from Rostov!
Tag check
```

## Use not correctly tagged modules
```bash
# get not correctly tagged version
$ go get github.com/alebedev87/gomodhello@0.0.4
 go: github.com/alebedev87/gomodhello 0.0.4 => v0.0.4-0.20200727134623-bb98021fc08a
# remove the comment to use the method from the version 0.0.4
$ sed -i 's|//||' main.go
$ go build main.go
# note the last line: "Can be w/o v"
$ ./main
Hello, world.
Concurrency is not parallelism.
Hello from Rostov!
Tag check
Can be w/o v
```
