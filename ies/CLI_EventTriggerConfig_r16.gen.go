// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CLI-EventTriggerConfig-r16 (line 13875).

var cLIEventTriggerConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "eventId-r16"},
		{Name: "reportInterval-r16"},
		{Name: "reportAmount-r16"},
		{Name: "maxReportCLI-r16"},
	},
}

var cLI_EventTriggerConfig_r16EventIdR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eventI1-r16"},
	},
}

const (
	CLI_EventTriggerConfig_r16_EventId_r16_EventI1_r16 = 0
)

const (
	CLI_EventTriggerConfig_r16_ReportAmount_r16_r1       = 0
	CLI_EventTriggerConfig_r16_ReportAmount_r16_r2       = 1
	CLI_EventTriggerConfig_r16_ReportAmount_r16_r4       = 2
	CLI_EventTriggerConfig_r16_ReportAmount_r16_r8       = 3
	CLI_EventTriggerConfig_r16_ReportAmount_r16_r16      = 4
	CLI_EventTriggerConfig_r16_ReportAmount_r16_r32      = 5
	CLI_EventTriggerConfig_r16_ReportAmount_r16_r64      = 6
	CLI_EventTriggerConfig_r16_ReportAmount_r16_Infinity = 7
)

var cLIEventTriggerConfigR16ReportAmountR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CLI_EventTriggerConfig_r16_ReportAmount_r16_r1, CLI_EventTriggerConfig_r16_ReportAmount_r16_r2, CLI_EventTriggerConfig_r16_ReportAmount_r16_r4, CLI_EventTriggerConfig_r16_ReportAmount_r16_r8, CLI_EventTriggerConfig_r16_ReportAmount_r16_r16, CLI_EventTriggerConfig_r16_ReportAmount_r16_r32, CLI_EventTriggerConfig_r16_ReportAmount_r16_r64, CLI_EventTriggerConfig_r16_ReportAmount_r16_Infinity},
}

var cLIEventTriggerConfigR16MaxReportCLIR16Constraints = per.Constrained(1, common.MaxCLI_Report_r16)

type CLI_EventTriggerConfig_r16 struct {
	EventId_r16 struct {
		Choice      int
		EventI1_r16 *struct {
			I1_Threshold_r16  MeasTriggerQuantityCLI_r16
			ReportOnLeave_r16 bool
			Hysteresis_r16    Hysteresis
			TimeToTrigger_r16 TimeToTrigger
		}
	}
	ReportInterval_r16 ReportInterval
	ReportAmount_r16   int64
	MaxReportCLI_r16   int64
}

func (ie *CLI_EventTriggerConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cLIEventTriggerConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. eventId-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(cLI_EventTriggerConfig_r16EventIdR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.EventId_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.EventId_r16.Choice {
		case CLI_EventTriggerConfig_r16_EventId_r16_EventI1_r16:
			c := (*ie.EventId_r16.EventI1_r16)
			if err := c.I1_Threshold_r16.Encode(e); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.ReportOnLeave_r16); err != nil {
				return err
			}
			if err := c.Hysteresis_r16.Encode(e); err != nil {
				return err
			}
			if err := c.TimeToTrigger_r16.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.EventId_r16.Choice), Constraint: "index out of range"}
		}
	}

	// 3. reportInterval-r16: ref
	{
		if err := ie.ReportInterval_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. reportAmount-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.ReportAmount_r16, cLIEventTriggerConfigR16ReportAmountR16Constraints); err != nil {
			return err
		}
	}

	// 5. maxReportCLI-r16: integer
	{
		if err := e.EncodeInteger(ie.MaxReportCLI_r16, cLIEventTriggerConfigR16MaxReportCLIR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CLI_EventTriggerConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cLIEventTriggerConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. eventId-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(cLI_EventTriggerConfig_r16EventIdR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.EventId_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CLI_EventTriggerConfig_r16_EventId_r16_EventI1_r16:
			ie.EventId_r16.EventI1_r16 = &struct {
				I1_Threshold_r16  MeasTriggerQuantityCLI_r16
				ReportOnLeave_r16 bool
				Hysteresis_r16    Hysteresis
				TimeToTrigger_r16 TimeToTrigger
			}{}
			c := (*ie.EventId_r16.EventI1_r16)
			{
				if err := c.I1_Threshold_r16.Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.ReportOnLeave_r16 = v
			}
			{
				if err := c.Hysteresis_r16.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.TimeToTrigger_r16.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. reportInterval-r16: ref
	{
		if err := ie.ReportInterval_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. reportAmount-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(cLIEventTriggerConfigR16ReportAmountR16Constraints)
		if err != nil {
			return err
		}
		ie.ReportAmount_r16 = v2
	}

	// 5. maxReportCLI-r16: integer
	{
		v3, err := d.DecodeInteger(cLIEventTriggerConfigR16MaxReportCLIR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxReportCLI_r16 = v3
	}

	return nil
}
