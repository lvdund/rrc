// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CA-ParametersNR-v16a0 (line 17432).

var cAParametersNRV16a0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-BlindDetectionMixedList-r16"},
	},
}

var cAParametersNRV16a0PdcchBlindDetectionMixedListR16Constraints = per.SizeRange(1, common.MaxNrofPdcch_BlindDetectionMixed_1_r16)

type CA_ParametersNR_V16a0 struct {
	Pdcch_BlindDetectionMixedList_r16 []PDCCH_BlindDetectionMixedList_r16
}

func (ie *CA_ParametersNR_V16a0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV16a0Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionMixedList-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(cAParametersNRV16a0PdcchBlindDetectionMixedListR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Pdcch_BlindDetectionMixedList_r16))); err != nil {
			return err
		}
		for i := range ie.Pdcch_BlindDetectionMixedList_r16 {
			if err := ie.Pdcch_BlindDetectionMixedList_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_V16a0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV16a0Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionMixedList-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(cAParametersNRV16a0PdcchBlindDetectionMixedListR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Pdcch_BlindDetectionMixedList_r16 = make([]PDCCH_BlindDetectionMixedList_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Pdcch_BlindDetectionMixedList_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
