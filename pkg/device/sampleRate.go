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
	"github.com/bhojpur/sdr/pkg/sdrerror"
)

// SetSampleRate sets the baseband sample rate of the chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - rate: the sample rate in samples per second
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetSampleRate(direction Direction, channel uint, rate float64) (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_setSampleRate(dev.device, C.int(direction), C.size_t(channel), C.double(rate))))
}

// GetSampleRate gets the baseband sample rate of the chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return the sample rate in samples per second
func (dev *SDRDevice) GetSampleRate(direction Direction, channel uint) float64 {

	return float64(C.SoapySDRDevice_getSampleRate(dev.device, C.int(direction), C.size_t(channel)))
}

// GetSampleRateRange gets the range of possible baseband sample rates.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return a list of sample rate ranges in samples per second
func (dev *SDRDevice) GetSampleRateRange(direction Direction, channel uint) []SDRRange {

	length := C.size_t(0)

	info := C.SoapySDRDevice_getSampleRateRange(dev.device, C.int(direction), C.size_t(channel), &length)
	defer rangeArrayClear(info)

	return rangeArray2Go(info, length)
}
