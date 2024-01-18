package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")

	presentTime := time.Now()
	fmt.Println(presentTime)
	// 2024-01-17 16:21:39.486733323 +0530 IST m=+0.000052713

	// Basic full date  "Mon Jan 2 15:04:05 MST 2006"

	fmt.Println(presentTime.Format("02-01-2006"))
	// 17-01-2024   => 17 jan 2006

	fmt.Println(presentTime.Format("15:04:05"))
	// 16:29:54   => 16hr 29min 54sec

	fmt.Println(presentTime.Format("Monday"))
	// Wednesday

	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04:05"))
	// 17-01-2024 Wednesday 16:31:43

	// creating date //
	createdDate := time.Date(2004, time.April, 8, 23, 30, 0, 0, time.Local)

	fmt.Println(createdDate)
	// 2004-04-08 23:30:00 +0530 IST

	fmt.Println(createdDate.Format("02-01-2006 Monday 15:04:05"))
	// 08-04-2004 Thursday 23:30:00

}

//************************************************************************************************
//******************************* GO BUILD *******************************************************
//************************************************************************************************

//ayush@lucas:~/Desktop/goWithGolang/05mytime$ go env

// GO111MODULE=''
// GOARCH='amd64'
// GOBIN=''
// GOCACHE='/home/ayush/.cache/go-build'
// GOENV='/home/ayush/.config/go/env'
// GOEXE=''
// GOEXPERIMENT=''
// GOFLAGS=''
// GOHOSTARCH='amd64'
// GOHOSTOS='linux'
// GOINSECURE=''
// GOMODCACHE='/home/ayush/go/pkg/mod'
// GONOPROXY=''
// GONOSUMDB=''
//**********************************
// GOOS='linux'
//**********************************
// GOPATH='/home/ayush/go'
// GOPRIVATE=''
// GOPROXY='https://proxy.golang.org,direct'
// GOROOT='/snap/go/10489'
// GOSUMDB='sum.golang.org'
// GOTMPDIR=''
// GOTOOLCHAIN='auto'
// GOTOOLDIR='/snap/go/10489/pkg/tool/linux_amd64'
// GOVCS=''
// GOVERSION='go1.21.6'
// GCCGO='gccgo'
// GOAMD64='v1'
// AR='ar'
// CC='gcc'
// CXX='g++'
// CGO_ENABLED='1'
// GOMOD='/home/ayush/Desktop/goWithGolang/05mytime/go.mod'
// GOWORK=''
// CGO_CFLAGS='-O2 -g'
// CGO_CPPFLAGS=''
// CGO_CXXFLAGS='-O2 -g'
// CGO_FFLAGS='-O2 -g'
// CGO_LDFLAGS='-O2 -g'
// PKG_CONFIG='pkg-config'
// GOGCCFLAGS='-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -ffile-prefix-map=/tmp/go-build3225025562=/tmp/go-build -gno-record-gcc-switches'

//************************************************************************************************
// ayush@lucas:~/Desktop/goWithGolang/05mytime$ GOOS="windows" go build
// Compiles a Go program specifically for the Windows operating system, even if you're currently running on a different OS like Linux or macOS.

// ayush@lucas:~/Desktop/goWithGolang/05mytime$ GOOS="darwin" go build
//Compiles a Go program specifically for Apple macOS (formerly known as Darwin), even if you're not currently on a macOS machine.

//both of the code above generates an executrable file

//************************************************************************************************

//*********************************************************************************************************
//******************************* MEMORY MANAGEMENT *******************************************************
//*********************************************************************************************************

// > memory allocation and deallocation happens automatically in GOLANG

// ~ new():
// > allocates memory but does not initialize it
// > you will get a memory address
// > zeroed storage (we can not put data initially)

// ~ make():
// > allocates memory and also initialize it
// > you will get a memory address
// > non-zeroed storage (we can put data initially)

// garbage collection(GC) happens automatically, anything which is out of scope or is nil becomes eligible for garbage collection

//The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. The default is GOGC=100. Setting GOGC=off disables the garbage collector entirely.

//func NumCPU
//func NumCPU() int
//NumCPU returns the number of logical CPUs usable by the current process.
//The set of available CPUs is checked by querying the operating system at process startup. Changes to operating system CPU allocation after process startup are not reflected.
