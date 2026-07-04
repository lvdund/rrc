// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NeedForGapNCSG-ConfigEUTRA-r17 (line 10475).

var needForGapNCSGConfigEUTRAR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "requestedTargetBandFilterNCSG-EUTRA-r17", Optional: true},
	},
}

var needForGapNCSGConfigEUTRAR17RequestedTargetBandFilterNCSGEUTRAR17Constraints = per.SizeRange(1, common.MaxBandsEUTRA)

type NeedForGapNCSG_ConfigEUTRA_r17 struct {
	RequestedTargetBandFilterNCSG_EUTRA_r17 []FreqBandIndicatorEUTRA
}

func (ie *NeedForGapNCSG_ConfigEUTRA_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(needForGapNCSGConfigEUTRAR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RequestedTargetBandFilterNCSG_EUTRA_r17 != nil}); err != nil {
		return err
	}

	// 2. requestedTargetBandFilterNCSG-EUTRA-r17: sequence-of
	{
		if ie.RequestedTargetBandFilterNCSG_EUTRA_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(needForGapNCSGConfigEUTRAR17RequestedTargetBandFilterNCSGEUTRAR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.RequestedTargetBandFilterNCSG_EUTRA_r17))); err != nil {
				return err
			}
			for i := range ie.RequestedTargetBandFilterNCSG_EUTRA_r17 {
				if err := ie.RequestedTargetBandFilterNCSG_EUTRA_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *NeedForGapNCSG_ConfigEUTRA_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(needForGapNCSGConfigEUTRAR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. requestedTargetBandFilterNCSG-EUTRA-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(needForGapNCSGConfigEUTRAR17RequestedTargetBandFilterNCSGEUTRAR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.RequestedTargetBandFilterNCSG_EUTRA_r17 = make([]FreqBandIndicatorEUTRA, n)
			for i := int64(0); i < n; i++ {
				if err := ie.RequestedTargetBandFilterNCSG_EUTRA_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
