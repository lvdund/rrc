// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AvailabilityIndicator-r16 (line 5087).

var availabilityIndicatorR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ai-RNTI-r16"},
		{Name: "dci-PayloadSizeAI-r16"},
		{Name: "availableCombToAddModList-r16", Optional: true},
		{Name: "availableCombToReleaseList-r16", Optional: true},
	},
}

var availabilityIndicatorR16DciPayloadSizeAIR16Constraints = per.Constrained(1, common.MaxAI_DCI_PayloadSize_r16)

var availabilityIndicatorR16AvailableCombToAddModListR16Constraints = per.SizeRange(1, common.MaxNrofDUCells_r16)

var availabilityIndicatorR16AvailableCombToReleaseListR16Constraints = per.SizeRange(1, common.MaxNrofDUCells_r16)

type AvailabilityIndicator_r16 struct {
	Ai_RNTI_r16                    AI_RNTI_r16
	Dci_PayloadSizeAI_r16          int64
	AvailableCombToAddModList_r16  []AvailabilityCombinationsPerCell_r16
	AvailableCombToReleaseList_r16 []AvailabilityCombinationsPerCellIndex_r16
}

func (ie *AvailabilityIndicator_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(availabilityIndicatorR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AvailableCombToAddModList_r16 != nil, ie.AvailableCombToReleaseList_r16 != nil}); err != nil {
		return err
	}

	// 3. ai-RNTI-r16: ref
	{
		if err := ie.Ai_RNTI_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. dci-PayloadSizeAI-r16: integer
	{
		if err := e.EncodeInteger(ie.Dci_PayloadSizeAI_r16, availabilityIndicatorR16DciPayloadSizeAIR16Constraints); err != nil {
			return err
		}
	}

	// 5. availableCombToAddModList-r16: sequence-of
	{
		if ie.AvailableCombToAddModList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(availabilityIndicatorR16AvailableCombToAddModListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.AvailableCombToAddModList_r16))); err != nil {
				return err
			}
			for i := range ie.AvailableCombToAddModList_r16 {
				if err := ie.AvailableCombToAddModList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. availableCombToReleaseList-r16: sequence-of
	{
		if ie.AvailableCombToReleaseList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(availabilityIndicatorR16AvailableCombToReleaseListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.AvailableCombToReleaseList_r16))); err != nil {
				return err
			}
			for i := range ie.AvailableCombToReleaseList_r16 {
				if err := ie.AvailableCombToReleaseList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *AvailabilityIndicator_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(availabilityIndicatorR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ai-RNTI-r16: ref
	{
		if err := ie.Ai_RNTI_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. dci-PayloadSizeAI-r16: integer
	{
		v1, err := d.DecodeInteger(availabilityIndicatorR16DciPayloadSizeAIR16Constraints)
		if err != nil {
			return err
		}
		ie.Dci_PayloadSizeAI_r16 = v1
	}

	// 5. availableCombToAddModList-r16: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(availabilityIndicatorR16AvailableCombToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AvailableCombToAddModList_r16 = make([]AvailabilityCombinationsPerCell_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AvailableCombToAddModList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. availableCombToReleaseList-r16: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(availabilityIndicatorR16AvailableCombToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AvailableCombToReleaseList_r16 = make([]AvailabilityCombinationsPerCellIndex_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AvailableCombToReleaseList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
