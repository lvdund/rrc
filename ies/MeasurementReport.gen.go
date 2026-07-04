// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasurementReport (line 729).

var measurementReportConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var measurementReportCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "measurementReport"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	MeasurementReport_CriticalExtensions_MeasurementReport        = 0
	MeasurementReport_CriticalExtensions_CriticalExtensionsFuture = 1
)

type MeasurementReport struct {
	CriticalExtensions struct {
		Choice                   int
		MeasurementReport        *MeasurementReport_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *MeasurementReport) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measurementReportConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(measurementReportCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case MeasurementReport_CriticalExtensions_MeasurementReport:
			if err := ie.CriticalExtensions.MeasurementReport.Encode(e); err != nil {
				return err
			}
		case MeasurementReport_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *MeasurementReport) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measurementReportConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(measurementReportCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MeasurementReport_CriticalExtensions_MeasurementReport:
			ie.CriticalExtensions.MeasurementReport = new(MeasurementReport_IEs)
			if err := ie.CriticalExtensions.MeasurementReport.Decode(d); err != nil {
				return err
			}
		case MeasurementReport_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
