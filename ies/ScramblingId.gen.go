// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ScramblingId (line 14329).
// ScramblingId ::=                    INTEGER(0..1023)

var scramblingIdConstraints = per.Constrained(0, 1023)

type ScramblingId struct {
	Value int64
}

func (v *ScramblingId) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, scramblingIdConstraints)
}

func (v *ScramblingId) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(scramblingIdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
