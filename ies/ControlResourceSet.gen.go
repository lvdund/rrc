// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ControlResourceSet (line 6761).

var controlResourceSetConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "controlResourceSetId"},
		{Name: "frequencyDomainResources"},
		{Name: "duration"},
		{Name: "cce-REG-MappingType"},
		{Name: "precoderGranularity"},
		{Name: "tci-StatesPDCCH-ToAddList", Optional: true},
		{Name: "tci-StatesPDCCH-ToReleaseList", Optional: true},
		{Name: "tci-PresentInDCI", Optional: true},
		{Name: "pdcch-DMRS-ScramblingID", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var controlResourceSetFrequencyDomainResourcesConstraints = per.FixedSize(45)

var controlResourceSetDurationConstraints = per.Constrained(1, common.MaxCoReSetDuration)

var controlResourceSetCceREGMappingTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "interleaved"},
		{Name: "nonInterleaved"},
	},
}

const (
	ControlResourceSet_Cce_REG_MappingType_Interleaved    = 0
	ControlResourceSet_Cce_REG_MappingType_NonInterleaved = 1
)

const (
	ControlResourceSet_PrecoderGranularity_SameAsREG_Bundle = 0
	ControlResourceSet_PrecoderGranularity_AllContiguousRBs = 1
)

var controlResourceSetPrecoderGranularityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ControlResourceSet_PrecoderGranularity_SameAsREG_Bundle, ControlResourceSet_PrecoderGranularity_AllContiguousRBs},
}

var controlResourceSetTciStatesPDCCHToAddListConstraints = per.SizeRange(1, common.MaxNrofTCI_StatesPDCCH)

var controlResourceSetTciStatesPDCCHToReleaseListConstraints = per.SizeRange(1, common.MaxNrofTCI_StatesPDCCH)

const (
	ControlResourceSet_Tci_PresentInDCI_Enabled = 0
)

var controlResourceSetTciPresentInDCIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ControlResourceSet_Tci_PresentInDCI_Enabled},
}

var controlResourceSetPdcchDMRSScramblingIDConstraints = per.Constrained(0, 65535)

var controlResourceSetRbOffsetR16Constraints = per.Constrained(0, 5)

var controlResourceSetTciPresentDCI12R16Constraints = per.Constrained(1, 3)

var controlResourceSetCoresetPoolIndexR16Constraints = per.Constrained(0, 1)

const (
	ControlResourceSet_Ext_FollowUnifiedTCI_State_r17_Enabled = 0
)

var controlResourceSetExtFollowUnifiedTCIStateR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ControlResourceSet_Ext_FollowUnifiedTCI_State_r17_Enabled},
}

const (
	ControlResourceSet_Ext_ApplyIndicatedTCI_State_r18_First  = 0
	ControlResourceSet_Ext_ApplyIndicatedTCI_State_r18_Second = 1
	ControlResourceSet_Ext_ApplyIndicatedTCI_State_r18_Both   = 2
	ControlResourceSet_Ext_ApplyIndicatedTCI_State_r18_None   = 3
)

var controlResourceSetExtApplyIndicatedTCIStateR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ControlResourceSet_Ext_ApplyIndicatedTCI_State_r18_First, ControlResourceSet_Ext_ApplyIndicatedTCI_State_r18_Second, ControlResourceSet_Ext_ApplyIndicatedTCI_State_r18_Both, ControlResourceSet_Ext_ApplyIndicatedTCI_State_r18_None},
}

var controlResourceSetCceREGMappingTypeInterleavedConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reg-BundleSize"},
		{Name: "interleaverSize"},
		{Name: "shiftIndex", Optional: true},
	},
}

const (
	ControlResourceSet_Cce_REG_MappingType_Interleaved_Reg_BundleSize_N2 = 0
	ControlResourceSet_Cce_REG_MappingType_Interleaved_Reg_BundleSize_N3 = 1
	ControlResourceSet_Cce_REG_MappingType_Interleaved_Reg_BundleSize_N6 = 2
)

var controlResourceSetCceREGMappingTypeInterleavedRegBundleSizeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ControlResourceSet_Cce_REG_MappingType_Interleaved_Reg_BundleSize_N2, ControlResourceSet_Cce_REG_MappingType_Interleaved_Reg_BundleSize_N3, ControlResourceSet_Cce_REG_MappingType_Interleaved_Reg_BundleSize_N6},
}

const (
	ControlResourceSet_Cce_REG_MappingType_Interleaved_InterleaverSize_N2 = 0
	ControlResourceSet_Cce_REG_MappingType_Interleaved_InterleaverSize_N3 = 1
	ControlResourceSet_Cce_REG_MappingType_Interleaved_InterleaverSize_N6 = 2
)

var controlResourceSetCceREGMappingTypeInterleavedInterleaverSizeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ControlResourceSet_Cce_REG_MappingType_Interleaved_InterleaverSize_N2, ControlResourceSet_Cce_REG_MappingType_Interleaved_InterleaverSize_N3, ControlResourceSet_Cce_REG_MappingType_Interleaved_InterleaverSize_N6},
}

type ControlResourceSet struct {
	ControlResourceSetId     ControlResourceSetId
	FrequencyDomainResources per.BitString
	Duration                 int64
	Cce_REG_MappingType      struct {
		Choice      int
		Interleaved *struct {
			Reg_BundleSize  int64
			InterleaverSize int64
			ShiftIndex      *int64
		}
		NonInterleaved *struct{}
	}
	PrecoderGranularity           int64
	Tci_StatesPDCCH_ToAddList     []TCI_StateId
	Tci_StatesPDCCH_ToReleaseList []TCI_StateId
	Tci_PresentInDCI              *int64
	Pdcch_DMRS_ScramblingID       *int64
	Rb_Offset_r16                 *int64
	Tci_PresentDCI_1_2_r16        *int64
	CoresetPoolIndex_r16          *int64
	ControlResourceSetId_v1610    *ControlResourceSetId_v1610
	FollowUnifiedTCI_State_r17    *int64
	ApplyIndicatedTCI_State_r18   *int64
}

func (ie *ControlResourceSet) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(controlResourceSetConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Rb_Offset_r16 != nil || ie.Tci_PresentDCI_1_2_r16 != nil || ie.CoresetPoolIndex_r16 != nil || ie.ControlResourceSetId_v1610 != nil
	hasExtGroup1 := ie.FollowUnifiedTCI_State_r17 != nil
	hasExtGroup2 := ie.ApplyIndicatedTCI_State_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Tci_StatesPDCCH_ToAddList != nil, ie.Tci_StatesPDCCH_ToReleaseList != nil, ie.Tci_PresentInDCI != nil, ie.Pdcch_DMRS_ScramblingID != nil}); err != nil {
		return err
	}

	// 3. controlResourceSetId: ref
	{
		if err := ie.ControlResourceSetId.Encode(e); err != nil {
			return err
		}
	}

	// 4. frequencyDomainResources: bit-string
	{
		if err := e.EncodeBitString(ie.FrequencyDomainResources, controlResourceSetFrequencyDomainResourcesConstraints); err != nil {
			return err
		}
	}

	// 5. duration: integer
	{
		if err := e.EncodeInteger(ie.Duration, controlResourceSetDurationConstraints); err != nil {
			return err
		}
	}

	// 6. cce-REG-MappingType: choice
	{
		choiceEnc := e.NewChoiceEncoder(controlResourceSetCceREGMappingTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Cce_REG_MappingType.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Cce_REG_MappingType.Choice {
		case ControlResourceSet_Cce_REG_MappingType_Interleaved:
			c := (*ie.Cce_REG_MappingType.Interleaved)
			controlResourceSetCceREGMappingTypeInterleavedSeq := e.NewSequenceEncoder(controlResourceSetCceREGMappingTypeInterleavedConstraints)
			if err := controlResourceSetCceREGMappingTypeInterleavedSeq.EncodePreamble([]bool{c.ShiftIndex != nil}); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Reg_BundleSize, controlResourceSetCceREGMappingTypeInterleavedRegBundleSizeConstraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.InterleaverSize, controlResourceSetCceREGMappingTypeInterleavedInterleaverSizeConstraints); err != nil {
				return err
			}
			if c.ShiftIndex != nil {
				if err := e.EncodeInteger((*c.ShiftIndex), per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)); err != nil {
					return err
				}
			}
		case ControlResourceSet_Cce_REG_MappingType_NonInterleaved:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Cce_REG_MappingType.Choice), Constraint: "index out of range"}
		}
	}

	// 7. precoderGranularity: enumerated
	{
		if err := e.EncodeEnumerated(ie.PrecoderGranularity, controlResourceSetPrecoderGranularityConstraints); err != nil {
			return err
		}
	}

	// 8. tci-StatesPDCCH-ToAddList: sequence-of
	{
		if ie.Tci_StatesPDCCH_ToAddList != nil {
			seqOf := e.NewSequenceOfEncoder(controlResourceSetTciStatesPDCCHToAddListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Tci_StatesPDCCH_ToAddList))); err != nil {
				return err
			}
			for i := range ie.Tci_StatesPDCCH_ToAddList {
				if err := ie.Tci_StatesPDCCH_ToAddList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 9. tci-StatesPDCCH-ToReleaseList: sequence-of
	{
		if ie.Tci_StatesPDCCH_ToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(controlResourceSetTciStatesPDCCHToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Tci_StatesPDCCH_ToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Tci_StatesPDCCH_ToReleaseList {
				if err := ie.Tci_StatesPDCCH_ToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 10. tci-PresentInDCI: enumerated
	{
		if ie.Tci_PresentInDCI != nil {
			if err := e.EncodeEnumerated(*ie.Tci_PresentInDCI, controlResourceSetTciPresentInDCIConstraints); err != nil {
				return err
			}
		}
	}

	// 11. pdcch-DMRS-ScramblingID: integer
	{
		if ie.Pdcch_DMRS_ScramblingID != nil {
			if err := e.EncodeInteger(*ie.Pdcch_DMRS_ScramblingID, controlResourceSetPdcchDMRSScramblingIDConstraints); err != nil {
				return err
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
					{Name: "rb-Offset-r16", Optional: true},
					{Name: "tci-PresentDCI-1-2-r16", Optional: true},
					{Name: "coresetPoolIndex-r16", Optional: true},
					{Name: "controlResourceSetId-v1610", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Rb_Offset_r16 != nil, ie.Tci_PresentDCI_1_2_r16 != nil, ie.CoresetPoolIndex_r16 != nil, ie.ControlResourceSetId_v1610 != nil}); err != nil {
				return err
			}

			if ie.Rb_Offset_r16 != nil {
				if err := ex.EncodeInteger(*ie.Rb_Offset_r16, controlResourceSetRbOffsetR16Constraints); err != nil {
					return err
				}
			}

			if ie.Tci_PresentDCI_1_2_r16 != nil {
				if err := ex.EncodeInteger(*ie.Tci_PresentDCI_1_2_r16, controlResourceSetTciPresentDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.CoresetPoolIndex_r16 != nil {
				if err := ex.EncodeInteger(*ie.CoresetPoolIndex_r16, controlResourceSetCoresetPoolIndexR16Constraints); err != nil {
					return err
				}
			}

			if ie.ControlResourceSetId_v1610 != nil {
				if err := ie.ControlResourceSetId_v1610.Encode(ex); err != nil {
					return err
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
					{Name: "followUnifiedTCI-State-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FollowUnifiedTCI_State_r17 != nil}); err != nil {
				return err
			}

			if ie.FollowUnifiedTCI_State_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.FollowUnifiedTCI_State_r17, controlResourceSetExtFollowUnifiedTCIStateR17Constraints); err != nil {
					return err
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
					{Name: "applyIndicatedTCI-State-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ApplyIndicatedTCI_State_r18 != nil}); err != nil {
				return err
			}

			if ie.ApplyIndicatedTCI_State_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ApplyIndicatedTCI_State_r18, controlResourceSetExtApplyIndicatedTCIStateR18Constraints); err != nil {
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

func (ie *ControlResourceSet) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(controlResourceSetConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. controlResourceSetId: ref
	{
		if err := ie.ControlResourceSetId.Decode(d); err != nil {
			return err
		}
	}

	// 4. frequencyDomainResources: bit-string
	{
		v1, err := d.DecodeBitString(controlResourceSetFrequencyDomainResourcesConstraints)
		if err != nil {
			return err
		}
		ie.FrequencyDomainResources = v1
	}

	// 5. duration: integer
	{
		v2, err := d.DecodeInteger(controlResourceSetDurationConstraints)
		if err != nil {
			return err
		}
		ie.Duration = v2
	}

	// 6. cce-REG-MappingType: choice
	{
		choiceDec := d.NewChoiceDecoder(controlResourceSetCceREGMappingTypeConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Cce_REG_MappingType.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case ControlResourceSet_Cce_REG_MappingType_Interleaved:
			ie.Cce_REG_MappingType.Interleaved = &struct {
				Reg_BundleSize  int64
				InterleaverSize int64
				ShiftIndex      *int64
			}{}
			c := (*ie.Cce_REG_MappingType.Interleaved)
			controlResourceSetCceREGMappingTypeInterleavedSeq := d.NewSequenceDecoder(controlResourceSetCceREGMappingTypeInterleavedConstraints)
			if err := controlResourceSetCceREGMappingTypeInterleavedSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeEnumerated(controlResourceSetCceREGMappingTypeInterleavedRegBundleSizeConstraints)
				if err != nil {
					return err
				}
				c.Reg_BundleSize = v
			}
			{
				v, err := d.DecodeEnumerated(controlResourceSetCceREGMappingTypeInterleavedInterleaverSizeConstraints)
				if err != nil {
					return err
				}
				c.InterleaverSize = v
			}
			if controlResourceSetCceREGMappingTypeInterleavedSeq.IsComponentPresent(2) {
				c.ShiftIndex = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1))
				if err != nil {
					return err
				}
				(*c.ShiftIndex) = v
			}
		case ControlResourceSet_Cce_REG_MappingType_NonInterleaved:
			ie.Cce_REG_MappingType.NonInterleaved = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		}
	}

	// 7. precoderGranularity: enumerated
	{
		v4, err := d.DecodeEnumerated(controlResourceSetPrecoderGranularityConstraints)
		if err != nil {
			return err
		}
		ie.PrecoderGranularity = v4
	}

	// 8. tci-StatesPDCCH-ToAddList: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(controlResourceSetTciStatesPDCCHToAddListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Tci_StatesPDCCH_ToAddList = make([]TCI_StateId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Tci_StatesPDCCH_ToAddList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. tci-StatesPDCCH-ToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(controlResourceSetTciStatesPDCCHToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Tci_StatesPDCCH_ToReleaseList = make([]TCI_StateId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Tci_StatesPDCCH_ToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. tci-PresentInDCI: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(controlResourceSetTciPresentInDCIConstraints)
			if err != nil {
				return err
			}
			ie.Tci_PresentInDCI = &idx
		}
	}

	// 11. pdcch-DMRS-ScramblingID: integer
	{
		if seq.IsComponentPresent(8) {
			v, err := d.DecodeInteger(controlResourceSetPdcchDMRSScramblingIDConstraints)
			if err != nil {
				return err
			}
			ie.Pdcch_DMRS_ScramblingID = &v
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
				{Name: "rb-Offset-r16", Optional: true},
				{Name: "tci-PresentDCI-1-2-r16", Optional: true},
				{Name: "coresetPoolIndex-r16", Optional: true},
				{Name: "controlResourceSetId-v1610", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(controlResourceSetRbOffsetR16Constraints)
			if err != nil {
				return err
			}
			ie.Rb_Offset_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(controlResourceSetTciPresentDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.Tci_PresentDCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(controlResourceSetCoresetPoolIndexR16Constraints)
			if err != nil {
				return err
			}
			ie.CoresetPoolIndex_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			ie.ControlResourceSetId_v1610 = new(ControlResourceSetId_v1610)
			if err := ie.ControlResourceSetId_v1610.Decode(dx); err != nil {
				return err
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
				{Name: "followUnifiedTCI-State-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(controlResourceSetExtFollowUnifiedTCIStateR17Constraints)
			if err != nil {
				return err
			}
			ie.FollowUnifiedTCI_State_r17 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "applyIndicatedTCI-State-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(controlResourceSetExtApplyIndicatedTCIStateR18Constraints)
			if err != nil {
				return err
			}
			ie.ApplyIndicatedTCI_State_r18 = &v
		}
	}

	return nil
}
