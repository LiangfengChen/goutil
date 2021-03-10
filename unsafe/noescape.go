package unsafe

import "unsafe"

// NoEscape hides a pointer from escape analysis.
// It can let you allocate objs in stack instead of heap
// PLEASE USE CAREFULLY!
func NoEscape(p interface{}) interface{} {
	x := uintptr(unsafe.Pointer(&p))
	return *(*interface{})(unsafe.Pointer(x))
}

