package modules

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

// #include <SoapySDR/Types.h>
import "C"
import "unsafe"

// stringArray2Go converts an array of C string to an array of Go String
func stringArray2Go(strings **C.char, length C.size_t) []string {

	results := make([]string, int(length))
	var charPtrTemplate *C.char

	// Read all the strings
	for i := 0; i < int(length); i++ {
		val := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(strings)) + uintptr(i)*unsafe.Sizeof(charPtrTemplate)))
		results[i] = C.GoString(*val)
	}

	return results
}

// args2Go converts a single C Args to Go Arg
func args2Go(args C.SoapySDRKwargs) map[string]string {

	results := make(map[string]string, args.size)

	keys := (**C.char)(unsafe.Pointer(args.keys))
	vals := (**C.char)(unsafe.Pointer(args.vals))

	// Read all the strings
	for i := 0; i < int(args.size); i++ {
		key := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(keys)) + uintptr(i)*unsafe.Sizeof(*keys)))
		val := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(vals)) + uintptr(i)*unsafe.Sizeof(*vals)))
		results[C.GoString(*key)] = C.GoString(*val)
	}

	return results
}
