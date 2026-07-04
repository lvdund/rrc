// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRB-Identity-v1900 (line 15282).
// SRB-Identity-v1900 ::=              INTEGER (6)

var sRBIdentityV1900Constraints = per.Constrained(6, 6)

type SRB_Identity_v1900 struct {
	Value int64
}

func (v *SRB_Identity_v1900) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sRBIdentityV1900Constraints)
}

func (v *SRB_Identity_v1900) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sRBIdentityV1900Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
