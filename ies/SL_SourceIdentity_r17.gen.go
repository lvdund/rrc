// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-SourceIdentity-r17 (line 28257).
// SL-SourceIdentity-r17 ::=   BIT STRING (SIZE (24))

var sLSourceIdentityR17Constraints = per.FixedSize(24)

type SL_SourceIdentity_r17 struct {
	Value per.BitString
}

func (v *SL_SourceIdentity_r17) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, sLSourceIdentityR17Constraints)
}

func (v *SL_SourceIdentity_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(sLSourceIdentityR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
