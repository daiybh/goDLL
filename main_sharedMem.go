package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

type godll struct {
	dllHandle           *syscall.DLL
	proc_openProvider   *syscall.Proc //__declspec(dllexport) tagProviderHandle* openProvider(int providerID);
	proc_popFrame       *syscall.Proc //__declspec(dllexport) Frame popFrame(tagProviderHandle* _ProviderHandle);
	proc_pushVideoFrame *syscall.Proc //__declspec(dllexport) bool pushVideoFrame(tagProviderHandle* _ProviderHandle, Frame frame, int nSize);
	proc_pushAudioFrame *syscall.Proc //__declspec(dllexport) bool pushAudioFrame(tagProviderHandle* _ProviderHandle, Frame frame, int nSize);

	sharememObj uintptr
	name        int
}

const sizeOfUintPtr = unsafe.Sizeof(uintptr(0))

func uintptrToBytes(u *uintptr) []byte {
	return (*[sizeOfUintPtr]byte)(unsafe.Pointer(u))[:]
}
func (d *godll) Init() {
	d.dllHandle = syscall.MustLoadDLL("SharedMemCoreLibDLL.dll")

	d.proc_openProvider = d.dllHandle.MustFindProc("openProvider")
	d.proc_popFrame = d.dllHandle.MustFindProc("popFrame")
	d.proc_pushVideoFrame = d.dllHandle.MustFindProc("pushVideoFrame")
	d.proc_pushAudioFrame = d.dllHandle.MustFindProc("pushAudioFrame")

}
func (d *godll) Open(cam int) {
	d.sharememObj, _, _ = d.proc_openProvider.Call(uintptr(cam))
	fmt.Println("result is:", d.sharememObj)
}

func (d *godll) PushVideo() {
	bufByte := make([]byte, 1920*1080*2)
	buffer := unsafe.Pointer(&bufByte[0])

	ret, _, _ := d.proc_pushVideoFrame.Call(d.sharememObj, uintptr(buffer), 1920*1080*2)
	fmt.Println("result is:", ret)
}

func (d *godll) pushAudioFrame() {

	bufByte := make([]byte, 1920*1080*2)
	buffer := unsafe.Pointer(&bufByte[0])

	ret, _, _ := d.proc_pushAudioFrame.Call(d.sharememObj, uintptr(buffer), 1920)
	fmt.Println("result is:", ret)
}

func (d *godll) Close(cam int) {
}

func main() {

	var mydll godll
	mydll.Init()
	mydll.Open(1)
	
	loop := 0
	for {
		fmt.Println("", loop)
		loop += 1
		time.Sleep(time.Duration(1) * time.Second)
	}
}
