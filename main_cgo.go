package main

import (
	"fmt"
)

/*
#cgo CFLAGS: -Iinclude
#cgo LDFLAGS: -Llib -lhello
#include "hello.h"
*/
import "C"
func main() {

	fmt.Println("heelo")

	ret:=int32(C.add(2,3))

	fmt.Println("aaaa ",ret)
}