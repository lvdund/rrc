// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SetOfCellsId-r18 (line 14872).
// SetOfCellsId-r18 ::=                   INTEGER (0..maxNrofSetsOfCells-1-r18)

var setOfCellsIdR18Constraints = per.Constrained(0, common.MaxNrofSetsOfCells_1_r18)

type SetOfCellsId_r18 struct {
	Value int64
}

func (v *SetOfCellsId_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, setOfCellsIdR18Constraints)
}

func (v *SetOfCellsId_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(setOfCellsIdR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
