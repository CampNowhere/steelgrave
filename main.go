package main

// #cgo CFLAGS: -I/usr/include/php/20151012
// #cgo CFLAGS: -I/usr/include/php/20151012/Zend
// #cgo CFLAGS: -I/usr/include/php/20151012/TSRM
// #cgo CFLAGS: -I/usr/include/php/20151012/sapi
// #cgo CFLAGS: -I/usr/include/php/20151012/main
// #cgo LDFLAGS: -lphp7
// #include <stdlib.h>
// #include "render_php.h"
import "C"
import (
	"fmt"
	"os"
	"time"
	"unsafe"
)

func main() {
	cs := C.CString("test.php")
	defer C.free(unsafe.Pointer(cs))
	fmt.Println("Steelgrave Web Server starting with PID", os.Getpid())
	stdoutOrig := os.Stdout
	os.Stdout, _ = os.OpenFile("/dev/null", os.O_WRONLY, 0777)
	C.render_php(cs)
	os.Stdout.Close()
	os.Stdout = stdoutOrig
	wait, _ := time.ParseDuration("30s")
	time.Sleep(wait)
}
