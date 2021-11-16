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

	FooInitProc *syscall.Proc //Foo FooInit(void);
	FooFreeProc *syscall.Proc //void FooFree(Foo);
	FooBarProc  *syscall.Proc //void FooBar(Foo);
	FooPopFrame *syscall.Proc //Frame popFrame(Foo)
	fooObj      uintptr
	name        int
}

const sizeOfUintPtr = unsafe.Sizeof(uintptr(0))

func uintptrToBytes(u *uintptr) []byte {
	return (*[sizeOfUintPtr]byte)(unsafe.Pointer(u))[:]
}
func (d *godll) Init() {
	d.helloDll = syscall.MustLoadDLL("bin/hello.dll")
	d.addProc = d.helloDll.MustFindProc("add")

	d.FooInitProc = d.helloDll.MustFindProc("FooInit")
	d.FooFreeProc = d.helloDll.MustFindProc("FooFree")
	d.FooBarProc = d.helloDll.MustFindProc("FooBar")
	d.FooPopFrame = d.helloDll.MustFindProc("popFrame")

	d.fooObj, _, _ = d.FooInitProc.Call()
	d.FooBarProc.Call(d.fooObj)

	ret, _, _ := d.FooPopFrame.Call(d.fooObj)
	aa:=uintptrToBytes(ret)
	

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
func callDLL() {
	dll32 := syscall.NewLazyDLL("bin/hello.dll")
	fmt.Println("call dll:", dll32.Name)
	g := dll32.NewProc("add")

	ret, _, _ := g.Call(uintptr(4), uintptr(5))
	fmt.Println("result is:", ret)
}

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Println("Email is ", u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {

	fmt.Println("heelo")
	var u user
	u.notify()
	u.changeEmail("1012083552@qq.com")
	u.notify()

	var mydll godll
	mydll.Init()
	mydll.Add(4, 5)
}
