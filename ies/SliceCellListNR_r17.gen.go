// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SliceCellListNR-r17 (line 8393).
// SliceCellListNR-r17 ::=           SEQUENCE (SIZE (1..maxCellSlice-r17)) OF PCI-Range

var sliceCellListNRR17SizeConstraints = per.SizeRange(1, common.MaxCellSlice_r17)

type SliceCellListNR_r17 []PCI_Range

func (ie *SliceCellListNR_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sliceCellListNRR17SizeConstraints)
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

func (ie *SliceCellListNR_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sliceCellListNRR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SliceCellListNR_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
