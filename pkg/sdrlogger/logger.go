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

// It groups functions of the logger API for SoapySDR devices. Implementations
// should use logger rather than stdio. The default log handler prints to stderr

/*
#cgo CFLAGS: -I . -g -Wall
#cgo LDFLAGS: -L . -lSoapySDR

#include <stdlib.h>
#include <stddef.h>
#include <SoapySDR/Logger.h>

// Declare the function logHandlerBridge_cgo, which is implemented in cfuncs.go.
// This declaration allows the function to be used as a parameter by
// SoapySDR_registerLogHandler
void logHandlerBridge_cgo(const SoapySDRLogLevel logLevel, const char *message); // Forward declaration.
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// SDRLogLevel is the priority level for log messages.
//
// The default log level threshold is Info. Log messages with lower priorities are dropped.
//
// The default threshold can be set via the SOAPY_SDR_LOG_LEVEL environment variable.
// Set SOAPY_SDR_LOG_LEVEL to the string value: "WARNING", "ERROR", "DEBUG", etc. or
// set it to the equivalent integer value.
type SDRLogLevel int

const (
	// Fatal represents a fatal error. The application will most likely terminate. This is the highest priority.
	Fatal SDRLogLevel = 1
	// Critical represents a critical error. The application might not be able to continue running successfully.
	Critical SDRLogLevel = 2
	// Error represents an error. An operation did not complete successfully, but the application as a whole is not affected.
	Error SDRLogLevel = 3
	// Warning represents a warning. An operation completed with an unexpected result.
	Warning SDRLogLevel = 4
	// Notice represents a notice, which is an information with just a higher priority.
	Notice SDRLogLevel = 5
	// Info represents an informational message, usually denoting the successful completion of an operation.
	Info SDRLogLevel = 6
	// Debug represents a debugging message.
	Debug SDRLogLevel = 7
	// Trace represents a tracing message. This is the lowest priority.
	Trace SDRLogLevel = 8
	// SSI represents a streaming status indicators such as "U" (underflow) and "O" (overflow).
	SSI SDRLogLevel = 9
)

// Log sends a message to the registered logger.
//
// Params:
//  - logLevel: a possible logging level
//  - message: a logger message string
func Log(level SDRLogLevel, message string) {

	cLevel := C.SoapySDRLogLevel(int(level))

	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))

	C.SoapySDR_log(cLevel, cMessage)
}

// Logf sends a message to the registered logger.
//
// Params:
//  - logLevel: a possible logging level
//  - format: a printf style format string
//  - a: the parameters of the printf function
func Logf(level SDRLogLevel, format string, a ...interface{}) {

	cLevel := C.SoapySDRLogLevel(int(level))

	// Cgo does not support variadic arguments, so the format and print is done locally before
	// sending the formatted string to the logger
	cFormat := C.CString(fmt.Sprintf(format, a...))
	defer C.free(unsafe.Pointer(cFormat))

	C.SoapySDR_log(cLevel, cFormat)
}

// Keep track of the current handler
var currentLogHandler func(level SDRLogLevel, message string)

// logHandlerBridge is the function that is called as a call back by Soapy for logging.
// This function is C-exported so it can be used from cfuncs.go
//
// Params:
//  - logLevel: a possible logging level
//  - message: a logger message string
//
//export logHandlerBridge
func logHandlerBridge(level C.SoapySDRLogLevel, message *C.char) {
	if currentLogHandler != nil {
		currentLogHandler(SDRLogLevel(level), C.GoString(message))
	}
}

// RegisterLogHandler registers a new system log handler.
//
// Platforms should call this to replace the default stdio handler.
//
// Params:
//  - logHandler: the function that will receive the log. Passing nil restores the default.
func RegisterLogHandler(logHandler func(level SDRLogLevel, message string)) {

	// Keep track of the current log handler
	currentLogHandler = logHandler

	// Clean the current logHandler defined is Soapy
	C.SoapySDR_registerLogHandler(C.SoapySDRLogHandler(unsafe.Pointer(nil)))

	// Inform the Soapy layer to use the bridge (if a local log handler is defined)
	if logHandler != nil {
		C.SoapySDR_registerLogHandler(C.SoapySDRLogHandler(unsafe.Pointer(C.logHandlerBridge_cgo)))
	}
}

// SetLogLevel sets the log level threshold. Log messages with lower priority are dropped.
//
// Params:
//  - level: the minimum log level
func SetLogLevel(level SDRLogLevel) {

	cLevel := C.SoapySDRLogLevel(int(level))

	C.SoapySDR_setLogLevel(cLevel)
}
