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

// SetFrontendMapping sets the frontend mapping of available DSP units to RF frontends.
//
// This mapping controls channel mapping and channel availability.
//
// Params:
//  - direction: the channel direction DirectionRX or DIRECTION_TX
//  - mapping: a vendor-specific mapping string
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetFrontendMapping(direction Direction, mapping string) (err sdrerror.SDRError) {

	cMapping := C.CString(mapping)
	defer C.free(unsafe.Pointer(cMapping))

	return sdrerror.Err(int(C.SoapySDRDevice_setFrontendMapping(dev.device, C.int(direction), cMapping)))
}

// GetFrontendMapping gets the mapping configuration string.
//
// Params:
//  - direction: the channel direction DirectionRX or DIRECTION_TX
//
// Return the vendor-specific mapping string
func (dev *SDRDevice) GetFrontendMapping(direction Direction) string {

	val := (*C.char)(C.SoapySDRDevice_getFrontendMapping(dev.device, C.int(direction)))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}

// GetNumChannels gets a number of channels given the streaming direction.
//
// Params:
//  - direction: the channel direction DirectionRX or DIRECTION_TX
//
// Return the number of channels
func (dev *SDRDevice) GetNumChannels(direction Direction) uint {

	return uint(C.SoapySDRDevice_getNumChannels(dev.device, C.int(direction)))
}

// GetChannelInfo gets channel info given the streaming direction.
//
// Params:
//  - direction: the channel direction DirectionRX or DIRECTION_TX
//  - channel: the channel number to get info for
//
// Return channel information
func (dev *SDRDevice) GetChannelInfo(direction Direction, channel uint) map[string]string {

	info := C.SoapySDRDevice_getChannelInfo(dev.device, C.int(direction), C.size_t(channel))
	defer argsClear(info)

	return args2Go(info)
}

// GetFullDuplex finds out if the specified channel is full or half duplex.
//
// Params:
//  - direction the channel direction DirectionRX or DIRECTION_TX
//  - channel an available channel on the device
//
// Return true for full duplex, false for half duplex
func (dev *SDRDevice) GetFullDuplex(direction Direction, channel uint) bool {

	return bool(C.SoapySDRDevice_getFullDuplex(dev.device, C.int(direction), C.size_t(channel)))
}
