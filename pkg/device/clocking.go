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

// SetMasterClockRate the frontend DC offset correction.
//
// Params:
//  - rate: the clock rate in Hz
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetMasterClockRate(rate float64) (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_setMasterClockRate(dev.device, C.double(rate))))
}

// GetMasterClockRate gets the master clock rate of the device.
//
// Return the clock rate in Hz
func (dev *SDRDevice) GetMasterClockRate() float64 {

	return float64(C.SoapySDRDevice_getMasterClockRate(dev.device))
}

// GetMasterClockRates gets the range of available master clock rates.
//
// Return a list of clock rate ranges in Hz
func (dev *SDRDevice) GetMasterClockRates() []SDRRange {

	length := C.size_t(0)

	info := C.SoapySDRDevice_getMasterClockRates(dev.device, &length)
	defer rangeArrayClear(info)

	return rangeArray2Go(info, length)
}

// ListClockSources gets the list of available clock sources.
//
// Return a list of available antenna names
func (dev *SDRDevice) ListClockSources() []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_listClockSources(dev.device, &length)
	defer stringArrayClear(info, length)

	return stringArray2Go(info, length)
}

// SetClockSource set the clock source on the device.
//
// Params:
//  - source: the name of a clock source
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetClockSource(source string) (err sdrerror.SDRError) {

	cSource := C.CString(source)
	defer C.free(unsafe.Pointer(cSource))

	return sdrerror.Err(int(C.SoapySDRDevice_setClockSource(dev.device, cSource)))
}

// GetClockSource gets the clock source of the device.
//
// Return the name of a clock source
func (dev *SDRDevice) GetClockSource() string {

	val := (*C.char)(C.SoapySDRDevice_getClockSource(dev.device))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}
