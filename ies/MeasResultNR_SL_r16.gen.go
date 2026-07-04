// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultNR-SL-r16 (line 10015).

var measResultNRSLR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultListCBR-NR-r16"},
	},
}

var measResultNRSLR16MeasResultListCBRNRR16Constraints = per.SizeRange(1, common.MaxNrofSL_PoolToMeasureNR_r16)

type MeasResultNR_SL_r16 struct {
	MeasResultListCBR_NR_r16 []MeasResultCBR_NR_r16
}

func (ie *MeasResultNR_SL_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultNRSLR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. measResultListCBR-NR-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(measResultNRSLR16MeasResultListCBRNRR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.MeasResultListCBR_NR_r16))); err != nil {
			return err
		}
		for i := range ie.MeasResultListCBR_NR_r16 {
			if err := ie.MeasResultListCBR_NR_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResultNR_SL_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultNRSLR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. measResultListCBR-NR-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(measResultNRSLR16MeasResultListCBRNRR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.MeasResultListCBR_NR_r16 = make([]MeasResultCBR_NR_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.MeasResultListCBR_NR_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
