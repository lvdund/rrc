// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AffectedCarrierFreqComb-r16 (line 2513).

var affectedCarrierFreqCombR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "affectedCarrierFreqComb-r16", Optional: true},
		{Name: "victimSystemType-r16"},
	},
}

var affectedCarrierFreqCombR16AffectedCarrierFreqCombR16Constraints = per.SizeRange(2, common.MaxNrofServingCells)

type AffectedCarrierFreqComb_r16 struct {
	AffectedCarrierFreqComb_r16 []ARFCN_ValueNR
	VictimSystemType_r16        VictimSystemType_r16
}

func (ie *AffectedCarrierFreqComb_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(affectedCarrierFreqCombR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AffectedCarrierFreqComb_r16 != nil}); err != nil {
		return err
	}

	// 2. affectedCarrierFreqComb-r16: sequence-of
	{
		if ie.AffectedCarrierFreqComb_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(affectedCarrierFreqCombR16AffectedCarrierFreqCombR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.AffectedCarrierFreqComb_r16))); err != nil {
				return err
			}
			for i := range ie.AffectedCarrierFreqComb_r16 {
				if err := ie.AffectedCarrierFreqComb_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. victimSystemType-r16: ref
	{
		if err := ie.VictimSystemType_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *AffectedCarrierFreqComb_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(affectedCarrierFreqCombR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. affectedCarrierFreqComb-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(affectedCarrierFreqCombR16AffectedCarrierFreqCombR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AffectedCarrierFreqComb_r16 = make([]ARFCN_ValueNR, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AffectedCarrierFreqComb_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. victimSystemType-r16: ref
	{
		if err := ie.VictimSystemType_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
