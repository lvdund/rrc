// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NZP-CSI-RS-ResourceSet (line 10847).

var nZPCSIRSResourceSetConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nzp-CSI-ResourceSetId"},
		{Name: "nzp-CSI-RS-Resources"},
		{Name: "repetition", Optional: true},
		{Name: "aperiodicTriggeringOffset", Optional: true},
		{Name: "trs-Info", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var nZPCSIRSResourceSetNzpCSIRSResourcesConstraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_ResourcesPerSet)

const (
	NZP_CSI_RS_ResourceSet_Repetition_On  = 0
	NZP_CSI_RS_ResourceSet_Repetition_Off = 1
)

var nZPCSIRSResourceSetRepetitionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NZP_CSI_RS_ResourceSet_Repetition_On, NZP_CSI_RS_ResourceSet_Repetition_Off},
}

var nZPCSIRSResourceSetAperiodicTriggeringOffsetConstraints = per.Constrained(0, 6)

const (
	NZP_CSI_RS_ResourceSet_Trs_Info_True = 0
)

var nZPCSIRSResourceSetTrsInfoConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NZP_CSI_RS_ResourceSet_Trs_Info_True},
}

var nZPCSIRSResourceSetAperiodicTriggeringOffsetR16Constraints = per.Constrained(0, 31)

const (
	NZP_CSI_RS_ResourceSet_Ext_Pdc_Info_r17_True = 0
)

var nZPCSIRSResourceSetExtPdcInfoR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NZP_CSI_RS_ResourceSet_Ext_Pdc_Info_r17_True},
}

var nZPCSIRSResourceSetAperiodicTriggeringOffsetR17Constraints = per.Constrained(0, 124)

var nZPCSIRSResourceSetAperiodicTriggeringOffsetL2R17Constraints = per.Constrained(0, 31)

const (
	NZP_CSI_RS_ResourceSet_Ext_ResourceType_r18_Periodic = 0
)

var nZPCSIRSResourceSetExtResourceTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NZP_CSI_RS_ResourceSet_Ext_ResourceType_r18_Periodic},
}

var nZPCSIRSResourceSetExtAdditionalOneSlotOffsetDopplerR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "twoResourcesPerGroup-r19"},
		{Name: "threeResourcesPerGroup-r19"},
		{Name: "fourResourcesPerGroup-r19"},
	},
}

const (
	NZP_CSI_RS_ResourceSet_Ext_AdditionalOneSlotOffsetDoppler_r19_TwoResourcesPerGroup_r19   = 0
	NZP_CSI_RS_ResourceSet_Ext_AdditionalOneSlotOffsetDoppler_r19_ThreeResourcesPerGroup_r19 = 1
	NZP_CSI_RS_ResourceSet_Ext_AdditionalOneSlotOffsetDoppler_r19_FourResourcesPerGroup_r19  = 2
)

const (
	NZP_CSI_RS_ResourceSet_Ext_Kdopp_r19_NumberOfResourceGroups_r19_N4  = 0
	NZP_CSI_RS_ResourceSet_Ext_Kdopp_r19_NumberOfResourceGroups_r19_N8  = 1
	NZP_CSI_RS_ResourceSet_Ext_Kdopp_r19_NumberOfResourceGroups_r19_N12 = 2
)

var nZPCSIRSResourceSetExtKdoppR19NumberOfResourceGroupsR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NZP_CSI_RS_ResourceSet_Ext_Kdopp_r19_NumberOfResourceGroups_r19_N4, NZP_CSI_RS_ResourceSet_Ext_Kdopp_r19_NumberOfResourceGroups_r19_N8, NZP_CSI_RS_ResourceSet_Ext_Kdopp_r19_NumberOfResourceGroups_r19_N12},
}

const (
	NZP_CSI_RS_ResourceSet_Ext_Kdopp_r19_NumberOfResourcesPerGroup_r19_N2 = 0
	NZP_CSI_RS_ResourceSet_Ext_Kdopp_r19_NumberOfResourcesPerGroup_r19_N3 = 1
	NZP_CSI_RS_ResourceSet_Ext_Kdopp_r19_NumberOfResourcesPerGroup_r19_N4 = 2
)

var nZPCSIRSResourceSetExtKdoppR19NumberOfResourcesPerGroupR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NZP_CSI_RS_ResourceSet_Ext_Kdopp_r19_NumberOfResourcesPerGroup_r19_N2, NZP_CSI_RS_ResourceSet_Ext_Kdopp_r19_NumberOfResourcesPerGroup_r19_N3, NZP_CSI_RS_ResourceSet_Ext_Kdopp_r19_NumberOfResourcesPerGroup_r19_N4},
}

type NZP_CSI_RS_ResourceSet struct {
	Nzp_CSI_ResourceSetId           NZP_CSI_RS_ResourceSetId
	Nzp_CSI_RS_Resources            []NZP_CSI_RS_ResourceId
	Repetition                      *int64
	AperiodicTriggeringOffset       *int64
	Trs_Info                        *int64
	AperiodicTriggeringOffset_r16   *int64
	Pdc_Info_r17                    *int64
	CmrGroupingAndPairing_r17       *CMRGroupingAndPairing_r17
	AperiodicTriggeringOffset_r17   *int64
	AperiodicTriggeringOffsetL2_r17 *int64
	ResourceType_r18                *int64
	Kdopp_r19                       *struct {
		NumberOfResourceGroups_r19    int64
		NumberOfResourcesPerGroup_r19 int64
	}
	AdditionalOneSlotOffsetDoppler_r19 *struct {
		Choice                     int
		TwoResourcesPerGroup_r19   *per.BitString
		ThreeResourcesPerGroup_r19 *per.BitString
		FourResourcesPerGroup_r19  *per.BitString
	}
}

func (ie *NZP_CSI_RS_ResourceSet) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nZPCSIRSResourceSetConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.AperiodicTriggeringOffset_r16 != nil
	hasExtGroup1 := ie.Pdc_Info_r17 != nil || ie.CmrGroupingAndPairing_r17 != nil || ie.AperiodicTriggeringOffset_r17 != nil || ie.AperiodicTriggeringOffsetL2_r17 != nil
	hasExtGroup2 := ie.ResourceType_r18 != nil
	hasExtGroup3 := ie.Kdopp_r19 != nil || ie.AdditionalOneSlotOffsetDoppler_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Repetition != nil, ie.AperiodicTriggeringOffset != nil, ie.Trs_Info != nil}); err != nil {
		return err
	}

	// 3. nzp-CSI-ResourceSetId: ref
	{
		if err := ie.Nzp_CSI_ResourceSetId.Encode(e); err != nil {
			return err
		}
	}

	// 4. nzp-CSI-RS-Resources: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(nZPCSIRSResourceSetNzpCSIRSResourcesConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Nzp_CSI_RS_Resources))); err != nil {
			return err
		}
		for i := range ie.Nzp_CSI_RS_Resources {
			if err := ie.Nzp_CSI_RS_Resources[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. repetition: enumerated
	{
		if ie.Repetition != nil {
			if err := e.EncodeEnumerated(*ie.Repetition, nZPCSIRSResourceSetRepetitionConstraints); err != nil {
				return err
			}
		}
	}

	// 6. aperiodicTriggeringOffset: integer
	{
		if ie.AperiodicTriggeringOffset != nil {
			if err := e.EncodeInteger(*ie.AperiodicTriggeringOffset, nZPCSIRSResourceSetAperiodicTriggeringOffsetConstraints); err != nil {
				return err
			}
		}
	}

	// 7. trs-Info: enumerated
	{
		if ie.Trs_Info != nil {
			if err := e.EncodeEnumerated(*ie.Trs_Info, nZPCSIRSResourceSetTrsInfoConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "aperiodicTriggeringOffset-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AperiodicTriggeringOffset_r16 != nil}); err != nil {
				return err
			}

			if ie.AperiodicTriggeringOffset_r16 != nil {
				if err := ex.EncodeInteger(*ie.AperiodicTriggeringOffset_r16, nZPCSIRSResourceSetAperiodicTriggeringOffsetR16Constraints); err != nil {
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
					{Name: "pdc-Info-r17", Optional: true},
					{Name: "cmrGroupingAndPairing-r17", Optional: true},
					{Name: "aperiodicTriggeringOffset-r17", Optional: true},
					{Name: "aperiodicTriggeringOffsetL2-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdc_Info_r17 != nil, ie.CmrGroupingAndPairing_r17 != nil, ie.AperiodicTriggeringOffset_r17 != nil, ie.AperiodicTriggeringOffsetL2_r17 != nil}); err != nil {
				return err
			}

			if ie.Pdc_Info_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdc_Info_r17, nZPCSIRSResourceSetExtPdcInfoR17Constraints); err != nil {
					return err
				}
			}

			if ie.CmrGroupingAndPairing_r17 != nil {
				if err := ie.CmrGroupingAndPairing_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.AperiodicTriggeringOffset_r17 != nil {
				if err := ex.EncodeInteger(*ie.AperiodicTriggeringOffset_r17, nZPCSIRSResourceSetAperiodicTriggeringOffsetR17Constraints); err != nil {
					return err
				}
			}

			if ie.AperiodicTriggeringOffsetL2_r17 != nil {
				if err := ex.EncodeInteger(*ie.AperiodicTriggeringOffsetL2_r17, nZPCSIRSResourceSetAperiodicTriggeringOffsetL2R17Constraints); err != nil {
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
					{Name: "resourceType-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ResourceType_r18 != nil}); err != nil {
				return err
			}

			if ie.ResourceType_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ResourceType_r18, nZPCSIRSResourceSetExtResourceTypeR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "kdopp-r19", Optional: true},
					{Name: "additionalOneSlotOffsetDoppler-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Kdopp_r19 != nil, ie.AdditionalOneSlotOffsetDoppler_r19 != nil}); err != nil {
				return err
			}

			if ie.Kdopp_r19 != nil {
				c := ie.Kdopp_r19
				if err := ex.EncodeEnumerated(c.NumberOfResourceGroups_r19, nZPCSIRSResourceSetExtKdoppR19NumberOfResourceGroupsR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.NumberOfResourcesPerGroup_r19, nZPCSIRSResourceSetExtKdoppR19NumberOfResourcesPerGroupR19Constraints); err != nil {
					return err
				}
			}

			if ie.AdditionalOneSlotOffsetDoppler_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(nZPCSIRSResourceSetExtAdditionalOneSlotOffsetDopplerR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.AdditionalOneSlotOffsetDoppler_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.AdditionalOneSlotOffsetDoppler_r19).Choice {
				case NZP_CSI_RS_ResourceSet_Ext_AdditionalOneSlotOffsetDoppler_r19_TwoResourcesPerGroup_r19:
					if err := ex.EncodeBitString((*(*ie.AdditionalOneSlotOffsetDoppler_r19).TwoResourcesPerGroup_r19), per.FixedSize(2)); err != nil {
						return err
					}
				case NZP_CSI_RS_ResourceSet_Ext_AdditionalOneSlotOffsetDoppler_r19_ThreeResourcesPerGroup_r19:
					if err := ex.EncodeBitString((*(*ie.AdditionalOneSlotOffsetDoppler_r19).ThreeResourcesPerGroup_r19), per.FixedSize(3)); err != nil {
						return err
					}
				case NZP_CSI_RS_ResourceSet_Ext_AdditionalOneSlotOffsetDoppler_r19_FourResourcesPerGroup_r19:
					if err := ex.EncodeBitString((*(*ie.AdditionalOneSlotOffsetDoppler_r19).FourResourcesPerGroup_r19), per.FixedSize(4)); err != nil {
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

func (ie *NZP_CSI_RS_ResourceSet) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nZPCSIRSResourceSetConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. nzp-CSI-ResourceSetId: ref
	{
		if err := ie.Nzp_CSI_ResourceSetId.Decode(d); err != nil {
			return err
		}
	}

	// 4. nzp-CSI-RS-Resources: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(nZPCSIRSResourceSetNzpCSIRSResourcesConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Nzp_CSI_RS_Resources = make([]NZP_CSI_RS_ResourceId, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Nzp_CSI_RS_Resources[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. repetition: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(nZPCSIRSResourceSetRepetitionConstraints)
			if err != nil {
				return err
			}
			ie.Repetition = &idx
		}
	}

	// 6. aperiodicTriggeringOffset: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(nZPCSIRSResourceSetAperiodicTriggeringOffsetConstraints)
			if err != nil {
				return err
			}
			ie.AperiodicTriggeringOffset = &v
		}
	}

	// 7. trs-Info: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(nZPCSIRSResourceSetTrsInfoConstraints)
			if err != nil {
				return err
			}
			ie.Trs_Info = &idx
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
				{Name: "aperiodicTriggeringOffset-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(nZPCSIRSResourceSetAperiodicTriggeringOffsetR16Constraints)
			if err != nil {
				return err
			}
			ie.AperiodicTriggeringOffset_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pdc-Info-r17", Optional: true},
				{Name: "cmrGroupingAndPairing-r17", Optional: true},
				{Name: "aperiodicTriggeringOffset-r17", Optional: true},
				{Name: "aperiodicTriggeringOffsetL2-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(nZPCSIRSResourceSetExtPdcInfoR17Constraints)
			if err != nil {
				return err
			}
			ie.Pdc_Info_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.CmrGroupingAndPairing_r17 = new(CMRGroupingAndPairing_r17)
			if err := ie.CmrGroupingAndPairing_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(nZPCSIRSResourceSetAperiodicTriggeringOffsetR17Constraints)
			if err != nil {
				return err
			}
			ie.AperiodicTriggeringOffset_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeInteger(nZPCSIRSResourceSetAperiodicTriggeringOffsetL2R17Constraints)
			if err != nil {
				return err
			}
			ie.AperiodicTriggeringOffsetL2_r17 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "resourceType-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(nZPCSIRSResourceSetExtResourceTypeR18Constraints)
			if err != nil {
				return err
			}
			ie.ResourceType_r18 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "kdopp-r19", Optional: true},
				{Name: "additionalOneSlotOffsetDoppler-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Kdopp_r19 = &struct {
				NumberOfResourceGroups_r19    int64
				NumberOfResourcesPerGroup_r19 int64
			}{}
			c := ie.Kdopp_r19
			{
				v, err := dx.DecodeEnumerated(nZPCSIRSResourceSetExtKdoppR19NumberOfResourceGroupsR19Constraints)
				if err != nil {
					return err
				}
				c.NumberOfResourceGroups_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(nZPCSIRSResourceSetExtKdoppR19NumberOfResourcesPerGroupR19Constraints)
				if err != nil {
					return err
				}
				c.NumberOfResourcesPerGroup_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.AdditionalOneSlotOffsetDoppler_r19 = &struct {
				Choice                     int
				TwoResourcesPerGroup_r19   *per.BitString
				ThreeResourcesPerGroup_r19 *per.BitString
				FourResourcesPerGroup_r19  *per.BitString
			}{}
			choiceDec := dx.NewChoiceDecoder(nZPCSIRSResourceSetExtAdditionalOneSlotOffsetDopplerR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.AdditionalOneSlotOffsetDoppler_r19).Choice = int(index)
			switch index {
			case NZP_CSI_RS_ResourceSet_Ext_AdditionalOneSlotOffsetDoppler_r19_TwoResourcesPerGroup_r19:
				(*ie.AdditionalOneSlotOffsetDoppler_r19).TwoResourcesPerGroup_r19 = new(per.BitString)
				v, err := dx.DecodeBitString(per.FixedSize(2))
				if err != nil {
					return err
				}
				(*(*ie.AdditionalOneSlotOffsetDoppler_r19).TwoResourcesPerGroup_r19) = v
			case NZP_CSI_RS_ResourceSet_Ext_AdditionalOneSlotOffsetDoppler_r19_ThreeResourcesPerGroup_r19:
				(*ie.AdditionalOneSlotOffsetDoppler_r19).ThreeResourcesPerGroup_r19 = new(per.BitString)
				v, err := dx.DecodeBitString(per.FixedSize(3))
				if err != nil {
					return err
				}
				(*(*ie.AdditionalOneSlotOffsetDoppler_r19).ThreeResourcesPerGroup_r19) = v
			case NZP_CSI_RS_ResourceSet_Ext_AdditionalOneSlotOffsetDoppler_r19_FourResourcesPerGroup_r19:
				(*ie.AdditionalOneSlotOffsetDoppler_r19).FourResourcesPerGroup_r19 = new(per.BitString)
				v, err := dx.DecodeBitString(per.FixedSize(4))
				if err != nil {
					return err
				}
				(*(*ie.AdditionalOneSlotOffsetDoppler_r19).FourResourcesPerGroup_r19) = v
			}
		}
	}

	return nil
}
