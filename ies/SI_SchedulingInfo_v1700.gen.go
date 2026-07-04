// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SI-SchedulingInfo-v1700 (line 15066).

var sISchedulingInfoV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "schedulingInfoList2-r17"},
		{Name: "dummy", Optional: true},
	},
}

var sISchedulingInfoV1700SchedulingInfoList2R17Constraints = per.SizeRange(1, common.MaxSI_Message)

type SI_SchedulingInfo_v1700 struct {
	SchedulingInfoList2_r17 []SchedulingInfo2_r17
	Dummy                   *SI_RequestConfig
}

func (ie *SI_SchedulingInfo_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sISchedulingInfoV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dummy != nil}); err != nil {
		return err
	}

	// 2. schedulingInfoList2-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sISchedulingInfoV1700SchedulingInfoList2R17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.SchedulingInfoList2_r17))); err != nil {
			return err
		}
		for i := range ie.SchedulingInfoList2_r17 {
			if err := ie.SchedulingInfoList2_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. dummy: ref
	{
		if ie.Dummy != nil {
			if err := ie.Dummy.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SI_SchedulingInfo_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sISchedulingInfoV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. schedulingInfoList2-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sISchedulingInfoV1700SchedulingInfoList2R17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.SchedulingInfoList2_r17 = make([]SchedulingInfo2_r17, n)
		for i := int64(0); i < n; i++ {
			if err := ie.SchedulingInfoList2_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. dummy: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Dummy = new(SI_RequestConfig)
			if err := ie.Dummy.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
