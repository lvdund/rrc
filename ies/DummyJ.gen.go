// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DummyJ (line 14756).

var dummyJConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxEnergyDetectionThreshold-r16"},
		{Name: "energyDetectionThresholdOffset-r16"},
		{Name: "ul-toDL-COT-SharingED-Threshold-r16", Optional: true},
		{Name: "absenceOfAnyOtherTechnology-r16", Optional: true},
	},
}

var dummyJMaxEnergyDetectionThresholdR16Constraints = per.Constrained(-85, -52)

var dummyJEnergyDetectionThresholdOffsetR16Constraints = per.Constrained(-20, -13)

var dummyJUlToDLCOTSharingEDThresholdR16Constraints = per.Constrained(-85, -52)

const (
	DummyJ_AbsenceOfAnyOtherTechnology_r16_True = 0
)

var dummyJAbsenceOfAnyOtherTechnologyR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyJ_AbsenceOfAnyOtherTechnology_r16_True},
}

type DummyJ struct {
	MaxEnergyDetectionThreshold_r16     int64
	EnergyDetectionThresholdOffset_r16  int64
	Ul_ToDL_COT_SharingED_Threshold_r16 *int64
	AbsenceOfAnyOtherTechnology_r16     *int64
}

func (ie *DummyJ) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dummyJConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ul_ToDL_COT_SharingED_Threshold_r16 != nil, ie.AbsenceOfAnyOtherTechnology_r16 != nil}); err != nil {
		return err
	}

	// 2. maxEnergyDetectionThreshold-r16: integer
	{
		if err := e.EncodeInteger(ie.MaxEnergyDetectionThreshold_r16, dummyJMaxEnergyDetectionThresholdR16Constraints); err != nil {
			return err
		}
	}

	// 3. energyDetectionThresholdOffset-r16: integer
	{
		if err := e.EncodeInteger(ie.EnergyDetectionThresholdOffset_r16, dummyJEnergyDetectionThresholdOffsetR16Constraints); err != nil {
			return err
		}
	}

	// 4. ul-toDL-COT-SharingED-Threshold-r16: integer
	{
		if ie.Ul_ToDL_COT_SharingED_Threshold_r16 != nil {
			if err := e.EncodeInteger(*ie.Ul_ToDL_COT_SharingED_Threshold_r16, dummyJUlToDLCOTSharingEDThresholdR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. absenceOfAnyOtherTechnology-r16: enumerated
	{
		if ie.AbsenceOfAnyOtherTechnology_r16 != nil {
			if err := e.EncodeEnumerated(*ie.AbsenceOfAnyOtherTechnology_r16, dummyJAbsenceOfAnyOtherTechnologyR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DummyJ) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dummyJConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. maxEnergyDetectionThreshold-r16: integer
	{
		v0, err := d.DecodeInteger(dummyJMaxEnergyDetectionThresholdR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxEnergyDetectionThreshold_r16 = v0
	}

	// 3. energyDetectionThresholdOffset-r16: integer
	{
		v1, err := d.DecodeInteger(dummyJEnergyDetectionThresholdOffsetR16Constraints)
		if err != nil {
			return err
		}
		ie.EnergyDetectionThresholdOffset_r16 = v1
	}

	// 4. ul-toDL-COT-SharingED-Threshold-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(dummyJUlToDLCOTSharingEDThresholdR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_ToDL_COT_SharingED_Threshold_r16 = &v
		}
	}

	// 5. absenceOfAnyOtherTechnology-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(dummyJAbsenceOfAnyOtherTechnologyR16Constraints)
			if err != nil {
				return err
			}
			ie.AbsenceOfAnyOtherTechnology_r16 = &idx
		}
	}

	return nil
}
