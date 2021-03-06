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

// ListRegisterInterfaces gets a list of available register interfaces by name.
//
// Return a list of available register interfaces
func (dev *SDRDevice) ListRegisterInterfaces() []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_listRegisterInterfaces(dev.device, &length)
	defer stringArrayClear(info, length)

	return stringArray2Go(info, length)
}

// WriteRegister writes a register on the device given the interface name. This can represent a register on a soft CPU,
// FPGA, IC; the interpretation is up the implementation to decide.
//
// Params:
//  - name: the name of a available register interface
//  - addr: the register address
//  - value: the register value
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteRegister(name string, addr uint32, value uint32) (err sdrerror.SDRError) {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cAddr := C.uint(addr)
	cValue := C.uint(value)

	return sdrerror.Err(int(C.SoapySDRDevice_writeRegister(dev.device, cName, cAddr, cValue)))
}

// ReadRegister reads a register on the device given the interface name.
//
// Params:
//  - name: the name of a available register interface
//  - addr: the register address
//
// Return an error or nil in case of success
func (dev *SDRDevice) ReadRegister(name string, addr uint32) uint32 {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cAddr := C.uint(addr)

	return uint32(C.SoapySDRDevice_readRegister(dev.device, cName, cAddr))
}

// WriteRegisters writes a memory block on the device given the interface name. This can represent a memory block on a
// soft CPU, FPGA, IC; the interpretation is up the implementation to decide.
//
// Params:
//  - name: the name of a available register interface
//  - addr: the register address
//  - value: the register value
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteRegisters(name string, addr uint32, value []uint32) (err sdrerror.SDRError) {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cAddr := C.uint(addr)
	cValue := (*C.uint)(unsafe.Pointer(&value[0]))
	cLength := C.size_t(len(value))

	return sdrerror.Err(int(C.SoapySDRDevice_writeRegisters(dev.device, cName, cAddr, cValue, cLength)))
}

// ReadRegisters reads a a memory block on the device given the interface name. Pass the number of words to be read
// in via length;
//
// Params:
//  - name: the name of a available register interface
//  - addr: the register address
//
// Return an error or nil in case of success
func (dev *SDRDevice) ReadRegisters(name string, addr uint32, length uint) []uint32 {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cAddr := C.uint(addr)
	cLength := C.size_t(length)

	cValue := C.SoapySDRDevice_readRegisters(dev.device, cName, cAddr, &cLength)
	defer C.free(unsafe.Pointer(cValue))

	var uintTemplate *C.uint

	results := make([]uint32, int(cLength))

	// Read all the strings
	for i := 0; i < int(cLength); i++ {
		val := (*C.uint)(unsafe.Pointer(uintptr(unsafe.Pointer(cValue)) + uintptr(i)*unsafe.Sizeof(uintTemplate)))
		results[i] = uint32(*val)
	}

	return results
}
