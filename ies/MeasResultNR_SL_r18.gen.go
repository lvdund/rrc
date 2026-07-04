// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultNR-SL-r18 (line 10026).

var measResultNRSLR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultListCBR-DedicatedSL-PRS-r18"},
	},
}

var measResultNRSLR18MeasResultListCBRDedicatedSLPRSR18Constraints = per.SizeRange(1, common.MaxNrofDedicatedSL_PRS_PoolToMeas_r18)

type MeasResultNR_SL_r18 struct {
	MeasResultListCBR_DedicatedSL_PRS_r18 []MeasResultCBR_DedicatedSL_PRS_r18
}

func (ie *MeasResultNR_SL_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultNRSLR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. measResultListCBR-DedicatedSL-PRS-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(measResultNRSLR18MeasResultListCBRDedicatedSLPRSR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.MeasResultListCBR_DedicatedSL_PRS_r18))); err != nil {
			return err
		}
		for i := range ie.MeasResultListCBR_DedicatedSL_PRS_r18 {
			if err := ie.MeasResultListCBR_DedicatedSL_PRS_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MeasResultNR_SL_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultNRSLR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. measResultListCBR-DedicatedSL-PRS-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(measResultNRSLR18MeasResultListCBRDedicatedSLPRSR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.MeasResultListCBR_DedicatedSL_PRS_r18 = make([]MeasResultCBR_DedicatedSL_PRS_r18, n)
		for i := int64(0); i < n; i++ {
			if err := ie.MeasResultListCBR_DedicatedSL_PRS_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
