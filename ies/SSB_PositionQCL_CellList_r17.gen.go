// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SSB-PositionQCL-CellList-r17 (line 9589).
// SSB-PositionQCL-CellList-r17 ::= SEQUENCE (SIZE (1..maxNrofCellMeas)) OF SSB-PositionQCL-Cell-r17

var sSBPositionQCLCellListR17SizeConstraints = per.SizeRange(1, common.MaxNrofCellMeas)

type SSB_PositionQCL_CellList_r17 []SSB_PositionQCL_Cell_r17

func (ie *SSB_PositionQCL_CellList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sSBPositionQCLCellListR17SizeConstraints)
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

func (ie *SSB_PositionQCL_CellList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sSBPositionQCLCellListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SSB_PositionQCL_CellList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
