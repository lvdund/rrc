// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRB-Identity-v1700 (line 15278).
// SRB-Identity-v1700 ::=              INTEGER (4)

var sRBIdentityV1700Constraints = per.Constrained(4, 4)

type SRB_Identity_v1700 struct {
	Value int64
}

func (v *SRB_Identity_v1700) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sRBIdentityV1700Constraints)
}

func (v *SRB_Identity_v1700) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sRBIdentityV1700Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
