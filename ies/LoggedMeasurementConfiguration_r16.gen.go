// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LoggedMeasurementConfiguration-r16 (line 534).

var loggedMeasurementConfigurationR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var loggedMeasurementConfiguration_r16CriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "loggedMeasurementConfiguration-r16"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	LoggedMeasurementConfiguration_r16_CriticalExtensions_LoggedMeasurementConfiguration_r16 = 0
	LoggedMeasurementConfiguration_r16_CriticalExtensions_CriticalExtensionsFuture           = 1
)

type LoggedMeasurementConfiguration_r16 struct {
	CriticalExtensions struct {
		Choice                             int
		LoggedMeasurementConfiguration_r16 *LoggedMeasurementConfiguration_r16_IEs
		CriticalExtensionsFuture           *struct{}
	}
}

func (ie *LoggedMeasurementConfiguration_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(loggedMeasurementConfigurationR16Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(loggedMeasurementConfiguration_r16CriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case LoggedMeasurementConfiguration_r16_CriticalExtensions_LoggedMeasurementConfiguration_r16:
			if err := ie.CriticalExtensions.LoggedMeasurementConfiguration_r16.Encode(e); err != nil {
				return err
			}
		case LoggedMeasurementConfiguration_r16_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *LoggedMeasurementConfiguration_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(loggedMeasurementConfigurationR16Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(loggedMeasurementConfiguration_r16CriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case LoggedMeasurementConfiguration_r16_CriticalExtensions_LoggedMeasurementConfiguration_r16:
			ie.CriticalExtensions.LoggedMeasurementConfiguration_r16 = new(LoggedMeasurementConfiguration_r16_IEs)
			if err := ie.CriticalExtensions.LoggedMeasurementConfiguration_r16.Decode(d); err != nil {
				return err
			}
		case LoggedMeasurementConfiguration_r16_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
