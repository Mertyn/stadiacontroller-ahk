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
	var scriptName = "stadiacontroller.ahk"
	var utfPtr, _ = syscall.UTF16PtrFromString(scriptName)
	var uintPtr = uintptr(unsafe.Pointer(utfPtr))

	ahktextdll.Call()
	addFile.Call(uintPtr)
}

func CallAHKFunction(name string) {
	var utfPtr, _ = syscall.UTF16PtrFromString(name)
	var uintPtr = uintptr(unsafe.Pointer(utfPtr))

	ahkFunction.Call(uintPtr)
}

func CloseAHK() {
	ahkTerminate.Call()
}
