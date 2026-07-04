// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NeedForGapsConfigNR-r16 (line 10446).

var needForGapsConfigNRR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "requestedTargetBandFilterNR-r16", Optional: true},
	},
}

var needForGapsConfigNRR16RequestedTargetBandFilterNRR16Constraints = per.SizeRange(1, common.MaxBands)

type NeedForGapsConfigNR_r16 struct {
	RequestedTargetBandFilterNR_r16 []FreqBandIndicatorNR
}

func (ie *NeedForGapsConfigNR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(needForGapsConfigNRR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RequestedTargetBandFilterNR_r16 != nil}); err != nil {
		return err
	}

	// 2. requestedTargetBandFilterNR-r16: sequence-of
	{
		if ie.RequestedTargetBandFilterNR_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(needForGapsConfigNRR16RequestedTargetBandFilterNRR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.RequestedTargetBandFilterNR_r16))); err != nil {
				return err
			}
			for i := range ie.RequestedTargetBandFilterNR_r16 {
				if err := ie.RequestedTargetBandFilterNR_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *NeedForGapsConfigNR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(needForGapsConfigNRR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. requestedTargetBandFilterNR-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(needForGapsConfigNRR16RequestedTargetBandFilterNRR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.RequestedTargetBandFilterNR_r16 = make([]FreqBandIndicatorNR, n)
			for i := int64(0); i < n; i++ {
				if err := ie.RequestedTargetBandFilterNR_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
