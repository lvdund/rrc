// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasGapId-r17 (line 9213).
// MeasGapId-r17 ::=                       INTEGER (1..maxNrofGapId-r17)

var measGapIdR17Constraints = per.Constrained(1, common.MaxNrofGapId_r17)

type MeasGapId_r17 struct {
	Value int64
}

func (v *MeasGapId_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, measGapIdR17Constraints)
}

func (v *MeasGapId_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(measGapIdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
