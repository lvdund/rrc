// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRB-Identity-r17 (line 10108).
// MRB-Identity-r17 ::=                    INTEGER (1..512)

var mRBIdentityR17Constraints = per.Constrained(1, 512)

type MRB_Identity_r17 struct {
	Value int64
}

func (v *MRB_Identity_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, mRBIdentityR17Constraints)
}

func (v *MRB_Identity_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(mRBIdentityR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
