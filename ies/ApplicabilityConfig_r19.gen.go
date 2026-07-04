// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ApplicabilityConfig-r19 (line 26544).

var applicabilityConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "applicabilityConfigCellId-r19"},
		{Name: "applicabilitySetConfigCSI-ToAddModList-r19", Optional: true},
		{Name: "applicabilitySetConfigCSI-ToReleaseList-r19", Optional: true},
	},
}

var applicabilityConfigR19ApplicabilitySetConfigCSIToAddModListR19Constraints = per.SizeRange(1, common.MaxNrofApplicabilitySetCSI_Configs_r19)

var applicabilityConfigR19ApplicabilitySetConfigCSIToReleaseListR19Constraints = per.SizeRange(1, common.MaxNrofApplicabilitySetCSI_Configs_r19)

type ApplicabilityConfig_r19 struct {
	ApplicabilityConfigCellId_r19               ServCellIndex
	ApplicabilitySetConfigCSI_ToAddModList_r19  []ApplicabilitySetConfigCSI_r19
	ApplicabilitySetConfigCSI_ToReleaseList_r19 []ApplicabilitySetConfigId_r19
}

func (ie *ApplicabilityConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(applicabilityConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ApplicabilitySetConfigCSI_ToAddModList_r19 != nil, ie.ApplicabilitySetConfigCSI_ToReleaseList_r19 != nil}); err != nil {
		return err
	}

	// 3. applicabilityConfigCellId-r19: ref
	{
		if err := ie.ApplicabilityConfigCellId_r19.Encode(e); err != nil {
			return err
		}
	}

	// 4. applicabilitySetConfigCSI-ToAddModList-r19: sequence-of
	{
		if ie.ApplicabilitySetConfigCSI_ToAddModList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(applicabilityConfigR19ApplicabilitySetConfigCSIToAddModListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.ApplicabilitySetConfigCSI_ToAddModList_r19))); err != nil {
				return err
			}
			for i := range ie.ApplicabilitySetConfigCSI_ToAddModList_r19 {
				if err := ie.ApplicabilitySetConfigCSI_ToAddModList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. applicabilitySetConfigCSI-ToReleaseList-r19: sequence-of
	{
		if ie.ApplicabilitySetConfigCSI_ToReleaseList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(applicabilityConfigR19ApplicabilitySetConfigCSIToReleaseListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.ApplicabilitySetConfigCSI_ToReleaseList_r19))); err != nil {
				return err
			}
			for i := range ie.ApplicabilitySetConfigCSI_ToReleaseList_r19 {
				if err := ie.ApplicabilitySetConfigCSI_ToReleaseList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *ApplicabilityConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(applicabilityConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. applicabilityConfigCellId-r19: ref
	{
		if err := ie.ApplicabilityConfigCellId_r19.Decode(d); err != nil {
			return err
		}
	}

	// 4. applicabilitySetConfigCSI-ToAddModList-r19: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(applicabilityConfigR19ApplicabilitySetConfigCSIToAddModListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ApplicabilitySetConfigCSI_ToAddModList_r19 = make([]ApplicabilitySetConfigCSI_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ApplicabilitySetConfigCSI_ToAddModList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. applicabilitySetConfigCSI-ToReleaseList-r19: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(applicabilityConfigR19ApplicabilitySetConfigCSIToReleaseListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ApplicabilitySetConfigCSI_ToReleaseList_r19 = make([]ApplicabilitySetConfigId_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ApplicabilitySetConfigCSI_ToReleaseList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
