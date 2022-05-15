package sdrlogger

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

/*
#include <stdio.h>
#include <SoapySDR/Logger.h>

// C code can call exported Go functions with their explicit name. But, if a
// C-program wants a function pointer, a gateway function has to be written.
// This is because we can't take the address of a Go function and give that to
// C-code since the CGO tool will generate a stub in C that should be called.
// The following example shows how to integrate with C code wanting a function
// pointer of a give type.
// In logger.go:
//   - forward declaration of logHandlerBridge_cgo, so it can be used as a
//		parameter to SoapySDR_registerLogHandler
//   - export of function logHandlerBridge, that is receiving the log data ultimately
// In cfuncs.go:
//   - implementation of logHandlerBridge_cgo, that is calling logHandlerBridge

// The gateway function, written in C. this function be used as a call back by
// SoapySDR_registerLogHandler. The Gateway function is simply calling the Go
// function(logHandlerBridge) with the same parameters
void logHandlerBridge_cgo(const SoapySDRLogLevel logLevel, const char *message)
{
	// printf("C.logHandlerBridge_cgo(): called with arg = %d, str = %s\n", logLevel, message);

	// Declare locally the function exported from logger.go
	void logHandlerBridge(const SoapySDRLogLevel, const char *);

	// Call the function in logger.go
	return logHandlerBridge(logLevel, message);
}
*/
import "C"
