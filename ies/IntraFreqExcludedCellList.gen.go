// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: IntraFreqExcludedCellList (line 3917).
// IntraFreqExcludedCellList ::=       SEQUENCE (SIZE (1..maxCellExcluded)) OF PCI-Range

var intraFreqExcludedCellListSizeConstraints = per.SizeRange(1, common.MaxCellExcluded)

type IntraFreqExcludedCellList []PCI_Range

func (ie *IntraFreqExcludedCellList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(intraFreqExcludedCellListSizeConstraints)
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

func (ie *IntraFreqExcludedCellList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(intraFreqExcludedCellListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(IntraFreqExcludedCellList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
