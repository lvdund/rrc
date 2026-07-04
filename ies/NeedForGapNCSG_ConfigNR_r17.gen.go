// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NeedForGapNCSG-ConfigNR-r17 (line 10482).

var needForGapNCSGConfigNRR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "requestedTargetBandFilterNCSG-NR-r17", Optional: true},
	},
}

var needForGapNCSGConfigNRR17RequestedTargetBandFilterNCSGNRR17Constraints = per.SizeRange(1, common.MaxBands)

type NeedForGapNCSG_ConfigNR_r17 struct {
	RequestedTargetBandFilterNCSG_NR_r17 []FreqBandIndicatorNR
}

func (ie *NeedForGapNCSG_ConfigNR_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(needForGapNCSGConfigNRR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RequestedTargetBandFilterNCSG_NR_r17 != nil}); err != nil {
		return err
	}

	// 2. requestedTargetBandFilterNCSG-NR-r17: sequence-of
	{
		if ie.RequestedTargetBandFilterNCSG_NR_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(needForGapNCSGConfigNRR17RequestedTargetBandFilterNCSGNRR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.RequestedTargetBandFilterNCSG_NR_r17))); err != nil {
				return err
			}
			for i := range ie.RequestedTargetBandFilterNCSG_NR_r17 {
				if err := ie.RequestedTargetBandFilterNCSG_NR_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *NeedForGapNCSG_ConfigNR_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(needForGapNCSGConfigNRR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. requestedTargetBandFilterNCSG-NR-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(needForGapNCSGConfigNRR17RequestedTargetBandFilterNCSGNRR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.RequestedTargetBandFilterNCSG_NR_r17 = make([]FreqBandIndicatorNR, n)
			for i := int64(0); i < n; i++ {
				if err := ie.RequestedTargetBandFilterNCSG_NR_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
