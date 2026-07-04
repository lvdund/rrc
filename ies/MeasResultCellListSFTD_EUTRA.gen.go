// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultCellListSFTD-EUTRA (line 9699).
// MeasResultCellListSFTD-EUTRA ::=          SEQUENCE (SIZE (1..maxCellSFTD)) OF MeasResultSFTD-EUTRA

var measResultCellListSFTDEUTRASizeConstraints = per.SizeRange(1, common.MaxCellSFTD)

type MeasResultCellListSFTD_EUTRA []MeasResultSFTD_EUTRA

func (ie *MeasResultCellListSFTD_EUTRA) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultCellListSFTDEUTRASizeConstraints)
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

func (ie *MeasResultCellListSFTD_EUTRA) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultCellListSFTDEUTRASizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultCellListSFTD_EUTRA, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
