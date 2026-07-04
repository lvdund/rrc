// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReportConfigNR-SL-r16 (line 13930).

var reportConfigNRSLR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reportType-r16"},
	},
}

var reportConfigNR_SL_r16ReportTypeR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "periodical-r16"},
		{Name: "eventTriggered-r16"},
	},
}

const (
	ReportConfigNR_SL_r16_ReportType_r16_Periodical_r16     = 0
	ReportConfigNR_SL_r16_ReportType_r16_EventTriggered_r16 = 1
)

type ReportConfigNR_SL_r16 struct {
	ReportType_r16 struct {
		Choice             int
		Periodical_r16     *PeriodicalReportConfigNR_SL_r16
		EventTriggered_r16 *EventTriggerConfigNR_SL_r16
	}
}

func (ie *ReportConfigNR_SL_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(reportConfigNRSLR16Constraints)
	_ = seq

	// 1. reportType-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(reportConfigNR_SL_r16ReportTypeR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReportType_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReportType_r16.Choice {
		case ReportConfigNR_SL_r16_ReportType_r16_Periodical_r16:
			if err := ie.ReportType_r16.Periodical_r16.Encode(e); err != nil {
				return err
			}
		case ReportConfigNR_SL_r16_ReportType_r16_EventTriggered_r16:
			if err := ie.ReportType_r16.EventTriggered_r16.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReportType_r16.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *ReportConfigNR_SL_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(reportConfigNRSLR16Constraints)
	_ = seq

	// 1. reportType-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(reportConfigNR_SL_r16ReportTypeR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReportType_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case ReportConfigNR_SL_r16_ReportType_r16_Periodical_r16:
			ie.ReportType_r16.Periodical_r16 = new(PeriodicalReportConfigNR_SL_r16)
			if err := ie.ReportType_r16.Periodical_r16.Decode(d); err != nil {
				return err
			}
		case ReportConfigNR_SL_r16_ReportType_r16_EventTriggered_r16:
			ie.ReportType_r16.EventTriggered_r16 = new(EventTriggerConfigNR_SL_r16)
			if err := ie.ReportType_r16.EventTriggered_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
