// SPDX-FileCopyrightText: Copyright 2023 Prasad Tengse
// SPDX-License-Identifier: MIT
package guid_test

import (
	"bytes"
	"testing"

	"github.com/tprasadtp/pkg/assert"
	"github.com/tprasadtp/pkg/guid"
	"github.com/tprasadtp/pkg/race"
)

const (
	zeroGUID = "00000000-0000-0000-0000-000000000000"
)

func TestNewGUID(t *testing.T) {
	t.Run("Version", func(t *testing.T) {
		v := guid.NewGUID()
		version := v.Data3 & 0xF000 >> 12
		if version != 4 {
			t.Errorf("version should be (4)")
		}
	})
	t.Run("Variant", func(t *testing.T) {
		v := guid.NewGUID()
		variant := v.Data4[0] & 0xc0
		if variant != 0x80 {
			t.Errorf("variant should be VariantRFC4122")
		}
	})
	t.Run("NotZero", func(t *testing.T) {
		v := guid.NewGUID()
		if v.IsZero() {
			t.Errorf("NewGUID is zero")
		}
	})

	t.Run("String-Allocs", func(t *testing.T) {
		if race.Enabled {
			t.Skipf("%s => skipping allocation tests in race mode", t.Name())
		}

		v := guid.NewGUID()
		var str string
		allocs := testing.AllocsPerRun(10, func() {
			str = v.String()
		})
		if allocs > 1 {
			t.Errorf("GUID.String() allocates > 1: %f", allocs)
		}
		_ = str
	})
}

func TestGUID_Encoding(t *testing.T) {
	t.Run("windows", func(t *testing.T) {
		expect := "6bb6f6f2-8a38-42c4-868f-be3a285b33a7"
		v := guid.MustParseGUID(expect)
		if v.String() != expect {
			t.Errorf("expected=%s, got=%s", expect, v.String())
		}
	})
	t.Run("linux", func(t *testing.T) {
		expect := "edf2af85-5405-414b-a0af-714bff56386c"
		v := guid.MustParseGUID(expect)
		if v.String() != expect {
			t.Errorf("expected=%s, got=%s", expect, v.String())
		}
	})
}

func TestGUID(t *testing.T) {
	t.Run("valid-string", func(t *testing.T) {
		v := guid.MustParseGUID("93b2c683-d8af-4cd0-8ca2-bf8177871a3b")
		expected := "93b2c683-d8af-4cd0-8ca2-bf8177871a3b"
		if s := v.String(); s != "93b2c683-d8af-4cd0-8ca2-bf8177871a3b" {
			t.Errorf("expected=%s, got=%s", expected, s)
		}
	})
	t.Run("zero-string", func(t *testing.T) {
		v := guid.MustParseGUID(zeroGUID)
		if s := v.String(); s != zeroGUID {
			t.Errorf("expected=%s, got=%s", zeroGUID, s)
		}
		if !v.IsZero() {
			t.Errorf("not zero")
		}
	})

	t.Run("valid-unmarshal-text", func(t *testing.T) {
		data := []byte(`93b2c683-d8af-4cd0-8ca2-bf8177871a3b`)
		expected := "93b2c683-d8af-4cd0-8ca2-bf8177871a3b"

		v := guid.GUID{}
		err := v.UnmarshalText(data)

		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		if v.String() != expected {
			t.Errorf("expected=%s, got=%s", expected, v.String())
		}
	})

	t.Run("invalid-too-short-quotes-unmarshal-text", func(t *testing.T) {
		data := []byte(`93b2c683-d8af-4cd0-8ca2-bf8177871a3`)

		v := guid.GUID{}
		err := v.UnmarshalText(data)

		if err == nil {
			t.Errorf("expected error: %s", err)
		}

		if v.String() != zeroGUID {
			t.Errorf("expected=%s, got=%s", zeroGUID, v.String())
		}
	})

	t.Run("valid-marshal-txt", func(t *testing.T) {
		v := guid.MustParseGUID("93b2c683-d8af-4cd0-8ca2-bf8177871a3b")
		expected := []byte(`93b2c683-d8af-4cd0-8ca2-bf8177871a3b`)
		s, err := v.MarshalText()

		if err != nil {
			t.Errorf("MarshalText() should never return error")
		}

		if !bytes.Equal(s, expected) {
			t.Errorf("expected=%s, got=%s", expected, s)
		}
	})

	t.Run("valid-marshal-json", func(t *testing.T) {
		guid := guid.MustParseGUID("93b2c683-d8af-4cd0-8ca2-bf8177871a3b")
		expected := []byte(`"93b2c683-d8af-4cd0-8ca2-bf8177871a3b"`)
		s, err := guid.MarshalJSON()

		if err != nil {
			t.Errorf("MarshalJSON() should never return error")
		}

		if !bytes.Equal(s, expected) {
			t.Errorf("expected=%s, got=%s", expected, s)
		}
	})

	t.Run("valid-unmarshal-json", func(t *testing.T) {
		data := []byte(`"93b2c683-d8af-4cd0-8ca2-bf8177871a3b"`)
		expected := "93b2c683-d8af-4cd0-8ca2-bf8177871a3b"

		v := guid.GUID{}
		err := v.UnmarshalJSON(data)

		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		if v.String() != expected {
			t.Errorf("expected=%s, got=%s", expected, v.String())
		}
	})

	t.Run("invalid-missing-quotes-unmarshal-json", func(t *testing.T) {
		data := []byte(`93b2c683-d8af-4cd0-8ca2-bf8177871a3b`)

		v := guid.GUID{}
		err := v.UnmarshalJSON(data)

		if err == nil {
			t.Errorf("expected error: %s", err)
		}

		if v.String() != zeroGUID {
			t.Errorf("expected=%s, got=%s", zeroGUID, v.String())
		}
	})

	t.Run("invalid-unbalanced-quotes-unmarshal-json", func(t *testing.T) {
		data := []byte(`"93b2c683-d8af-4cd0-8ca2-bf8177871a3b`)

		v := guid.GUID{}
		err := v.UnmarshalJSON(data)

		if err == nil {
			t.Errorf("expected error: %s", err)
		}

		if v.String() != zeroGUID {
			t.Errorf("expected=%s, got=%s", zeroGUID, v.String())
		}
	})

	t.Run("invalid-too-short-unmarshal-json", func(t *testing.T) {
		data := []byte(`"93b2c683-d8af-4cd0-8ca2-bf8177871a3"`)

		v := guid.GUID{}
		err := v.UnmarshalJSON(data)

		if err == nil {
			t.Errorf("expected error: %s", err)
		}

		if v.String() != zeroGUID {
			t.Errorf("expected=%s, got=%s", zeroGUID, v.String())
		}
	})
}

func TestParse_Invalid(t *testing.T) {
	type testCase struct {
		name  string
		input string
		valid bool
	}
	tt := []testCase{
		{
			name: "empty-string",
		},
		{
			name:  "too-short",
			input: "93b2c683-d8af-4cd0-8ca2-bf8177871a3",
		},
		{
			name:  "too-short-long",
			input: "93b2c683-d8af-4cd0-8ca2-bf8177871a3aa",
		},
		{
			name:  "missing-first-dash",
			input: "93b2c683?d8af-4cd0-8ca2-bf8177871a3b",
		},
		{
			name:  "missing-second-dash",
			input: "93b2c683-d8af?4cd0-8ca2-bf8177871a3b",
		},
		{
			name:  "missing-third-dash",
			input: "93b2c683-d8af-4cd0?8ca2-bf8177871a3b",
		},
		{
			name:  "missing-fourth-dash",
			input: "93b2c683-d8af-4cd0-8ca2?bf8177871a3b",
		},
		{
			name:  "invalid-first-segment",
			input: "93?2c683-d8af-4cd0-8ca2-bf8177871a3b",
		},
		{
			name:  "invalid-second-segment",
			input: "93b2c683-d8?f-4cd0-8ca2-bf8177871a3b",
		},
		{
			name:  "invalid-third-segment",
			input: "93b2c683-d8af-4c?0-8ca2-bf8177871a3b",
		},
		{
			name:  "invalid-fourth-segment",
			input: "93b2c683-d8af-4cd0-8c?2-bf8177871a3b",
		},
		{
			name:  "invalid-fifth-segment-a",
			input: "93b2c683-d8af-4cd0-8ca2-?f8177871a3b",
		},
		{
			name:  "invalid-fifth-segment-b",
			input: "93b2c683-d8af-4cd0-8ca2-bf?177871a3b",
		},
		{
			name:  "invalid-fifth-segment-c",
			input: "93b2c683-d8af-4cd0-8ca2-bf81?7871a3b",
		},
		{
			name:  "invalid-fifth-segment-d",
			input: "93b2c683-d8af-4cd0-8ca2-bf8177?71a3b",
		},
		{
			name:  "invalid-fifth-segment-e",
			input: "93b2c683-d8af-4cd0-8ca2-bf817787?a3b",
		},
		{
			name:  "invalid-fifth-segment-f",
			input: "93b2c683-d8af-4cd0-8ca2-bf8177871a3?",
		},
		{
			name:  "valid",
			input: "93b2c683-d8af-4cd0-8ca2-bf8177871a3b",
			valid: true,
		},
		{
			name:  "zero",
			input: zeroGUID,
			valid: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			v, err := guid.ParseGUID(tc.input)
			if tc.valid {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				if v.String() != tc.input {
					t.Errorf("expected=%s, got=%s", tc.input, v.String())
				}
			} else {
				if err == nil {
					t.Errorf("expected error")
				}
				if v.String() != zeroGUID {
					t.Errorf("expected=%s, got=%s", zeroGUID, v.String())
				}
			}
		})
	}
}

func TestMustParse(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		assert.Panics(t, func() {
			guid.MustParseGUID("")
		})
	})
	t.Run("Valid", func(t *testing.T) {
		v := guid.MustParseGUID("93b2c683-d8af-4cd0-8ca2-bf8177871a3b")
		expected := "93b2c683-d8af-4cd0-8ca2-bf8177871a3b"
		if v.String() != expected {
			t.Errorf("expected=%s, got=%s", expected, v.String())
		}
	})
}
