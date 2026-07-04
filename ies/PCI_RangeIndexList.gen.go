// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PCI-RangeIndexList (line 10984).
// PCI-RangeIndexList ::=              SEQUENCE (SIZE (1..maxNrofPCI-Ranges)) OF PCI-RangeIndex

var pCIRangeIndexListSizeConstraints = per.SizeRange(1, common.MaxNrofPCI_Ranges)

type PCI_RangeIndexList []PCI_RangeIndex

func (ie *PCI_RangeIndexList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pCIRangeIndexListSizeConstraints)
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

func (ie *PCI_RangeIndexList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pCIRangeIndexListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PCI_RangeIndexList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
