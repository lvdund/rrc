// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUSCH-PowerControl-v1610 (line 12628).

var pUSCHPowerControlV1610Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pathlossReferenceRSToAddModListSizeExt-v1610", Optional: true},
		{Name: "pathlossReferenceRSToReleaseListSizeExt-v1610", Optional: true},
		{Name: "p0-PUSCH-SetList-r16", Optional: true},
		{Name: "olpc-ParameterSet", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var pUSCHPowerControlV1610PathlossReferenceRSToAddModListSizeExtV1610Constraints = per.SizeRange(1, common.MaxNrofPUSCH_PathlossReferenceRSsDiff_r16)

var pUSCHPowerControlV1610PathlossReferenceRSToReleaseListSizeExtV1610Constraints = per.SizeRange(1, common.MaxNrofPUSCH_PathlossReferenceRSsDiff_r16)

var pUSCHPowerControlV1610P0PUSCHSetListR16Constraints = per.SizeRange(1, common.MaxNrofSRI_PUSCH_Mappings)

var pUSCHPowerControlV1610ExtSriPUSCHMappingToAddModList2R17Constraints = per.SizeRange(1, common.MaxNrofSRI_PUSCH_Mappings)

var pUSCHPowerControlV1610ExtSriPUSCHMappingToReleaseList2R17Constraints = per.SizeRange(1, common.MaxNrofSRI_PUSCH_Mappings)

var pUSCHPowerControlV1610ExtP0PUSCHSetList2R17Constraints = per.SizeRange(1, common.MaxNrofSRI_PUSCH_Mappings)

var pUSCHPowerControlV1610ExtDummyConstraints = per.SizeRange(1, common.MaxNrofPUSCH_PathlossReferenceRSs_r16)

var pUSCHPowerControlV1610OlpcParameterSetConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "olpc-ParameterSetDCI-0-1-r16", Optional: true},
		{Name: "olpc-ParameterSetDCI-0-2-r16", Optional: true},
	},
}

type PUSCH_PowerControl_v1610 struct {
	PathlossReferenceRSToAddModListSizeExt_v1610  []PUSCH_PathlossReferenceRS_r16
	PathlossReferenceRSToReleaseListSizeExt_v1610 []PUSCH_PathlossReferenceRS_Id_v1610
	P0_PUSCH_SetList_r16                          []P0_PUSCH_Set_r16
	Olpc_ParameterSet                             *struct {
		Olpc_ParameterSetDCI_0_1_r16 *int64
		Olpc_ParameterSetDCI_0_2_r16 *int64
	}
	Sri_PUSCH_MappingToAddModList2_r17  []SRI_PUSCH_PowerControl
	Sri_PUSCH_MappingToReleaseList2_r17 []SRI_PUSCH_PowerControlId
	P0_PUSCH_SetList2_r17               []P0_PUSCH_Set_r16
	Dummy                               []DummyPathlossReferenceRS_v1710
	Msg3_Alpha_SBFD_r19                 *Alpha
}

func (ie *PUSCH_PowerControl_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHPowerControlV1610Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sri_PUSCH_MappingToAddModList2_r17 != nil || ie.Sri_PUSCH_MappingToReleaseList2_r17 != nil || ie.P0_PUSCH_SetList2_r17 != nil || ie.Dummy != nil
	hasExtGroup1 := ie.Msg3_Alpha_SBFD_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PathlossReferenceRSToAddModListSizeExt_v1610 != nil, ie.PathlossReferenceRSToReleaseListSizeExt_v1610 != nil, ie.P0_PUSCH_SetList_r16 != nil, ie.Olpc_ParameterSet != nil}); err != nil {
		return err
	}

	// 3. pathlossReferenceRSToAddModListSizeExt-v1610: sequence-of
	{
		if ie.PathlossReferenceRSToAddModListSizeExt_v1610 != nil {
			seqOf := e.NewSequenceOfEncoder(pUSCHPowerControlV1610PathlossReferenceRSToAddModListSizeExtV1610Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.PathlossReferenceRSToAddModListSizeExt_v1610))); err != nil {
				return err
			}
			for i := range ie.PathlossReferenceRSToAddModListSizeExt_v1610 {
				if err := ie.PathlossReferenceRSToAddModListSizeExt_v1610[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. pathlossReferenceRSToReleaseListSizeExt-v1610: sequence-of
	{
		if ie.PathlossReferenceRSToReleaseListSizeExt_v1610 != nil {
			seqOf := e.NewSequenceOfEncoder(pUSCHPowerControlV1610PathlossReferenceRSToReleaseListSizeExtV1610Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.PathlossReferenceRSToReleaseListSizeExt_v1610))); err != nil {
				return err
			}
			for i := range ie.PathlossReferenceRSToReleaseListSizeExt_v1610 {
				if err := ie.PathlossReferenceRSToReleaseListSizeExt_v1610[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. p0-PUSCH-SetList-r16: sequence-of
	{
		if ie.P0_PUSCH_SetList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(pUSCHPowerControlV1610P0PUSCHSetListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.P0_PUSCH_SetList_r16))); err != nil {
				return err
			}
			for i := range ie.P0_PUSCH_SetList_r16 {
				if err := ie.P0_PUSCH_SetList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. olpc-ParameterSet: sequence
	{
		if ie.Olpc_ParameterSet != nil {
			c := ie.Olpc_ParameterSet
			pUSCHPowerControlV1610OlpcParameterSetSeq := e.NewSequenceEncoder(pUSCHPowerControlV1610OlpcParameterSetConstraints)
			if err := pUSCHPowerControlV1610OlpcParameterSetSeq.EncodePreamble([]bool{c.Olpc_ParameterSetDCI_0_1_r16 != nil, c.Olpc_ParameterSetDCI_0_2_r16 != nil}); err != nil {
				return err
			}
			if c.Olpc_ParameterSetDCI_0_1_r16 != nil {
				if err := e.EncodeInteger((*c.Olpc_ParameterSetDCI_0_1_r16), per.Constrained(1, 2)); err != nil {
					return err
				}
			}
			if c.Olpc_ParameterSetDCI_0_2_r16 != nil {
				if err := e.EncodeInteger((*c.Olpc_ParameterSetDCI_0_2_r16), per.Constrained(1, 2)); err != nil {
					return err
				}
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
					{Name: "sri-PUSCH-MappingToAddModList2-r17", Optional: true},
					{Name: "sri-PUSCH-MappingToReleaseList2-r17", Optional: true},
					{Name: "p0-PUSCH-SetList2-r17", Optional: true},
					{Name: "dummy", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sri_PUSCH_MappingToAddModList2_r17 != nil, ie.Sri_PUSCH_MappingToReleaseList2_r17 != nil, ie.P0_PUSCH_SetList2_r17 != nil, ie.Dummy != nil}); err != nil {
				return err
			}

			if ie.Sri_PUSCH_MappingToAddModList2_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUSCHPowerControlV1610ExtSriPUSCHMappingToAddModList2R17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sri_PUSCH_MappingToAddModList2_r17))); err != nil {
					return err
				}
				for i := range ie.Sri_PUSCH_MappingToAddModList2_r17 {
					if err := ie.Sri_PUSCH_MappingToAddModList2_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sri_PUSCH_MappingToReleaseList2_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUSCHPowerControlV1610ExtSriPUSCHMappingToReleaseList2R17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sri_PUSCH_MappingToReleaseList2_r17))); err != nil {
					return err
				}
				for i := range ie.Sri_PUSCH_MappingToReleaseList2_r17 {
					if err := ie.Sri_PUSCH_MappingToReleaseList2_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.P0_PUSCH_SetList2_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUSCHPowerControlV1610ExtP0PUSCHSetList2R17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.P0_PUSCH_SetList2_r17))); err != nil {
					return err
				}
				for i := range ie.P0_PUSCH_SetList2_r17 {
					if err := ie.P0_PUSCH_SetList2_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Dummy != nil {
				seqOf := ex.NewSequenceOfEncoder(pUSCHPowerControlV1610ExtDummyConstraints)
				if err := seqOf.EncodeLength(int64(len(ie.Dummy))); err != nil {
					return err
				}
				for i := range ie.Dummy {
					if err := ie.Dummy[i].Encode(ex); err != nil {
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
					{Name: "msg3-Alpha-SBFD-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Msg3_Alpha_SBFD_r19 != nil}); err != nil {
				return err
			}

			if ie.Msg3_Alpha_SBFD_r19 != nil {
				if err := ie.Msg3_Alpha_SBFD_r19.Encode(ex); err != nil {
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

func (ie *PUSCH_PowerControl_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHPowerControlV1610Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pathlossReferenceRSToAddModListSizeExt-v1610: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(pUSCHPowerControlV1610PathlossReferenceRSToAddModListSizeExtV1610Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PathlossReferenceRSToAddModListSizeExt_v1610 = make([]PUSCH_PathlossReferenceRS_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PathlossReferenceRSToAddModListSizeExt_v1610[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. pathlossReferenceRSToReleaseListSizeExt-v1610: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(pUSCHPowerControlV1610PathlossReferenceRSToReleaseListSizeExtV1610Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PathlossReferenceRSToReleaseListSizeExt_v1610 = make([]PUSCH_PathlossReferenceRS_Id_v1610, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PathlossReferenceRSToReleaseListSizeExt_v1610[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. p0-PUSCH-SetList-r16: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(pUSCHPowerControlV1610P0PUSCHSetListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.P0_PUSCH_SetList_r16 = make([]P0_PUSCH_Set_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.P0_PUSCH_SetList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. olpc-ParameterSet: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.Olpc_ParameterSet = &struct {
				Olpc_ParameterSetDCI_0_1_r16 *int64
				Olpc_ParameterSetDCI_0_2_r16 *int64
			}{}
			c := ie.Olpc_ParameterSet
			pUSCHPowerControlV1610OlpcParameterSetSeq := d.NewSequenceDecoder(pUSCHPowerControlV1610OlpcParameterSetConstraints)
			if err := pUSCHPowerControlV1610OlpcParameterSetSeq.DecodePreamble(); err != nil {
				return err
			}
			if pUSCHPowerControlV1610OlpcParameterSetSeq.IsComponentPresent(0) {
				c.Olpc_ParameterSetDCI_0_1_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				(*c.Olpc_ParameterSetDCI_0_1_r16) = v
			}
			if pUSCHPowerControlV1610OlpcParameterSetSeq.IsComponentPresent(1) {
				c.Olpc_ParameterSetDCI_0_2_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				(*c.Olpc_ParameterSetDCI_0_2_r16) = v
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
				{Name: "sri-PUSCH-MappingToAddModList2-r17", Optional: true},
				{Name: "sri-PUSCH-MappingToReleaseList2-r17", Optional: true},
				{Name: "p0-PUSCH-SetList2-r17", Optional: true},
				{Name: "dummy", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(pUSCHPowerControlV1610ExtSriPUSCHMappingToAddModList2R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sri_PUSCH_MappingToAddModList2_r17 = make([]SRI_PUSCH_PowerControl, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sri_PUSCH_MappingToAddModList2_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(pUSCHPowerControlV1610ExtSriPUSCHMappingToReleaseList2R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sri_PUSCH_MappingToReleaseList2_r17 = make([]SRI_PUSCH_PowerControlId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sri_PUSCH_MappingToReleaseList2_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(pUSCHPowerControlV1610ExtP0PUSCHSetList2R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.P0_PUSCH_SetList2_r17 = make([]P0_PUSCH_Set_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.P0_PUSCH_SetList2_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(pUSCHPowerControlV1610ExtDummyConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Dummy = make([]DummyPathlossReferenceRS_v1710, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Dummy[i].Decode(dx); err != nil {
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
				{Name: "msg3-Alpha-SBFD-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Msg3_Alpha_SBFD_r19 = new(Alpha)
			if err := ie.Msg3_Alpha_SBFD_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
