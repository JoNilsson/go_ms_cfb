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

import (
	"bytes"
	"encoding/binary"
	"io"
	"unicode/utf16"
)

type offsetReader struct {
	r      io.ReaderAt
	offset int64
}

func (o *offsetReader) Read(b []byte) (int, error) {
	return o.r.ReadAt(b, o.offset)
}

func bytesToUint32s(b []byte) ([]uint32, error) {
	var ret []uint32
	r := bytes.NewReader(b)
	for {
		var u uint32
		if err := binary.Read(r, binary.LittleEndian, &u); err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		ret = append(ret, u)
	}
	return ret, nil
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func decodeUTF16String(b []byte) string {
	u := make([]uint16, len(b)/2)
	for i := range u {
		u[i] = (uint16(b[i*2+1]) << 8) | uint16(b[i*2])
	}
	return string(utf16.Decode(u))
}

func isEqualBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
