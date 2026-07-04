// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EventType-r16 (line 589).

var eventTypeR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "outOfCoverage"},
		{Name: "eventL1"},
	},
}

const (
	EventType_r16_OutOfCoverage = 0
	EventType_r16_EventL1       = 1
)

type EventType_r16 struct {
	Choice        int
	OutOfCoverage *struct{}
	EventL1       *struct {
		L1_Threshold  MeasTriggerQuantity
		Hysteresis    Hysteresis
		TimeToTrigger TimeToTrigger
	}
}

func (ie *EventType_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(eventTypeR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case EventType_r16_OutOfCoverage:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case EventType_r16_EventL1:
		c := (*ie.EventL1)
		if err := c.L1_Threshold.Encode(e); err != nil {
			return err
		}
		if err := c.Hysteresis.Encode(e); err != nil {
			return err
		}
		if err := c.TimeToTrigger.Encode(e); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "EventType-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *EventType_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(eventTypeR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "EventType-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case EventType_r16_OutOfCoverage:
		ie.OutOfCoverage = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case EventType_r16_EventL1:
		ie.EventL1 = &struct {
			L1_Threshold  MeasTriggerQuantity
			Hysteresis    Hysteresis
			TimeToTrigger TimeToTrigger
		}{}
		c := (*ie.EventL1)
		{
			if err := c.L1_Threshold.Decode(d); err != nil {
				return err
			}
		}
		{
			if err := c.Hysteresis.Decode(d); err != nil {
				return err
			}
		}
		{
			if err := c.TimeToTrigger.Decode(d); err != nil {
				return err
			}
		}
	default:
		return &per.DecodeError{TypeName: "EventType-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
