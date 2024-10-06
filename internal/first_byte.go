package internal

/*
#cgo LDFLAGS: ./lib/libfirstbyte.a -ldl
#include "../lib/lib.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func FirstByte(s string) rune {
	cstr_word := C.CString(s)
	defer C.free(unsafe.Pointer(cstr_word))
	c_char := C.first_byte(cstr_word)
	return rune(c_char)
}
