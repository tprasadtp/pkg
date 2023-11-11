// SPDX-FileCopyrightText: Copyright 2023 Prasad Tengse
// SPDX-License-Identifier: MIT

// Package guid provides utilities to marshal and unmarshal GUIDs.
//
// This representation of GUID is compatible with [golang.org/x/sys/windows]
// and can be used wherever syscall interface/func expects [golang.org/x/sys/windows.GUID].
// Unlike [github.com/google/uuid], encoding is always little endian.
package guid

import (
	"bytes"
	"crypto/rand"
	"encoding"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_ encoding.TextMarshaler   = (*GUID)(nil)
	_ encoding.TextUnmarshaler = (*GUID)(nil)
	_ json.Marshaler           = (*GUID)(nil)
	_ json.Unmarshaler         = (*GUID)(nil)
)

// Generate new [RFC 4122] v4 GUID.
func NewGUID() GUID {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(fmt.Sprintf("log(guid): failed to generate random bytes: %s", err))
	}

	var g GUID
	g.Data1 = binary.LittleEndian.Uint32(b[0:4])
	g.Data2 = binary.LittleEndian.Uint16(b[4:6])
	g.Data3 = binary.LittleEndian.Uint16(b[6:8])
	g.Data3 = (g.Data3 & 0x0fff) | (uint16(4) << 12) // set UUID version to v4
	copy(g.Data4[:], b[8:16])
	g.Data4[0] = (g.Data4[0] & 0x3f) | 0x80 // set type to RFC 4122
	return g
}

// IsZero returns true if GUID is empty.
func (g GUID) IsZero() bool {
	return g.Data1 == 0 && g.Data2 == 0 && g.Data3 == 0 && bytes.Equal(g.Data4[:], make([]byte, 8))
}

// String formats GUID into `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` format.
func (g GUID) String() string {
	buf := make([]byte, 0, 36)
	return string(g.AppendString(buf))
}

// AppendString appends hexadecimal encoded string representation of GUID to buf.
func (g GUID) AppendString(buf []byte) []byte {
	guidBytes := make([]byte, 0, 16)
	guidBytes = binary.LittleEndian.AppendUint32(guidBytes, g.Data1)
	guidBytes = binary.LittleEndian.AppendUint16(guidBytes, g.Data2)
	guidBytes = binary.LittleEndian.AppendUint16(guidBytes, g.Data3)
	guidBytes = append(guidBytes, g.Data4[:]...)

	hexBuf := make([]byte, 36)

	hex.Encode(hexBuf[:8], guidBytes[:4])
	hexBuf[8] = '-'

	hex.Encode(hexBuf[9:13], guidBytes[4:6])
	hexBuf[13] = '-'

	hex.Encode(hexBuf[14:18], guidBytes[6:8])
	hexBuf[18] = '-'

	hex.Encode(hexBuf[19:23], guidBytes[8:10])
	hexBuf[23] = '-'

	hex.Encode(hexBuf[24:], guidBytes[10:])

	return append(buf, hexBuf...)
}

// MarshalText returns the text representation of the GUID.
func (g GUID) MarshalText() ([]byte, error) {
	return []byte(g.String()), nil
}

// UnmarshalText takes the text representation of a GUID, and unmarshal it into this GUID.
func (g *GUID) UnmarshalText(text []byte) error {
	v, err := ParseGUID(text)
	if err != nil {
		return err
	}
	*g = v
	return nil
}

// MarshalText returns the text representation of the GUID.
func (g GUID) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, g.String()), nil
}

// UnmarshalJSON takes the json representation of a GUID, and unmarshal it into this GUID.
func (g *GUID) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return fmt.Errorf("invalid GUID %q", s)
	}
	v, err := ParseGUID(s)
	if err != nil {
		return err
	}
	*g = v
	return nil
}

// ParseGUID parses a string/byte slice containing a GUID and returns the GUID.
// The following formats are currently supported are
//   - `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`
//   - `urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`
func ParseGUID[T ~string | ~[]byte](input T) (GUID, error) {
	b := bytes.TrimPrefix([]byte(input), []byte("urn:uuid:"))

	// Hex encoded 16 bytes is 32 bits + 4 for - separators.
	if len(b) != 36 {
		return GUID{}, fmt.Errorf("invalid GUID %q", b)
	}

	// Ensure separators are at correct positions.
	if b[8] != '-' || b[13] != '-' || b[18] != '-' || b[23] != '-' {
		return GUID{}, fmt.Errorf("invalid GUID %q", b)
	}

	guidBytes := make([]byte, 16)
	_, err := hex.Decode(guidBytes[:4], b[:8])
	if err != nil {
		return GUID{}, fmt.Errorf("invalid GUID(Data1): %w", err)
	}

	_, err = hex.Decode(guidBytes[4:6], b[9:13])
	if err != nil {
		return GUID{}, fmt.Errorf("invalid GUID(Data2): %w", err)
	}

	_, err = hex.Decode(guidBytes[6:8], b[14:18])
	if err != nil {
		return GUID{}, fmt.Errorf("invalid GUID(Data3): %w", err)
	}

	_, err = hex.Decode(guidBytes[8:10], b[19:23])
	if err != nil {
		return GUID{}, fmt.Errorf("invalid GUID(Data4a): %w", err)
	}

	_, err = hex.Decode(guidBytes[10:], b[24:])
	if err != nil {
		return GUID{}, fmt.Errorf("invalid GUID(Data4b): %w", err)
	}

	g := GUID{
		Data1: binary.LittleEndian.Uint32(guidBytes[:4]),
		Data2: binary.LittleEndian.Uint16(guidBytes[4:6]),
		Data3: binary.LittleEndian.Uint16(guidBytes[6:8]),
	}
	copy(g.Data4[:], guidBytes[8:])

	return g, nil
}

// MustParseGUID parses input as GUID, but upon errors panics.
func MustParseGUID[T ~string | ~[]byte](input T) GUID {
	v, err := ParseGUID(input)
	if err != nil {
		panic(err)
	}
	return v
}
