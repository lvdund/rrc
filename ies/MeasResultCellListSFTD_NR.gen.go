// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultCellListSFTD-NR (line 9687).
// MeasResultCellListSFTD-NR ::=          SEQUENCE (SIZE (1..maxCellSFTD)) OF MeasResultCellSFTD-NR

var measResultCellListSFTDNRSizeConstraints = per.SizeRange(1, common.MaxCellSFTD)

type MeasResultCellListSFTD_NR []MeasResultCellSFTD_NR

func (ie *MeasResultCellListSFTD_NR) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultCellListSFTDNRSizeConstraints)
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

func (ie *MeasResultCellListSFTD_NR) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultCellListSFTDNRSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultCellListSFTD_NR, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
