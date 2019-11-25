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
)

// Unallocated free sectors are marked in the FAT as FREESECT (0xFFFFFFFF). 


type rawFileHeader struct {
	HeaderSignature              [8]byte
	HeaderCLSID                  [16]byte
	MinorVersion                 uint16
	MajorVersion                 uint16
	ByteOrder                    uint16
	SectorShift                  uint16
	MiniSectorShift              uint16
	Reserved                     [6]byte
	NumberOfDirectorySectors     uint32
	NumberOfFATSectors           uint32
	FirstDirectorySectorLocation uint32
	TransactionSignatureNumber   uint32
	MiniStreamCutoffSize         uint32
	FirstMiniFATSectorLocation   uint32
	NumberOfMiniFATSectors       uint32
	FirstDIFATSectorLocation     uint32
	NumberOfDIFATSectors         uint32
	DIFAT                        [109]uint32
}

type FileHeader struct {
	raw rawFileHeader
}

func readFileHeader(r io.Reader) (*FileHeader, error) {
	var h FileHeader
	if err := binary.Read(r, binary.LittleEndian, &h.raw); err != nil {
		return nil, err
	}
	if err := h.validate(); err != nil {
		return nil, err
	}
	return &h, nil
}

func (h *FileHeader) validate() error {
	// TODO: return reason of validation error
	if !isEqualBytes(h.raw.HeaderSignature[:], []byte("\xd0\xcf\x11\xe0\xa1\xb1\x1a\xe1")) {
		return ErrValidation
	} else if !isEqualBytes(h.raw.HeaderCLSID[:], []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")) {
		return ErrValidation
	} else if h.raw.MinorVersion != 0x3e {
		return ErrValidation
	} else if h.raw.MajorVersion != 0x3 && h.raw.MajorVersion != 0x4 {
		return ErrValidation
	} else if h.raw.ByteOrder != 0xfffe {
		return ErrValidation
	} else if h.raw.SectorShift != 0x9 && h.raw.SectorShift != 0xc {
		return ErrValidation
	} else if h.raw.MiniSectorShift != 0x6 {
		return ErrValidation
	} else if !isEqualBytes(h.raw.Reserved[:], []byte("\x00\x00\x00\x00\x00\x00")) {
		return ErrValidation
	} else if h.raw.MiniStreamCutoffSize != 0x1000 {
		return ErrValidation
	}
	return nil
}

func (h *FileHeader) SectorSize() uint32 {
	return 1 << h.raw.SectorShift
}

func (h *FileHeader) MiniSectorSize() uint32 {
	return 1 << h.raw.MiniSectorShift
}
