// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: IntraFreqAllowedCellList-r16 (line 3919).
// IntraFreqAllowedCellList-r16 ::=    SEQUENCE (SIZE (1..maxCellAllowed)) OF PCI-Range

var intraFreqAllowedCellListR16SizeConstraints = per.SizeRange(1, common.MaxCellAllowed)

type IntraFreqAllowedCellList_r16 []PCI_Range

func (ie *IntraFreqAllowedCellList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(intraFreqAllowedCellListR16SizeConstraints)
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

func (ie *IntraFreqAllowedCellList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(intraFreqAllowedCellListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(IntraFreqAllowedCellList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
