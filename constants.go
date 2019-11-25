/*
Copyright (c) 2019 Jo Nilsson

Permission is hereby granted, free of charge, to any person obtaining a
copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
DEALINGS IN THE SOFTWARE.
*/

package go_ms_cfb

//go:generate stringer -type ObjectType
type ObjectType uint8

const (
	maxRegSect uint32 = 0xfffffffa
	difatSect  uint32 = 0xfffffffc
	fatSect    uint32 = 0xfffffffd
	endOfChain uint32 = 0xfffffffe
	freeSect   uint32 = 0xffffffff

	maxRegSID uint32 = 0xfffffffa
	noStream  uint32 = 0xffffffff

	UnknownObject     ObjectType = 0x00
	StorageObject     ObjectType = 0x01
	StreamObject      ObjectType = 0x02
	RootStorageObject ObjectType = 0x05
)
