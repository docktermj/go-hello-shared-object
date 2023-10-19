package main

import (
	"github.com/docktermj/go-hello-shared-object-custom/encryption"
)

const (
	SIGNATURE = "LoVeToHaCk"
)

var customEncryption CustomEncryption = CustomEncryption{
	UserMethods: &encryption.Encrypt{},
}
