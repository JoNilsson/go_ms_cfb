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

import "io"

type SectorReader struct {
	r              io.ReaderAt
	sectorSize     uint32
	startSector    uint32
	offset         int64
	maxOffset      int64
	sectors        []uint32
	offsetResolver func(uint32) int64
}

func newSectorReader(r io.ReaderAt, sectorSize, startSector uint32, fat []uint32, offsetResolver func(uint32) int64) (*SectorReader, error) {
	if startSector < 0 {
		return nil, ErrWrongSector
	} else if s := fat[startSector]; s > maxRegSect && s != endOfChain {
		return nil, ErrInvalidSectorChain
	}

	var sectors []uint32
	s := startSector
	for s != endOfChain {
		sectors = append(sectors, s)
		s = fat[s]
	}

	return &SectorReader{r, sectorSize, startSector, 0, int64(sectorSize) * int64(len(sectors)), sectors, offsetResolver}, nil
}

func (sr *SectorReader) ReadAt(b []byte, offset int64) (int, error) {
	if offset < 0 {
		return 0, ErrInvalidOffset
	} else if offset >= sr.maxOffset {
		return 0, io.EOF
	}

	eof := false
	if max := sr.maxOffset - offset; int64(len(b)) > max {
		b = b[:max]
		eof = true
	}

	read := 0
	for read < len(b) {
		if offset >= sr.maxOffset {
			return read, io.EOF
		}
		sectorOffset := int(offset % int64(sr.sectorSize))

		o := sr.offsetResolver(sr.sectors[offset/int64(sr.sectorSize)])

		data := make([]byte, sr.sectorSize)
		if n, err := sr.r.ReadAt(data, o); err != nil {
			return 0, err
		} else if n != int(sr.sectorSize) {
			return 0, ErrInsufficientData
		}

		limit := min(len(b)-read, int(sr.sectorSize)-sectorOffset)
		copy(b[read:read+limit], data[sectorOffset:sectorOffset+limit])
		read += limit
		offset += int64(limit)
	}

	var err error
	if eof {
		err = io.EOF
	}

	return read, err
}

func (sr *SectorReader) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	default:
		return 0, ErrInvalidSeek
	case io.SeekStart:
	case io.SeekCurrent:
		offset += sr.offset
	case io.SeekEnd:
		offset += sr.maxOffset
	}

	if offset < 0 || offset > sr.maxOffset {
		return 0, ErrInvalidOffset
	}

	sr.offset = offset
	return offset, nil
}

func (sr *SectorReader) Read(b []byte) (int, error) {
	if sr.offset >= sr.maxOffset {
		return 0, io.EOF
	}
	// TODO: do we really need this code below?
	if max := sr.maxOffset - sr.offset; int64(len(b)) > max {
		b = b[:max]
	}
	n, err := sr.ReadAt(b, sr.offset)
	if err != nil { // TODO: check if err is io.EOF
		return 0, err
	}
	sr.offset += int64(n)
	return n, nil
}
