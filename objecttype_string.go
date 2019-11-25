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

// Code below generated by "stringer -type ObjectType"; DO NOT EDIT.

package go_ms_cfb

import "strconv"

const (
	_ObjectType_name_0 = "UnknownObjectStorageObjectStreamObject"
	_ObjectType_name_1 = "RootStorageObject"
)

var (
	_ObjectType_index_0 = [...]uint8{0, 13, 26, 38}
)

func (i ObjectType) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _ObjectType_name_0[_ObjectType_index_0[i]:_ObjectType_index_0[i+1]]
	case i == 5:
		return _ObjectType_name_1
	default:
		return "ObjectType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
