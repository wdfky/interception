package interception

import (
	"syscall"
	"unsafe"
)

type MouseStroke struct {
	State       uint16
	Flags       uint16
	Rolling     int16
	X           int32
	Y           int32
	Information uint32
	isInjected  chan bool
}

type KeyBoardStroke struct {
	Code        uint16
	State       uint16
	Information uint32
	isInjected  chan bool
}

type Interception struct {
	handle syscall.Handle
	ctx    uintptr
}

func New() *Interception {
	h, err := syscall.LoadLibrary("./interception.dll")
	if err != nil {
		panic(err)
	}
	proc, err := syscall.GetProcAddress(h, "interception_create_context")
	if err != nil {
		panic("CreateContext " + err.Error())
	}
	ctx, _, errno := syscall.SyscallN(uintptr(proc))
	if errno != 0 {
		panic("CreateContext " + errno.Error())
	}
	return &Interception{handle: h, ctx: ctx}
}
func (i *Interception) Destroy() {
	proc, err := syscall.GetProcAddress(i.handle, "interception_destroy_context")
	if err != nil {
		panic("DestroyContext " + err.Error())
	}
	_, _, errno := syscall.SyscallN(uintptr(proc), uintptr(i.ctx))
	if errno != 0 {
		panic("DestroyContext " + errno.Error())
	}
	syscall.FreeLibrary(i.handle)
}

func (i *Interception) CreateContext() {
	proc, err := syscall.GetProcAddress(i.handle, "interception_create_context")
	if err != nil {
		panic("CreateContext " + err.Error())
	}
	ctx, _, errno := syscall.SyscallN(uintptr(proc))
	if errno != 0 {
		panic("CreateContext " + errno.Error())
	}
	i.ctx = ctx
}

func (i *Interception) IsMouse() uintptr {
	proc, err := syscall.GetProcAddress(i.handle, "interception_is_mouse")
	if err != nil {
		panic("IsMouse " + err.Error())
	}

	return proc
}

func (i *Interception) IsKeyBoard() uintptr {
	proc, err := syscall.GetProcAddress(i.handle, "interception_is_keyboard")
	if err != nil {
		panic("IsKeyBoard " + err.Error())
	}

	return proc
}

func (i *Interception) SetFilter(callback uintptr, filter Filter) {
	proc, err := syscall.GetProcAddress(i.handle, "interception_set_filter")
	if err != nil {
		panic("SetFilter " + err.Error())
	}
	_, _, errno := syscall.SyscallN(uintptr(proc), uintptr(i.ctx), uintptr(callback), uintptr(filter))
	if errno != 0 {
		panic("SetFilter " + errno.Error())
	}
}

func (i *Interception) Wait() int {
	proc, err := syscall.GetProcAddress(i.handle, "interception_wait")
	if err != nil {
		panic("Wait " + err.Error())
	}
	r, _, errno := syscall.SyscallN(uintptr(proc), uintptr(i.ctx))
	if errno != 0 {
		panic("Wait " + errno.Error())
	}
	return int(r)
}

func (i *Interception) Receive(device int, stroke interface{}) int {
	proc, err := syscall.GetProcAddress(i.handle, "interception_receive")
	if err != nil {
		panic("Receive " + err.Error())
	}
	switch stroke.(type) {
	case *MouseStroke:
		r, _, errno := syscall.SyscallN(uintptr(proc), uintptr(i.ctx), uintptr(device), uintptr(unsafe.Pointer(stroke.(*MouseStroke))), uintptr(1))
		if errno != 0 {
			panic("Receive " + errno.Error())
		}
		return int(r)
	case *KeyBoardStroke:
		r, _, errno := syscall.SyscallN(uintptr(proc), uintptr(i.ctx), uintptr(device), uintptr(unsafe.Pointer(stroke.(*KeyBoardStroke))), uintptr(1))
		if errno != 0 {
			panic("Receive " + errno.Error())
		}
		return int(r)
	default:
		panic("Type error stroke must be *MouseStroke or *KeyBoardStroke")
	}

}

func (i *Interception) Send(device int, stroke interface{}) int {
	proc, err := syscall.GetProcAddress(i.handle, "interception_send")
	if err != nil {
		panic("Send " + err.Error())
	}
	switch stroke.(type) {
	case *MouseStroke:
		r, _, errno := syscall.SyscallN(uintptr(proc), uintptr(i.ctx), uintptr(device), uintptr(unsafe.Pointer(stroke.(*MouseStroke))), uintptr(1))
		if errno != 0 {
			panic("Send " + errno.Error())
		}
		return int(r)
	case *KeyBoardStroke:
		r, _, errno := syscall.SyscallN(uintptr(proc), uintptr(i.ctx), uintptr(device), uintptr(unsafe.Pointer(stroke.(*KeyBoardStroke))), uintptr(1))
		if errno != 0 {
			panic("Send " + errno.Error())
		}
		return int(r)
	default:
		panic("Type error stroke must be *MouseStroke or *KeyBoardStroke")
	}
}
