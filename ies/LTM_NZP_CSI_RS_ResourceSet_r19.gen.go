// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-NZP-CSI-RS-ResourceSet-r19 (line 8931).

var lTMNZPCSIRSResourceSetR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ltm-CSI-RS-ResourceList-r19"},
		{Name: "ltm-CandidateIdList-r19"},
		{Name: "repetition-r19", Optional: true},
	},
}

var lTMNZPCSIRSResourceSetR19LtmCSIRSResourceListR19Constraints = per.SizeRange(1, common.MaxNrofLTM_CSI_ResourcesPerSet_r18)

var lTMNZPCSIRSResourceSetR19LtmCandidateIdListR19Constraints = per.SizeRange(1, common.MaxNrofLTM_CSI_ResourcesPerSet_r18)

const (
	LTM_NZP_CSI_RS_ResourceSet_r19_Repetition_r19_Off = 0
)

var lTMNZPCSIRSResourceSetR19RepetitionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_NZP_CSI_RS_ResourceSet_r19_Repetition_r19_Off},
}

type LTM_NZP_CSI_RS_ResourceSet_r19 struct {
	Ltm_CSI_RS_ResourceList_r19 []NZP_CSI_RS_ResourceId
	Ltm_CandidateIdList_r19     []LTM_CandidateId_r18
	Repetition_r19              *int64
}

func (ie *LTM_NZP_CSI_RS_ResourceSet_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMNZPCSIRSResourceSetR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Repetition_r19 != nil}); err != nil {
		return err
	}

	// 3. ltm-CSI-RS-ResourceList-r19: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(lTMNZPCSIRSResourceSetR19LtmCSIRSResourceListR19Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Ltm_CSI_RS_ResourceList_r19))); err != nil {
			return err
		}
		for i := range ie.Ltm_CSI_RS_ResourceList_r19 {
			if err := ie.Ltm_CSI_RS_ResourceList_r19[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ltm-CandidateIdList-r19: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(lTMNZPCSIRSResourceSetR19LtmCandidateIdListR19Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Ltm_CandidateIdList_r19))); err != nil {
			return err
		}
		for i := range ie.Ltm_CandidateIdList_r19 {
			if err := ie.Ltm_CandidateIdList_r19[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. repetition-r19: enumerated
	{
		if ie.Repetition_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Repetition_r19, lTMNZPCSIRSResourceSetR19RepetitionR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LTM_NZP_CSI_RS_ResourceSet_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMNZPCSIRSResourceSetR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ltm-CSI-RS-ResourceList-r19: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(lTMNZPCSIRSResourceSetR19LtmCSIRSResourceListR19Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Ltm_CSI_RS_ResourceList_r19 = make([]NZP_CSI_RS_ResourceId, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Ltm_CSI_RS_ResourceList_r19[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ltm-CandidateIdList-r19: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(lTMNZPCSIRSResourceSetR19LtmCandidateIdListR19Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Ltm_CandidateIdList_r19 = make([]LTM_CandidateId_r18, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Ltm_CandidateIdList_r19[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. repetition-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(lTMNZPCSIRSResourceSetR19RepetitionR19Constraints)
			if err != nil {
				return err
			}
			ie.Repetition_r19 = &idx
		}
	}

	return nil
}
