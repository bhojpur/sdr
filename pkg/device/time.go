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

// ListTimeSources gets the list of available time sources.
//
// Return a list of time source names
func (dev *SDRDevice) ListTimeSources() []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_listTimeSources(dev.device, &length)
	defer stringArrayClear(info, length)

	return stringArray2Go(info, length)
}

// SetTimeSource set the time source on the device.
//
// Params:
//  - source: the name of a time source
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetTimeSource(source string) (err sdrerror.SDRError) {

	cSource := C.CString(source)
	defer C.free(unsafe.Pointer(cSource))

	return sdrerror.Err(int(C.SoapySDRDevice_setTimeSource(dev.device, cSource)))
}

// GetTimeSource gets the time source of the device.
//
// Return the name of a time source
func (dev *SDRDevice) GetTimeSource() string {

	val := (*C.char)(C.SoapySDRDevice_getTimeSource(dev.device))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}

// HasHardwareTime checks if the device have a hardware clock
//
// Params:
//  - what: optional argument
//
// Return true if the hardware clock exists
func (dev *SDRDevice) HasHardwareTime(what string) bool {

	cWhat := C.CString(what)
	defer C.free(unsafe.Pointer(cWhat))

	return bool(C.SoapySDRDevice_hasHardwareTime(dev.device, cWhat))
}

// GetHardwareTime reads the time from the hardware clock on the device.
//
// Params:
//  - what: optional argument. The what argument can refer to a specific time counter.
//
// Return the time in nanoseconds
func (dev *SDRDevice) GetHardwareTime(what string) uint {

	cWhat := C.CString(what)
	defer C.free(unsafe.Pointer(cWhat))

	return uint(C.SoapySDRDevice_getHardwareTime(dev.device, cWhat))
}

// SetHardwareTime writes the time to the hardware clock on the device.
//
// Params:
//  - timeNs: time in nanoseconds
//  - what: optional argument. The what argument can refer to a specific time counter.
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetHardwareTime(timeNs uint, what string) (err sdrerror.SDRError) {

	cWhat := C.CString(what)
	defer C.free(unsafe.Pointer(cWhat))

	cTimeNs := C.longlong(timeNs)

	return sdrerror.Err(int(C.SoapySDRDevice_setHardwareTime(dev.device, cTimeNs, cWhat)))
}
