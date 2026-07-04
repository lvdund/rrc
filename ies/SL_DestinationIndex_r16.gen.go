// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-DestinationIndex-r16 (line 26975).
// SL-DestinationIndex-r16  ::=             INTEGER (0..maxNrofSL-Dest-1-r16)

var sLDestinationIndexR16Constraints = per.Constrained(0, common.MaxNrofSL_Dest_1_r16)

type SL_DestinationIndex_r16 struct {
	Value int64
}

func (v *SL_DestinationIndex_r16) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, sLDestinationIndexR16Constraints)
}

func (v *SL_DestinationIndex_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(sLDestinationIndexR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
