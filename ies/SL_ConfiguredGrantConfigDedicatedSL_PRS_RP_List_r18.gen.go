// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-ConfiguredGrantConfigDedicatedSL-PRS-RP-List-r18 (line 28220).

var sLConfiguredGrantConfigDedicatedSLPRSRPListR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ConfiguredGrantConfigDedicatedSL-PRS-RP-ToReleaseList-r18", Optional: true},
		{Name: "sl-ConfiguredGrantConfigDedicatedSL-PRS-RP-ToAddModList-r18", Optional: true},
	},
}

var sLConfiguredGrantConfigDedicatedSLPRSRPListR18SlConfiguredGrantConfigDedicatedSLPRSRPToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofCG_SL_r16)

var sLConfiguredGrantConfigDedicatedSLPRSRPListR18SlConfiguredGrantConfigDedicatedSLPRSRPToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofCG_SL_r16)

type SL_ConfiguredGrantConfigDedicatedSL_PRS_RP_List_r18 struct {
	Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToReleaseList_r18 []SL_ConfigIndexCG_r16
	Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToAddModList_r18  []SL_ConfiguredGrantConfigDedicatedSL_PRS_RP_r18
}

func (ie *SL_ConfiguredGrantConfigDedicatedSL_PRS_RP_List_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLConfiguredGrantConfigDedicatedSLPRSRPListR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToReleaseList_r18 != nil, ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToAddModList_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-ConfiguredGrantConfigDedicatedSL-PRS-RP-ToReleaseList-r18: sequence-of
	{
		if ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLConfiguredGrantConfigDedicatedSLPRSRPListR18SlConfiguredGrantConfigDedicatedSLPRSRPToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToReleaseList_r18 {
				if err := ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-ConfiguredGrantConfigDedicatedSL-PRS-RP-ToAddModList-r18: sequence-of
	{
		if ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLConfiguredGrantConfigDedicatedSLPRSRPListR18SlConfiguredGrantConfigDedicatedSLPRSRPToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToAddModList_r18 {
				if err := ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_ConfiguredGrantConfigDedicatedSL_PRS_RP_List_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLConfiguredGrantConfigDedicatedSLPRSRPListR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-ConfiguredGrantConfigDedicatedSL-PRS-RP-ToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLConfiguredGrantConfigDedicatedSLPRSRPListR18SlConfiguredGrantConfigDedicatedSLPRSRPToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToReleaseList_r18 = make([]SL_ConfigIndexCG_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-ConfiguredGrantConfigDedicatedSL-PRS-RP-ToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLConfiguredGrantConfigDedicatedSLPRSRPListR18SlConfiguredGrantConfigDedicatedSLPRSRPToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToAddModList_r18 = make([]SL_ConfiguredGrantConfigDedicatedSL_PRS_RP_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_ConfiguredGrantConfigDedicatedSL_PRS_RP_ToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
