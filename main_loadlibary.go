package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

type godll struct {
	helloDll *syscall.DLL
	addProc  *syscall.Proc
	pushProc *syscall.Proc //int push(char* buf, int size)

	FooInitProc  *syscall.Proc //Foo FooInit(void);
	FooFreeProc  *syscall.Proc //void FooFree(Foo);
	FooBarProc   *syscall.Proc //void FooBar(Foo);
	FooPopFrame  *syscall.Proc //Frame popFrame(Foo)
	FooPushFrame *syscall.Proc //void pushFrame(Foo,Frame);
	fooObj       uintptr
	name         int
}

const sizeOfUintPtr = unsafe.Sizeof(uintptr(0))

func uintptrToBytes(u *uintptr) []byte {
	return (*[sizeOfUintPtr]byte)(unsafe.Pointer(u))[:]
}
func (d *godll) Init() {
	d.helloDll = syscall.MustLoadDLL("bin/helloDLL.dll")
	d.addProc = d.helloDll.MustFindProc("add")

	d.FooInitProc = d.helloDll.MustFindProc("FooInit")
	d.FooFreeProc = d.helloDll.MustFindProc("FooFree")
	d.FooBarProc = d.helloDll.MustFindProc("FooBar")
	d.FooPopFrame = d.helloDll.MustFindProc("popFrame")
	d.FooPushFrame = d.helloDll.MustFindProc("pushFrame")

	d.fooObj, _, _ = d.FooInitProc.Call()
	d.FooBarProc.Call(d.fooObj)

	for i := 0; i < 20; i++ {
		ret, _, _ := d.FooPopFrame.Call(d.fooObj)
		fmt.Println("ret", ret)
		lowPointer := (*[2]int)(unsafe.Pointer(ret))
		lowPointer[0] = i
		lowPointer[1] = 0xaa

		d.FooPushFrame.Call(d.fooObj, ret, 2)
	}
}
func (d godll) Push(a int, b int) int {
	ret, _, _ := d.addProc.Call(uintptr(a), uintptr(b))
	fmt.Println("result is:", ret)
	return int(ret)
}
func (d godll) Add(a int, b int) int {
	ret, _, _ := d.addProc.Call(uintptr(a), uintptr(b))
	fmt.Println("result is:", ret)
	return int(ret)
}
func main() {

	fmt.Println("heelo")

	var mydll godll
	mydll.Init()
	mydll.Add(4, 5)
}
