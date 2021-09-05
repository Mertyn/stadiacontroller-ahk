package stadiacontroller

import "C"

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	ahk        = windows.NewLazyDLL("AutoHotkey.dll")
	ahktextdll = ahk.NewProc("ahktextdll")

	addFile     = ahk.NewProc("addFile")
	ahkFunction = ahk.NewProc("ahkFunction")

	ahkTerminate = ahk.NewProc("ahkTerminate")
)

func InitAHK() {
	scriptName := "stadiacontroller.ahk"
	utfPtr, _ := syscall.UTF16PtrFromString(scriptName)
	uintPtr := uintptr(unsafe.Pointer(utfPtr))

	ahktextdll.Call()
	addFile.Call(uintPtr)
}

func CallAHKFunction(name string) {
	utfPtr, _ := syscall.UTF16PtrFromString(name)
	uintPtr := uintptr(unsafe.Pointer(utfPtr))

	ahkFunction.Call(uintPtr)
}

func CloseAHK() {
	ahkTerminate.Call()
}
