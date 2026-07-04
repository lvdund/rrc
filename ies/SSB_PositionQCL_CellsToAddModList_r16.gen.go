// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SSB-PositionQCL-CellsToAddModList-r16 (line 9582).
// SSB-PositionQCL-CellsToAddModList-r16 ::= SEQUENCE (SIZE (1..maxNrofCellMeas)) OF SSB-PositionQCL-CellsToAddMod-r16

var sSBPositionQCLCellsToAddModListR16SizeConstraints = per.SizeRange(1, common.MaxNrofCellMeas)

type SSB_PositionQCL_CellsToAddModList_r16 []SSB_PositionQCL_CellsToAddMod_r16

func (ie *SSB_PositionQCL_CellsToAddModList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sSBPositionQCLCellsToAddModListR16SizeConstraints)
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

func (ie *SSB_PositionQCL_CellsToAddModList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sSBPositionQCLCellsToAddModListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SSB_PositionQCL_CellsToAddModList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
