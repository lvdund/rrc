// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MCC-MNC-Digit (line 11894).
// MCC-MNC-Digit ::=                   INTEGER (0..9)

var mCCMNCDigitConstraints = per.Constrained(0, 9)

type MCC_MNC_Digit struct {
	Value int64
}

func (v *MCC_MNC_Digit) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, mCCMNCDigitConstraints)
}

func (v *MCC_MNC_Digit) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(mCCMNCDigitConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
