// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Sensor-NameList-r16 (line 26650).

var sensorNameListR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measUncomBarPre-r16", Optional: true},
		{Name: "measUeSpeed", Optional: true},
		{Name: "measUeOrientation", Optional: true},
	},
}

const (
	Sensor_NameList_r16_MeasUncomBarPre_r16_True = 0
)

var sensorNameListR16MeasUncomBarPreR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Sensor_NameList_r16_MeasUncomBarPre_r16_True},
}

const (
	Sensor_NameList_r16_MeasUeSpeed_True = 0
)

var sensorNameListR16MeasUeSpeedConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Sensor_NameList_r16_MeasUeSpeed_True},
}

const (
	Sensor_NameList_r16_MeasUeOrientation_True = 0
)

var sensorNameListR16MeasUeOrientationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Sensor_NameList_r16_MeasUeOrientation_True},
}

type Sensor_NameList_r16 struct {
	MeasUncomBarPre_r16 *int64
	MeasUeSpeed         *int64
	MeasUeOrientation   *int64
}

func (ie *Sensor_NameList_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sensorNameListR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasUncomBarPre_r16 != nil, ie.MeasUeSpeed != nil, ie.MeasUeOrientation != nil}); err != nil {
		return err
	}

	// 2. measUncomBarPre-r16: enumerated
	{
		if ie.MeasUncomBarPre_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MeasUncomBarPre_r16, sensorNameListR16MeasUncomBarPreR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. measUeSpeed: enumerated
	{
		if ie.MeasUeSpeed != nil {
			if err := e.EncodeEnumerated(*ie.MeasUeSpeed, sensorNameListR16MeasUeSpeedConstraints); err != nil {
				return err
			}
		}
	}

	// 4. measUeOrientation: enumerated
	{
		if ie.MeasUeOrientation != nil {
			if err := e.EncodeEnumerated(*ie.MeasUeOrientation, sensorNameListR16MeasUeOrientationConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Sensor_NameList_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sensorNameListR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. measUncomBarPre-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sensorNameListR16MeasUncomBarPreR16Constraints)
			if err != nil {
				return err
			}
			ie.MeasUncomBarPre_r16 = &idx
		}
	}

	// 3. measUeSpeed: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sensorNameListR16MeasUeSpeedConstraints)
			if err != nil {
				return err
			}
			ie.MeasUeSpeed = &idx
		}
	}

	// 4. measUeOrientation: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sensorNameListR16MeasUeOrientationConstraints)
			if err != nil {
				return err
			}
			ie.MeasUeOrientation = &idx
		}
	}

	return nil
}
