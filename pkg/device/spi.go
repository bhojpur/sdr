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

// TransactSPI performs a SPI transaction and return the result.
//
// Its up to the implementation to set the clock rate, and read edge, and the write edge of the SPI core. SPI slaves
// without a readback pin will return 0.
//
// If the device contains multiple SPI masters, the address bits can encode which master.
//
// Params:
//  - addr: an address of an available SPI slave
//  - data: the SPI data, numBits-1 is first out
//  - numBits: the number of bits to clock out
//
// Return the readback data, numBits-1 is first in
func (dev *SDRDevice) TransactSPI(addr int32, data uint32, numBits uint32) uint32 {

	return uint32(C.SoapySDRDevice_transactSPI(dev.device, C.int(addr), C.uint(data), C.size_t(numBits)))
}
