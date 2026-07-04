// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ValidityCellList (line 9270).
// ValidityCellList ::= SEQUENCE (SIZE (1.. maxCellMeasIdle-r16)) OF PCI-Range

var validityCellListSizeConstraints = per.SizeRange(1, common.MaxCellMeasIdle_r16)

type ValidityCellList []PCI_Range

func (ie *ValidityCellList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(validityCellListSizeConstraints)
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

func (ie *ValidityCellList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(validityCellListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ValidityCellList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
