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

// ListUARTs enumerate the available UART devices.
//
// Return a list of names of available UARTs
func (dev *SDRDevice) ListUARTs() []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_listUARTs(dev.device, &length)
	defer stringArrayClear(info, length)

	return stringArray2Go(info, length)
}

// WriteUART writes data to a UART device.
//
// Its up to the implementation to set the baud rate, carriage return settings, flushing on newline.
//
// Params:
//  - which: the name of an available UART
//  - data: an array of byte to send (packed as a string for convenience)
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteUART(which string, data string) (err sdrerror.SDRError) {

	cWhich := C.CString(which)
	defer C.free(unsafe.Pointer(cWhich))

	cData := C.CString(data)
	defer C.free(unsafe.Pointer(cData))

	return sdrerror.Err(int(C.SoapySDRDevice_writeUART(dev.device, cWhich, cData)))
}

// ReadUART read bytes from a UART until timeout or newline.
//
// Its up to the implementation to set the baud rate, carriage return settings, flushing on newline.
//
// Params:
//  - which: the name of an available UART
//  - timeoutUs: a timeout in microseconds
//
// Return an array of byte packed as a string fdr convenience
func (dev *SDRDevice) ReadUART(which string, timeoutUs uint) string {

	cWhich := C.CString(which)
	defer C.free(unsafe.Pointer(cWhich))

	val := (*C.char)(C.SoapySDRDevice_readUART(dev.device, cWhich, C.long(timeoutUs)))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}
