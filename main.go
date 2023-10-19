package main

/*
#include <stdlib.h>
#include <string.h>
*/

import "C"

import (
	"github.com/docktermj/go-hello-shared-object/sharedobject"
)

// int G2Encryption_InitPlugin(const struct CParameterList* configParams, char *error, const size_t maxErrorSize, size_t* errorSize);
//
//export G2Encryption_InitPlugin
func G2Encryption_InitPlugin(nothing *C.int, error_msg *C.char, maxErrorSize C.size_t, errorSize *C.size_t) C.int {
	sharedObject := sharedobject.SharedObject{}
	return sharedObject.InitPlugin(nothing, error_msg, maxErrorSize, errorSize)
}
