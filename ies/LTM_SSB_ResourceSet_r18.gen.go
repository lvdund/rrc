// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-SSB-ResourceSet-r18 (line 8925).

var lTMSSBResourceSetR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ltm-SSB-ResourceList-r18"},
		{Name: "ltm-CandidateIdList-r18"},
	},
}

var lTMSSBResourceSetR18LtmSSBResourceListR18Constraints = per.SizeRange(1, common.MaxNrofLTM_CSI_ResourcesPerSet_r18)

var lTMSSBResourceSetR18LtmCandidateIdListR18Constraints = per.SizeRange(1, common.MaxNrofLTM_CSI_ResourcesPerSet_r18)

type LTM_SSB_ResourceSet_r18 struct {
	Ltm_SSB_ResourceList_r18 []SSB_Index
	Ltm_CandidateIdList_r18  []LTM_CandidateId_r18
}

func (ie *LTM_SSB_ResourceSet_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMSSBResourceSetR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. ltm-SSB-ResourceList-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(lTMSSBResourceSetR18LtmSSBResourceListR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Ltm_SSB_ResourceList_r18))); err != nil {
			return err
		}
		for i := range ie.Ltm_SSB_ResourceList_r18 {
			if err := ie.Ltm_SSB_ResourceList_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ltm-CandidateIdList-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(lTMSSBResourceSetR18LtmCandidateIdListR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Ltm_CandidateIdList_r18))); err != nil {
			return err
		}
		for i := range ie.Ltm_CandidateIdList_r18 {
			if err := ie.Ltm_CandidateIdList_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LTM_SSB_ResourceSet_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMSSBResourceSetR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. ltm-SSB-ResourceList-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(lTMSSBResourceSetR18LtmSSBResourceListR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Ltm_SSB_ResourceList_r18 = make([]SSB_Index, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Ltm_SSB_ResourceList_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ltm-CandidateIdList-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(lTMSSBResourceSetR18LtmCandidateIdListR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Ltm_CandidateIdList_r18 = make([]LTM_CandidateId_r18, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Ltm_CandidateIdList_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
