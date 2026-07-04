// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-ConfigDedicatedNR-r16 (line 26945).

var sLConfigDedicatedNRR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PHY-MAC-RLC-Config-r16", Optional: true},
		{Name: "sl-RadioBearerToReleaseList-r16", Optional: true},
		{Name: "sl-RadioBearerToAddModList-r16", Optional: true},
		{Name: "sl-MeasConfigInfoToReleaseList-r16", Optional: true},
		{Name: "sl-MeasConfigInfoToAddModList-r16", Optional: true},
		{Name: "t400-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
	},
}

var sLConfigDedicatedNRR16SlRadioBearerToReleaseListR16Constraints = per.SizeRange(1, common.MaxNrofSLRB_r16)

var sLConfigDedicatedNRR16SlRadioBearerToAddModListR16Constraints = per.SizeRange(1, common.MaxNrofSLRB_r16)

var sLConfigDedicatedNRR16SlMeasConfigInfoToReleaseListR16Constraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

var sLConfigDedicatedNRR16SlMeasConfigInfoToAddModListR16Constraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

const (
	SL_ConfigDedicatedNR_r16_T400_r16_Ms100  = 0
	SL_ConfigDedicatedNR_r16_T400_r16_Ms200  = 1
	SL_ConfigDedicatedNR_r16_T400_r16_Ms300  = 2
	SL_ConfigDedicatedNR_r16_T400_r16_Ms400  = 3
	SL_ConfigDedicatedNR_r16_T400_r16_Ms600  = 4
	SL_ConfigDedicatedNR_r16_T400_r16_Ms1000 = 5
	SL_ConfigDedicatedNR_r16_T400_r16_Ms1500 = 6
	SL_ConfigDedicatedNR_r16_T400_r16_Ms2000 = 7
)

var sLConfigDedicatedNRR16T400R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_ConfigDedicatedNR_r16_T400_r16_Ms100, SL_ConfigDedicatedNR_r16_T400_r16_Ms200, SL_ConfigDedicatedNR_r16_T400_r16_Ms300, SL_ConfigDedicatedNR_r16_T400_r16_Ms400, SL_ConfigDedicatedNR_r16_T400_r16_Ms600, SL_ConfigDedicatedNR_r16_T400_r16_Ms1000, SL_ConfigDedicatedNR_r16_T400_r16_Ms1500, SL_ConfigDedicatedNR_r16_T400_r16_Ms2000},
}

var sLConfigDedicatedNRR16ExtSlPHYMACRLCConfigV1700Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_ConfigDedicatedNR_r16_Ext_Sl_PHY_MAC_RLC_Config_v1700_Release = 0
	SL_ConfigDedicatedNR_r16_Ext_Sl_PHY_MAC_RLC_Config_v1700_Setup   = 1
)

var sLConfigDedicatedNRR16ExtSlDiscConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_ConfigDedicatedNR_r16_Ext_Sl_DiscConfig_r17_Release = 0
	SL_ConfigDedicatedNR_r16_Ext_Sl_DiscConfig_r17_Setup   = 1
)

type SL_ConfigDedicatedNR_r16 struct {
	Sl_PHY_MAC_RLC_Config_r16          *SL_PHY_MAC_RLC_Config_r16
	Sl_RadioBearerToReleaseList_r16    []SLRB_Uu_ConfigIndex_r16
	Sl_RadioBearerToAddModList_r16     []SL_RadioBearerConfig_r16
	Sl_MeasConfigInfoToReleaseList_r16 []SL_DestinationIndex_r16
	Sl_MeasConfigInfoToAddModList_r16  []SL_MeasConfigInfo_r16
	T400_r16                           *int64
	Sl_PHY_MAC_RLC_Config_v1700        *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_PHY_MAC_RLC_Config_v1700
	}
	Sl_DiscConfig_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_DiscConfig_r17
	}
	Sl_DiscConfig_v1800 *SL_DiscConfig_v1800
	Sl_DiscConfig_v1830 *SL_DiscConfig_v1830
	Sl_DiscConfig_v1840 *SL_DiscConfig_v1840
	Sl_DiscConfig_v1900 *SL_DiscConfig_v1900
}

func (ie *SL_ConfigDedicatedNR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLConfigDedicatedNRR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_PHY_MAC_RLC_Config_v1700 != nil || ie.Sl_DiscConfig_r17 != nil
	hasExtGroup1 := ie.Sl_DiscConfig_v1800 != nil
	hasExtGroup2 := ie.Sl_DiscConfig_v1830 != nil
	hasExtGroup3 := ie.Sl_DiscConfig_v1840 != nil
	hasExtGroup4 := ie.Sl_DiscConfig_v1900 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PHY_MAC_RLC_Config_r16 != nil, ie.Sl_RadioBearerToReleaseList_r16 != nil, ie.Sl_RadioBearerToAddModList_r16 != nil, ie.Sl_MeasConfigInfoToReleaseList_r16 != nil, ie.Sl_MeasConfigInfoToAddModList_r16 != nil, ie.T400_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-PHY-MAC-RLC-Config-r16: ref
	{
		if ie.Sl_PHY_MAC_RLC_Config_r16 != nil {
			if err := ie.Sl_PHY_MAC_RLC_Config_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-RadioBearerToReleaseList-r16: sequence-of
	{
		if ie.Sl_RadioBearerToReleaseList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLConfigDedicatedNRR16SlRadioBearerToReleaseListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_RadioBearerToReleaseList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_RadioBearerToReleaseList_r16 {
				if err := ie.Sl_RadioBearerToReleaseList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-RadioBearerToAddModList-r16: sequence-of
	{
		if ie.Sl_RadioBearerToAddModList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLConfigDedicatedNRR16SlRadioBearerToAddModListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_RadioBearerToAddModList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_RadioBearerToAddModList_r16 {
				if err := ie.Sl_RadioBearerToAddModList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. sl-MeasConfigInfoToReleaseList-r16: sequence-of
	{
		if ie.Sl_MeasConfigInfoToReleaseList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLConfigDedicatedNRR16SlMeasConfigInfoToReleaseListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_MeasConfigInfoToReleaseList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_MeasConfigInfoToReleaseList_r16 {
				if err := ie.Sl_MeasConfigInfoToReleaseList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-MeasConfigInfoToAddModList-r16: sequence-of
	{
		if ie.Sl_MeasConfigInfoToAddModList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLConfigDedicatedNRR16SlMeasConfigInfoToAddModListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_MeasConfigInfoToAddModList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_MeasConfigInfoToAddModList_r16 {
				if err := ie.Sl_MeasConfigInfoToAddModList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. t400-r16: enumerated
	{
		if ie.T400_r16 != nil {
			if err := e.EncodeEnumerated(*ie.T400_r16, sLConfigDedicatedNRR16T400R16Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-PHY-MAC-RLC-Config-v1700", Optional: true},
					{Name: "sl-DiscConfig-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_PHY_MAC_RLC_Config_v1700 != nil, ie.Sl_DiscConfig_r17 != nil}); err != nil {
				return err
			}

			if ie.Sl_PHY_MAC_RLC_Config_v1700 != nil {
				choiceEnc := ex.NewChoiceEncoder(sLConfigDedicatedNRR16ExtSlPHYMACRLCConfigV1700Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_PHY_MAC_RLC_Config_v1700).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_PHY_MAC_RLC_Config_v1700).Choice {
				case SL_ConfigDedicatedNR_r16_Ext_Sl_PHY_MAC_RLC_Config_v1700_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SL_ConfigDedicatedNR_r16_Ext_Sl_PHY_MAC_RLC_Config_v1700_Setup:
					if err := (*ie.Sl_PHY_MAC_RLC_Config_v1700).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_DiscConfig_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(sLConfigDedicatedNRR16ExtSlDiscConfigR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_DiscConfig_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_DiscConfig_r17).Choice {
				case SL_ConfigDedicatedNR_r16_Ext_Sl_DiscConfig_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SL_ConfigDedicatedNR_r16_Ext_Sl_DiscConfig_r17_Setup:
					if err := (*ie.Sl_DiscConfig_r17).Setup.Encode(ex); err != nil {
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
					{Name: "sl-DiscConfig-v1800", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_DiscConfig_v1800 != nil}); err != nil {
				return err
			}

			if ie.Sl_DiscConfig_v1800 != nil {
				if err := ie.Sl_DiscConfig_v1800.Encode(ex); err != nil {
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
					{Name: "sl-DiscConfig-v1830", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_DiscConfig_v1830 != nil}); err != nil {
				return err
			}

			if ie.Sl_DiscConfig_v1830 != nil {
				if err := ie.Sl_DiscConfig_v1830.Encode(ex); err != nil {
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
					{Name: "sl-DiscConfig-v1840", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_DiscConfig_v1840 != nil}); err != nil {
				return err
			}

			if ie.Sl_DiscConfig_v1840 != nil {
				if err := ie.Sl_DiscConfig_v1840.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 4:
		if hasExtGroup4 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-DiscConfig-v1900", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_DiscConfig_v1900 != nil}); err != nil {
				return err
			}

			if ie.Sl_DiscConfig_v1900 != nil {
				if err := ie.Sl_DiscConfig_v1900.Encode(ex); err != nil {
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

func (ie *SL_ConfigDedicatedNR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLConfigDedicatedNRR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-PHY-MAC-RLC-Config-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_PHY_MAC_RLC_Config_r16 = new(SL_PHY_MAC_RLC_Config_r16)
			if err := ie.Sl_PHY_MAC_RLC_Config_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-RadioBearerToReleaseList-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLConfigDedicatedNRR16SlRadioBearerToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RadioBearerToReleaseList_r16 = make([]SLRB_Uu_ConfigIndex_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_RadioBearerToReleaseList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-RadioBearerToAddModList-r16: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(sLConfigDedicatedNRR16SlRadioBearerToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RadioBearerToAddModList_r16 = make([]SL_RadioBearerConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_RadioBearerToAddModList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. sl-MeasConfigInfoToReleaseList-r16: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(sLConfigDedicatedNRR16SlMeasConfigInfoToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_MeasConfigInfoToReleaseList_r16 = make([]SL_DestinationIndex_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_MeasConfigInfoToReleaseList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-MeasConfigInfoToAddModList-r16: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(sLConfigDedicatedNRR16SlMeasConfigInfoToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_MeasConfigInfoToAddModList_r16 = make([]SL_MeasConfigInfo_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_MeasConfigInfoToAddModList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. t400-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(sLConfigDedicatedNRR16T400R16Constraints)
			if err != nil {
				return err
			}
			ie.T400_r16 = &idx
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
				{Name: "sl-PHY-MAC-RLC-Config-v1700", Optional: true},
				{Name: "sl-DiscConfig-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_PHY_MAC_RLC_Config_v1700 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_PHY_MAC_RLC_Config_v1700
			}{}
			choiceDec := dx.NewChoiceDecoder(sLConfigDedicatedNRR16ExtSlPHYMACRLCConfigV1700Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_PHY_MAC_RLC_Config_v1700).Choice = int(index)
			switch index {
			case SL_ConfigDedicatedNR_r16_Ext_Sl_PHY_MAC_RLC_Config_v1700_Release:
				(*ie.Sl_PHY_MAC_RLC_Config_v1700).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SL_ConfigDedicatedNR_r16_Ext_Sl_PHY_MAC_RLC_Config_v1700_Setup:
				(*ie.Sl_PHY_MAC_RLC_Config_v1700).Setup = new(SL_PHY_MAC_RLC_Config_v1700)
				if err := (*ie.Sl_PHY_MAC_RLC_Config_v1700).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sl_DiscConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_DiscConfig_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(sLConfigDedicatedNRR16ExtSlDiscConfigR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_DiscConfig_r17).Choice = int(index)
			switch index {
			case SL_ConfigDedicatedNR_r16_Ext_Sl_DiscConfig_r17_Release:
				(*ie.Sl_DiscConfig_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SL_ConfigDedicatedNR_r16_Ext_Sl_DiscConfig_r17_Setup:
				(*ie.Sl_DiscConfig_r17).Setup = new(SL_DiscConfig_r17)
				if err := (*ie.Sl_DiscConfig_r17).Setup.Decode(dx); err != nil {
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
				{Name: "sl-DiscConfig-v1800", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_DiscConfig_v1800 = new(SL_DiscConfig_v1800)
			if err := ie.Sl_DiscConfig_v1800.Decode(dx); err != nil {
				return err
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
				{Name: "sl-DiscConfig-v1830", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_DiscConfig_v1830 = new(SL_DiscConfig_v1830)
			if err := ie.Sl_DiscConfig_v1830.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sl-DiscConfig-v1840", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_DiscConfig_v1840 = new(SL_DiscConfig_v1840)
			if err := ie.Sl_DiscConfig_v1840.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sl-DiscConfig-v1900", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_DiscConfig_v1900 = new(SL_DiscConfig_v1900)
			if err := ie.Sl_DiscConfig_v1900.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
