// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: InterFreqTargetInfo-r16 (line 26070).

var interFreqTargetInfoR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-CarrierFreq-r16"},
		{Name: "cellList-r16", Optional: true},
	},
}

var interFreqTargetInfoR16CellListR16Constraints = per.SizeRange(1, 32)

type InterFreqTargetInfo_r16 struct {
	Dl_CarrierFreq_r16 ARFCN_ValueNR
	CellList_r16       []PhysCellId
}

func (ie *InterFreqTargetInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(interFreqTargetInfoR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CellList_r16 != nil}); err != nil {
		return err
	}

	// 2. dl-CarrierFreq-r16: ref
	{
		if err := ie.Dl_CarrierFreq_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. cellList-r16: sequence-of
	{
		if ie.CellList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(interFreqTargetInfoR16CellListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.CellList_r16))); err != nil {
				return err
			}
			for i := range ie.CellList_r16 {
				if err := ie.CellList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *InterFreqTargetInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(interFreqTargetInfoR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dl-CarrierFreq-r16: ref
	{
		if err := ie.Dl_CarrierFreq_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. cellList-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(interFreqTargetInfoR16CellListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CellList_r16 = make([]PhysCellId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CellList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
