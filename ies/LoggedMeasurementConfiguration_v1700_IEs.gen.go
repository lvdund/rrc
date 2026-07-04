// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LoggedMeasurementConfiguration-v1700-IEs (line 561).

var loggedMeasurementConfigurationV1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sigLoggedMeasType-r17", Optional: true},
		{Name: "earlyMeasIndication-r17", Optional: true},
		{Name: "areaConfiguration-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	LoggedMeasurementConfiguration_v1700_IEs_SigLoggedMeasType_r17_True = 0
)

var loggedMeasurementConfigurationV1700IEsSigLoggedMeasTypeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LoggedMeasurementConfiguration_v1700_IEs_SigLoggedMeasType_r17_True},
}

const (
	LoggedMeasurementConfiguration_v1700_IEs_EarlyMeasIndication_r17_True = 0
)

var loggedMeasurementConfigurationV1700IEsEarlyMeasIndicationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LoggedMeasurementConfiguration_v1700_IEs_EarlyMeasIndication_r17_True},
}

type LoggedMeasurementConfiguration_v1700_IEs struct {
	SigLoggedMeasType_r17   *int64
	EarlyMeasIndication_r17 *int64
	AreaConfiguration_r17   *AreaConfiguration_r17
	NonCriticalExtension    *LoggedMeasurementConfiguration_v1800_IEs
}

func (ie *LoggedMeasurementConfiguration_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(loggedMeasurementConfigurationV1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SigLoggedMeasType_r17 != nil, ie.EarlyMeasIndication_r17 != nil, ie.AreaConfiguration_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. sigLoggedMeasType-r17: enumerated
	{
		if ie.SigLoggedMeasType_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SigLoggedMeasType_r17, loggedMeasurementConfigurationV1700IEsSigLoggedMeasTypeR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. earlyMeasIndication-r17: enumerated
	{
		if ie.EarlyMeasIndication_r17 != nil {
			if err := e.EncodeEnumerated(*ie.EarlyMeasIndication_r17, loggedMeasurementConfigurationV1700IEsEarlyMeasIndicationR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. areaConfiguration-r17: ref
	{
		if ie.AreaConfiguration_r17 != nil {
			if err := ie.AreaConfiguration_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LoggedMeasurementConfiguration_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(loggedMeasurementConfigurationV1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sigLoggedMeasType-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(loggedMeasurementConfigurationV1700IEsSigLoggedMeasTypeR17Constraints)
			if err != nil {
				return err
			}
			ie.SigLoggedMeasType_r17 = &idx
		}
	}

	// 3. earlyMeasIndication-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(loggedMeasurementConfigurationV1700IEsEarlyMeasIndicationR17Constraints)
			if err != nil {
				return err
			}
			ie.EarlyMeasIndication_r17 = &idx
		}
	}

	// 4. areaConfiguration-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.AreaConfiguration_r17 = new(AreaConfiguration_r17)
			if err := ie.AreaConfiguration_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(3) {
			ie.NonCriticalExtension = new(LoggedMeasurementConfiguration_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
