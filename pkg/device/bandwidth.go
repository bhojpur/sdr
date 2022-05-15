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

// SetBandwidth sets the baseband filter width of the chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - bw: the baseband filter width in Hz
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetBandwidth(direction Direction, channel uint, bw float64) (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_setBandwidth(dev.device, C.int(direction), C.size_t(channel), C.double(bw))))
}

// GetBandwidth gets the baseband filter width of the chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return the baseband filter width in Hz
func (dev *SDRDevice) GetBandwidth(direction Direction, channel uint) float64 {

	return float64(C.SoapySDRDevice_getBandwidth(dev.device, C.int(direction), C.size_t(channel)))
}

// GetBandwidthRanges gets the range of possible baseband filter widths.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return a list of bandwidth ranges in Hz
func (dev *SDRDevice) GetBandwidthRanges(direction Direction, channel uint) []SDRRange {

	length := C.size_t(0)

	info := C.SoapySDRDevice_getBandwidthRange(dev.device, C.int(direction), C.size_t(channel), &length)
	defer rangeArrayClear(info)

	return rangeArray2Go(info, length)
}
