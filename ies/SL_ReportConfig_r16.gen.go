// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ReportConfig-r16 (line 27882).

var sLReportConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ReportType-r16"},
	},
}

var sL_ReportConfig_r16SlReportTypeR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-Periodical-r16"},
		{Name: "sl-EventTriggered-r16"},
	},
}

const (
	SL_ReportConfig_r16_Sl_ReportType_r16_Sl_Periodical_r16     = 0
	SL_ReportConfig_r16_Sl_ReportType_r16_Sl_EventTriggered_r16 = 1
)

type SL_ReportConfig_r16 struct {
	Sl_ReportType_r16 struct {
		Choice                int
		Sl_Periodical_r16     *SL_PeriodicalReportConfig_r16
		Sl_EventTriggered_r16 *SL_EventTriggerConfig_r16
	}
}

func (ie *SL_ReportConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLReportConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sl-ReportType-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(sL_ReportConfig_r16SlReportTypeR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Sl_ReportType_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Sl_ReportType_r16.Choice {
		case SL_ReportConfig_r16_Sl_ReportType_r16_Sl_Periodical_r16:
			if err := ie.Sl_ReportType_r16.Sl_Periodical_r16.Encode(e); err != nil {
				return err
			}
		case SL_ReportConfig_r16_Sl_ReportType_r16_Sl_EventTriggered_r16:
			if err := ie.Sl_ReportType_r16.Sl_EventTriggered_r16.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Sl_ReportType_r16.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SL_ReportConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLReportConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sl-ReportType-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(sL_ReportConfig_r16SlReportTypeR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Sl_ReportType_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SL_ReportConfig_r16_Sl_ReportType_r16_Sl_Periodical_r16:
			ie.Sl_ReportType_r16.Sl_Periodical_r16 = new(SL_PeriodicalReportConfig_r16)
			if err := ie.Sl_ReportType_r16.Sl_Periodical_r16.Decode(d); err != nil {
				return err
			}
		case SL_ReportConfig_r16_Sl_ReportType_r16_Sl_EventTriggered_r16:
			ie.Sl_ReportType_r16.Sl_EventTriggered_r16 = new(SL_EventTriggerConfig_r16)
			if err := ie.Sl_ReportType_r16.Sl_EventTriggered_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
