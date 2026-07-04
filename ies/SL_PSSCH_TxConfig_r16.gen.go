// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PSSCH-TxConfig-r16 (line 27718).

var sLPSSCHTxConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-TypeTxSync-r16", Optional: true},
		{Name: "sl-ThresUE-Speed-r16"},
		{Name: "sl-ParametersAboveThres-r16"},
		{Name: "sl-ParametersBelowThres-r16"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph60  = 0
	SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph80  = 1
	SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph100 = 2
	SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph120 = 3
	SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph140 = 4
	SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph160 = 5
	SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph180 = 6
	SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph200 = 7
)

var sLPSSCHTxConfigR16SlThresUESpeedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph60, SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph80, SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph100, SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph120, SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph140, SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph160, SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph180, SL_PSSCH_TxConfig_r16_Sl_ThresUE_Speed_r16_Kmph200},
}

type SL_PSSCH_TxConfig_r16 struct {
	Sl_TypeTxSync_r16             *SL_TypeTxSync_r16
	Sl_ThresUE_Speed_r16          int64
	Sl_ParametersAboveThres_r16   SL_PSSCH_TxParameters_r16
	Sl_ParametersBelowThres_r16   SL_PSSCH_TxParameters_r16
	Sl_ParametersAboveThres_v1650 *SL_MinMaxMCS_List_r16
	Sl_ParametersBelowThres_v1650 *SL_MinMaxMCS_List_r16
}

func (ie *SL_PSSCH_TxConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPSSCHTxConfigR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_ParametersAboveThres_v1650 != nil || ie.Sl_ParametersBelowThres_v1650 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_TypeTxSync_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-TypeTxSync-r16: ref
	{
		if ie.Sl_TypeTxSync_r16 != nil {
			if err := ie.Sl_TypeTxSync_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-ThresUE-Speed-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_ThresUE_Speed_r16, sLPSSCHTxConfigR16SlThresUESpeedR16Constraints); err != nil {
			return err
		}
	}

	// 5. sl-ParametersAboveThres-r16: ref
	{
		if err := ie.Sl_ParametersAboveThres_r16.Encode(e); err != nil {
			return err
		}
	}

	// 6. sl-ParametersBelowThres-r16: ref
	{
		if err := ie.Sl_ParametersBelowThres_r16.Encode(e); err != nil {
			return err
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
					{Name: "sl-ParametersAboveThres-v1650", Optional: true},
					{Name: "sl-ParametersBelowThres-v1650", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_ParametersAboveThres_v1650 != nil, ie.Sl_ParametersBelowThres_v1650 != nil}); err != nil {
				return err
			}

			if ie.Sl_ParametersAboveThres_v1650 != nil {
				if err := ie.Sl_ParametersAboveThres_v1650.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_ParametersBelowThres_v1650 != nil {
				if err := ie.Sl_ParametersBelowThres_v1650.Encode(ex); err != nil {
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

func (ie *SL_PSSCH_TxConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPSSCHTxConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-TypeTxSync-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_TypeTxSync_r16 = new(SL_TypeTxSync_r16)
			if err := ie.Sl_TypeTxSync_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-ThresUE-Speed-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(sLPSSCHTxConfigR16SlThresUESpeedR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_ThresUE_Speed_r16 = v1
	}

	// 5. sl-ParametersAboveThres-r16: ref
	{
		if err := ie.Sl_ParametersAboveThres_r16.Decode(d); err != nil {
			return err
		}
	}

	// 6. sl-ParametersBelowThres-r16: ref
	{
		if err := ie.Sl_ParametersBelowThres_r16.Decode(d); err != nil {
			return err
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
				{Name: "sl-ParametersAboveThres-v1650", Optional: true},
				{Name: "sl-ParametersBelowThres-v1650", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_ParametersAboveThres_v1650 = new(SL_MinMaxMCS_List_r16)
			if err := ie.Sl_ParametersAboveThres_v1650.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sl_ParametersBelowThres_v1650 = new(SL_MinMaxMCS_List_r16)
			if err := ie.Sl_ParametersBelowThres_v1650.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
