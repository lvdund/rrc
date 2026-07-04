// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AffectedCarrierFreqRangeComb-r18 (line 2747).

var affectedCarrierFreqRangeCombR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "affectedCarrierFreqRangeComb-r18"},
		{Name: "interferenceDirection-r18"},
		{Name: "victimSystemType-r18", Optional: true},
	},
}

var affectedCarrierFreqRangeCombR18AffectedCarrierFreqRangeCombR18Constraints = per.SizeRange(2, common.MaxNrofServingCells)

const (
	AffectedCarrierFreqRangeComb_r18_InterferenceDirection_r18_Nr    = 0
	AffectedCarrierFreqRangeComb_r18_InterferenceDirection_r18_Other = 1
	AffectedCarrierFreqRangeComb_r18_InterferenceDirection_r18_Both  = 2
	AffectedCarrierFreqRangeComb_r18_InterferenceDirection_r18_Spare = 3
)

var affectedCarrierFreqRangeCombR18InterferenceDirectionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AffectedCarrierFreqRangeComb_r18_InterferenceDirection_r18_Nr, AffectedCarrierFreqRangeComb_r18_InterferenceDirection_r18_Other, AffectedCarrierFreqRangeComb_r18_InterferenceDirection_r18_Both, AffectedCarrierFreqRangeComb_r18_InterferenceDirection_r18_Spare},
}

type AffectedCarrierFreqRangeComb_r18 struct {
	AffectedCarrierFreqRangeComb_r18 []AffectedFreqRange_r18
	InterferenceDirection_r18        int64
	VictimSystemType_r18             *VictimSystemType_r16
}

func (ie *AffectedCarrierFreqRangeComb_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(affectedCarrierFreqRangeCombR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.VictimSystemType_r18 != nil}); err != nil {
		return err
	}

	// 2. affectedCarrierFreqRangeComb-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(affectedCarrierFreqRangeCombR18AffectedCarrierFreqRangeCombR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.AffectedCarrierFreqRangeComb_r18))); err != nil {
			return err
		}
		for i := range ie.AffectedCarrierFreqRangeComb_r18 {
			if err := ie.AffectedCarrierFreqRangeComb_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. interferenceDirection-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.InterferenceDirection_r18, affectedCarrierFreqRangeCombR18InterferenceDirectionR18Constraints); err != nil {
			return err
		}
	}

	// 4. victimSystemType-r18: ref
	{
		if ie.VictimSystemType_r18 != nil {
			if err := ie.VictimSystemType_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *AffectedCarrierFreqRangeComb_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(affectedCarrierFreqRangeCombR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. affectedCarrierFreqRangeComb-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(affectedCarrierFreqRangeCombR18AffectedCarrierFreqRangeCombR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.AffectedCarrierFreqRangeComb_r18 = make([]AffectedFreqRange_r18, n)
		for i := int64(0); i < n; i++ {
			if err := ie.AffectedCarrierFreqRangeComb_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. interferenceDirection-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(affectedCarrierFreqRangeCombR18InterferenceDirectionR18Constraints)
		if err != nil {
			return err
		}
		ie.InterferenceDirection_r18 = v1
	}

	// 4. victimSystemType-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.VictimSystemType_r18 = new(VictimSystemType_r16)
			if err := ie.VictimSystemType_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
