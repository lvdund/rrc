// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: InterFreqCAG-CellListPerPLMN-r16 (line 4094).

var interFreqCAGCellListPerPLMNR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-IdentityIndex-r16"},
		{Name: "cag-CellList-r16"},
	},
}

var interFreqCAGCellListPerPLMNR16PlmnIdentityIndexR16Constraints = per.Constrained(1, common.MaxPLMN)

var interFreqCAGCellListPerPLMNR16CagCellListR16Constraints = per.SizeRange(1, common.MaxCAG_Cell_r16)

type InterFreqCAG_CellListPerPLMN_r16 struct {
	Plmn_IdentityIndex_r16 int64
	Cag_CellList_r16       []PCI_Range
}

func (ie *InterFreqCAG_CellListPerPLMN_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(interFreqCAGCellListPerPLMNR16Constraints)
	_ = seq

	// 1. plmn-IdentityIndex-r16: integer
	{
		if err := e.EncodeInteger(ie.Plmn_IdentityIndex_r16, interFreqCAGCellListPerPLMNR16PlmnIdentityIndexR16Constraints); err != nil {
			return err
		}
	}

	// 2. cag-CellList-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(interFreqCAGCellListPerPLMNR16CagCellListR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Cag_CellList_r16))); err != nil {
			return err
		}
		for i := range ie.Cag_CellList_r16 {
			if err := ie.Cag_CellList_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *InterFreqCAG_CellListPerPLMN_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(interFreqCAGCellListPerPLMNR16Constraints)
	_ = seq

	// 1. plmn-IdentityIndex-r16: integer
	{
		v0, err := d.DecodeInteger(interFreqCAGCellListPerPLMNR16PlmnIdentityIndexR16Constraints)
		if err != nil {
			return err
		}
		ie.Plmn_IdentityIndex_r16 = v0
	}

	// 2. cag-CellList-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(interFreqCAGCellListPerPLMNR16CagCellListR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Cag_CellList_r16 = make([]PCI_Range, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Cag_CellList_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
