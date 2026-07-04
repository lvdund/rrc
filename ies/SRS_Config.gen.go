// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-Config (line 15310).

var sRSConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-ResourceSetToReleaseList", Optional: true},
		{Name: "srs-ResourceSetToAddModList", Optional: true},
		{Name: "srs-ResourceToReleaseList", Optional: true},
		{Name: "srs-ResourceToAddModList", Optional: true},
		{Name: "tpc-Accumulation", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var sRSConfigSrsResourceSetToReleaseListConstraints = per.SizeRange(1, common.MaxNrofSRS_ResourceSets)

var sRSConfigSrsResourceSetToAddModListConstraints = per.SizeRange(1, common.MaxNrofSRS_ResourceSets)

var sRSConfigSrsResourceToReleaseListConstraints = per.SizeRange(1, common.MaxNrofSRS_Resources)

var sRSConfigSrsResourceToAddModListConstraints = per.SizeRange(1, common.MaxNrofSRS_Resources)

const (
	SRS_Config_Tpc_Accumulation_Disabled = 0
)

var sRSConfigTpcAccumulationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Config_Tpc_Accumulation_Disabled},
}

var sRSConfigSrsRequestDCI12R16Constraints = per.Constrained(1, 2)

var sRSConfigSrsRequestDCI02R16Constraints = per.Constrained(1, 2)

var sRSConfigExtSrsResourceSetToAddModListDCI02R16Constraints = per.SizeRange(1, common.MaxNrofSRS_ResourceSets)

var sRSConfigExtSrsResourceSetToReleaseListDCI02R16Constraints = per.SizeRange(1, common.MaxNrofSRS_ResourceSets)

var sRSConfigExtSrsPosResourceSetToReleaseListR16Constraints = per.SizeRange(1, common.MaxNrofSRS_PosResourceSets_r16)

var sRSConfigExtSrsPosResourceSetToAddModListR16Constraints = per.SizeRange(1, common.MaxNrofSRS_PosResourceSets_r16)

var sRSConfigExtSrsPosResourceToReleaseListR16Constraints = per.SizeRange(1, common.MaxNrofSRS_PosResources_r16)

var sRSConfigExtSrsPosResourceToAddModListR16Constraints = per.SizeRange(1, common.MaxNrofSRS_PosResources_r16)

const (
	SRS_Config_Ext_Dci_TriggeringPosResourceSetLink_r18_Enabled = 0
)

var sRSConfigExtDciTriggeringPosResourceSetLinkR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Config_Ext_Dci_TriggeringPosResourceSetLink_r18_Enabled},
}

const (
	SRS_Config_Ext_Srs_TwoSeparatePowerControlAdjustmentStates_r19_Enabled = 0
)

var sRSConfigExtSrsTwoSeparatePowerControlAdjustmentStatesR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Config_Ext_Srs_TwoSeparatePowerControlAdjustmentStates_r19_Enabled},
}

const (
	SRS_Config_Ext_Tpc_OfSRS_ClosedLoopIndexInDCI_1_1_r19_Enabled = 0
)

var sRSConfigExtTpcOfSRSClosedLoopIndexInDCI11R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Config_Ext_Tpc_OfSRS_ClosedLoopIndexInDCI_1_1_r19_Enabled},
}

const (
	SRS_Config_Ext_Srs_ClosedLoopIndexIndicatorInDCI_1_1_r19_Enabled = 0
)

var sRSConfigExtSrsClosedLoopIndexIndicatorInDCI11R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_Config_Ext_Srs_ClosedLoopIndexIndicatorInDCI_1_1_r19_Enabled},
}

type SRS_Config struct {
	Srs_ResourceSetToReleaseList                    []SRS_ResourceSetId
	Srs_ResourceSetToAddModList                     []SRS_ResourceSet
	Srs_ResourceToReleaseList                       []SRS_ResourceId
	Srs_ResourceToAddModList                        []SRS_Resource
	Tpc_Accumulation                                *int64
	Srs_RequestDCI_1_2_r16                          *int64
	Srs_RequestDCI_0_2_r16                          *int64
	Srs_ResourceSetToAddModListDCI_0_2_r16          []SRS_ResourceSet
	Srs_ResourceSetToReleaseListDCI_0_2_r16         []SRS_ResourceSetId
	Srs_PosResourceSetToReleaseList_r16             []SRS_PosResourceSetId_r16
	Srs_PosResourceSetToAddModList_r16              []SRS_PosResourceSet_r16
	Srs_PosResourceToReleaseList_r16                []SRS_PosResourceId_r16
	Srs_PosResourceToAddModList_r16                 []SRS_PosResource_r16
	Dci_TriggeringPosResourceSetLink_r18            *int64
	Srs_TwoSeparatePowerControlAdjustmentStates_r19 *int64
	Tpc_OfSRS_ClosedLoopIndexInDCI_1_1_r19          *int64
	Srs_ClosedLoopIndexIndicatorInDCI_1_1_r19       *int64
}

func (ie *SRS_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Srs_RequestDCI_1_2_r16 != nil || ie.Srs_RequestDCI_0_2_r16 != nil || ie.Srs_ResourceSetToAddModListDCI_0_2_r16 != nil || ie.Srs_ResourceSetToReleaseListDCI_0_2_r16 != nil || ie.Srs_PosResourceSetToReleaseList_r16 != nil || ie.Srs_PosResourceSetToAddModList_r16 != nil || ie.Srs_PosResourceToReleaseList_r16 != nil || ie.Srs_PosResourceToAddModList_r16 != nil
	hasExtGroup1 := ie.Dci_TriggeringPosResourceSetLink_r18 != nil
	hasExtGroup2 := ie.Srs_TwoSeparatePowerControlAdjustmentStates_r19 != nil || ie.Tpc_OfSRS_ClosedLoopIndexInDCI_1_1_r19 != nil || ie.Srs_ClosedLoopIndexIndicatorInDCI_1_1_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_ResourceSetToReleaseList != nil, ie.Srs_ResourceSetToAddModList != nil, ie.Srs_ResourceToReleaseList != nil, ie.Srs_ResourceToAddModList != nil, ie.Tpc_Accumulation != nil}); err != nil {
		return err
	}

	// 3. srs-ResourceSetToReleaseList: sequence-of
	{
		if ie.Srs_ResourceSetToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(sRSConfigSrsResourceSetToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_ResourceSetToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Srs_ResourceSetToReleaseList {
				if err := ie.Srs_ResourceSetToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. srs-ResourceSetToAddModList: sequence-of
	{
		if ie.Srs_ResourceSetToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(sRSConfigSrsResourceSetToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_ResourceSetToAddModList))); err != nil {
				return err
			}
			for i := range ie.Srs_ResourceSetToAddModList {
				if err := ie.Srs_ResourceSetToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. srs-ResourceToReleaseList: sequence-of
	{
		if ie.Srs_ResourceToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(sRSConfigSrsResourceToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_ResourceToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Srs_ResourceToReleaseList {
				if err := ie.Srs_ResourceToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. srs-ResourceToAddModList: sequence-of
	{
		if ie.Srs_ResourceToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(sRSConfigSrsResourceToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_ResourceToAddModList))); err != nil {
				return err
			}
			for i := range ie.Srs_ResourceToAddModList {
				if err := ie.Srs_ResourceToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. tpc-Accumulation: enumerated
	{
		if ie.Tpc_Accumulation != nil {
			if err := e.EncodeEnumerated(*ie.Tpc_Accumulation, sRSConfigTpcAccumulationConstraints); err != nil {
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
					{Name: "srs-RequestDCI-1-2-r16", Optional: true},
					{Name: "srs-RequestDCI-0-2-r16", Optional: true},
					{Name: "srs-ResourceSetToAddModListDCI-0-2-r16", Optional: true},
					{Name: "srs-ResourceSetToReleaseListDCI-0-2-r16", Optional: true},
					{Name: "srs-PosResourceSetToReleaseList-r16", Optional: true},
					{Name: "srs-PosResourceSetToAddModList-r16", Optional: true},
					{Name: "srs-PosResourceToReleaseList-r16", Optional: true},
					{Name: "srs-PosResourceToAddModList-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srs_RequestDCI_1_2_r16 != nil, ie.Srs_RequestDCI_0_2_r16 != nil, ie.Srs_ResourceSetToAddModListDCI_0_2_r16 != nil, ie.Srs_ResourceSetToReleaseListDCI_0_2_r16 != nil, ie.Srs_PosResourceSetToReleaseList_r16 != nil, ie.Srs_PosResourceSetToAddModList_r16 != nil, ie.Srs_PosResourceToReleaseList_r16 != nil, ie.Srs_PosResourceToAddModList_r16 != nil}); err != nil {
				return err
			}

			if ie.Srs_RequestDCI_1_2_r16 != nil {
				if err := ex.EncodeInteger(*ie.Srs_RequestDCI_1_2_r16, sRSConfigSrsRequestDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_RequestDCI_0_2_r16 != nil {
				if err := ex.EncodeInteger(*ie.Srs_RequestDCI_0_2_r16, sRSConfigSrsRequestDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_ResourceSetToAddModListDCI_0_2_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(sRSConfigExtSrsResourceSetToAddModListDCI02R16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Srs_ResourceSetToAddModListDCI_0_2_r16))); err != nil {
					return err
				}
				for i := range ie.Srs_ResourceSetToAddModListDCI_0_2_r16 {
					if err := ie.Srs_ResourceSetToAddModListDCI_0_2_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Srs_ResourceSetToReleaseListDCI_0_2_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(sRSConfigExtSrsResourceSetToReleaseListDCI02R16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Srs_ResourceSetToReleaseListDCI_0_2_r16))); err != nil {
					return err
				}
				for i := range ie.Srs_ResourceSetToReleaseListDCI_0_2_r16 {
					if err := ie.Srs_ResourceSetToReleaseListDCI_0_2_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Srs_PosResourceSetToReleaseList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(sRSConfigExtSrsPosResourceSetToReleaseListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Srs_PosResourceSetToReleaseList_r16))); err != nil {
					return err
				}
				for i := range ie.Srs_PosResourceSetToReleaseList_r16 {
					if err := ie.Srs_PosResourceSetToReleaseList_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Srs_PosResourceSetToAddModList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(sRSConfigExtSrsPosResourceSetToAddModListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Srs_PosResourceSetToAddModList_r16))); err != nil {
					return err
				}
				for i := range ie.Srs_PosResourceSetToAddModList_r16 {
					if err := ie.Srs_PosResourceSetToAddModList_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Srs_PosResourceToReleaseList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(sRSConfigExtSrsPosResourceToReleaseListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Srs_PosResourceToReleaseList_r16))); err != nil {
					return err
				}
				for i := range ie.Srs_PosResourceToReleaseList_r16 {
					if err := ie.Srs_PosResourceToReleaseList_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Srs_PosResourceToAddModList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(sRSConfigExtSrsPosResourceToAddModListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Srs_PosResourceToAddModList_r16))); err != nil {
					return err
				}
				for i := range ie.Srs_PosResourceToAddModList_r16 {
					if err := ie.Srs_PosResourceToAddModList_r16[i].Encode(ex); err != nil {
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
					{Name: "dci-TriggeringPosResourceSetLink-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dci_TriggeringPosResourceSetLink_r18 != nil}); err != nil {
				return err
			}

			if ie.Dci_TriggeringPosResourceSetLink_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Dci_TriggeringPosResourceSetLink_r18, sRSConfigExtDciTriggeringPosResourceSetLinkR18Constraints); err != nil {
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
					{Name: "srs-TwoSeparatePowerControlAdjustmentStates-r19", Optional: true},
					{Name: "tpc-OfSRS-ClosedLoopIndexInDCI-1-1-r19", Optional: true},
					{Name: "srs-ClosedLoopIndexIndicatorInDCI-1-1-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srs_TwoSeparatePowerControlAdjustmentStates_r19 != nil, ie.Tpc_OfSRS_ClosedLoopIndexInDCI_1_1_r19 != nil, ie.Srs_ClosedLoopIndexIndicatorInDCI_1_1_r19 != nil}); err != nil {
				return err
			}

			if ie.Srs_TwoSeparatePowerControlAdjustmentStates_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_TwoSeparatePowerControlAdjustmentStates_r19, sRSConfigExtSrsTwoSeparatePowerControlAdjustmentStatesR19Constraints); err != nil {
					return err
				}
			}

			if ie.Tpc_OfSRS_ClosedLoopIndexInDCI_1_1_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Tpc_OfSRS_ClosedLoopIndexInDCI_1_1_r19, sRSConfigExtTpcOfSRSClosedLoopIndexInDCI11R19Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_ClosedLoopIndexIndicatorInDCI_1_1_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_ClosedLoopIndexIndicatorInDCI_1_1_r19, sRSConfigExtSrsClosedLoopIndexIndicatorInDCI11R19Constraints); err != nil {
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

func (ie *SRS_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. srs-ResourceSetToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sRSConfigSrsResourceSetToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_ResourceSetToReleaseList = make([]SRS_ResourceSetId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_ResourceSetToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. srs-ResourceSetToAddModList: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sRSConfigSrsResourceSetToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_ResourceSetToAddModList = make([]SRS_ResourceSet, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_ResourceSetToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. srs-ResourceToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(sRSConfigSrsResourceToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_ResourceToReleaseList = make([]SRS_ResourceId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_ResourceToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. srs-ResourceToAddModList: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(sRSConfigSrsResourceToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_ResourceToAddModList = make([]SRS_Resource, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_ResourceToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. tpc-Accumulation: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(sRSConfigTpcAccumulationConstraints)
			if err != nil {
				return err
			}
			ie.Tpc_Accumulation = &idx
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
				{Name: "srs-RequestDCI-1-2-r16", Optional: true},
				{Name: "srs-RequestDCI-0-2-r16", Optional: true},
				{Name: "srs-ResourceSetToAddModListDCI-0-2-r16", Optional: true},
				{Name: "srs-ResourceSetToReleaseListDCI-0-2-r16", Optional: true},
				{Name: "srs-PosResourceSetToReleaseList-r16", Optional: true},
				{Name: "srs-PosResourceSetToAddModList-r16", Optional: true},
				{Name: "srs-PosResourceToReleaseList-r16", Optional: true},
				{Name: "srs-PosResourceToAddModList-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(sRSConfigSrsRequestDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.Srs_RequestDCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(sRSConfigSrsRequestDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.Srs_RequestDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(sRSConfigExtSrsResourceSetToAddModListDCI02R16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_ResourceSetToAddModListDCI_0_2_r16 = make([]SRS_ResourceSet, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_ResourceSetToAddModListDCI_0_2_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(sRSConfigExtSrsResourceSetToReleaseListDCI02R16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_ResourceSetToReleaseListDCI_0_2_r16 = make([]SRS_ResourceSetId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_ResourceSetToReleaseListDCI_0_2_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			seqOf := dx.NewSequenceOfDecoder(sRSConfigExtSrsPosResourceSetToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_PosResourceSetToReleaseList_r16 = make([]SRS_PosResourceSetId_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_PosResourceSetToReleaseList_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			seqOf := dx.NewSequenceOfDecoder(sRSConfigExtSrsPosResourceSetToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_PosResourceSetToAddModList_r16 = make([]SRS_PosResourceSet_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_PosResourceSetToAddModList_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(6) {
			seqOf := dx.NewSequenceOfDecoder(sRSConfigExtSrsPosResourceToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_PosResourceToReleaseList_r16 = make([]SRS_PosResourceId_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_PosResourceToReleaseList_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(7) {
			seqOf := dx.NewSequenceOfDecoder(sRSConfigExtSrsPosResourceToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_PosResourceToAddModList_r16 = make([]SRS_PosResource_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_PosResourceToAddModList_r16[i].Decode(dx); err != nil {
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
				{Name: "dci-TriggeringPosResourceSetLink-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sRSConfigExtDciTriggeringPosResourceSetLinkR18Constraints)
			if err != nil {
				return err
			}
			ie.Dci_TriggeringPosResourceSetLink_r18 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "srs-TwoSeparatePowerControlAdjustmentStates-r19", Optional: true},
				{Name: "tpc-OfSRS-ClosedLoopIndexInDCI-1-1-r19", Optional: true},
				{Name: "srs-ClosedLoopIndexIndicatorInDCI-1-1-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sRSConfigExtSrsTwoSeparatePowerControlAdjustmentStatesR19Constraints)
			if err != nil {
				return err
			}
			ie.Srs_TwoSeparatePowerControlAdjustmentStates_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(sRSConfigExtTpcOfSRSClosedLoopIndexInDCI11R19Constraints)
			if err != nil {
				return err
			}
			ie.Tpc_OfSRS_ClosedLoopIndexInDCI_1_1_r19 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(sRSConfigExtSrsClosedLoopIndexIndicatorInDCI11R19Constraints)
			if err != nil {
				return err
			}
			ie.Srs_ClosedLoopIndexIndicatorInDCI_1_1_r19 = &v
		}
	}

	return nil
}
