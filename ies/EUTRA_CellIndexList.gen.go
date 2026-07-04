// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: EUTRA-CellIndexList (line 9406).
// EUTRA-CellIndexList ::=                     SEQUENCE (SIZE (1..maxCellMeasEUTRA)) OF EUTRA-CellIndex

var eUTRACellIndexListSizeConstraints = per.SizeRange(1, common.MaxCellMeasEUTRA)

type EUTRA_CellIndexList []EUTRA_CellIndex

func (ie *EUTRA_CellIndexList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(eUTRACellIndexListSizeConstraints)
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

func (ie *EUTRA_CellIndexList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(eUTRACellIndexListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(EUTRA_CellIndexList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
