// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SlotFormatCombinationsPerCell (line 15137).

var slotFormatCombinationsPerCellConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "servingCellId"},
		{Name: "subcarrierSpacing"},
		{Name: "subcarrierSpacing2", Optional: true},
		{Name: "slotFormatCombinations", Optional: true},
		{Name: "positionInDCI", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var slotFormatCombinationsPerCellSlotFormatCombinationsConstraints = per.SizeRange(1, common.MaxNrofSlotFormatCombinationsPerSet)

var slotFormatCombinationsPerCellPositionInDCIConstraints = per.Constrained(0, common.MaxSFI_DCI_PayloadSize_1)

const (
	SlotFormatCombinationsPerCell_Ext_EnableConfiguredUL_r16_Enabled = 0
)

var slotFormatCombinationsPerCellExtEnableConfiguredULR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SlotFormatCombinationsPerCell_Ext_EnableConfiguredUL_r16_Enabled},
}

type SlotFormatCombinationsPerCell struct {
	ServingCellId          ServCellIndex
	SubcarrierSpacing      SubcarrierSpacing
	SubcarrierSpacing2     *SubcarrierSpacing
	SlotFormatCombinations []SlotFormatCombination
	PositionInDCI          *int64
	EnableConfiguredUL_r16 *int64
}

func (ie *SlotFormatCombinationsPerCell) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(slotFormatCombinationsPerCellConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.EnableConfiguredUL_r16 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SubcarrierSpacing2 != nil, ie.SlotFormatCombinations != nil, ie.PositionInDCI != nil}); err != nil {
		return err
	}

	// 3. servingCellId: ref
	{
		if err := ie.ServingCellId.Encode(e); err != nil {
			return err
		}
	}

	// 4. subcarrierSpacing: ref
	{
		if err := ie.SubcarrierSpacing.Encode(e); err != nil {
			return err
		}
	}

	// 5. subcarrierSpacing2: ref
	{
		if ie.SubcarrierSpacing2 != nil {
			if err := ie.SubcarrierSpacing2.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. slotFormatCombinations: sequence-of
	{
		if ie.SlotFormatCombinations != nil {
			seqOf := e.NewSequenceOfEncoder(slotFormatCombinationsPerCellSlotFormatCombinationsConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SlotFormatCombinations))); err != nil {
				return err
			}
			for i := range ie.SlotFormatCombinations {
				if err := ie.SlotFormatCombinations[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. positionInDCI: integer
	{
		if ie.PositionInDCI != nil {
			if err := e.EncodeInteger(*ie.PositionInDCI, slotFormatCombinationsPerCellPositionInDCIConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "enableConfiguredUL-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EnableConfiguredUL_r16 != nil}); err != nil {
				return err
			}

			if ie.EnableConfiguredUL_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnableConfiguredUL_r16, slotFormatCombinationsPerCellExtEnableConfiguredULR16Constraints); err != nil {
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

func (ie *SlotFormatCombinationsPerCell) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(slotFormatCombinationsPerCellConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. servingCellId: ref
	{
		if err := ie.ServingCellId.Decode(d); err != nil {
			return err
		}
	}

	// 4. subcarrierSpacing: ref
	{
		if err := ie.SubcarrierSpacing.Decode(d); err != nil {
			return err
		}
	}

	// 5. subcarrierSpacing2: ref
	{
		if seq.IsComponentPresent(2) {
			ie.SubcarrierSpacing2 = new(SubcarrierSpacing)
			if err := ie.SubcarrierSpacing2.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. slotFormatCombinations: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(slotFormatCombinationsPerCellSlotFormatCombinationsConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SlotFormatCombinations = make([]SlotFormatCombination, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SlotFormatCombinations[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. positionInDCI: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(slotFormatCombinationsPerCellPositionInDCIConstraints)
			if err != nil {
				return err
			}
			ie.PositionInDCI = &v
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
				{Name: "enableConfiguredUL-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(slotFormatCombinationsPerCellExtEnableConfiguredULR16Constraints)
			if err != nil {
				return err
			}
			ie.EnableConfiguredUL_r16 = &v
		}
	}

	return nil
}
