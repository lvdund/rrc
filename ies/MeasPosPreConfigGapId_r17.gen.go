// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasPosPreConfigGapId-r17 (line 9208).
// MeasPosPreConfigGapId-r17 ::= INTEGER (1..maxNrofPreConfigPosGapId-r17)

var measPosPreConfigGapIdR17Constraints = per.Constrained(1, common.MaxNrofPreConfigPosGapId_r17)

type MeasPosPreConfigGapId_r17 struct {
	Value int64
}

func (v *MeasPosPreConfigGapId_r17) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, measPosPreConfigGapIdR17Constraints)
}

func (v *MeasPosPreConfigGapId_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(measPosPreConfigGapIdR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
