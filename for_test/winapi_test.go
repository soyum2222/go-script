package for_test

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
	"testing"
	"time"
	"unsafe"
)

var (
	user32   = syscall.NewLazyDLL("User32.dll")
	kernel32 = syscall.NewLazyDLL("Kernel32.dll")
)

func TestGetCurrentProcessId(t *testing.T) {

	f := kernel32.NewProc("GetCurrentProcessId")
	r1, r2, _ := f.Call()
	//if err != nil {
	//	panic(err)
	//}

	var i *int

	i = &([]int{255}[0])
	fmt.Println(i)

	fmt.Println(r1, r2)
	time.Sleep(time.Hour)

}

func TestVirtualProtect(t *testing.T) {

	VirtualProtectEx := kernel32.NewProc("VirtualProtectEx")

	openprocess := kernel32.NewProc("OpenProcess")

	ReadProcessMemory := kernel32.NewProc("ReadProcessMemory")

	or1, or2, oerr := openprocess.Call(windows.PROCESS_CREATE_THREAD|windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_WRITE|windows.PROCESS_VM_READ,
		0, uintptr(3452))
	fmt.Println(or1, or2, oerr)

	var i int16
	r1, r2, err := VirtualProtectEx.Call(or1, uintptr(0xc00000a430), 8, syscall.PAGE_READWRITE, uintptr(unsafe.Pointer(&i)))
	fmt.Println(r1, r2, err)

	var ii int
	rr1, rr2, err := ReadProcessMemory.Call(or1, 0xc00000a430, uintptr(unsafe.Pointer(&ii)), 8, uintptr(unsafe.Pointer(nil)))
	fmt.Println(rr1, rr2, err)
	fmt.Println(ii)
}

func TestFindWindowsApi(t *testing.T) {

	f := user32.NewProc("FindWindowW")

	r1, r2, err := syscall.Syscall(f.Addr(), 2,
		uintptr(unsafe.Pointer(nil)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Steam"))),
		0)

	if err != 0 {
		panic(err)
	}

	fmt.Println(r1, r2)
}

func TestGetWindowThreadProcessId(t *testing.T) {

	FindWindowW := user32.NewProc("FindWindowW")

	handle, _, err := syscall.Syscall(FindWindowW.Addr(), 2,
		uintptr(unsafe.Pointer(nil)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("WeGame"))),
		0)

	if err != 0 {
		panic(err)
	}

	f := user32.NewProc("GetWindowThreadProcessId")

	r1, r2, err := syscall.Syscall(f.Addr(), 2, handle, uintptr(unsafe.Pointer(nil)), 0)
	if err != 0 {
		panic(err)
	}

	fmt.Println(r1, r2)

}
