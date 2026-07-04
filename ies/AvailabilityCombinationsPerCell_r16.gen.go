// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AvailabilityCombinationsPerCell-r16 (line 5049).

var availabilityCombinationsPerCellR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "availabilityCombinationsPerCellIndex-r16"},
		{Name: "iab-DU-CellIdentity-r16"},
		{Name: "positionInDCI-AI-r16", Optional: true},
		{Name: "availabilityCombinations-r16"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var availabilityCombinationsPerCellR16PositionInDCIAIR16Constraints = per.Constrained(0, common.MaxAI_DCI_PayloadSize_1_r16)

var availabilityCombinationsPerCellR16AvailabilityCombinationsR16Constraints = per.SizeRange(1, common.MaxNrofAvailabilityCombinationsPerSet_r16)

var availabilityCombinationsPerCellR16ExtAvailabilityCombinationsRBGroupsR17Constraints = per.SizeRange(1, common.MaxNrofAvailabilityCombinationsPerSet_r16)

var availabilityCombinationsPerCellR16PositionInDCIAIRBGroupsV1720Constraints = per.Constrained(0, common.MaxAI_DCI_PayloadSize_1_r16)

type AvailabilityCombinationsPerCell_r16 struct {
	AvailabilityCombinationsPerCellIndex_r16 AvailabilityCombinationsPerCellIndex_r16
	Iab_DU_CellIdentity_r16                  CellIdentity
	PositionInDCI_AI_r16                     *int64
	AvailabilityCombinations_r16             []AvailabilityCombination_r16
	AvailabilityCombinationsRB_Groups_r17    []AvailabilityCombinationRB_Groups_r17
	PositionInDCI_AI_RBGroups_v1720          *int64
}

func (ie *AvailabilityCombinationsPerCell_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(availabilityCombinationsPerCellR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.AvailabilityCombinationsRB_Groups_r17 != nil
	hasExtGroup1 := ie.PositionInDCI_AI_RBGroups_v1720 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PositionInDCI_AI_r16 != nil}); err != nil {
		return err
	}

	// 3. availabilityCombinationsPerCellIndex-r16: ref
	{
		if err := ie.AvailabilityCombinationsPerCellIndex_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. iab-DU-CellIdentity-r16: ref
	{
		if err := ie.Iab_DU_CellIdentity_r16.Encode(e); err != nil {
			return err
		}
	}

	// 5. positionInDCI-AI-r16: integer
	{
		if ie.PositionInDCI_AI_r16 != nil {
			if err := e.EncodeInteger(*ie.PositionInDCI_AI_r16, availabilityCombinationsPerCellR16PositionInDCIAIR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. availabilityCombinations-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(availabilityCombinationsPerCellR16AvailabilityCombinationsR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.AvailabilityCombinations_r16))); err != nil {
			return err
		}
		for i := range ie.AvailabilityCombinations_r16 {
			if err := ie.AvailabilityCombinations_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "availabilityCombinationsRB-Groups-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AvailabilityCombinationsRB_Groups_r17 != nil}); err != nil {
				return err
			}

			if ie.AvailabilityCombinationsRB_Groups_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(availabilityCombinationsPerCellR16ExtAvailabilityCombinationsRBGroupsR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.AvailabilityCombinationsRB_Groups_r17))); err != nil {
					return err
				}
				for i := range ie.AvailabilityCombinationsRB_Groups_r17 {
					if err := ie.AvailabilityCombinationsRB_Groups_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "positionInDCI-AI-RBGroups-v1720", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PositionInDCI_AI_RBGroups_v1720 != nil}); err != nil {
				return err
			}

			if ie.PositionInDCI_AI_RBGroups_v1720 != nil {
				if err := ex.EncodeInteger(*ie.PositionInDCI_AI_RBGroups_v1720, availabilityCombinationsPerCellR16PositionInDCIAIRBGroupsV1720Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *AvailabilityCombinationsPerCell_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(availabilityCombinationsPerCellR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. availabilityCombinationsPerCellIndex-r16: ref
	{
		if err := ie.AvailabilityCombinationsPerCellIndex_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. iab-DU-CellIdentity-r16: ref
	{
		if err := ie.Iab_DU_CellIdentity_r16.Decode(d); err != nil {
			return err
		}
	}

	// 5. positionInDCI-AI-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(availabilityCombinationsPerCellR16PositionInDCIAIR16Constraints)
			if err != nil {
				return err
			}
			ie.PositionInDCI_AI_r16 = &v
		}
	}

	// 6. availabilityCombinations-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(availabilityCombinationsPerCellR16AvailabilityCombinationsR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.AvailabilityCombinations_r16 = make([]AvailabilityCombination_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.AvailabilityCombinations_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "availabilityCombinationsRB-Groups-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(availabilityCombinationsPerCellR16ExtAvailabilityCombinationsRBGroupsR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AvailabilityCombinationsRB_Groups_r17 = make([]AvailabilityCombinationRB_Groups_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AvailabilityCombinationsRB_Groups_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "positionInDCI-AI-RBGroups-v1720", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(availabilityCombinationsPerCellR16PositionInDCIAIRBGroupsV1720Constraints)
			if err != nil {
				return err
			}
			ie.PositionInDCI_AI_RBGroups_v1720 = &v
		}
	}

	return nil
}
