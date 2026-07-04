// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-v1680 (line 16727).

var bandCombinationV1680Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "intrabandConcurrentOperationPowerClass-r16", Optional: true},
	},
}

var bandCombinationV1680IntrabandConcurrentOperationPowerClassR16Constraints = per.SizeRange(1, common.MaxBandComb)

type BandCombination_v1680 struct {
	IntrabandConcurrentOperationPowerClass_r16 []IntraBandPowerClass_r16
}

func (ie *BandCombination_v1680) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1680Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IntrabandConcurrentOperationPowerClass_r16 != nil}); err != nil {
		return err
	}

	// 2. intrabandConcurrentOperationPowerClass-r16: sequence-of
	{
		if ie.IntrabandConcurrentOperationPowerClass_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationV1680IntrabandConcurrentOperationPowerClassR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.IntrabandConcurrentOperationPowerClass_r16))); err != nil {
				return err
			}
			for i := range ie.IntrabandConcurrentOperationPowerClass_r16 {
				if err := ie.IntrabandConcurrentOperationPowerClass_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1680) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1680Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. intrabandConcurrentOperationPowerClass-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationV1680IntrabandConcurrentOperationPowerClassR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.IntrabandConcurrentOperationPowerClass_r16 = make([]IntraBandPowerClass_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.IntrabandConcurrentOperationPowerClass_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
