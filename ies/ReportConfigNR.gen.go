// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReportConfigNR (line 13528).

var reportConfigNRConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reportType"},
	},
}

var reportConfigNRReportTypeConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "periodical"},
		{Name: "eventTriggered"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "reportCGI"},
		{Name: "reportSFTD"},
		{Name: "condTriggerConfig-r16"},
		{Name: "cli-Periodical-r16"},
		{Name: "cli-EventTriggered-r16"},
		{Name: "rxTxPeriodical-r17"},
		{Name: "reportOnScellActivation-r18"},
	},
}

const (
	ReportConfigNR_ReportType_Periodical     = 0
	ReportConfigNR_ReportType_EventTriggered = 1
)

type ReportConfigNR struct {
	ReportType struct {
		Choice         int
		Periodical     *PeriodicalReportConfig
		EventTriggered *EventTriggerConfig
	}
}

func (ie *ReportConfigNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(reportConfigNRConstraints)
	_ = seq

	// 1. reportType: choice
	{
		choiceEnc := e.NewChoiceEncoder(reportConfigNRReportTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReportType.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReportType.Choice {
		case ReportConfigNR_ReportType_Periodical:
			if err := ie.ReportType.Periodical.Encode(e); err != nil {
				return err
			}
		case ReportConfigNR_ReportType_EventTriggered:
			if err := ie.ReportType.EventTriggered.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReportType.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *ReportConfigNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(reportConfigNRConstraints)
	_ = seq

	// 1. reportType: choice
	{
		choiceDec := d.NewChoiceDecoder(reportConfigNRReportTypeConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReportType.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case ReportConfigNR_ReportType_Periodical:
			ie.ReportType.Periodical = new(PeriodicalReportConfig)
			if err := ie.ReportType.Periodical.Decode(d); err != nil {
				return err
			}
		case ReportConfigNR_ReportType_EventTriggered:
			ie.ReportType.EventTriggered = new(EventTriggerConfig)
			if err := ie.ReportType.EventTriggered.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
