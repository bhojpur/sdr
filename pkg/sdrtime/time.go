package sdrtime

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

// It groups utility functions to convert time and ticks.

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lSoapySDR
// #include <SoapySDR/Time.h>
import "C"

// TicksToTimeNs converts a tick count into a time in nanoseconds using the tick rate.
//
// Params:
//  - ticks: a integer tick count
//  - rate: the ticks per second
//
// Return the time in nanoseconds
func TicksToTimeNs(ticks int, rate float64) int {

	return int(C.SoapySDR_ticksToTimeNs(C.longlong(ticks), C.double(rate)))
}

// TimeNsToTicks converts a time in nanoseconds into a tick count using the tick rate.
//
// Params:
//  - timeNs: time in nanoseconds
//  - rate: the ticks per second
//
// Return the integer tick count
func TimeNsToTicks(timeNs int, rate float64) int {

	return int(C.SoapySDR_timeNsToTicks(C.longlong(timeNs), C.double(rate)))
}
