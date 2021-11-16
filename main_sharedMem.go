package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

type SharedMemDLL struct {
	dllHandle           *syscall.DLL
	proc_openProvider   *syscall.Proc //__declspec(dllexport) tagProviderHandle* openProvider(int providerID);
	proc_popFrame       *syscall.Proc //__declspec(dllexport) Frame popFrame(tagProviderHandle* _ProviderHandle);
	proc_pushVideoFrame *syscall.Proc //__declspec(dllexport) bool pushVideoFrame(tagProviderHandle* _ProviderHandle, Frame frame, int nSize);
	proc_pushAudioFrame *syscall.Proc //__declspec(dllexport) bool pushAudioFrame(tagProviderHandle* _ProviderHandle, Frame frame, int nSize);

	sharememObj uintptr
	name        int
}

func (d *SharedMemDLL) Init() {
	d.dllHandle = syscall.MustLoadDLL("SharedMemCoreLibDLL.dll")

	d.proc_openProvider = d.dllHandle.MustFindProc("openProvider")
	d.proc_popFrame = d.dllHandle.MustFindProc("popFrame")
	d.proc_pushVideoFrame = d.dllHandle.MustFindProc("pushVideoFrame")
	d.proc_pushAudioFrame = d.dllHandle.MustFindProc("pushAudioFrame")

}
func (d *SharedMemDLL) Open(cam int) {
	d.sharememObj, _, _ = d.proc_openProvider.Call(uintptr(cam))
	fmt.Println("result is:", d.sharememObj)
}

func (d *SharedMemDLL) PushVideo(bufByte []byte, size int) {
	buffer := unsafe.Pointer(&bufByte[0])

	ret, _, _ := d.proc_pushVideoFrame.Call(d.sharememObj, uintptr(buffer), uintptr(size))
	fmt.Println("result is:", ret)
}

func (d *SharedMemDLL) pushAudioFrame(bufByte []byte, size int) {

	buffer := unsafe.Pointer(&bufByte[0])

	ret, _, _ := d.proc_pushAudioFrame.Call(d.sharememObj, uintptr(buffer), uintptr(size))
	fmt.Println("result is:", ret)
}

func (d *SharedMemDLL) Close(cam int) {
}

func main() {

	var mydll SharedMemDLL
	mydll.Init()
	mydll.Open(1)

	loop := 0

	audiobufByte := make([]byte, 1920*4*16)
	videobufByte := make([]byte, 1920*1080*2)
	for {
		fmt.Println("", loop)
		loop += 1
		time.Sleep(time.Duration(1) * time.Second)
		mydll.pushAudioFrame(audiobufByte, 1920*4*16)
		mydll.PushVideo(videobufByte, 1920*1080*2)
	}
}
