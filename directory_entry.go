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
	"encoding/binary"
	"io"
	"strings"
)

type rawDirectoryEntry struct {
	DirectoryEntryName       [64]byte
	DirectoryEntryNameLength uint16
	ObjectType               ObjectType
	ColorFlag                uint8
	LeftSiblingID            uint32
	RightSiblingID           uint32
	ChildID                  uint32
	CLSID                    [16]byte
	StateBits                uint32
	CreationTime             uint64
	ModifiedTime             uint64
	StartingSectorLocation   uint32
	StreamSize               uint64
}

type DirectoryEntry struct {
	f *File

	id       int
	children []*DirectoryEntry
	path     []string

	raw rawDirectoryEntry
}

func readDirectoryEntry(r io.Reader) (*DirectoryEntry, error) {
	var d DirectoryEntry
	if err := binary.Read(r, binary.LittleEndian, &d.raw); err != nil {
		return nil, err
	}
	if err := d.validate(); err != nil {
		return nil, err
	}
	return &d, nil
}

func (d *DirectoryEntry) validate() error {
	if strings.ContainsAny(d.Name(), "/\\:!") {
		return ErrValidation
	}
	return nil
}

func (d *DirectoryEntry) Name() string {
	if d.raw.DirectoryEntryNameLength < 2 {
		return ""
	}
	return decodeUTF16String(d.raw.DirectoryEntryName[:d.raw.DirectoryEntryNameLength-2])
}

func (d *DirectoryEntry) Path() string {
	return strings.Join(d.path, "/")
}

func (d *DirectoryEntry) Type() ObjectType {
	return d.raw.ObjectType
}

func (d *DirectoryEntry) Size() uint64 {
	return d.raw.StreamSize
}

func (d *DirectoryEntry) StartingSector() uint32 {
	return d.raw.StartingSectorLocation
}

func (d *DirectoryEntry) object() (Object, error) {
	switch d.Type() {
	default:
		return nil, ErrInvalidObject
	case StorageObject:
		return d.storage()
	case StreamObject:
		return d.stream()
	}
}
