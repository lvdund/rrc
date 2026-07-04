// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CellsToAddModListExt-v1800 (line 9545).
// CellsToAddModListExt-v1800 ::=      SEQUENCE (SIZE (1..maxNrofCellMeas)) OF CellsToAddModExt-v1800

var cellsToAddModListExtV1800SizeConstraints = per.SizeRange(1, common.MaxNrofCellMeas)

type CellsToAddModListExt_v1800 []CellsToAddModExt_v1800

func (ie *CellsToAddModListExt_v1800) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cellsToAddModListExtV1800SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := (*ie)[i].Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CellsToAddModListExt_v1800) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cellsToAddModListExtV1800SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CellsToAddModListExt_v1800, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
