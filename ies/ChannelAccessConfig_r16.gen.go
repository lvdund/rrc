// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ChannelAccessConfig-r16 (line 14763).

var channelAccessConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "energyDetectionConfig-r16", Optional: true},
		{Name: "ul-toDL-COT-SharingED-Threshold-r16", Optional: true},
		{Name: "absenceOfAnyOtherTechnology-r16", Optional: true},
	},
}

var channelAccessConfig_r16EnergyDetectionConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "maxEnergyDetectionThreshold-r16"},
		{Name: "energyDetectionThresholdOffset-r16"},
	},
}

const (
	ChannelAccessConfig_r16_EnergyDetectionConfig_r16_MaxEnergyDetectionThreshold_r16    = 0
	ChannelAccessConfig_r16_EnergyDetectionConfig_r16_EnergyDetectionThresholdOffset_r16 = 1
)

var channelAccessConfigR16UlToDLCOTSharingEDThresholdR16Constraints = per.Constrained(-85, -52)

const (
	ChannelAccessConfig_r16_AbsenceOfAnyOtherTechnology_r16_True = 0
)

var channelAccessConfigR16AbsenceOfAnyOtherTechnologyR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ChannelAccessConfig_r16_AbsenceOfAnyOtherTechnology_r16_True},
}

type ChannelAccessConfig_r16 struct {
	EnergyDetectionConfig_r16 *struct {
		Choice                             int
		MaxEnergyDetectionThreshold_r16    *int64
		EnergyDetectionThresholdOffset_r16 *int64
	}
	Ul_ToDL_COT_SharingED_Threshold_r16 *int64
	AbsenceOfAnyOtherTechnology_r16     *int64
}

func (ie *ChannelAccessConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(channelAccessConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.EnergyDetectionConfig_r16 != nil, ie.Ul_ToDL_COT_SharingED_Threshold_r16 != nil, ie.AbsenceOfAnyOtherTechnology_r16 != nil}); err != nil {
		return err
	}

	// 2. energyDetectionConfig-r16: choice
	{
		if ie.EnergyDetectionConfig_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(channelAccessConfig_r16EnergyDetectionConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.EnergyDetectionConfig_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.EnergyDetectionConfig_r16).Choice {
			case ChannelAccessConfig_r16_EnergyDetectionConfig_r16_MaxEnergyDetectionThreshold_r16:
				if err := e.EncodeInteger((*(*ie.EnergyDetectionConfig_r16).MaxEnergyDetectionThreshold_r16), per.Constrained(-85, -52)); err != nil {
					return err
				}
			case ChannelAccessConfig_r16_EnergyDetectionConfig_r16_EnergyDetectionThresholdOffset_r16:
				if err := e.EncodeInteger((*(*ie.EnergyDetectionConfig_r16).EnergyDetectionThresholdOffset_r16), per.Constrained(-13, 20)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.EnergyDetectionConfig_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. ul-toDL-COT-SharingED-Threshold-r16: integer
	{
		if ie.Ul_ToDL_COT_SharingED_Threshold_r16 != nil {
			if err := e.EncodeInteger(*ie.Ul_ToDL_COT_SharingED_Threshold_r16, channelAccessConfigR16UlToDLCOTSharingEDThresholdR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. absenceOfAnyOtherTechnology-r16: enumerated
	{
		if ie.AbsenceOfAnyOtherTechnology_r16 != nil {
			if err := e.EncodeEnumerated(*ie.AbsenceOfAnyOtherTechnology_r16, channelAccessConfigR16AbsenceOfAnyOtherTechnologyR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ChannelAccessConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(channelAccessConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. energyDetectionConfig-r16: choice
	{
		if seq.IsComponentPresent(0) {
			ie.EnergyDetectionConfig_r16 = &struct {
				Choice                             int
				MaxEnergyDetectionThreshold_r16    *int64
				EnergyDetectionThresholdOffset_r16 *int64
			}{}
			choiceDec := d.NewChoiceDecoder(channelAccessConfig_r16EnergyDetectionConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.EnergyDetectionConfig_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case ChannelAccessConfig_r16_EnergyDetectionConfig_r16_MaxEnergyDetectionThreshold_r16:
				(*ie.EnergyDetectionConfig_r16).MaxEnergyDetectionThreshold_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(-85, -52))
				if err != nil {
					return err
				}
				(*(*ie.EnergyDetectionConfig_r16).MaxEnergyDetectionThreshold_r16) = v
			case ChannelAccessConfig_r16_EnergyDetectionConfig_r16_EnergyDetectionThresholdOffset_r16:
				(*ie.EnergyDetectionConfig_r16).EnergyDetectionThresholdOffset_r16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(-13, 20))
				if err != nil {
					return err
				}
				(*(*ie.EnergyDetectionConfig_r16).EnergyDetectionThresholdOffset_r16) = v
			}
		}
	}

	// 3. ul-toDL-COT-SharingED-Threshold-r16: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(channelAccessConfigR16UlToDLCOTSharingEDThresholdR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_ToDL_COT_SharingED_Threshold_r16 = &v
		}
	}

	// 4. absenceOfAnyOtherTechnology-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(channelAccessConfigR16AbsenceOfAnyOtherTechnologyR16Constraints)
			if err != nil {
				return err
			}
			ie.AbsenceOfAnyOtherTechnology_r16 = &idx
		}
	}

	return nil
}
