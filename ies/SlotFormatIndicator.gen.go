// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SlotFormatIndicator (line 15161).

var slotFormatIndicatorConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sfi-RNTI"},
		{Name: "dci-PayloadSize"},
		{Name: "slotFormatCombToAddModList", Optional: true},
		{Name: "slotFormatCombToReleaseList", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var slotFormatIndicatorDciPayloadSizeConstraints = per.Constrained(1, common.MaxSFI_DCI_PayloadSize)

var slotFormatIndicatorSlotFormatCombToAddModListConstraints = per.SizeRange(1, common.MaxNrofAggregatedCellsPerCellGroup)

var slotFormatIndicatorSlotFormatCombToReleaseListConstraints = per.SizeRange(1, common.MaxNrofAggregatedCellsPerCellGroup)

var slotFormatIndicatorExtAvailableRBSetsToAddModListR16Constraints = per.SizeRange(1, common.MaxNrofAggregatedCellsPerCellGroup)

var slotFormatIndicatorExtAvailableRBSetsToReleaseListR16Constraints = per.SizeRange(1, common.MaxNrofAggregatedCellsPerCellGroup)

var slotFormatIndicatorExtSwitchTriggerToAddModListR16Constraints = per.SizeRange(1, 4)

var slotFormatIndicatorExtSwitchTriggerToReleaseListR16Constraints = per.SizeRange(1, 4)

var slotFormatIndicatorExtCoDurationsPerCellToAddModListR16Constraints = per.SizeRange(1, common.MaxNrofAggregatedCellsPerCellGroup)

var slotFormatIndicatorExtCoDurationsPerCellToReleaseListR16Constraints = per.SizeRange(1, common.MaxNrofAggregatedCellsPerCellGroup)

var slotFormatIndicatorExtSwitchTriggerToAddModListSizeExtR16Constraints = per.SizeRange(1, common.MaxNrofAggregatedCellsPerCellGroupMinus4_r16)

var slotFormatIndicatorExtSwitchTriggerToReleaseListSizeExtR16Constraints = per.SizeRange(1, common.MaxNrofAggregatedCellsPerCellGroupMinus4_r16)

var slotFormatIndicatorExtCoDurationsPerCellToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofAggregatedCellsPerCellGroup)

type SlotFormatIndicator struct {
	Sfi_RNTI                              RNTI_Value
	Dci_PayloadSize                       int64
	SlotFormatCombToAddModList            []SlotFormatCombinationsPerCell
	SlotFormatCombToReleaseList           []ServCellIndex
	AvailableRB_SetsToAddModList_r16      []AvailableRB_SetsPerCell_r16
	AvailableRB_SetsToReleaseList_r16     []ServCellIndex
	SwitchTriggerToAddModList_r16         []SearchSpaceSwitchTrigger_r16
	SwitchTriggerToReleaseList_r16        []ServCellIndex
	Co_DurationsPerCellToAddModList_r16   []CO_DurationsPerCell_r16
	Co_DurationsPerCellToReleaseList_r16  []ServCellIndex
	SwitchTriggerToAddModListSizeExt_r16  []SearchSpaceSwitchTrigger_r16
	SwitchTriggerToReleaseListSizeExt_r16 []ServCellIndex
	Co_DurationsPerCellToAddModList_r17   []CO_DurationsPerCell_r17
}

func (ie *SlotFormatIndicator) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(slotFormatIndicatorConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.AvailableRB_SetsToAddModList_r16 != nil || ie.AvailableRB_SetsToReleaseList_r16 != nil || ie.SwitchTriggerToAddModList_r16 != nil || ie.SwitchTriggerToReleaseList_r16 != nil || ie.Co_DurationsPerCellToAddModList_r16 != nil || ie.Co_DurationsPerCellToReleaseList_r16 != nil
	hasExtGroup1 := ie.SwitchTriggerToAddModListSizeExt_r16 != nil || ie.SwitchTriggerToReleaseListSizeExt_r16 != nil
	hasExtGroup2 := ie.Co_DurationsPerCellToAddModList_r17 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SlotFormatCombToAddModList != nil, ie.SlotFormatCombToReleaseList != nil}); err != nil {
		return err
	}

	// 3. sfi-RNTI: ref
	{
		if err := ie.Sfi_RNTI.Encode(e); err != nil {
			return err
		}
	}

	// 4. dci-PayloadSize: integer
	{
		if err := e.EncodeInteger(ie.Dci_PayloadSize, slotFormatIndicatorDciPayloadSizeConstraints); err != nil {
			return err
		}
	}

	// 5. slotFormatCombToAddModList: sequence-of
	{
		if ie.SlotFormatCombToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(slotFormatIndicatorSlotFormatCombToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SlotFormatCombToAddModList))); err != nil {
				return err
			}
			for i := range ie.SlotFormatCombToAddModList {
				if err := ie.SlotFormatCombToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. slotFormatCombToReleaseList: sequence-of
	{
		if ie.SlotFormatCombToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(slotFormatIndicatorSlotFormatCombToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.SlotFormatCombToReleaseList))); err != nil {
				return err
			}
			for i := range ie.SlotFormatCombToReleaseList {
				if err := ie.SlotFormatCombToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "availableRB-SetsToAddModList-r16", Optional: true},
					{Name: "availableRB-SetsToReleaseList-r16", Optional: true},
					{Name: "switchTriggerToAddModList-r16", Optional: true},
					{Name: "switchTriggerToReleaseList-r16", Optional: true},
					{Name: "co-DurationsPerCellToAddModList-r16", Optional: true},
					{Name: "co-DurationsPerCellToReleaseList-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AvailableRB_SetsToAddModList_r16 != nil, ie.AvailableRB_SetsToReleaseList_r16 != nil, ie.SwitchTriggerToAddModList_r16 != nil, ie.SwitchTriggerToReleaseList_r16 != nil, ie.Co_DurationsPerCellToAddModList_r16 != nil, ie.Co_DurationsPerCellToReleaseList_r16 != nil}); err != nil {
				return err
			}

			if ie.AvailableRB_SetsToAddModList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(slotFormatIndicatorExtAvailableRBSetsToAddModListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.AvailableRB_SetsToAddModList_r16))); err != nil {
					return err
				}
				for i := range ie.AvailableRB_SetsToAddModList_r16 {
					if err := ie.AvailableRB_SetsToAddModList_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.AvailableRB_SetsToReleaseList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(slotFormatIndicatorExtAvailableRBSetsToReleaseListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.AvailableRB_SetsToReleaseList_r16))); err != nil {
					return err
				}
				for i := range ie.AvailableRB_SetsToReleaseList_r16 {
					if err := ie.AvailableRB_SetsToReleaseList_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SwitchTriggerToAddModList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(slotFormatIndicatorExtSwitchTriggerToAddModListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SwitchTriggerToAddModList_r16))); err != nil {
					return err
				}
				for i := range ie.SwitchTriggerToAddModList_r16 {
					if err := ie.SwitchTriggerToAddModList_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SwitchTriggerToReleaseList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(slotFormatIndicatorExtSwitchTriggerToReleaseListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SwitchTriggerToReleaseList_r16))); err != nil {
					return err
				}
				for i := range ie.SwitchTriggerToReleaseList_r16 {
					if err := ie.SwitchTriggerToReleaseList_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Co_DurationsPerCellToAddModList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(slotFormatIndicatorExtCoDurationsPerCellToAddModListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Co_DurationsPerCellToAddModList_r16))); err != nil {
					return err
				}
				for i := range ie.Co_DurationsPerCellToAddModList_r16 {
					if err := ie.Co_DurationsPerCellToAddModList_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Co_DurationsPerCellToReleaseList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(slotFormatIndicatorExtCoDurationsPerCellToReleaseListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Co_DurationsPerCellToReleaseList_r16))); err != nil {
					return err
				}
				for i := range ie.Co_DurationsPerCellToReleaseList_r16 {
					if err := ie.Co_DurationsPerCellToReleaseList_r16[i].Encode(ex); err != nil {
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
					{Name: "switchTriggerToAddModListSizeExt-r16", Optional: true},
					{Name: "switchTriggerToReleaseListSizeExt-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SwitchTriggerToAddModListSizeExt_r16 != nil, ie.SwitchTriggerToReleaseListSizeExt_r16 != nil}); err != nil {
				return err
			}

			if ie.SwitchTriggerToAddModListSizeExt_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(slotFormatIndicatorExtSwitchTriggerToAddModListSizeExtR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SwitchTriggerToAddModListSizeExt_r16))); err != nil {
					return err
				}
				for i := range ie.SwitchTriggerToAddModListSizeExt_r16 {
					if err := ie.SwitchTriggerToAddModListSizeExt_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SwitchTriggerToReleaseListSizeExt_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(slotFormatIndicatorExtSwitchTriggerToReleaseListSizeExtR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SwitchTriggerToReleaseListSizeExt_r16))); err != nil {
					return err
				}
				for i := range ie.SwitchTriggerToReleaseListSizeExt_r16 {
					if err := ie.SwitchTriggerToReleaseListSizeExt_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "co-DurationsPerCellToAddModList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Co_DurationsPerCellToAddModList_r17 != nil}); err != nil {
				return err
			}

			if ie.Co_DurationsPerCellToAddModList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(slotFormatIndicatorExtCoDurationsPerCellToAddModListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Co_DurationsPerCellToAddModList_r17))); err != nil {
					return err
				}
				for i := range ie.Co_DurationsPerCellToAddModList_r17 {
					if err := ie.Co_DurationsPerCellToAddModList_r17[i].Encode(ex); err != nil {
						return err
					}
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

func (ie *SlotFormatIndicator) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(slotFormatIndicatorConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sfi-RNTI: ref
	{
		if err := ie.Sfi_RNTI.Decode(d); err != nil {
			return err
		}
	}

	// 4. dci-PayloadSize: integer
	{
		v1, err := d.DecodeInteger(slotFormatIndicatorDciPayloadSizeConstraints)
		if err != nil {
			return err
		}
		ie.Dci_PayloadSize = v1
	}

	// 5. slotFormatCombToAddModList: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(slotFormatIndicatorSlotFormatCombToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SlotFormatCombToAddModList = make([]SlotFormatCombinationsPerCell, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SlotFormatCombToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. slotFormatCombToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(slotFormatIndicatorSlotFormatCombToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SlotFormatCombToReleaseList = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SlotFormatCombToReleaseList[i].Decode(d); err != nil {
					return err
				}
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
				{Name: "availableRB-SetsToAddModList-r16", Optional: true},
				{Name: "availableRB-SetsToReleaseList-r16", Optional: true},
				{Name: "switchTriggerToAddModList-r16", Optional: true},
				{Name: "switchTriggerToReleaseList-r16", Optional: true},
				{Name: "co-DurationsPerCellToAddModList-r16", Optional: true},
				{Name: "co-DurationsPerCellToReleaseList-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(slotFormatIndicatorExtAvailableRBSetsToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AvailableRB_SetsToAddModList_r16 = make([]AvailableRB_SetsPerCell_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AvailableRB_SetsToAddModList_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(slotFormatIndicatorExtAvailableRBSetsToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AvailableRB_SetsToReleaseList_r16 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AvailableRB_SetsToReleaseList_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(slotFormatIndicatorExtSwitchTriggerToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SwitchTriggerToAddModList_r16 = make([]SearchSpaceSwitchTrigger_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SwitchTriggerToAddModList_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(slotFormatIndicatorExtSwitchTriggerToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SwitchTriggerToReleaseList_r16 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SwitchTriggerToReleaseList_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			seqOf := dx.NewSequenceOfDecoder(slotFormatIndicatorExtCoDurationsPerCellToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Co_DurationsPerCellToAddModList_r16 = make([]CO_DurationsPerCell_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Co_DurationsPerCellToAddModList_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			seqOf := dx.NewSequenceOfDecoder(slotFormatIndicatorExtCoDurationsPerCellToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Co_DurationsPerCellToReleaseList_r16 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Co_DurationsPerCellToReleaseList_r16[i].Decode(dx); err != nil {
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
				{Name: "switchTriggerToAddModListSizeExt-r16", Optional: true},
				{Name: "switchTriggerToReleaseListSizeExt-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(slotFormatIndicatorExtSwitchTriggerToAddModListSizeExtR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SwitchTriggerToAddModListSizeExt_r16 = make([]SearchSpaceSwitchTrigger_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SwitchTriggerToAddModListSizeExt_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(slotFormatIndicatorExtSwitchTriggerToReleaseListSizeExtR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SwitchTriggerToReleaseListSizeExt_r16 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SwitchTriggerToReleaseListSizeExt_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "co-DurationsPerCellToAddModList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(slotFormatIndicatorExtCoDurationsPerCellToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Co_DurationsPerCellToAddModList_r17 = make([]CO_DurationsPerCell_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Co_DurationsPerCellToAddModList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
