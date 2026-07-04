// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: IAB-ResourceConfigID-r17 (line 5781).
// IAB-ResourceConfigID-r17 ::=        INTEGER(0..maxNrofIABResourceConfig-1-r17)

var iABResourceConfigIDR17Constraints = per.Constrained(0, common.MaxNrofIABResourceConfig_1_r17)

type IAB_ResourceConfigID_r17 struct {
	Value int64
}

func (v *IAB_ResourceConfigID_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, iABResourceConfigIDR17Constraints)
}

func (v *IAB_ResourceConfigID_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(iABResourceConfigIDR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
