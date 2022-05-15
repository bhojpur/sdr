package device

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lSoapySDR
// #include <stdlib.h>
// #include <stddef.h>
// #include <SoapySDR/Device.h>
// #include <SoapySDR/Formats.h>
// #include <SoapySDR/Types.h>
import "C"
import (
	"unsafe"

	"github.com/bhojpur/sdr/pkg/sdrerror"
)

// WriteI2C writes to an available I2C slave.
//
// If the device contains multiple I2C masters, the address bits can encode which master.
//
// Params:
//  - addr: the address of the slave
//  - data: an array of bytes write out
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteI2C(addr int32, data []uint8) (err sdrerror.SDRError) {

	cAddr := C.int(addr)
	cData := (*C.char)(unsafe.Pointer(&data[0]))
	cNumBytes := C.size_t(len(data))

	return sdrerror.Err(int(C.SoapySDRDevice_writeI2C(dev.device, cAddr, cData, cNumBytes)))
}

// ReadI2C reads from an available I2C slave.
//
// If the device contains multiple I2C masters, the address bits can encode which master.
//
// Params:
//  - addr: the address of the slave
//  - numBytes: the number of bytes to read
//
// Return the bytes actually read.
func (dev *SDRDevice) ReadI2C(addr int32, numBytes uint) (data []uint8) {

	cAddr := C.int(addr)
	cNumBytes := C.size_t(len(data))

	cData := C.SoapySDRDevice_readI2C(dev.device, cAddr, &cNumBytes)
	defer C.free(unsafe.Pointer(cData))

	data = make([]uint8, int(cNumBytes))

	for i := 0; i < int(cNumBytes); i++ {

		// Get the data from the returned array
		valPtr := (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(cData)) + uintptr(i)))
		val := uint8(*valPtr)

		// Fill the slice to return
		data[i] = val
	}

	return data
}
