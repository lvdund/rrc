// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MTCH-NeighbourCell-r18 (line 28558).
// MTCH-NeighbourCell-r18 ::= BIT STRING (SIZE(maxNeighCellMBS-r17))

var mTCHNeighbourCellR18Constraints = per.FixedSize(common.MaxNeighCellMBS_r17)

type MTCH_NeighbourCell_r18 struct {
	Value per.BitString
}

func (v *MTCH_NeighbourCell_r18) Encode(e *per.Encoder) error {
	return e.EncodeBitString(v.Value, mTCHNeighbourCellR18Constraints)
}

func (v *MTCH_NeighbourCell_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeBitString(mTCHNeighbourCellR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
