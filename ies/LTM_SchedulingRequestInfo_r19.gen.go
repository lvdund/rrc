// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-SchedulingRequestInfo-r19 (line 5869).

var lTMSchedulingRequestInfoR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ltm-SchedulingRequestConfig-r19"},
		{Name: "ltm-SchedulingRequestResourceConfig-r19"},
	},
}

var lTMSchedulingRequestInfoR19LtmSchedulingRequestResourceConfigR19Constraints = per.SizeRange(1, common.MaxNrofSR_Resources)

type LTM_SchedulingRequestInfo_r19 struct {
	Ltm_SchedulingRequestConfig_r19         SchedulingRequestConfig
	Ltm_SchedulingRequestResourceConfig_r19 []SchedulingRequestResourceConfig
}

func (ie *LTM_SchedulingRequestInfo_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMSchedulingRequestInfoR19Constraints)
	_ = seq

	// 1. ltm-SchedulingRequestConfig-r19: ref
	{
		if err := ie.Ltm_SchedulingRequestConfig_r19.Encode(e); err != nil {
			return err
		}
	}

	// 2. ltm-SchedulingRequestResourceConfig-r19: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(lTMSchedulingRequestInfoR19LtmSchedulingRequestResourceConfigR19Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Ltm_SchedulingRequestResourceConfig_r19))); err != nil {
			return err
		}
		for i := range ie.Ltm_SchedulingRequestResourceConfig_r19 {
			if err := ie.Ltm_SchedulingRequestResourceConfig_r19[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LTM_SchedulingRequestInfo_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMSchedulingRequestInfoR19Constraints)
	_ = seq

	// 1. ltm-SchedulingRequestConfig-r19: ref
	{
		if err := ie.Ltm_SchedulingRequestConfig_r19.Decode(d); err != nil {
			return err
		}
	}

	// 2. ltm-SchedulingRequestResourceConfig-r19: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(lTMSchedulingRequestInfoR19LtmSchedulingRequestResourceConfigR19Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Ltm_SchedulingRequestResourceConfig_r19 = make([]SchedulingRequestResourceConfig, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Ltm_SchedulingRequestResourceConfig_r19[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
