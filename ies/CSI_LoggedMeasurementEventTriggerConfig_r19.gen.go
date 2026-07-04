// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-LoggedMeasurementEventTriggerConfig-r19 (line 6986).

var cSILoggedMeasurementEventTriggerConfigR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "eventId-r19"},
	},
}

var cSI_LoggedMeasurementEventTriggerConfig_r19EventIdR19Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eventA1-r19"},
		{Name: "eventA2-r19"},
	},
}

const (
	CSI_LoggedMeasurementEventTriggerConfig_r19_EventId_r19_EventA1_r19 = 0
	CSI_LoggedMeasurementEventTriggerConfig_r19_EventId_r19_EventA2_r19 = 1
)

type CSI_LoggedMeasurementEventTriggerConfig_r19 struct {
	EventId_r19 struct {
		Choice      int
		EventA1_r19 *struct {
			A1_Threshold_r19  MeasTriggerQuantity
			Hysteresis_r19    Hysteresis
			TimeToTrigger_r19 TimeToTrigger
		}
		EventA2_r19 *struct {
			A2_Threshold_r19  MeasTriggerQuantity
			Hysteresis_r19    Hysteresis
			TimeToTrigger_r19 TimeToTrigger
		}
	}
}

func (ie *CSI_LoggedMeasurementEventTriggerConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSILoggedMeasurementEventTriggerConfigR19Constraints)
	_ = seq

	// 1. eventId-r19: choice
	{
		choiceEnc := e.NewChoiceEncoder(cSI_LoggedMeasurementEventTriggerConfig_r19EventIdR19Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.EventId_r19.Choice), false, nil); err != nil {
			return err
		}
		switch ie.EventId_r19.Choice {
		case CSI_LoggedMeasurementEventTriggerConfig_r19_EventId_r19_EventA1_r19:
			c := (*ie.EventId_r19.EventA1_r19)
			if err := c.A1_Threshold_r19.Encode(e); err != nil {
				return err
			}
			if err := c.Hysteresis_r19.Encode(e); err != nil {
				return err
			}
			if err := c.TimeToTrigger_r19.Encode(e); err != nil {
				return err
			}
		case CSI_LoggedMeasurementEventTriggerConfig_r19_EventId_r19_EventA2_r19:
			c := (*ie.EventId_r19.EventA2_r19)
			if err := c.A2_Threshold_r19.Encode(e); err != nil {
				return err
			}
			if err := c.Hysteresis_r19.Encode(e); err != nil {
				return err
			}
			if err := c.TimeToTrigger_r19.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.EventId_r19.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *CSI_LoggedMeasurementEventTriggerConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSILoggedMeasurementEventTriggerConfigR19Constraints)
	_ = seq

	// 1. eventId-r19: choice
	{
		choiceDec := d.NewChoiceDecoder(cSI_LoggedMeasurementEventTriggerConfig_r19EventIdR19Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.EventId_r19.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CSI_LoggedMeasurementEventTriggerConfig_r19_EventId_r19_EventA1_r19:
			ie.EventId_r19.EventA1_r19 = &struct {
				A1_Threshold_r19  MeasTriggerQuantity
				Hysteresis_r19    Hysteresis
				TimeToTrigger_r19 TimeToTrigger
			}{}
			c := (*ie.EventId_r19.EventA1_r19)
			{
				if err := c.A1_Threshold_r19.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.Hysteresis_r19.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.TimeToTrigger_r19.Decode(d); err != nil {
					return err
				}
			}
		case CSI_LoggedMeasurementEventTriggerConfig_r19_EventId_r19_EventA2_r19:
			ie.EventId_r19.EventA2_r19 = &struct {
				A2_Threshold_r19  MeasTriggerQuantity
				Hysteresis_r19    Hysteresis
				TimeToTrigger_r19 TimeToTrigger
			}{}
			c := (*ie.EventId_r19.EventA2_r19)
			{
				if err := c.A2_Threshold_r19.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.Hysteresis_r19.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.TimeToTrigger_r19.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
