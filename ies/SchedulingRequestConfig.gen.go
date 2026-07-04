// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SchedulingRequestConfig (line 14250).

var schedulingRequestConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "schedulingRequestToAddModList", Optional: true},
		{Name: "schedulingRequestToReleaseList", Optional: true},
	},
}

var schedulingRequestConfigSchedulingRequestToAddModListConstraints = per.SizeRange(1, common.MaxNrofSR_ConfigPerCellGroup)

var schedulingRequestConfigSchedulingRequestToReleaseListConstraints = per.SizeRange(1, common.MaxNrofSR_ConfigPerCellGroup)

type SchedulingRequestConfig struct {
	SchedulingRequestToAddModList  []SchedulingRequestToAddMod
	SchedulingRequestToReleaseList []SchedulingRequestId
}

func (ie *SchedulingRequestConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(schedulingRequestConfigConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SchedulingRequestToAddModList != nil, ie.SchedulingRequestToReleaseList != nil}); err != nil {
		return err
	}

	// 2. schedulingRequestToAddModList: sequence-of
	{
		if ie.SchedulingRequestToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(schedulingRequestConfigSchedulingRequestToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SchedulingRequestToAddModList))); err != nil {
				return err
			}
			for i := range ie.SchedulingRequestToAddModList {
				if err := ie.SchedulingRequestToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. schedulingRequestToReleaseList: sequence-of
	{
		if ie.SchedulingRequestToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(schedulingRequestConfigSchedulingRequestToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SchedulingRequestToReleaseList))); err != nil {
				return err
			}
			for i := range ie.SchedulingRequestToReleaseList {
				if err := ie.SchedulingRequestToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SchedulingRequestConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(schedulingRequestConfigConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. schedulingRequestToAddModList: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(schedulingRequestConfigSchedulingRequestToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SchedulingRequestToAddModList = make([]SchedulingRequestToAddMod, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SchedulingRequestToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. schedulingRequestToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(schedulingRequestConfigSchedulingRequestToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SchedulingRequestToReleaseList = make([]SchedulingRequestId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SchedulingRequestToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
