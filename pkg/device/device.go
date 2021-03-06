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
	"errors"
	"unsafe"

	"github.com/bhojpur/sdr/pkg/sdrerror"
)

// LastStatus returns the last status code after a Device API call.
//
// The status code is cleared on entry to each Device call. When an device API call throws, the C bindings catch
// the exception, and set a non-zero last status code. Use LastStatus() to determine success/failure for
// Device calls without integer status return codes.
func LastStatus() int {

	return int(C.SoapySDRDevice_lastStatus())
}

// LastError returns the last error message after a device call fails.
//
// When an device API call throws, the C bindings catch the exception, store its message in thread-safe storage,
// and return a non-zero status code to indicate failure. Use lastError() to access the exception's error message.
func LastError() string {

	// Do not free as it is internal string of Soapy
	return C.GoString(C.SoapySDRDevice_lastError())
}

// Enumerate returns a list of available devices on the system.
//
// Params:
//  - args: device construction key/value argument filters, for example {"driver":"hackrf"}. Can be set to nil if no
// filter is needed
//
// Return a list of information, each unique to a device
func Enumerate(args map[string]string) []map[string]string {

	length := C.size_t(0)

	cArgs, cArgsLength := go2Args(args)
	defer argsListClear(cArgs, cArgsLength)

	enumerateData := C.SoapySDRDevice_enumerate(cArgs, &length)
	defer argsListClear(enumerateData, length)

	return argsList2Go(enumerateData, length)
}

// EnumerateStrArgs returns a list of available devices on the system.
//
// Params:
//  - args: a markup string of key/value argument filters. Markup format for args: "keyA=valA, keyB=valB". Can be set to
// nil if no filter is needed
//
// Return a list of information, each unique to a device
func EnumerateStrArgs(args string) []map[string]string {

	length := C.size_t(0)

	cArgs := C.CString(args)
	defer C.free(unsafe.Pointer(cArgs))

	enumerateData := C.SoapySDRDevice_enumerateStrArgs(cArgs, &length)
	defer argsListClear(enumerateData, length)

	return argsList2Go(enumerateData, length)
}

// Make makes a new Device object given device construction args.
//
// The device pointer will be stored in a table so subsequent calls with the same arguments will produce the same
// device. For every call to make, there should be a matched call to unmake.
//
// Params:
//  - args: device construction key/value argument map
//
// Return a pointer to a new Device object or null for error
func Make(args map[string]string) (device *SDRDevice, err error) {

	cArgs, cArgsLength := go2Args(args)
	defer argsListClear(cArgs, cArgsLength)

	dev := C.SoapySDRDevice_make(cArgs)
	if dev == nil {
		return nil, errors.New(LastError())
	}

	return &SDRDevice{
		device: dev,
	}, nil
}

// MakeStrArgs makes a new Device object given device construction args.
//
// The device pointer will be stored in a table so subsequent calls with the same arguments will produce the same
// device. For every call to make, there should be a matched call to unmake.
//
// Params:
//  - args: a markup string of key/value arguments
//
// Return a pointer to a new Device object or null for error
func MakeStrArgs(args string) (device *SDRDevice, err error) {

	cArgs := C.CString(args)
	defer C.free(unsafe.Pointer(cArgs))

	dev := C.SoapySDRDevice_makeStrArgs(cArgs)
	if dev == nil {
		return nil, errors.New(LastError())
	}

	return &SDRDevice{
		device: dev,
	}, nil
}

// Unmake unmakes or releases a device object handle.
//
// Params:
//  - device: a pointer to a device object
//
// Return an error or nil in case of success
func (dev *SDRDevice) Unmake() (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_unmake(dev.device)))
}

// MakeList creates a list of devices from a list of construction arguments.
//
// This is a convenience call to parallelize device construction,
// and is fundamentally a parallel for loop of make(Kwargs).
//
// Params:
//  - argsList: a list of device arguments per each device
//
// Return a list of device pointers per each specified argument
func MakeList(argsList []map[string]string) (devices []*SDRDevice, err error) {

	cArgs, cLength := go2ArgsList(argsList)
	defer argsListClear(cArgs, cLength)

	dev := C.SoapySDRDevice_make_list(cArgs, cLength)
	if dev == nil {
		return nil, errors.New(LastError())
	}
	defer devicesClear(dev)

	return devices2Go(dev, cLength), nil
}

// UnmakeList unmakes or releases a list of device handles.
//
// This is a convenience call to parallelize device destruction,
// and is fundamentally a parallel for loop of unmake(Device *).
//
// Params:
//  - devices: a list of pointer to sdr devices
//
// Return an error or nil in case of success
func UnmakeList(devices []*SDRDevice) (err sdrerror.SDRError) {

	cDevices, cLength := go2Devices(devices)
	defer devicesClear(cDevices)

	return sdrerror.Err(int(C.SoapySDRDevice_unmake_list(cDevices, cLength)))
}
