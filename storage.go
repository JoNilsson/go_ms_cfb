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
	"fmt"
	"io"
)

type Storage struct {
	d *DirectoryEntry
}

func (d *DirectoryEntry) storage() (*Storage, error) {
	if d.Type() != StorageObject {
		return nil, ErrWrongObjectType
	}
	return &Storage{d}, nil
}

func (s *Storage) String() string {
	return fmt.Sprintf("Storage{Path:%q}", s.Path())
}

func (s *Storage) Name() string {
	return s.d.Name()
}

func (s *Storage) Path() string {
	return s.d.Path()
}

func (s *Storage) Type() ObjectType {
	return StorageObject
}

func (s *Storage) Size() uint64 {
	return 0
}

func (s *Storage) ReadAt(b []byte, offset int64) (int, error) {
	return 0, io.EOF
}

func (s *Storage) Seek(offset int64, whence int) (int64, error) {
	// TODO: error check
	return 0, nil
}

func (s *Storage) Read(b []byte) (int, error) {
	return 0, io.EOF
}
