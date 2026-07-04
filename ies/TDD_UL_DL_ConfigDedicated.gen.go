// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TDD-UL-DL-ConfigDedicated (line 16095).

var tDDULDLConfigDedicatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "slotSpecificConfigurationsToAddModList", Optional: true},
		{Name: "slotSpecificConfigurationsToReleaseList", Optional: true},
	},
}

var tDDULDLConfigDedicatedSlotSpecificConfigurationsToAddModListConstraints = per.SizeRange(1, common.MaxNrofSlots)

var tDDULDLConfigDedicatedSlotSpecificConfigurationsToReleaseListConstraints = per.SizeRange(1, common.MaxNrofSlots)

type TDD_UL_DL_ConfigDedicated struct {
	SlotSpecificConfigurationsToAddModList  []TDD_UL_DL_SlotConfig
	SlotSpecificConfigurationsToReleaseList []TDD_UL_DL_SlotIndex
}

func (ie *TDD_UL_DL_ConfigDedicated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tDDULDLConfigDedicatedConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SlotSpecificConfigurationsToAddModList != nil, ie.SlotSpecificConfigurationsToReleaseList != nil}); err != nil {
		return err
	}

	// 3. slotSpecificConfigurationsToAddModList: sequence-of
	{
		if ie.SlotSpecificConfigurationsToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(tDDULDLConfigDedicatedSlotSpecificConfigurationsToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SlotSpecificConfigurationsToAddModList))); err != nil {
				return err
			}
			for i := range ie.SlotSpecificConfigurationsToAddModList {
				if err := ie.SlotSpecificConfigurationsToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. slotSpecificConfigurationsToReleaseList: sequence-of
	{
		if ie.SlotSpecificConfigurationsToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(tDDULDLConfigDedicatedSlotSpecificConfigurationsToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SlotSpecificConfigurationsToReleaseList))); err != nil {
				return err
			}
			for i := range ie.SlotSpecificConfigurationsToReleaseList {
				if err := ie.SlotSpecificConfigurationsToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *TDD_UL_DL_ConfigDedicated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tDDULDLConfigDedicatedConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. slotSpecificConfigurationsToAddModList: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(tDDULDLConfigDedicatedSlotSpecificConfigurationsToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SlotSpecificConfigurationsToAddModList = make([]TDD_UL_DL_SlotConfig, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SlotSpecificConfigurationsToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. slotSpecificConfigurationsToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(tDDULDLConfigDedicatedSlotSpecificConfigurationsToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SlotSpecificConfigurationsToReleaseList = make([]TDD_UL_DL_SlotIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SlotSpecificConfigurationsToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
