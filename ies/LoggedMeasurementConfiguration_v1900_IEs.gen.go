// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LoggedMeasurementConfiguration-v1900-IEs (line 573).

var loggedMeasurementConfigurationV1900IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "areaConfigurationNTN-List-r19", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type LoggedMeasurementConfiguration_v1900_IEs struct {
	AreaConfigurationNTN_List_r19 *AreaConfigurationNTN_List_r19
	NonCriticalExtension          *struct{}
}

func (ie *LoggedMeasurementConfiguration_v1900_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(loggedMeasurementConfigurationV1900IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AreaConfigurationNTN_List_r19 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. areaConfigurationNTN-List-r19: ref
	{
		if ie.AreaConfigurationNTN_List_r19 != nil {
			if err := ie.AreaConfigurationNTN_List_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *LoggedMeasurementConfiguration_v1900_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(loggedMeasurementConfigurationV1900IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. areaConfigurationNTN-List-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.AreaConfigurationNTN_List_r19 = new(AreaConfigurationNTN_List_r19)
			if err := ie.AreaConfigurationNTN_List_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
