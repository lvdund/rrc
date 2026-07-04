// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRB-Identity-v1800 (line 15280).
// SRB-Identity-v1800 ::=              INTEGER (5)

var sRBIdentityV1800Constraints = per.Constrained(5, 5)

type SRB_Identity_v1800 struct {
	Value int64
}

func (v *SRB_Identity_v1800) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sRBIdentityV1800Constraints)
}

func (v *SRB_Identity_v1800) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sRBIdentityV1800Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
