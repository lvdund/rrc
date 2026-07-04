// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: EUTRA-FreqNeighCellList (line 4170).
// EUTRA-FreqNeighCellList ::=         SEQUENCE (SIZE (1..maxCellEUTRA)) OF EUTRA-FreqNeighCellInfo

var eUTRAFreqNeighCellListSizeConstraints = per.SizeRange(1, common.MaxCellEUTRA)

type EUTRA_FreqNeighCellList []EUTRA_FreqNeighCellInfo

func (ie *EUTRA_FreqNeighCellList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(eUTRAFreqNeighCellListSizeConstraints)
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

func (ie *EUTRA_FreqNeighCellList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(eUTRAFreqNeighCellListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(EUTRA_FreqNeighCellList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
