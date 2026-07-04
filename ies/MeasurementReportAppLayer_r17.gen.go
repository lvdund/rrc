// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasurementReportAppLayer-r17 (line 746).

var measurementReportAppLayerR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var measurementReportAppLayer_r17CriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "measurementReportAppLayer-r17"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	MeasurementReportAppLayer_r17_CriticalExtensions_MeasurementReportAppLayer_r17 = 0
	MeasurementReportAppLayer_r17_CriticalExtensions_CriticalExtensionsFuture      = 1
)

type MeasurementReportAppLayer_r17 struct {
	CriticalExtensions struct {
		Choice                        int
		MeasurementReportAppLayer_r17 *MeasurementReportAppLayer_r17_IEs
		CriticalExtensionsFuture      *struct{}
	}
}

func (ie *MeasurementReportAppLayer_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measurementReportAppLayerR17Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(measurementReportAppLayer_r17CriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case MeasurementReportAppLayer_r17_CriticalExtensions_MeasurementReportAppLayer_r17:
			if err := ie.CriticalExtensions.MeasurementReportAppLayer_r17.Encode(e); err != nil {
				return err
			}
		case MeasurementReportAppLayer_r17_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *MeasurementReportAppLayer_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measurementReportAppLayerR17Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(measurementReportAppLayer_r17CriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MeasurementReportAppLayer_r17_CriticalExtensions_MeasurementReportAppLayer_r17:
			ie.CriticalExtensions.MeasurementReportAppLayer_r17 = new(MeasurementReportAppLayer_r17_IEs)
			if err := ie.CriticalExtensions.MeasurementReportAppLayer_r17.Decode(d); err != nil {
				return err
			}
		case MeasurementReportAppLayer_r17_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
