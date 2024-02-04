package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main()  {
	fmt.Println("hello mod in golang")
	greeter()

	r := mux.NewRouter()
    r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
	
}


func greeter()  {
	fmt.Println("hey there, MOD users")
}

func serveHome(w http.ResponseWriter, r *http.Request)   {
	w.Write([]byte("<h1>Welcome to GO series</h1>"))
}



//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go mod init github.com/ayush844/mymodules
//go: creating new go.mod: module github.com/ayush844/mymodules
//go: to add module requirements and sums:
//        go mod tidy
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go get -u github.com/gorilla/mux
//go: downloading github.com/gorilla/mux v1.8.1
//go: added github.com/gorilla/mux v1.8.1
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go env
//GO111MODULE=''
//GOARCH='amd64'
//GOBIN=''
//GOCACHE='/home/ayush/.cache/go-build'
//GOENV='/home/ayush/.config/go/env'
//GOEXE=''
//GOEXPERIMENT=''
//GOFLAGS=''
//GOHOSTARCH='amd64'
//GOHOSTOS='linux'
//GOINSECURE=''
//GOMODCACHE='/home/ayush/go/pkg/mod'
//GONOPROXY=''
//GONOSUMDB=''
//GOOS='linux'
//GOPATH='/home/ayush/go'
//GOPRIVATE=''
//GOPROXY='https://proxy.golang.org,direct'
//GOROOT='/snap/go/10489'
//GOSUMDB='sum.golang.org'
//GOTMPDIR=''
//GOTOOLCHAIN='auto'
//GOTOOLDIR='/snap/go/10489/pkg/tool/linux_amd64'
//GOVCS=''
//GOVERSION='go1.21.6'
//GCCGO='gccgo'
//GOAMD64='v1'
//AR='ar'
//CC='gcc'
//CXX='g++'
//CGO_ENABLED='1'
//GOMOD='/home/ayush/Desktop/goWithGolang/22myModules/go.mod'
//GOWORK='/home/ayush/Desktop/goWithGolang/go.work'
//CGO_CFLAGS='-O2 -g'
//CGO_CPPFLAGS=''
//CGO_CXXFLAGS='-O2 -g'
//CGO_FFLAGS='-O2 -g'
//CGO_LDFLAGS='-O2 -g'
//PKG_CONFIG='pkg-config'
//GOGCCFLAGS='-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build1558192965=/tmp///go-build -gno-record-gcc-switches'
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go build .
//current directory is contained in a module that is not one of the workspace modules listed in go.work. You can add the module //to the workspace using:
//        go work use .
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go build .
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go run main.go
//hello mod in golang
//hey there, MOD users
//^Csignal: interrupt
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go mod tidy
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go mod verify
//variables : missing ziphash: open hash: no such file or directory
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go list
//github.com/ayush844/mymodules
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go list all
//arrays
//bufio
//bytes
//compress/flate
//compress/gzip
//container/list
//context
//conversions
//crypto
//crypto/aes
//crypto/cipher
//crypto/des
//crypto/dsa
//crypto/ecdh
//crypto/ecdsa
//crypto/ed25519
//crypto/elliptic
//crypto/hmac
//crypto/internal/alias
//crypto/internal/bigmod
//crypto/internal/boring
//crypto/internal/boring/bbig
//crypto/internal/boring/sig
//crypto/internal/edwards25519
//crypto/internal/edwards25519/field
//crypto/internal/nistec
//crypto/internal/nistec/fiat
//crypto/internal/randutil
//crypto/md5
//crypto/rand
//crypto/rc4
//crypto/rsa
//crypto/sha1
//crypto/sha256
//crypto/sha512
//crypto/subtle
//crypto/tls
//crypto/x509
//crypto/x509/pkix
//defer
//embed
//encoding
//encoding/asn1
//encoding/base64
//encoding/binary
//encoding/hex
//encoding/json
//encoding/pem
//errors
//files
//fmt
//functions
//getRequests
//github.com/ayush844/mymodules
//github.com/gorilla/mux
//hash
//hash/crc32
//hello
//if_else_statement
//internal/abi
//internal/bisect
//internal/bytealg
//internal/coverage/rtcov
//internal/cpu
//internal/fmtsort
//internal/goarch
//internal/godebug
//internal/godebugs
//internal/goexperiment
//internal/goos
//internal/intern
//internal/itoa
//internal/nettrace
//internal/oserror
//internal/poll
//internal/race
//internal/reflectlite
//internal/safefilepath
//internal/singleflight
//internal/syscall/execenv
//internal/syscall/unix
//internal/testlog
//internal/unsafeheader
//io
//io/fs
//io/ioutil
//log
//log/internal
//loops
//mapping
//math
//math/big
//math/bits
//math/rand
//methods
//mime
//mime/multipart
//mime/quotedprintable
//myjson
//net
//net/http
//net/http/httptrace
//net/http/internal
//net/http/internal/ascii
//net/netip
//net/textproto
//net/url
//os
//path
//path/filepath
//pointers
//reflect
//regexp
//regexp/syntax
//requests
//runtime
//runtime/internal/atomic
//runtime/internal/math
//runtime/internal/sys
//runtime/internal/syscall
//slice
//sort
//strconv
//strings
//structs
//switchcase
//sync
//sync/atomic
//syscall
//time
//timemanagement
//unicode
//unicode/utf16
//unicode/utf8
//unsafe
//url
//userinput
//variables
//vendor/golang.org/x/crypto/chacha20
//vendor/golang.org/x/crypto/chacha20poly1305
//vendor/golang.org/x/crypto/cryptobyte
//vendor/golang.org/x/crypto/cryptobyte/asn1
//vendor/golang.org/x/crypto/hkdf
//vendor/golang.org/x/crypto/internal/alias
//vendor/golang.org/x/crypto/internal/poly1305
//vendor/golang.org/x/net/dns/dnsmessage
//vendor/golang.org/x/net/http/httpguts
//vendor/golang.org/x/net/http/httpproxy
//vendor/golang.org/x/net/http2/hpack
//vendor/golang.org/x/net/idna
//vendor/golang.org/x/sys/cpu
//vendor/golang.org/x/text/secure/bidirule
//vendor/golang.org/x/text/transform
//vendor/golang.org/x/text/unicode/bidi
//vendor/golang.org/x/text/unicode/norm
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go list -m all
//arrays
//conversions
//defer
//files
//functions
//getRequests
//github.com/ayush844/mymodules
//hello
//if_else_statement
//loops
//mapping
//methods
//myjson
//pointers
//requests
//slice
//structs
//switchcase
//timemanagement
//url
//userinput
//variables
//github.com/gorilla/mux v1.8.1
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go list -m -versions github.com/gorilla/mux
//github.com/gorilla/mux v1.2.0 v1.3.0 v1.4.0 v1.5.0 v1.6.0 v1.6.1 v1.6.2 v1.7.0 v1.7.1 v1.7.2 v1.7.3 v1.7.4 v1.8.0 v1.8.1
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go mod tidy
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go mod why github.com/gorilla/mux
//# github.com/gorilla/mux
//github.com/ayush844/mymodules
//github.com/gorilla/mux
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go mod graph
//arrays go@1.21.6
//conversions go@1.21.6
//defer go@1.21.6
//files go@1.21.6
//functions go@1.21.6
//getRequests go@1.21.6
//github.com/ayush844/mymodules github.com/gorilla/mux@v1.8.1
//github.com/ayush844/mymodules go@1.21.6
//go@1.21.6 toolchain@go1.21.6
//hello go@1.21.5
//if_else_statement go@1.21.6
//loops go@1.21.6
//mapping go@1.21.6
//methods go@1.21.6
//myjson go@1.21.6
//pointers go@1.21.6
//requests go@1.21.6
//slice go@1.21.6
//structs go@1.21.6
//switchcase go@1.21.6
//timemanagement go@1.21.6
//url go@1.21.6
//userinput go@1.21.6
//variables go@1.21.5
//go@1.21.5 toolchain@go1.21.5
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go mod vendor
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ go run -mod=vendor main.go
//go: -mod may only be set to readonly when in workspace mode, but it is set to "vendor"
//        Remove the -mod flag to use the default readonly value,
//        or set GOWORK=off to disable workspace mode.
//ayush@lucas:~/Desktop/goWithGolang/22myModules$ 