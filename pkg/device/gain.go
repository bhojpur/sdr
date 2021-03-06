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

// ListGains lists available amplification elements.
//
// Elements should be in order RF to baseband.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//
// Return a list of gain string names
func (dev *SDRDevice) ListGains(direction Direction, channel uint) []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_listGains(dev.device, C.int(direction), C.size_t(channel), &length)
	defer stringArrayClear(info, length)

	return stringArray2Go(info, length)
}

// HasGainMode returns if the device support automatic gain control
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//
// Return true for automatic gain control
func (dev *SDRDevice) HasGainMode(direction Direction, channel uint) bool {

	return bool(C.SoapySDRDevice_hasGainMode(dev.device, C.int(direction), C.size_t(channel)))
}

// SetGainMode sets the automatic gain mode on the chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//  - automatic: true for automatic gain setting
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetGainMode(direction Direction, channel uint, automatic bool) (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_setGainMode(dev.device, C.int(direction), C.size_t(channel), C.bool(automatic))))
}

// GetGainMode gets the automatic gain mode on the chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//
// Return true for automatic gain setting
func (dev *SDRDevice) GetGainMode(direction Direction, channel uint) bool {

	return bool(C.SoapySDRDevice_getGainMode(dev.device, C.int(direction), C.size_t(channel)))
}

// SetGain sets the overall amplification in a chain.
//
// The gain will be distributed automatically across available element.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - value: the new amplification value in dB
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetGain(direction Direction, channel uint, gain float64) (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_setGain(dev.device, C.int(direction), C.size_t(channel), C.double(gain))))
}

// SetGainElement sets the value of a amplification element in a chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - name: the name of an amplification element
//  - value: the new amplification value in dB
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetGainElement(direction Direction, channel uint, name string, gain float64) (err sdrerror.SDRError) {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return sdrerror.Err(int(C.SoapySDRDevice_setGainElement(dev.device, C.int(direction), C.size_t(channel), cName, C.double(gain))))
}

// GetGain gets the overall value of the gain elements in a chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return the value of the gain in dB
func (dev *SDRDevice) GetGain(direction Direction, channel uint) float64 {

	return float64(C.SoapySDRDevice_getGain(dev.device, C.int(direction), C.size_t(channel)))
}

// GetGainElement gets the value of an individual amplification element in a chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - name: the name of an amplification element
//
// Return the value of the gain in dB
func (dev *SDRDevice) GetGainElement(direction Direction, channel uint, name string) float64 {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return float64(C.SoapySDRDevice_getGainElement(dev.device, C.int(direction), C.size_t(channel), cName))
}

// GetGainRange gets the overall range of possible gain values.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return a list of gain ranges in dB
func (dev *SDRDevice) GetGainRange(direction Direction, channel uint) SDRRange {

	cRange := C.SoapySDRDevice_getGainRange(dev.device, C.int(direction), C.size_t(channel))

	return SDRRange{
		Minimum: float64(cRange.minimum),
		Maximum: float64(cRange.maximum),
		Step:    float64(cRange.step),
	}
}

// GetGainElementRange gets the range of possible gain values for a specific element.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - name: the name of an amplification element
//
// Return a list of gain ranges in dB
func (dev *SDRDevice) GetGainElementRange(direction Direction, channel uint, name string) SDRRange {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cRange := C.SoapySDRDevice_getGainElementRange(dev.device, C.int(direction), C.size_t(channel), cName)

	return SDRRange{
		Minimum: float64(cRange.minimum),
		Maximum: float64(cRange.maximum),
		Step:    float64(cRange.step),
	}
}
