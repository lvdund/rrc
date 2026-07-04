// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EventTriggerConfigNR-SL-r16 (line 13937).

var eventTriggerConfigNRSLR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "eventId-r16"},
		{Name: "reportInterval-r16"},
		{Name: "reportAmount-r16"},
		{Name: "reportQuantity-r16"},
	},
}

var eventTriggerConfigNR_SL_r16EventIdR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eventC1"},
		{Name: "eventC2-r16"},
	},
}

const (
	EventTriggerConfigNR_SL_r16_EventId_r16_EventC1     = 0
	EventTriggerConfigNR_SL_r16_EventId_r16_EventC2_r16 = 1
)

const (
	EventTriggerConfigNR_SL_r16_ReportAmount_r16_r1       = 0
	EventTriggerConfigNR_SL_r16_ReportAmount_r16_r2       = 1
	EventTriggerConfigNR_SL_r16_ReportAmount_r16_r4       = 2
	EventTriggerConfigNR_SL_r16_ReportAmount_r16_r8       = 3
	EventTriggerConfigNR_SL_r16_ReportAmount_r16_r16      = 4
	EventTriggerConfigNR_SL_r16_ReportAmount_r16_r32      = 5
	EventTriggerConfigNR_SL_r16_ReportAmount_r16_r64      = 6
	EventTriggerConfigNR_SL_r16_ReportAmount_r16_Infinity = 7
)

var eventTriggerConfigNRSLR16ReportAmountR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EventTriggerConfigNR_SL_r16_ReportAmount_r16_r1, EventTriggerConfigNR_SL_r16_ReportAmount_r16_r2, EventTriggerConfigNR_SL_r16_ReportAmount_r16_r4, EventTriggerConfigNR_SL_r16_ReportAmount_r16_r8, EventTriggerConfigNR_SL_r16_ReportAmount_r16_r16, EventTriggerConfigNR_SL_r16_ReportAmount_r16_r32, EventTriggerConfigNR_SL_r16_ReportAmount_r16_r64, EventTriggerConfigNR_SL_r16_ReportAmount_r16_Infinity},
}

type EventTriggerConfigNR_SL_r16 struct {
	EventId_r16 struct {
		Choice  int
		EventC1 *struct {
			C1_Threshold_r16  SL_CBR_r16
			Hysteresis_r16    Hysteresis
			TimeToTrigger_r16 TimeToTrigger
		}
		EventC2_r16 *struct {
			C2_Threshold_r16  SL_CBR_r16
			Hysteresis_r16    Hysteresis
			TimeToTrigger_r16 TimeToTrigger
		}
	}
	ReportInterval_r16 ReportInterval
	ReportAmount_r16   int64
	ReportQuantity_r16 MeasReportQuantity_r16
}

func (ie *EventTriggerConfigNR_SL_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eventTriggerConfigNRSLR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. eventId-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(eventTriggerConfigNR_SL_r16EventIdR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.EventId_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.EventId_r16.Choice {
		case EventTriggerConfigNR_SL_r16_EventId_r16_EventC1:
			c := (*ie.EventId_r16.EventC1)
			if err := c.C1_Threshold_r16.Encode(e); err != nil {
				return err
			}
			if err := c.Hysteresis_r16.Encode(e); err != nil {
				return err
			}
			if err := c.TimeToTrigger_r16.Encode(e); err != nil {
				return err
			}
		case EventTriggerConfigNR_SL_r16_EventId_r16_EventC2_r16:
			c := (*ie.EventId_r16.EventC2_r16)
			if err := c.C2_Threshold_r16.Encode(e); err != nil {
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
		if err := e.EncodeEnumerated(ie.ReportAmount_r16, eventTriggerConfigNRSLR16ReportAmountR16Constraints); err != nil {
			return err
		}
	}

	// 5. reportQuantity-r16: ref
	{
		if err := ie.ReportQuantity_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *EventTriggerConfigNR_SL_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eventTriggerConfigNRSLR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. eventId-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(eventTriggerConfigNR_SL_r16EventIdR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.EventId_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case EventTriggerConfigNR_SL_r16_EventId_r16_EventC1:
			ie.EventId_r16.EventC1 = &struct {
				C1_Threshold_r16  SL_CBR_r16
				Hysteresis_r16    Hysteresis
				TimeToTrigger_r16 TimeToTrigger
			}{}
			c := (*ie.EventId_r16.EventC1)
			{
				if err := c.C1_Threshold_r16.Decode(d); err != nil {
					return err
				}
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
		case EventTriggerConfigNR_SL_r16_EventId_r16_EventC2_r16:
			ie.EventId_r16.EventC2_r16 = &struct {
				C2_Threshold_r16  SL_CBR_r16
				Hysteresis_r16    Hysteresis
				TimeToTrigger_r16 TimeToTrigger
			}{}
			c := (*ie.EventId_r16.EventC2_r16)
			{
				if err := c.C2_Threshold_r16.Decode(d); err != nil {
					return err
				}
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
		v2, err := d.DecodeEnumerated(eventTriggerConfigNRSLR16ReportAmountR16Constraints)
		if err != nil {
			return err
		}
		ie.ReportAmount_r16 = v2
	}

	// 5. reportQuantity-r16: ref
	{
		if err := ie.ReportQuantity_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
