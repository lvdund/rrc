// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-ServingCellConfig (line 11159).

var pDCCHServingCellConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "slotFormatIndicator", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var pDCCH_ServingCellConfigSlotFormatIndicatorConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDCCH_ServingCellConfig_SlotFormatIndicator_Release = 0
	PDCCH_ServingCellConfig_SlotFormatIndicator_Setup   = 1
)

var pDCCHServingCellConfigExtAvailabilityIndicatorR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDCCH_ServingCellConfig_Ext_AvailabilityIndicator_r16_Release = 0
	PDCCH_ServingCellConfig_Ext_AvailabilityIndicator_r16_Setup   = 1
)

var pDCCHServingCellConfigSearchSpaceSwitchTimerR16Constraints = per.Constrained(1, 80)

var pDCCHServingCellConfigSearchSpaceSwitchTimerV1710Constraints = per.Constrained(81, 1280)

type PDCCH_ServingCellConfig struct {
	SlotFormatIndicator *struct {
		Choice  int
		Release *struct{}
		Setup   *SlotFormatIndicator
	}
	AvailabilityIndicator_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *AvailabilityIndicator_r16
	}
	SearchSpaceSwitchTimer_r16   *int64
	SearchSpaceSwitchTimer_v1710 *int64
}

func (ie *PDCCH_ServingCellConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHServingCellConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.AvailabilityIndicator_r16 != nil || ie.SearchSpaceSwitchTimer_r16 != nil
	hasExtGroup1 := ie.SearchSpaceSwitchTimer_v1710 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SlotFormatIndicator != nil}); err != nil {
		return err
	}

	// 3. slotFormatIndicator: choice
	{
		if ie.SlotFormatIndicator != nil {
			choiceEnc := e.NewChoiceEncoder(pDCCH_ServingCellConfigSlotFormatIndicatorConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.SlotFormatIndicator).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.SlotFormatIndicator).Choice {
			case PDCCH_ServingCellConfig_SlotFormatIndicator_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PDCCH_ServingCellConfig_SlotFormatIndicator_Setup:
				if err := (*ie.SlotFormatIndicator).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.SlotFormatIndicator).Choice), Constraint: "index out of range"}
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
					{Name: "availabilityIndicator-r16", Optional: true},
					{Name: "searchSpaceSwitchTimer-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AvailabilityIndicator_r16 != nil, ie.SearchSpaceSwitchTimer_r16 != nil}); err != nil {
				return err
			}

			if ie.AvailabilityIndicator_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDCCHServingCellConfigExtAvailabilityIndicatorR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.AvailabilityIndicator_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.AvailabilityIndicator_r16).Choice {
				case PDCCH_ServingCellConfig_Ext_AvailabilityIndicator_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDCCH_ServingCellConfig_Ext_AvailabilityIndicator_r16_Setup:
					if err := (*ie.AvailabilityIndicator_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SearchSpaceSwitchTimer_r16 != nil {
				if err := ex.EncodeInteger(*ie.SearchSpaceSwitchTimer_r16, pDCCHServingCellConfigSearchSpaceSwitchTimerR16Constraints); err != nil {
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
					{Name: "searchSpaceSwitchTimer-v1710", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SearchSpaceSwitchTimer_v1710 != nil}); err != nil {
				return err
			}

			if ie.SearchSpaceSwitchTimer_v1710 != nil {
				if err := ex.EncodeInteger(*ie.SearchSpaceSwitchTimer_v1710, pDCCHServingCellConfigSearchSpaceSwitchTimerV1710Constraints); err != nil {
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

func (ie *PDCCH_ServingCellConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHServingCellConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. slotFormatIndicator: choice
	{
		if seq.IsComponentPresent(0) {
			ie.SlotFormatIndicator = &struct {
				Choice  int
				Release *struct{}
				Setup   *SlotFormatIndicator
			}{}
			choiceDec := d.NewChoiceDecoder(pDCCH_ServingCellConfigSlotFormatIndicatorConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SlotFormatIndicator).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PDCCH_ServingCellConfig_SlotFormatIndicator_Release:
				(*ie.SlotFormatIndicator).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PDCCH_ServingCellConfig_SlotFormatIndicator_Setup:
				(*ie.SlotFormatIndicator).Setup = new(SlotFormatIndicator)
				if err := (*ie.SlotFormatIndicator).Setup.Decode(d); err != nil {
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
				{Name: "availabilityIndicator-r16", Optional: true},
				{Name: "searchSpaceSwitchTimer-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.AvailabilityIndicator_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *AvailabilityIndicator_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pDCCHServingCellConfigExtAvailabilityIndicatorR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.AvailabilityIndicator_r16).Choice = int(index)
			switch index {
			case PDCCH_ServingCellConfig_Ext_AvailabilityIndicator_r16_Release:
				(*ie.AvailabilityIndicator_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDCCH_ServingCellConfig_Ext_AvailabilityIndicator_r16_Setup:
				(*ie.AvailabilityIndicator_r16).Setup = new(AvailabilityIndicator_r16)
				if err := (*ie.AvailabilityIndicator_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(pDCCHServingCellConfigSearchSpaceSwitchTimerR16Constraints)
			if err != nil {
				return err
			}
			ie.SearchSpaceSwitchTimer_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "searchSpaceSwitchTimer-v1710", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(pDCCHServingCellConfigSearchSpaceSwitchTimerV1710Constraints)
			if err != nil {
				return err
			}
			ie.SearchSpaceSwitchTimer_v1710 = &v
		}
	}

	return nil
}
