// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: EUTRA-FreqExcludedCellList (line 4168).
// EUTRA-FreqExcludedCellList ::=      SEQUENCE (SIZE (1..maxEUTRA-CellExcluded)) OF EUTRA-PhysCellIdRange

var eUTRAFreqExcludedCellListSizeConstraints = per.SizeRange(1, common.MaxEUTRA_CellExcluded)

type EUTRA_FreqExcludedCellList []EUTRA_PhysCellIdRange

func (ie *EUTRA_FreqExcludedCellList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(eUTRAFreqExcludedCellListSizeConstraints)
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

func (ie *EUTRA_FreqExcludedCellList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(eUTRAFreqExcludedCellListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(EUTRA_FreqExcludedCellList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
