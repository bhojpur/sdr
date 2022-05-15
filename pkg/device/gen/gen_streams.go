//go:build ignore
// +build ignore

package main

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

// This program generates streams.go. It can be invoked by running `go generate`

import "C"
import (
	"log"
	"os"
	"text/template"
	"time"
	"unsafe"
)

func main() {

	type Detail struct {
		StreamObjectName string
		SoapyFormat      string
		GoType           string
		CType            string
		DataSize         uint
	}

	var TemplateCU8 C.uchar
	var TemplateCS8 C.uchar
	var TemplateCU16 C.uint
	var TemplateCS16 C.int
	var TemplateCF32 C.complexfloat
	var TemplateCF64 C.complexdouble

	details := []Detail{
		{"SDRStreamCU8", "CU8", "uint8", "C.uchar", uint(unsafe.Sizeof(TemplateCU8))},
		{"SDRStreamCS8", "CS8", "int8", "C.char", uint(unsafe.Sizeof(TemplateCS8))},
		{"SDRStreamCU16", "CU16", "uint16", "C.uint", uint(unsafe.Sizeof(TemplateCU16))},
		{"SDRStreamCS16", "CS16", "int16", "C.int", uint(unsafe.Sizeof(TemplateCS16))},
		{"SDRStreamCF32", "CF32", "complex64", "C.complexfloat", uint(unsafe.Sizeof(TemplateCF32))},
		{"SDRStreamCF64", "CF64", "complex128", "C.complexdouble", uint(unsafe.Sizeof(TemplateCF64))},
	}

	f, err := os.Create("streams.go")
	die(err)
	defer f.Close()

	builderTemplate.Execute(
		f,
		struct {
			Timestamp time.Time
			Details   []Detail
		}{
			Timestamp: time.Now(),
			Details:   details,
		})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var builderTemplate = template.Must(template.New("").Parse(`// Code generated by Bhojpur SDR (go generate); DO NOT EDIT.
// This file was generated by gen_streams.go at {{ .Timestamp }}

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

// It regroups all the functions for accessing devices and streams.

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
	"github.com/bhojpur/sdr/pkg/sdrerror"
	"unsafe"
)

/* ********************************************************************************** */
/*                                                                                    */
/*                       DEFINITIONS OF STRUCTURES                                    */
/*                                                                                    */
/* ********************************************************************************** */
{{ range .Details }}
// {{ .StreamObjectName }} is a stream for accessing data in {{ .SoapyFormat }} format
type {{ .StreamObjectName }} struct {
	device         *C.SoapySDRDevice
	stream         *C.SoapySDRStream
	nbChannels     uint
	readBuffer     **C.void
	writeBuffer    **C.void
}
{{ end }}

{{ range .Details }}
/* ********************************************************************************** */
/*                                                                                    */
/*                     FUNCTIONS OF STREAMS {{ .StreamObjectName }}                             */
/*                                                                                    */
/* ********************************************************************************** */

/* ********************************************************************************** */
/*                          CREATION OF STREAMS                                       */
/* ********************************************************************************** */

// Setup{{ .StreamObjectName }} initializes a stream given a list of channels and stream arguments.
//
// The implementation may change switches or power-up components.
// All stream API calls should be usable with the new stream object
// after Setup{{ .StreamObjectName }}() is complete, regardless of the activity state.
//
// The API allows any number of simultaneous TX and RX streams, but many dual-channel
// devices are limited to one stream in each direction, using either one or both channels.
// This call will return an error if an unsupported combination is requested,
// or if a requested channel in this direction is already in use by another stream.
//
// When multiple channels are added to a stream, they are typically expected to have
// the same sample rate. See SetSampleRate().
//
// Params:
//  - direction: the channel direction ('DirectionRX' or 'DirectionTX')
//  - channels: a list of channels or empty for automatic. When multiple channels are added to a stream, they are 
//    typically expected to have the same sample rate. See SetSampleRate(). Warning: Contrary to SoapySDR API, the 
//    channels must be explicitly defined. Hence the channels slice can not be given empty.
//  - args: stream args or empty for defaults
//
// Args:
// Recommended keys to use in the args dictionary:
//   - "WIRE" - format of the samples between device and host
//
// Return the stream pointer and an error. The returned stream is not required to have internal locking,
// and may not be used concurrently from multiple threads.
func (dev *SDRDevice) Setup{{ .StreamObjectName }}(direction Direction, channels []uint, args map[string]string) (stream *{{ .StreamObjectName }}, err error) {

	if len(channels) == 0 {
		return nil, errors.New("the channels must be given explicitly during stream setup")
	}

	cFormat := C.CString("{{ .SoapyFormat }}")
	defer C.free(unsafe.Pointer(cFormat))

	cArgs, cArgsLength := go2Args(args)
	defer argsListClear(cArgs, cArgsLength)

	cChannels, cChannelsLength := go2SizeTList(channels)
	defer C.free(unsafe.Pointer(cChannels))

	val := C.SoapySDRDevice_setupStream(dev.device, C.int(direction), cFormat, cChannels, cChannelsLength, cArgs)

	if val == nil {
		return nil, errors.New(LastError())
	}

	var voidPtrTemplate *C.void

	nbChannels := uint(len(channels))
	
	// Allocate the buffers
	readBuffers := (**C.void)(C.malloc(C.size_t(nbChannels * uint(unsafe.Sizeof(voidPtrTemplate)))))
	writeBuffers := (**C.void)(C.malloc(C.size_t(nbChannels * uint(unsafe.Sizeof(voidPtrTemplate)))))

	return &{{ .StreamObjectName }}{
		device:      dev.device,
		stream:      val,
		nbChannels:  nbChannels,
		readBuffer:  readBuffers,
		writeBuffer: writeBuffers,
	}, nil
}

/* ********************************************************************************** */
/*                                GETTER AND SETTER                                   */            
/* ********************************************************************************** */
// getDevice returns the internal device  
func (stream *{{ .StreamObjectName }}) getDevice() *C.SoapySDRDevice {
	return stream.device
}

// getStream returns the internal stream
func (stream *{{ .StreamObjectName }}) getStream() *C.SoapySDRStream {
	return stream.stream
}

// getNbChannels returns the number of channels used by the stream
func (stream *{{ .StreamObjectName }}) getNbChannels() uint {
	return stream.nbChannels
}

/* ********************************************************************************** */
/*                                STREAMS FUNCTIONS                                   */            
/* ********************************************************************************** */

// Close closes an open stream created by setupStream
//
// Params:
//  - stream: the opaque pointer to a stream handle
//
// Return an error or nil in case of success
func (stream *{{ .StreamObjectName }}) Close() (err sdrerror.SDRError) {

	// Free the buffers
	C.free(unsafe.Pointer(stream.readBuffer))
	C.free(unsafe.Pointer(stream.writeBuffer))
	// Set the buffers to nil, in case someone try to reuse the stream
	stream.readBuffer = nil
	stream.writeBuffer = nil

	return sdrerror.Err(int(C.SoapySDRDevice_closeStream(stream.device, stream.stream)))
}

// GetMTU gets the stream's maximum transmission unit (MTU) in number of elements.
//
// The MTU specifies the maximum payload transfer in a stream operation. This value can be used as a stream buffer
// allocation size that can best optimize throughput given the underlying stream implementation.
//
// Return the MTU in number of stream elements (never zero)
func (stream *{{ .StreamObjectName }}) GetMTU() int {

	return int(C.SoapySDRDevice_getStreamMTU(stream.device, stream.stream))
}

// Activate activates a stream.
//
// Call activate to prepare a stream before using read/write(). The implementation control switches or stimulate data
// flow.
//
// Params:
//  - flags: optional flag indicators about the stream. The StreamFlagEndBurst flag can signal end on the finite burst.
//    Not all implementations will support the full range of options. In this case, the implementation returns
//    ErrorNotSupported.
//  - timeNs: optional activation time in nanoseconds. The timeNs is only valid when the flags have StreamFlagHasTime.
//  - numElems: optional element count for burst control. The numElems count can be used to request a finite burst size.
//
// Return an error or nil in case of success
func (stream *{{ .StreamObjectName }}) Activate(flags StreamFlag, timeNs int, numElems int) (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_activateStream(stream.device, stream.stream, C.int(flags), C.longlong(timeNs), C.size_t(numElems))))
}

// Deactivate deactivates a stream.
//
// Call deactivate when not using using read/write(). The implementation control switches or halt data flow.
//
// Params:
//  - flags: optional flag indicators about the stream. Not all implementations will support the full range of options.
//    In this case, the implementation returns ErrorNotSupported.
//  - timeNs: optional deactivation time in nanoseconds. The timeNs is only valid when the flags have StreamFlagHasTime.
//
// Return an error or nil in case of success
func (stream *{{ .StreamObjectName }}) Deactivate(flags StreamFlag, timeNs int) (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_deactivateStream(stream.device, stream.stream, C.int(flags), C.longlong(timeNs))))
}

// GetNumDirectAccessBuffers returns how many direct access buffers can the stream provide.
//
// This is the number of times the user can call acquire() on a stream without making subsequent calls to
// release(). A return value of 0 means that direct access is not supported.
//
// Return the number of direct access buffers or 0
func (stream *{{ .StreamObjectName }}) GetNumDirectAccessBuffers() uint {

	return getNumDirectAccessBuffers(stream)
}

/* ********************************************************************************** */
/*                                READ WRITE FUNCTIONS                                */            
/* ********************************************************************************** */

// Read reads elements from a stream for reception. The elements are written in the given buffer which must be allocated
// before call.
//
//
// This is a multi-channel call, and buffs should be a slice of slice of {{ .GoType }}, where each slice of {{ .GoType }} will
// be filled with data from a different channel.
//
// Params:
// Params:
//  - buffs: an array of buffers num chans in size. The number of buffers must match the number of channels of the 
//    stream. The buffers MUST already be fully allocated before the call.
//  - nbElems: the number of data to read. Note that the buffer must be large enough to hold the data. For example
//    complex data stored in non complex buffer (such as CS8) will use 2 elements of the buffer for 1 single data.
//  - outputFlags: The flag indicators of the result by channel. The number of flags must match the number of channels 
//    of the stream.
//  - timeoutUs: the timeout in microseconds
//
// Return the flag indicators about the result by channel, the buffer's timestamp in nanoseconds, the number of elements read per buffer, the errorCode and the error message
func (stream *{{ .StreamObjectName }}) Read(buffers [][]{{ .GoType }}, nbElems uint, outputFlags []int, timeoutUs uint) (timeNs uint, numElemsRead uint, err error) {

	if uint(len(buffers)) != stream.nbChannels {
		return 0, 0, errors.New("the read buffer must have the same number of channels as the stream")
	}

	if uint(len(outputFlags)) != stream.nbChannels {
		return 0, 0, errors.New("the flags must have the same number of channels as the stream")
	}

	var voidPtrTemplate *C.void

	// Convert the given buffers to C pointers 
	for channelIdx := uint(0); channelIdx < stream.nbChannels; channelIdx++ {

		// Get the pointer to the buffer for the channel
		ptrPtrBuffer := (**C.void)(unsafe.Pointer(uintptr(unsafe.Pointer(stream.readBuffer)) + uintptr(channelIdx)*unsafe.Sizeof(voidPtrTemplate)))
		*ptrPtrBuffer = (*C.void)(unsafe.Pointer(&buffers[channelIdx][0]))		
	}

	cFlags := (*C.int)(unsafe.Pointer(&outputFlags[0]))
	cTimeNs := C.longlong(0)

	// Make the actual read
	result := int(
		C.SoapySDRDevice_readStream(
			stream.device,
			stream.stream,
			(*unsafe.Pointer)(unsafe.Pointer(stream.readBuffer)),
			C.size_t(nbElems),
			cFlags,
			&cTimeNs,
			C.long(timeoutUs)))

	if result < 0 {
		return uint(cTimeNs), 0, sdrerror.Err(int(result))
	}

	return uint(cTimeNs), uint(result), nil
}

// WriteStream writes elements to a stream for transmission.
//
// This is a multi-channel call, and buffs should be a slice of slice of {{ .GoType }}, where each slice of {{ .GoType }}
// is sent to a channel.
//
// Params:
//  - buffs: an array of void* buffers num chans in size. The number of buffers must match the number of channels of the 
//    stream.
//  - nbElems: the number of data to write. Note that the buffer must be large enough to hold the data. For example
//    complex data stored in non complex buffer (such as CS8) will use 2 elements of the buffer for 1 single data.
//  - flags: input flags, may be updated with the value of the output flags (device specific). The number of flags must
//    match the number of channels of the stream.
//  - timeNs: the buffer's timestamp in nanoseconds
//  - timeoutUs: the timeout in microseconds
//
// Return the number of elements written per buffer or 0 in case of an error (even if some data were sent before the 
// error)
func (stream *{{ .StreamObjectName }}) Write(buffers [][]{{ .GoType }}, nbElems uint, flags []int, timeNs uint, timeoutUs uint) (NbElemsWritten uint, err error) {

	if uint(len(buffers)) != stream.nbChannels {
		return 0, errors.New("the write buffer must have the same number of channels as the stream")
	}

	if uint(len(flags)) != stream.nbChannels {
		return 0, errors.New("the write flags must have the same number of channels as the stream")
	}

	var voidPtrTemplate *C.void

	// Convert the given buffer to C 
	for channelIdx := uint(0); channelIdx < stream.nbChannels; channelIdx++ {

		// Get the pointer to the buffer for the channel
		ptrPtrBuffer := (**C.void)(unsafe.Pointer(uintptr(unsafe.Pointer(stream.readBuffer)) + uintptr(channelIdx)*unsafe.Sizeof(voidPtrTemplate)))
		*ptrPtrBuffer = (*C.void)(unsafe.Pointer(&buffers[channelIdx][0]))
	}

	cFlags := (*C.int)(unsafe.Pointer(&flags[0]))

	// Make the actual write
	result := int(
		C.SoapySDRDevice_writeStream(
			stream.device,
			stream.stream,
			(*unsafe.Pointer)(unsafe.Pointer(stream.writeBuffer)),
			C.size_t(nbElems), 
			cFlags, 
			C.longlong(timeNs), 
			C.long(timeoutUs)))

	if result < 0 {
		return 0, sdrerror.Err(int(result))
	}

	return uint(result), nil
}

// ReadStreamStatus reads status information about a stream.
//
// This call is typically used on a transmit stream to report time errors, underflows, and burst completion.
//
// Client code may continually poll readStreamStatus() in a loop. Implementations of readStreamStatus() should wait in
// the call for a status change event or until the timeout expiration. When stream status is not implemented on a
// particular stream, readStreamStatus() should return SOAPY_SDR_NOT_SUPPORTED. Client code may use this indication to
// disable a polling loop.
//
// Params:
//  - chanMask to which channels this status applies
//  - flags optional input flags and output flags
//  - timeNs the buffer's timestamp in nanoseconds
//  - timeoutUs the timeout in microseconds
//
// Return the buffer's timestamp in nanoseconds in case of success, an error otherwise
func (stream *{{ .StreamObjectName }}) ReadStreamStatus(chanMask []uint, flags []int, timeoutUs uint) (timeNs uint, err error) {

	return readStreamStatus(stream, chanMask, flags, timeoutUs)
}
{{ end }}

`))
