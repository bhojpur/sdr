package version

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
// #include <SoapySDR/Version.h>
import "C"

// GetABIVersion gets the ABI version string that the library was built against
//
// Return the ABI version
func GetABIVersion() string {

	// Note: The string must not be freed
	return C.GoString((*C.char)(C.SoapySDR_getABIVersion()))
}

// GetAPIVersion get the SoapySDR library API version as a string. The format of
// the version string is major.minor.increment
//
// Return the API version
func GetAPIVersion() string {

	// Note: The string must not be freed
	return C.GoString((*C.char)(C.SoapySDR_getAPIVersion()))
}

// GetLibVersion gets the library version and build information string. The format
// of the version string is major.minor.patch-buildInfo. This function is commonly
// used to identify the software back-end to the user for command-line utilities
// and graphical applications.
//
// Return the library version
func GetLibVersion() string {

	// Note: The string must not be freed
	return C.GoString((*C.char)(C.SoapySDR_getAPIVersion()))
}
