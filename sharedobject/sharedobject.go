package sharedobject

/*
#include <stdlib.h>
#include <string.h>
*/

import "C"

import (
	"github.com/docktermj/go-hello-shared-object/encrypt"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type SharedObject struct {
	UserMethods encrypt.Encrypt
}

func (sharedObject *SharedObject) InitPlugin(nothing *C.int, error_msg *C.char, maxErrorSize C.size_t, errorSize *C.size_t) C.int {
	var err error = nil
	err = sharedObject.UserMethods.InitEncryption()
	if err != nil {
		return 1
		// errLen := C.size_t(len(errStr))
		// s := C.CString(errStr)
		// defer C.free(unsafe.Pointer(s))
		// C.memcpy(unsafe.Pointer(error_msg), unsafe.Pointer(s), errLen)
		// *errorSize = errLen
		// return C.int(errCode)
	}
	return 0
}
