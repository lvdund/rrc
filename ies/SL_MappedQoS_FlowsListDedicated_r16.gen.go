// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-MappedQoS-FlowsListDedicated-r16 (line 28241).

var sLMappedQoSFlowsListDedicatedR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-MappedQoS-FlowsToAddList-r16", Optional: true},
		{Name: "sl-MappedQoS-FlowsToReleaseList-r16", Optional: true},
	},
}

var sLMappedQoSFlowsListDedicatedR16SlMappedQoSFlowsToAddListR16Constraints = per.SizeRange(1, common.MaxNrofSL_QFIs_r16)

var sLMappedQoSFlowsListDedicatedR16SlMappedQoSFlowsToReleaseListR16Constraints = per.SizeRange(1, common.MaxNrofSL_QFIs_r16)

type SL_MappedQoS_FlowsListDedicated_r16 struct {
	Sl_MappedQoS_FlowsToAddList_r16     []SL_QoS_FlowIdentity_r16
	Sl_MappedQoS_FlowsToReleaseList_r16 []SL_QoS_FlowIdentity_r16
}

func (ie *SL_MappedQoS_FlowsListDedicated_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLMappedQoSFlowsListDedicatedR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_MappedQoS_FlowsToAddList_r16 != nil, ie.Sl_MappedQoS_FlowsToReleaseList_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-MappedQoS-FlowsToAddList-r16: sequence-of
	{
		if ie.Sl_MappedQoS_FlowsToAddList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLMappedQoSFlowsListDedicatedR16SlMappedQoSFlowsToAddListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_MappedQoS_FlowsToAddList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_MappedQoS_FlowsToAddList_r16 {
				if err := ie.Sl_MappedQoS_FlowsToAddList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-MappedQoS-FlowsToReleaseList-r16: sequence-of
	{
		if ie.Sl_MappedQoS_FlowsToReleaseList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLMappedQoSFlowsListDedicatedR16SlMappedQoSFlowsToReleaseListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_MappedQoS_FlowsToReleaseList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_MappedQoS_FlowsToReleaseList_r16 {
				if err := ie.Sl_MappedQoS_FlowsToReleaseList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_MappedQoS_FlowsListDedicated_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLMappedQoSFlowsListDedicatedR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-MappedQoS-FlowsToAddList-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLMappedQoSFlowsListDedicatedR16SlMappedQoSFlowsToAddListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_MappedQoS_FlowsToAddList_r16 = make([]SL_QoS_FlowIdentity_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_MappedQoS_FlowsToAddList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-MappedQoS-FlowsToReleaseList-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLMappedQoSFlowsListDedicatedR16SlMappedQoSFlowsToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_MappedQoS_FlowsToReleaseList_r16 = make([]SL_QoS_FlowIdentity_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_MappedQoS_FlowsToReleaseList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
