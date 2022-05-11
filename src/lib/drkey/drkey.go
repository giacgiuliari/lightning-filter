// Copyright (c) 2021, [fullname]
// All rights reserved.

package main

import (
	"C"
	"unsafe"
)

//export GetDelegationSecret
func GetDelegationSecret(sciondAddr *C.char, srcIA, dstIA uint64, valTime int64,
	validityNotBefore, validityNotAfter *int64, key unsafe.Pointer) int {
	return -1
}

func main() {}
