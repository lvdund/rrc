// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReportConfigToAddMod (line 13974).

var reportConfigToAddModConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reportConfigId"},
		{Name: "reportConfig"},
	},
}

var reportConfigToAddModReportConfigConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "reportConfigNR"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "reportConfigInterRAT"},
		{Name: "reportConfigNR-SL-r16"},
	},
}

const (
	ReportConfigToAddMod_ReportConfig_ReportConfigNR = 0
)

type ReportConfigToAddMod struct {
	ReportConfigId ReportConfigId
	ReportConfig   struct {
		Choice         int
		ReportConfigNR *ReportConfigNR
	}
}

func (ie *ReportConfigToAddMod) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(reportConfigToAddModConstraints)
	_ = seq

	// 1. reportConfigId: ref
	{
		if err := ie.ReportConfigId.Encode(e); err != nil {
			return err
		}
	}

	// 2. reportConfig: choice
	{
		choiceEnc := e.NewChoiceEncoder(reportConfigToAddModReportConfigConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReportConfig.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReportConfig.Choice {
		case ReportConfigToAddMod_ReportConfig_ReportConfigNR:
			if err := ie.ReportConfig.ReportConfigNR.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReportConfig.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *ReportConfigToAddMod) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(reportConfigToAddModConstraints)
	_ = seq

	// 1. reportConfigId: ref
	{
		if err := ie.ReportConfigId.Decode(d); err != nil {
			return err
		}
	}

	// 2. reportConfig: choice
	{
		choiceDec := d.NewChoiceDecoder(reportConfigToAddModReportConfigConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReportConfig.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case ReportConfigToAddMod_ReportConfig_ReportConfigNR:
			ie.ReportConfig.ReportConfigNR = new(ReportConfigNR)
			if err := ie.ReportConfig.ReportConfigNR.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
