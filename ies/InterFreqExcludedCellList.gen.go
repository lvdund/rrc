// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: InterFreqExcludedCellList (line 4090).
// InterFreqExcludedCellList ::=       SEQUENCE (SIZE (1..maxCellExcluded)) OF PCI-Range

var interFreqExcludedCellListSizeConstraints = per.SizeRange(1, common.MaxCellExcluded)

type InterFreqExcludedCellList []PCI_Range

func (ie *InterFreqExcludedCellList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(interFreqExcludedCellListSizeConstraints)
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

func (ie *InterFreqExcludedCellList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(interFreqExcludedCellListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(InterFreqExcludedCellList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
