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

// ListAntennas gets a list of available antennas to select on a given chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel:  an available channel on the device
//
// Return a list of available antenna names
func (dev *SDRDevice) ListAntennas(direction Direction, channel uint) []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_listAntennas(dev.device, C.int(direction), C.size_t(channel), &length)
	defer stringArrayClear(info, length)

	return stringArray2Go(info, length)
}

// SetAntennas sets the selected antenna on a chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - name: the name of an available antenna
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetAntennas(direction Direction, channel uint, name string) (err sdrerror.SDRError) {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return sdrerror.Err(int(C.SoapySDRDevice_setAntenna(dev.device, C.int(direction), C.size_t(channel), cName)))
}

// GetAntennas gets the selected antenna on a chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return the name of an available antenna
func (dev *SDRDevice) GetAntennas(direction Direction, channel uint) string {

	val := (*C.char)(C.SoapySDRDevice_getAntenna(dev.device, C.int(direction), C.size_t(channel)))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}
