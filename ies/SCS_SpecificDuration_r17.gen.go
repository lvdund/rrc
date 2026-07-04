// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SCS-SpecificDuration-r17 (line 11039).
// SCS-SpecificDuration-r17   ::=      INTEGER (1..166)

var sCSSpecificDurationR17Constraints = per.Constrained(1, 166)

type SCS_SpecificDuration_r17 struct {
	Value int64
}

func (v *SCS_SpecificDuration_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sCSSpecificDurationR17Constraints)
}

func (v *SCS_SpecificDuration_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sCSSpecificDurationR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
