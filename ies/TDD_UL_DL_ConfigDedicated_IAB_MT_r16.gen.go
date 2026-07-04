// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TDD-UL-DL-ConfigDedicated-IAB-MT-r16 (line 16101).

var tDDULDLConfigDedicatedIABMTR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "slotSpecificConfigurationsToAddModList-IAB-MT-r16", Optional: true},
		{Name: "slotSpecificConfigurationsToReleaseList-IAB-MT-r16", Optional: true},
	},
}

var tDDULDLConfigDedicatedIABMTR16SlotSpecificConfigurationsToAddModListIABMTR16Constraints = per.SizeRange(1, common.MaxNrofSlots)

var tDDULDLConfigDedicatedIABMTR16SlotSpecificConfigurationsToReleaseListIABMTR16Constraints = per.SizeRange(1, common.MaxNrofSlots)

type TDD_UL_DL_ConfigDedicated_IAB_MT_r16 struct {
	SlotSpecificConfigurationsToAddModList_IAB_MT_r16  []TDD_UL_DL_SlotConfig_IAB_MT_r16
	SlotSpecificConfigurationsToReleaseList_IAB_MT_r16 []TDD_UL_DL_SlotIndex
}

func (ie *TDD_UL_DL_ConfigDedicated_IAB_MT_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tDDULDLConfigDedicatedIABMTR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SlotSpecificConfigurationsToAddModList_IAB_MT_r16 != nil, ie.SlotSpecificConfigurationsToReleaseList_IAB_MT_r16 != nil}); err != nil {
		return err
	}

	// 3. slotSpecificConfigurationsToAddModList-IAB-MT-r16: sequence-of
	{
		if ie.SlotSpecificConfigurationsToAddModList_IAB_MT_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(tDDULDLConfigDedicatedIABMTR16SlotSpecificConfigurationsToAddModListIABMTR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SlotSpecificConfigurationsToAddModList_IAB_MT_r16))); err != nil {
				return err
			}
			for i := range ie.SlotSpecificConfigurationsToAddModList_IAB_MT_r16 {
				if err := ie.SlotSpecificConfigurationsToAddModList_IAB_MT_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. slotSpecificConfigurationsToReleaseList-IAB-MT-r16: sequence-of
	{
		if ie.SlotSpecificConfigurationsToReleaseList_IAB_MT_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(tDDULDLConfigDedicatedIABMTR16SlotSpecificConfigurationsToReleaseListIABMTR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SlotSpecificConfigurationsToReleaseList_IAB_MT_r16))); err != nil {
				return err
			}
			for i := range ie.SlotSpecificConfigurationsToReleaseList_IAB_MT_r16 {
				if err := ie.SlotSpecificConfigurationsToReleaseList_IAB_MT_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *TDD_UL_DL_ConfigDedicated_IAB_MT_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tDDULDLConfigDedicatedIABMTR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. slotSpecificConfigurationsToAddModList-IAB-MT-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(tDDULDLConfigDedicatedIABMTR16SlotSpecificConfigurationsToAddModListIABMTR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SlotSpecificConfigurationsToAddModList_IAB_MT_r16 = make([]TDD_UL_DL_SlotConfig_IAB_MT_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SlotSpecificConfigurationsToAddModList_IAB_MT_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. slotSpecificConfigurationsToReleaseList-IAB-MT-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(tDDULDLConfigDedicatedIABMTR16SlotSpecificConfigurationsToReleaseListIABMTR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SlotSpecificConfigurationsToReleaseList_IAB_MT_r16 = make([]TDD_UL_DL_SlotIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SlotSpecificConfigurationsToReleaseList_IAB_MT_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
