// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LTM-TCI-Info-r18 (line 8965).

var lTMTCIInfoR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ltm-DL-OrJointTCI-StateToAddModList-r18", Optional: true},
		{Name: "ltm-DL-OrJointTCI-StateToReleaseList-r18", Optional: true},
		{Name: "ltm-UL-TCI-StateToAddModList-r18", Optional: true},
		{Name: "ltm-UL-TCI-StateToReleaseList-r18", Optional: true},
		{Name: "ltm-NZP-CSI-RS-ResourceToAddModList-r18", Optional: true},
		{Name: "ltm-NZP-CSI-RS-ResourceToReleaseList-r18", Optional: true},
		{Name: "ltm-NZP-CSI-RS-ResourceSetToAddModList-r18", Optional: true},
		{Name: "ltm-NZP-CSI-RS-ResourceSetToReleaseList-r18", Optional: true},
		{Name: "pathlossReferenceRS-ToAddModList-r18", Optional: true},
		{Name: "pathlossReferenceRS-ToReleaseList-r18", Optional: true},
		{Name: "unifiedTCI-StateType-r18", Optional: true},
	},
}

var lTMTCIInfoR18LtmDLOrJointTCIStateToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofCandidateTCI_State_r18)

var lTMTCIInfoR18LtmDLOrJointTCIStateToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofCandidateTCI_State_r18)

var lTMTCIInfoR18LtmULTCIStateToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofCandidateUL_TCI_r18)

var lTMTCIInfoR18LtmULTCIStateToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofCandidateUL_TCI_r18)

var lTMTCIInfoR18LtmNZPCSIRSResourceToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_Resources)

var lTMTCIInfoR18LtmNZPCSIRSResourceToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_Resources)

var lTMTCIInfoR18LtmNZPCSIRSResourceSetToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_ResourceSets)

var lTMTCIInfoR18LtmNZPCSIRSResourceSetToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_ResourceSets)

var lTMTCIInfoR18PathlossReferenceRSToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofPathlossReferenceRSs_r17)

var lTMTCIInfoR18PathlossReferenceRSToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofPathlossReferenceRSs_r17)

const (
	LTM_TCI_Info_r18_UnifiedTCI_StateType_r18_Separate = 0
	LTM_TCI_Info_r18_UnifiedTCI_StateType_r18_Joint    = 1
)

var lTMTCIInfoR18UnifiedTCIStateTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_TCI_Info_r18_UnifiedTCI_StateType_r18_Separate, LTM_TCI_Info_r18_UnifiedTCI_StateType_r18_Joint},
}

type LTM_TCI_Info_r18 struct {
	Ltm_DL_OrJointTCI_StateToAddModList_r18     []CandidateTCI_State_r18
	Ltm_DL_OrJointTCI_StateToReleaseList_r18    []TCI_StateId
	Ltm_UL_TCI_StateToAddModList_r18            []CandidateTCI_UL_State_r18
	Ltm_UL_TCI_StateToReleaseList_r18           []TCI_UL_StateId_r17
	Ltm_NZP_CSI_RS_ResourceToAddModList_r18     []NZP_CSI_RS_Resource
	Ltm_NZP_CSI_RS_ResourceToReleaseList_r18    []NZP_CSI_RS_ResourceId
	Ltm_NZP_CSI_RS_ResourceSetToAddModList_r18  []NZP_CSI_RS_ResourceSet
	Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r18 []NZP_CSI_RS_ResourceSetId
	PathlossReferenceRS_ToAddModList_r18        []PathlossReferenceRS_r17
	PathlossReferenceRS_ToReleaseList_r18       []PathlossReferenceRS_Id_r17
	UnifiedTCI_StateType_r18                    *int64
}

func (ie *LTM_TCI_Info_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMTCIInfoR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ltm_DL_OrJointTCI_StateToAddModList_r18 != nil, ie.Ltm_DL_OrJointTCI_StateToReleaseList_r18 != nil, ie.Ltm_UL_TCI_StateToAddModList_r18 != nil, ie.Ltm_UL_TCI_StateToReleaseList_r18 != nil, ie.Ltm_NZP_CSI_RS_ResourceToAddModList_r18 != nil, ie.Ltm_NZP_CSI_RS_ResourceToReleaseList_r18 != nil, ie.Ltm_NZP_CSI_RS_ResourceSetToAddModList_r18 != nil, ie.Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r18 != nil, ie.PathlossReferenceRS_ToAddModList_r18 != nil, ie.PathlossReferenceRS_ToReleaseList_r18 != nil, ie.UnifiedTCI_StateType_r18 != nil}); err != nil {
		return err
	}

	// 3. ltm-DL-OrJointTCI-StateToAddModList-r18: sequence-of
	{
		if ie.Ltm_DL_OrJointTCI_StateToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(lTMTCIInfoR18LtmDLOrJointTCIStateToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ltm_DL_OrJointTCI_StateToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.Ltm_DL_OrJointTCI_StateToAddModList_r18 {
				if err := ie.Ltm_DL_OrJointTCI_StateToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. ltm-DL-OrJointTCI-StateToReleaseList-r18: sequence-of
	{
		if ie.Ltm_DL_OrJointTCI_StateToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(lTMTCIInfoR18LtmDLOrJointTCIStateToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ltm_DL_OrJointTCI_StateToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.Ltm_DL_OrJointTCI_StateToReleaseList_r18 {
				if err := ie.Ltm_DL_OrJointTCI_StateToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. ltm-UL-TCI-StateToAddModList-r18: sequence-of
	{
		if ie.Ltm_UL_TCI_StateToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(lTMTCIInfoR18LtmULTCIStateToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ltm_UL_TCI_StateToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.Ltm_UL_TCI_StateToAddModList_r18 {
				if err := ie.Ltm_UL_TCI_StateToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. ltm-UL-TCI-StateToReleaseList-r18: sequence-of
	{
		if ie.Ltm_UL_TCI_StateToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(lTMTCIInfoR18LtmULTCIStateToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ltm_UL_TCI_StateToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.Ltm_UL_TCI_StateToReleaseList_r18 {
				if err := ie.Ltm_UL_TCI_StateToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. ltm-NZP-CSI-RS-ResourceToAddModList-r18: sequence-of
	{
		if ie.Ltm_NZP_CSI_RS_ResourceToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(lTMTCIInfoR18LtmNZPCSIRSResourceToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ltm_NZP_CSI_RS_ResourceToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.Ltm_NZP_CSI_RS_ResourceToAddModList_r18 {
				if err := ie.Ltm_NZP_CSI_RS_ResourceToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. ltm-NZP-CSI-RS-ResourceToReleaseList-r18: sequence-of
	{
		if ie.Ltm_NZP_CSI_RS_ResourceToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(lTMTCIInfoR18LtmNZPCSIRSResourceToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ltm_NZP_CSI_RS_ResourceToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.Ltm_NZP_CSI_RS_ResourceToReleaseList_r18 {
				if err := ie.Ltm_NZP_CSI_RS_ResourceToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 9. ltm-NZP-CSI-RS-ResourceSetToAddModList-r18: sequence-of
	{
		if ie.Ltm_NZP_CSI_RS_ResourceSetToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(lTMTCIInfoR18LtmNZPCSIRSResourceSetToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ltm_NZP_CSI_RS_ResourceSetToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.Ltm_NZP_CSI_RS_ResourceSetToAddModList_r18 {
				if err := ie.Ltm_NZP_CSI_RS_ResourceSetToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 10. ltm-NZP-CSI-RS-ResourceSetToReleaseList-r18: sequence-of
	{
		if ie.Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(lTMTCIInfoR18LtmNZPCSIRSResourceSetToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r18 {
				if err := ie.Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 11. pathlossReferenceRS-ToAddModList-r18: sequence-of
	{
		if ie.PathlossReferenceRS_ToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(lTMTCIInfoR18PathlossReferenceRSToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.PathlossReferenceRS_ToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.PathlossReferenceRS_ToAddModList_r18 {
				if err := ie.PathlossReferenceRS_ToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 12. pathlossReferenceRS-ToReleaseList-r18: sequence-of
	{
		if ie.PathlossReferenceRS_ToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(lTMTCIInfoR18PathlossReferenceRSToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.PathlossReferenceRS_ToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.PathlossReferenceRS_ToReleaseList_r18 {
				if err := ie.PathlossReferenceRS_ToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 13. unifiedTCI-StateType-r18: enumerated
	{
		if ie.UnifiedTCI_StateType_r18 != nil {
			if err := e.EncodeEnumerated(*ie.UnifiedTCI_StateType_r18, lTMTCIInfoR18UnifiedTCIStateTypeR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LTM_TCI_Info_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMTCIInfoR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ltm-DL-OrJointTCI-StateToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(lTMTCIInfoR18LtmDLOrJointTCIStateToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ltm_DL_OrJointTCI_StateToAddModList_r18 = make([]CandidateTCI_State_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ltm_DL_OrJointTCI_StateToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. ltm-DL-OrJointTCI-StateToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(lTMTCIInfoR18LtmDLOrJointTCIStateToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ltm_DL_OrJointTCI_StateToReleaseList_r18 = make([]TCI_StateId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ltm_DL_OrJointTCI_StateToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. ltm-UL-TCI-StateToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(lTMTCIInfoR18LtmULTCIStateToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ltm_UL_TCI_StateToAddModList_r18 = make([]CandidateTCI_UL_State_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ltm_UL_TCI_StateToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. ltm-UL-TCI-StateToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(lTMTCIInfoR18LtmULTCIStateToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ltm_UL_TCI_StateToReleaseList_r18 = make([]TCI_UL_StateId_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ltm_UL_TCI_StateToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. ltm-NZP-CSI-RS-ResourceToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(lTMTCIInfoR18LtmNZPCSIRSResourceToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ltm_NZP_CSI_RS_ResourceToAddModList_r18 = make([]NZP_CSI_RS_Resource, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ltm_NZP_CSI_RS_ResourceToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. ltm-NZP-CSI-RS-ResourceToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(lTMTCIInfoR18LtmNZPCSIRSResourceToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ltm_NZP_CSI_RS_ResourceToReleaseList_r18 = make([]NZP_CSI_RS_ResourceId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ltm_NZP_CSI_RS_ResourceToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. ltm-NZP-CSI-RS-ResourceSetToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(lTMTCIInfoR18LtmNZPCSIRSResourceSetToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ltm_NZP_CSI_RS_ResourceSetToAddModList_r18 = make([]NZP_CSI_RS_ResourceSet, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ltm_NZP_CSI_RS_ResourceSetToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. ltm-NZP-CSI-RS-ResourceSetToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(7) {
			seqOf := d.NewSequenceOfDecoder(lTMTCIInfoR18LtmNZPCSIRSResourceSetToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r18 = make([]NZP_CSI_RS_ResourceSetId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ltm_NZP_CSI_RS_ResourceSetToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 11. pathlossReferenceRS-ToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(8) {
			seqOf := d.NewSequenceOfDecoder(lTMTCIInfoR18PathlossReferenceRSToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PathlossReferenceRS_ToAddModList_r18 = make([]PathlossReferenceRS_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PathlossReferenceRS_ToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 12. pathlossReferenceRS-ToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(9) {
			seqOf := d.NewSequenceOfDecoder(lTMTCIInfoR18PathlossReferenceRSToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PathlossReferenceRS_ToReleaseList_r18 = make([]PathlossReferenceRS_Id_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PathlossReferenceRS_ToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 13. unifiedTCI-StateType-r18: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(lTMTCIInfoR18UnifiedTCIStateTypeR18Constraints)
			if err != nil {
				return err
			}
			ie.UnifiedTCI_StateType_r18 = &idx
		}
	}

	return nil
}
