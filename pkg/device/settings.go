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

// GetSettingInfo describes the allowed keys and values used for settings.
//
// Return a list of argument info structures
func (dev *SDRDevice) GetSettingInfo() []SDRArgInfo {

	length := C.size_t(0)

	info := C.SoapySDRDevice_getSettingInfo(dev.device, &length)
	defer argInfoListClear(info, length)

	return argInfoList2Go(info, length)
}

// WriteSetting writes an arbitrary setting on the device.
//
// The interpretation is up the implementation.
//
// Params:
//  - key: the setting identifier
//  - value: the setting value
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteSetting(key string, value string) (err sdrerror.SDRError) {

	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	return sdrerror.Err(int(C.SoapySDRDevice_writeSetting(dev.device, cKey, cValue)))
}

// Read an arbitrary setting on the device.
// \param device a pointer to a device instance
// \param key the setting identifier
// \return the setting value
//SOAPY_SDR_API char *SoapySDRDevice_readSetting(const SoapySDRDevice *device, const char *key);

// ReadSetting reads an arbitrary setting on the device.
//
// The interpretation is up the implementation.
//
// Params:
//  - key: the setting identifier
//
// Return the setting value
func (dev *SDRDevice) ReadSetting(key string) string {

	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	val := (*C.char)(C.SoapySDRDevice_readSetting(dev.device, cKey))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}

// GetChannelSettingInfo describes the allowed keys and values used for channel settings.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return a list of argument info structures
func (dev *SDRDevice) GetChannelSettingInfo(direction Direction, channel uint) []SDRArgInfo {

	cDirection := C.int(direction)
	cChannel := C.size_t(channel)
	length := C.size_t(0)

	info := C.SoapySDRDevice_getChannelSettingInfo(dev.device, cDirection, cChannel, &length)
	defer argInfoListClear(info, length)

	return argInfoList2Go(info, length)
}

// WriteChannelSetting writes an arbitrary channel setting on the device.
//
// The interpretation is up the implementation.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - key: the setting identifier
//  - value: the setting value
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteChannelSetting(direction Direction, channel uint, key string, value string) (err sdrerror.SDRError) {

	cDirection := C.int(direction)
	cChannel := C.size_t(channel)

	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	return sdrerror.Err(int(C.SoapySDRDevice_writeChannelSetting(dev.device, cDirection, cChannel, cKey, cValue)))
}

// ReadChannelSetting an arbitrary channel setting on the device.
//
// The interpretation is up the implementation.
//
// Params:
//  - key: the setting identifier
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return the setting value
func (dev *SDRDevice) ReadChannelSetting(direction Direction, channel uint, key string) string {

	cDirection := C.int(direction)
	cChannel := C.size_t(channel)

	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	val := (*C.char)(C.SoapySDRDevice_readChannelSetting(dev.device, cDirection, cChannel, cKey))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}
