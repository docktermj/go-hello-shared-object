package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "unsafe"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type EncryptionMethods interface {
	InitEncryption() error
	CloseEncryption() error
	Encrypt(clearText string, maxClearTextLen int) (string, error)
	EncryptDeterministic(clearText string, maxClearTextLen int) (string, error)
	Decrypt(clearText string, maxClearTextLen int) (string, error)
	DecryptDeterministic(clearText string, maxClearTextLen int) (string, error)
}

type CustomEncryption struct {
	UserMethods EncryptionMethods
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func findErrorCode(errorString string) int {
	var result int = 0
	return result

}

// ----------------------------------------------------------------------------
// Interface functions
// ----------------------------------------------------------------------------

// int G2Encryption_InitPlugin(const struct CParameterList* configParams, char *error, const size_t maxErrorSize, size_t* errorSize);
//
//export G2Encryption_InitPlugin
func G2Encryption_InitPlugin(nothing *C.int, error_msg *C.char, maxErrorSize C.size_t, errorSize *C.size_t) C.int {
	err := customEncryption.UserMethods.InitEncryption()
	if err != nil {
		errStr := err.Error()
		errLen := C.size_t(len(errStr))
		s := C.CString(errStr)
		defer C.free(unsafe.Pointer(s))
		C.memcpy(unsafe.Pointer(error_msg), unsafe.Pointer(s), errLen)
		*errorSize = errLen
		return C.int(findErrorCode(errStr))
	}
	return 0
}

// int G2Encryption_ClosePlugin(char *error, const size_t maxErrorSize, size_t* errorSize);
//
//export G2Encryption_ClosePlugin
func G2Encryption_ClosePlugin(error_msg *C.char, maxErrorSize C.size_t, errorSize *C.size_t) C.int {
	err := customEncryption.UserMethods.CloseEncryption()
	if err != nil {
		errStr := err.Error()
		errLen := C.size_t(len(errStr))
		s := C.CString(errStr)
		defer C.free(unsafe.Pointer(s))
		C.memcpy(unsafe.Pointer(error_msg), unsafe.Pointer(s), errLen)
		*errorSize = errLen
		return C.int(findErrorCode(errStr))
	}
	return 0
}

// int G2Encryption_GetSignature(char *signature, const size_t maxSignatureSize, size_t* signatureSize, char *error, const size_t maxErrorSize, size_t* errorSize);
//
//export G2Encryption_GetSignature
func G2Encryption_GetSignature(signature *C.char, maxSignatureSize C.size_t, signatureSize *C.size_t, error_msg *C.char, maxErrorSize C.size_t, errorSize *C.size_t) C.int {
	cypherText, err := customEncryption.UserMethods.EncryptDeterministic(SIGNATURE, int(maxSignatureSize))
	if err != nil {
		errStr := err.Error()
		errLen := C.size_t(len(errStr))
		s := C.CString(errStr)
		defer C.free(unsafe.Pointer(s))
		C.memcpy(unsafe.Pointer(error_msg), unsafe.Pointer(s), errLen)
		*errorSize = errLen
		return C.int(findErrorCode(errStr))
	}
	cypherTextLen := C.size_t(len(cypherText))
	s := C.CString(cypherText)
	defer C.free(unsafe.Pointer(s))
	C.memcpy(unsafe.Pointer(signature), unsafe.Pointer(s), cypherTextLen)
	*signatureSize = cypherTextLen
	return 0
}

// int G2Encryption_ValidateSignatureCompatibility(const char *signatureToValidate, const size_t signatureToValidateSize, char *error, const size_t maxErrorSize, size_t* errorSize);
//
//export G2Encryption_ValidateSignatureCompatibility
func G2Encryption_ValidateSignatureCompatibility(signatureToValidate *C.char, signatureToValidateSize C.size_t, error_msg *C.char, maxErrorSize C.size_t, errorSize *C.size_t) C.int {
	cypherText, err := customEncryption.UserMethods.EncryptDeterministic(SIGNATURE, 100*1024)
	if err != nil {
		errStr := err.Error()
		errLen := C.size_t(len(errStr))
		s := C.CString(errStr)
		defer C.free(unsafe.Pointer(s))
		C.memcpy(unsafe.Pointer(error_msg), unsafe.Pointer(s), errLen)
		*errorSize = errLen
		return C.int(findErrorCode(errStr))
	}
	sigLen := C.size_t(len(cypherText))
	if signatureToValidateSize != sigLen {
		return -1
	}
	s := C.CString(cypherText)
	defer C.free(unsafe.Pointer(s))
	return C.memcmp(unsafe.Pointer(signatureToValidate), unsafe.Pointer(s), sigLen)
}

// int G2Encryption_EncryptDataField(const char *input, const size_t inputSize, char *result, const size_t maxResultSize, size_t* resultSize, char *error, const size_t maxErrorSize, size_t* errorSize);
//
//export G2Encryption_EncryptDataField
func G2Encryption_EncryptDataField(input *C.char, inputSize C.size_t, result *C.char, maxResultSize C.size_t, resultSize *C.size_t, error_msg *C.char, maxErrorSize C.size_t, errorSize *C.size_t) C.int {
	cypherText, err := customEncryption.UserMethods.Encrypt(C.GoStringN(input, C.int(inputSize)), int(maxResultSize))
	if err != nil {
		errStr := err.Error()
		errLen := C.size_t(len(errStr))
		s := C.CString(errStr)
		defer C.free(unsafe.Pointer(s))
		C.memcpy(unsafe.Pointer(error_msg), unsafe.Pointer(s), errLen)
		*errorSize = errLen
		return C.int(findErrorCode(errStr))
	}
	cypherTextLen := C.size_t(len(cypherText))
	s := C.CString(cypherText)
	defer C.free(unsafe.Pointer(s))
	C.memcpy(unsafe.Pointer(result), unsafe.Pointer(s), cypherTextLen)
	*resultSize = cypherTextLen
	return 0
}

// int G2Encryption_EncryptDataFieldDeterministic(const char *input, const size_t inputSize, char *result, const size_t maxResultSize, size_t* resultSize, char *error, const size_t maxErrorSize, size_t* errorSize);
//
//export G2Encryption_EncryptDataFieldDeterministic
func G2Encryption_EncryptDataFieldDeterministic(input *C.char, inputSize C.size_t, result *C.char, maxResultSize C.size_t, resultSize *C.size_t, error_msg *C.char, maxErrorSize C.size_t, errorSize *C.size_t) C.int {
	cypherText, err := customEncryption.UserMethods.EncryptDeterministic(C.GoStringN(input, C.int(inputSize)), int(maxResultSize))
	if err != nil {
		errStr := err.Error()
		errLen := C.size_t(len(errStr))
		s := C.CString(errStr)
		defer C.free(unsafe.Pointer(s))
		C.memcpy(unsafe.Pointer(error_msg), unsafe.Pointer(s), errLen)
		*errorSize = errLen
		return C.int(findErrorCode(errStr))
	}
	cypherTextLen := C.size_t(len(cypherText))
	s := C.CString(cypherText)
	defer C.free(unsafe.Pointer(s))
	C.memcpy(unsafe.Pointer(result), unsafe.Pointer(s), cypherTextLen)
	*resultSize = cypherTextLen
	return 0
}

// int G2Encryption_DecryptDataFieldDeterministic(const char *input, const size_t inputSize, char *result, const size_t maxResultSize, size_t* resultSize, char *error, const size_t maxErrorSize, size_t* errorSize);
//
//export G2Encryption_DecryptDataFieldDeterministic
func G2Encryption_DecryptDataFieldDeterministic(input *C.char, inputSize C.size_t, result *C.char, maxResultSize C.size_t, resultSize *C.size_t, error_msg *C.char, maxErrorSize C.size_t, errorSize *C.size_t) C.int {
	clearText, err := customEncryption.UserMethods.DecryptDeterministic(C.GoStringN(input, C.int(inputSize)), int(maxResultSize))
	if err != nil {
		errStr := err.Error()
		errLen := C.size_t(len(errStr))
		s := C.CString(errStr)
		defer C.free(unsafe.Pointer(s))
		C.memcpy(unsafe.Pointer(error_msg), unsafe.Pointer(s), errLen)
		*errorSize = errLen
		return C.int(findErrorCode(errStr))
	}
	clearTextLen := C.size_t(len(clearText))
	s := C.CString(clearText)
	defer C.free(unsafe.Pointer(s))
	C.memcpy(unsafe.Pointer(result), unsafe.Pointer(s), clearTextLen)
	*resultSize = clearTextLen
	return 0
}

func main() {}
