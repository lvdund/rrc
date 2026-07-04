// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Sensor-LocationInfo-r16 (line 14594).

var sensorLocationInfoR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sensor-MeasurementInformation-r16", Optional: true},
		{Name: "sensor-MotionInformation-r16", Optional: true},
	},
}

var sensorLocationInfoR16SensorMeasurementInformationR16Constraints = per.SizeConstraints{}

var sensorLocationInfoR16SensorMotionInformationR16Constraints = per.SizeConstraints{}

type Sensor_LocationInfo_r16 struct {
	Sensor_MeasurementInformation_r16 []byte
	Sensor_MotionInformation_r16      []byte
}

func (ie *Sensor_LocationInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sensorLocationInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sensor_MeasurementInformation_r16 != nil, ie.Sensor_MotionInformation_r16 != nil}); err != nil {
		return err
	}

	// 3. sensor-MeasurementInformation-r16: octet-string
	{
		if ie.Sensor_MeasurementInformation_r16 != nil {
			if err := e.EncodeOctetString(ie.Sensor_MeasurementInformation_r16, sensorLocationInfoR16SensorMeasurementInformationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sensor-MotionInformation-r16: octet-string
	{
		if ie.Sensor_MotionInformation_r16 != nil {
			if err := e.EncodeOctetString(ie.Sensor_MotionInformation_r16, sensorLocationInfoR16SensorMotionInformationR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Sensor_LocationInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sensorLocationInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sensor-MeasurementInformation-r16: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(sensorLocationInfoR16SensorMeasurementInformationR16Constraints)
			if err != nil {
				return err
			}
			ie.Sensor_MeasurementInformation_r16 = v
		}
	}

	// 4. sensor-MotionInformation-r16: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(sensorLocationInfoR16SensorMotionInformationR16Constraints)
			if err != nil {
				return err
			}
			ie.Sensor_MotionInformation_r16 = v
		}
	}

	return nil
}
