// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LocationMeasurementIndication (line 518).

var locationMeasurementIndicationConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var locationMeasurementIndicationCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "locationMeasurementIndication"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	LocationMeasurementIndication_CriticalExtensions_LocationMeasurementIndication = 0
	LocationMeasurementIndication_CriticalExtensions_CriticalExtensionsFuture      = 1
)

type LocationMeasurementIndication struct {
	CriticalExtensions struct {
		Choice                        int
		LocationMeasurementIndication *LocationMeasurementIndication_IEs
		CriticalExtensionsFuture      *struct{}
	}
}

func (ie *LocationMeasurementIndication) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(locationMeasurementIndicationConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(locationMeasurementIndicationCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case LocationMeasurementIndication_CriticalExtensions_LocationMeasurementIndication:
			if err := ie.CriticalExtensions.LocationMeasurementIndication.Encode(e); err != nil {
				return err
			}
		case LocationMeasurementIndication_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *LocationMeasurementIndication) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(locationMeasurementIndicationConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(locationMeasurementIndicationCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case LocationMeasurementIndication_CriticalExtensions_LocationMeasurementIndication:
			ie.CriticalExtensions.LocationMeasurementIndication = new(LocationMeasurementIndication_IEs)
			if err := ie.CriticalExtensions.LocationMeasurementIndication.Decode(d); err != nil {
				return err
			}
		case LocationMeasurementIndication_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
