## sort inlining example

Here is the output of `run.sh`

```
+ go env
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/myitcv/gostuff"
GORACE=""
GOROOT="/home/myitcv/gos"
GOTOOLDIR="/home/myitcv/gos/pkg/tool/linux_amd64"
CC="gcc"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build555042588=/tmp/go-build -gno-record-gcc-switches"
CXX="g++"
CGO_ENABLED="1"
+ echo

+ go build -gcflags=-m
# github.com/myitcv/sort_inlining_example
./main.go:11: can inline dataLess
./main.go:25: can inline dataSwap
./main.go:164: inlining call to dataLess
./main.go:167: inlining call to dataLess
./main.go:170: inlining call to dataSwap
./main.go:187: inlining call to dataSwap
./main.go:198: inlining call to dataLess
./main.go:199: inlining call to dataSwap
./main.go:202: inlining call to dataLess
./main.go:203: inlining call to dataSwap
./main.go:205: inlining call to dataLess
./main.go:206: inlining call to dataSwap
./main.go:80: inlining call to dataLess
./main.go:84: inlining call to dataLess
./main.go:86: inlining call to dataLess
./main.go:92: inlining call to dataSwap
./main.go:102: inlining call to dataLess
./main.go:103: inlining call to dataSwap
./main.go:107: inlining call to dataLess
./main.go:114: inlining call to dataLess
./main.go:115: inlining call to dataSwap
./main.go:128: inlining call to dataLess
./main.go:130: inlining call to dataLess
./main.go:136: inlining call to dataSwap
./main.go:142: inlining call to dataSwap
./main.go:149: inlining call to dataLess
./main.go:150: inlining call to dataSwap
./main.go:51: inlining call to dataLess
./main.go:52: inlining call to dataSwap
./main.go:214: inlining call to dataSwap
./main.go:157: leaking param content: data
./main.go:175: leaking param content: data
./main.go:196: leaking param content: data
./main.go:59: leaking param content: data
./main.go:147: leaking param content: data
./main.go:29: leaking param content: data
./main.go:15: leaking param content: vs
./main.go:8: x escapes to heap
./main.go:6: []string literal escapes to heap
./main.go:8: main ... argument does not escape
./main.go:11: dataLess vs does not escape
./main.go:25: leaking param content: data
./main.go:212: leaking param content: data
+ echo

+ go test -bench BenchmarkSortString1K
testing: warning: no tests to run
BenchmarkSortString1K-8   	   10000	    104990 ns/op
PASS
ok  	github.com/myitcv/sort_inlining_example	1.801s
+ echo

+ go test -bench BenchmarkSortString1K sort
BenchmarkSortString1K-8   	   10000	    156416 ns/op
PASS
ok  	sort	2.939s
```
