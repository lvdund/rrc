// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasurementValidityDuration-r18 (line 10078).
// MeasurementValidityDuration-r18 ::= ENUMERATED {s5, s10, s20, s50, s100, spare3, spare2, spare1}

const (
	MeasurementValidityDuration_r18_S5     = 0
	MeasurementValidityDuration_r18_S10    = 1
	MeasurementValidityDuration_r18_S20    = 2
	MeasurementValidityDuration_r18_S50    = 3
	MeasurementValidityDuration_r18_S100   = 4
	MeasurementValidityDuration_r18_Spare3 = 5
	MeasurementValidityDuration_r18_Spare2 = 6
	MeasurementValidityDuration_r18_Spare1 = 7
)

var measurementValidityDurationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasurementValidityDuration_r18_S5, MeasurementValidityDuration_r18_S10, MeasurementValidityDuration_r18_S20, MeasurementValidityDuration_r18_S50, MeasurementValidityDuration_r18_S100, MeasurementValidityDuration_r18_Spare3, MeasurementValidityDuration_r18_Spare2, MeasurementValidityDuration_r18_Spare1},
}

type MeasurementValidityDuration_r18 struct {
	Value int64
}

func (v *MeasurementValidityDuration_r18) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, measurementValidityDurationR18Constraints)
}

func (v *MeasurementValidityDuration_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(measurementValidityDurationR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
