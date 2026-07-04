// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReportConfigInterRAT (line 13375).

var reportConfigInterRATConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reportType"},
	},
}

var reportConfigInterRATReportTypeConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "periodical"},
		{Name: "eventTriggered"},
		{Name: "reportCGI"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "reportSFTD"},
	},
}

const (
	ReportConfigInterRAT_ReportType_Periodical     = 0
	ReportConfigInterRAT_ReportType_EventTriggered = 1
	ReportConfigInterRAT_ReportType_ReportCGI      = 2
)

type ReportConfigInterRAT struct {
	ReportType struct {
		Choice         int
		Periodical     *PeriodicalReportConfigInterRAT
		EventTriggered *EventTriggerConfigInterRAT
		ReportCGI      *ReportCGI_EUTRA
	}
}

func (ie *ReportConfigInterRAT) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(reportConfigInterRATConstraints)
	_ = seq

	// 1. reportType: choice
	{
		choiceEnc := e.NewChoiceEncoder(reportConfigInterRATReportTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReportType.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReportType.Choice {
		case ReportConfigInterRAT_ReportType_Periodical:
			if err := ie.ReportType.Periodical.Encode(e); err != nil {
				return err
			}
		case ReportConfigInterRAT_ReportType_EventTriggered:
			if err := ie.ReportType.EventTriggered.Encode(e); err != nil {
				return err
			}
		case ReportConfigInterRAT_ReportType_ReportCGI:
			if err := ie.ReportType.ReportCGI.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReportType.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *ReportConfigInterRAT) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(reportConfigInterRATConstraints)
	_ = seq

	// 1. reportType: choice
	{
		choiceDec := d.NewChoiceDecoder(reportConfigInterRATReportTypeConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReportType.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case ReportConfigInterRAT_ReportType_Periodical:
			ie.ReportType.Periodical = new(PeriodicalReportConfigInterRAT)
			if err := ie.ReportType.Periodical.Decode(d); err != nil {
				return err
			}
		case ReportConfigInterRAT_ReportType_EventTriggered:
			ie.ReportType.EventTriggered = new(EventTriggerConfigInterRAT)
			if err := ie.ReportType.EventTriggered.Decode(d); err != nil {
				return err
			}
		case ReportConfigInterRAT_ReportType_ReportCGI:
			ie.ReportType.ReportCGI = new(ReportCGI_EUTRA)
			if err := ie.ReportType.ReportCGI.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
