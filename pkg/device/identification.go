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
// #include <SoapySDR/Types.h>
import "C"
import "unsafe"

// GetDriverKey returns a key that uniquely identifies the device driver.
//
// This key identifies the underlying implementation. Several variants of a product may share a driver.
func (dev *SDRDevice) GetDriverKey() (driverKey string) {

	val := (*C.char)(C.SoapySDRDevice_getDriverKey(dev.device))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}

// GetHardwareKey returns a key that uniquely identifies the hardware.
//
// This key should be meaningful to the user to optimize for the underlying hardware.
func (dev *SDRDevice) GetHardwareKey() (hardwareKey string) {

	val := (*C.char)(C.SoapySDRDevice_getHardwareKey(dev.device))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}

// GetHardwareInfo queries a dictionary of available device information.
//
// This dictionary can any number of values like vendor name, product name, revisions, serials...
// This information can be displayed to the user to help identify the instantiated device.
func (dev *SDRDevice) GetHardwareInfo() (hardwareInfo map[string]string) {

	info := (C.SoapySDRKwargs)(C.SoapySDRDevice_getHardwareInfo(dev.device))
	defer argsClear(info)

	return args2Go(info)
}
