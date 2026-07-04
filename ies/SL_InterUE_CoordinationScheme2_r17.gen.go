// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-InterUE-CoordinationScheme2-r17 (line 27375).

var sLInterUECoordinationScheme2R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-IUC-Scheme2-r17", Optional: true},
		{Name: "sl-RB-SetPSFCH-r17", Optional: true},
		{Name: "sl-TypeUE-A-r17", Optional: true},
		{Name: "sl-PSFCH-Occasion-r17", Optional: true},
		{Name: "sl-SlotLevelResourceExclusion-r17", Optional: true},
		{Name: "sl-OptionForCondition2-A-1-r17", Optional: true},
		{Name: "sl-IndicationUE-B-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	SL_InterUE_CoordinationScheme2_r17_Sl_IUC_Scheme2_r17_Enabled = 0
)

var sLInterUECoordinationScheme2R17SlIUCScheme2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_InterUE_CoordinationScheme2_r17_Sl_IUC_Scheme2_r17_Enabled},
}

var sLInterUECoordinationScheme2R17SlRBSetPSFCHR17Constraints = per.SizeRange(10, 275)

const (
	SL_InterUE_CoordinationScheme2_r17_Sl_TypeUE_A_r17_Enabled = 0
)

var sLInterUECoordinationScheme2R17SlTypeUEAR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_InterUE_CoordinationScheme2_r17_Sl_TypeUE_A_r17_Enabled},
}

var sLInterUECoordinationScheme2R17SlPSFCHOccasionR17Constraints = per.Constrained(0, 1)

const (
	SL_InterUE_CoordinationScheme2_r17_Sl_SlotLevelResourceExclusion_r17_Enabled = 0
)

var sLInterUECoordinationScheme2R17SlSlotLevelResourceExclusionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_InterUE_CoordinationScheme2_r17_Sl_SlotLevelResourceExclusion_r17_Enabled},
}

var sLInterUECoordinationScheme2R17SlOptionForCondition2A1R17Constraints = per.Constrained(0, 1)

const (
	SL_InterUE_CoordinationScheme2_r17_Sl_IndicationUE_B_r17_Enabled  = 0
	SL_InterUE_CoordinationScheme2_r17_Sl_IndicationUE_B_r17_Disabled = 1
)

var sLInterUECoordinationScheme2R17SlIndicationUEBR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_InterUE_CoordinationScheme2_r17_Sl_IndicationUE_B_r17_Enabled, SL_InterUE_CoordinationScheme2_r17_Sl_IndicationUE_B_r17_Disabled},
}

var sLInterUECoordinationScheme2R17SlDeltaRSRPThreshV1720Constraints = per.Constrained(-30, 30)

type SL_InterUE_CoordinationScheme2_r17 struct {
	Sl_IUC_Scheme2_r17                *int64
	Sl_RB_SetPSFCH_r17                *per.BitString
	Sl_TypeUE_A_r17                   *int64
	Sl_PSFCH_Occasion_r17             *int64
	Sl_SlotLevelResourceExclusion_r17 *int64
	Sl_OptionForCondition2_A_1_r17    *int64
	Sl_IndicationUE_B_r17             *int64
	Sl_DeltaRSRP_Thresh_v1720         *int64
}

func (ie *SL_InterUE_CoordinationScheme2_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLInterUECoordinationScheme2R17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_DeltaRSRP_Thresh_v1720 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_IUC_Scheme2_r17 != nil, ie.Sl_RB_SetPSFCH_r17 != nil, ie.Sl_TypeUE_A_r17 != nil, ie.Sl_PSFCH_Occasion_r17 != nil, ie.Sl_SlotLevelResourceExclusion_r17 != nil, ie.Sl_OptionForCondition2_A_1_r17 != nil, ie.Sl_IndicationUE_B_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-IUC-Scheme2-r17: enumerated
	{
		if ie.Sl_IUC_Scheme2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_IUC_Scheme2_r17, sLInterUECoordinationScheme2R17SlIUCScheme2R17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-RB-SetPSFCH-r17: bit-string
	{
		if ie.Sl_RB_SetPSFCH_r17 != nil {
			if err := e.EncodeBitString(*ie.Sl_RB_SetPSFCH_r17, sLInterUECoordinationScheme2R17SlRBSetPSFCHR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-TypeUE-A-r17: enumerated
	{
		if ie.Sl_TypeUE_A_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_TypeUE_A_r17, sLInterUECoordinationScheme2R17SlTypeUEAR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-PSFCH-Occasion-r17: integer
	{
		if ie.Sl_PSFCH_Occasion_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_PSFCH_Occasion_r17, sLInterUECoordinationScheme2R17SlPSFCHOccasionR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sl-SlotLevelResourceExclusion-r17: enumerated
	{
		if ie.Sl_SlotLevelResourceExclusion_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_SlotLevelResourceExclusion_r17, sLInterUECoordinationScheme2R17SlSlotLevelResourceExclusionR17Constraints); err != nil {
				return err
			}
		}
	}

	// 8. sl-OptionForCondition2-A-1-r17: integer
	{
		if ie.Sl_OptionForCondition2_A_1_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_OptionForCondition2_A_1_r17, sLInterUECoordinationScheme2R17SlOptionForCondition2A1R17Constraints); err != nil {
				return err
			}
		}
	}

	// 9. sl-IndicationUE-B-r17: enumerated
	{
		if ie.Sl_IndicationUE_B_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_IndicationUE_B_r17, sLInterUECoordinationScheme2R17SlIndicationUEBR17Constraints); err != nil {
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
					{Name: "sl-DeltaRSRP-Thresh-v1720", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_DeltaRSRP_Thresh_v1720 != nil}); err != nil {
				return err
			}

			if ie.Sl_DeltaRSRP_Thresh_v1720 != nil {
				if err := ex.EncodeInteger(*ie.Sl_DeltaRSRP_Thresh_v1720, sLInterUECoordinationScheme2R17SlDeltaRSRPThreshV1720Constraints); err != nil {
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

func (ie *SL_InterUE_CoordinationScheme2_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLInterUECoordinationScheme2R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-IUC-Scheme2-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sLInterUECoordinationScheme2R17SlIUCScheme2R17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_IUC_Scheme2_r17 = &idx
		}
	}

	// 4. sl-RB-SetPSFCH-r17: bit-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeBitString(sLInterUECoordinationScheme2R17SlRBSetPSFCHR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_RB_SetPSFCH_r17 = &v
		}
	}

	// 5. sl-TypeUE-A-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLInterUECoordinationScheme2R17SlTypeUEAR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TypeUE_A_r17 = &idx
		}
	}

	// 6. sl-PSFCH-Occasion-r17: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(sLInterUECoordinationScheme2R17SlPSFCHOccasionR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PSFCH_Occasion_r17 = &v
		}
	}

	// 7. sl-SlotLevelResourceExclusion-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(sLInterUECoordinationScheme2R17SlSlotLevelResourceExclusionR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SlotLevelResourceExclusion_r17 = &idx
		}
	}

	// 8. sl-OptionForCondition2-A-1-r17: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(sLInterUECoordinationScheme2R17SlOptionForCondition2A1R17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_OptionForCondition2_A_1_r17 = &v
		}
	}

	// 9. sl-IndicationUE-B-r17: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(sLInterUECoordinationScheme2R17SlIndicationUEBR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_IndicationUE_B_r17 = &idx
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
				{Name: "sl-DeltaRSRP-Thresh-v1720", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(sLInterUECoordinationScheme2R17SlDeltaRSRPThreshV1720Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DeltaRSRP_Thresh_v1720 = &v
		}
	}

	return nil
}
