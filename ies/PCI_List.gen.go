// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PCI-List (line 10957).
// PCI-List ::=                        SEQUENCE (SIZE (1..maxNrofCellMeas)) OF PhysCellId

var pCIListSizeConstraints = per.SizeRange(1, common.MaxNrofCellMeas)

type PCI_List []PhysCellId

func (ie *PCI_List) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pCIListSizeConstraints)
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

func (ie *PCI_List) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pCIListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PCI_List, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
