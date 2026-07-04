// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AffectedCarrierFreqRange-r18 (line 2739).

var affectedCarrierFreqRangeR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "affectedFreqRange-r18"},
		{Name: "interferenceDirection-r18"},
		{Name: "victimSystemType-r18", Optional: true},
	},
}

const (
	AffectedCarrierFreqRange_r18_InterferenceDirection_r18_Nr    = 0
	AffectedCarrierFreqRange_r18_InterferenceDirection_r18_Other = 1
	AffectedCarrierFreqRange_r18_InterferenceDirection_r18_Both  = 2
	AffectedCarrierFreqRange_r18_InterferenceDirection_r18_Spare = 3
)

var affectedCarrierFreqRangeR18InterferenceDirectionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AffectedCarrierFreqRange_r18_InterferenceDirection_r18_Nr, AffectedCarrierFreqRange_r18_InterferenceDirection_r18_Other, AffectedCarrierFreqRange_r18_InterferenceDirection_r18_Both, AffectedCarrierFreqRange_r18_InterferenceDirection_r18_Spare},
}

type AffectedCarrierFreqRange_r18 struct {
	AffectedFreqRange_r18     AffectedFreqRange_r18
	InterferenceDirection_r18 int64
	VictimSystemType_r18      *VictimSystemType_r16
}

func (ie *AffectedCarrierFreqRange_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(affectedCarrierFreqRangeR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.VictimSystemType_r18 != nil}); err != nil {
		return err
	}

	// 2. affectedFreqRange-r18: ref
	{
		if err := ie.AffectedFreqRange_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. interferenceDirection-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.InterferenceDirection_r18, affectedCarrierFreqRangeR18InterferenceDirectionR18Constraints); err != nil {
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

func (ie *AffectedCarrierFreqRange_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(affectedCarrierFreqRangeR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. affectedFreqRange-r18: ref
	{
		if err := ie.AffectedFreqRange_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. interferenceDirection-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(affectedCarrierFreqRangeR18InterferenceDirectionR18Constraints)
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
