// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LoggedMeasurementConfiguration-v1800-IEs (line 568).

var loggedMeasurementConfigurationV1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "areaConfiguration-v1800", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type LoggedMeasurementConfiguration_v1800_IEs struct {
	AreaConfiguration_v1800 *AreaConfiguration_v1800
	NonCriticalExtension    *LoggedMeasurementConfiguration_v1900_IEs
}

func (ie *LoggedMeasurementConfiguration_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(loggedMeasurementConfigurationV1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AreaConfiguration_v1800 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. areaConfiguration-v1800: ref
	{
		if ie.AreaConfiguration_v1800 != nil {
			if err := ie.AreaConfiguration_v1800.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LoggedMeasurementConfiguration_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(loggedMeasurementConfigurationV1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. areaConfiguration-v1800: ref
	{
		if seq.IsComponentPresent(0) {
			ie.AreaConfiguration_v1800 = new(AreaConfiguration_v1800)
			if err := ie.AreaConfiguration_v1800.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(LoggedMeasurementConfiguration_v1900_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
