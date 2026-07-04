// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-EventTriggerConfig-r16 (line 27899).

var sLEventTriggerConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-EventId-r16"},
		{Name: "sl-ReportInterval-r16"},
		{Name: "sl-ReportAmount-r16"},
		{Name: "sl-ReportQuantity-r16"},
		{Name: "sl-RS-Type-r16"},
	},
}

var sL_EventTriggerConfig_r16SlEventIdR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eventS1-r16"},
		{Name: "eventS2-r16"},
	},
}

const (
	SL_EventTriggerConfig_r16_Sl_EventId_r16_EventS1_r16 = 0
	SL_EventTriggerConfig_r16_Sl_EventId_r16_EventS2_r16 = 1
)

const (
	SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_r1       = 0
	SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_r2       = 1
	SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_r4       = 2
	SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_r8       = 3
	SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_r16      = 4
	SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_r32      = 5
	SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_r64      = 6
	SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_Infinity = 7
)

var sLEventTriggerConfigR16SlReportAmountR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_r1, SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_r2, SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_r4, SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_r8, SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_r16, SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_r32, SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_r64, SL_EventTriggerConfig_r16_Sl_ReportAmount_r16_Infinity},
}

var sLEventTriggerConfigR16SlEventIdR16EventS1R16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "s1-Threshold-r16"},
		{Name: "sl-ReportOnLeave-r16"},
		{Name: "sl-Hysteresis-r16"},
		{Name: "sl-TimeToTrigger-r16"},
	},
}

var sLEventTriggerConfigR16SlEventIdR16EventS2R16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "s2-Threshold-r16"},
		{Name: "sl-ReportOnLeave-r16"},
		{Name: "sl-Hysteresis-r16"},
		{Name: "sl-TimeToTrigger-r16"},
	},
}

type SL_EventTriggerConfig_r16 struct {
	Sl_EventId_r16 struct {
		Choice      int
		EventS1_r16 *struct {
			S1_Threshold_r16     SL_MeasTriggerQuantity_r16
			Sl_ReportOnLeave_r16 bool
			Sl_Hysteresis_r16    Hysteresis
			Sl_TimeToTrigger_r16 TimeToTrigger
		}
		EventS2_r16 *struct {
			S2_Threshold_r16     SL_MeasTriggerQuantity_r16
			Sl_ReportOnLeave_r16 bool
			Sl_Hysteresis_r16    Hysteresis
			Sl_TimeToTrigger_r16 TimeToTrigger
		}
	}
	Sl_ReportInterval_r16 ReportInterval
	Sl_ReportAmount_r16   int64
	Sl_ReportQuantity_r16 SL_MeasReportQuantity_r16
	Sl_RS_Type_r16        SL_RS_Type_r16
}

func (ie *SL_EventTriggerConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLEventTriggerConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sl-EventId-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(sL_EventTriggerConfig_r16SlEventIdR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Sl_EventId_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Sl_EventId_r16.Choice {
		case SL_EventTriggerConfig_r16_Sl_EventId_r16_EventS1_r16:
			c := (*ie.Sl_EventId_r16.EventS1_r16)
			sLEventTriggerConfigR16SlEventIdR16EventS1R16Seq := e.NewSequenceEncoder(sLEventTriggerConfigR16SlEventIdR16EventS1R16Constraints)
			if err := sLEventTriggerConfigR16SlEventIdR16EventS1R16Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := c.S1_Threshold_r16.Encode(e); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.Sl_ReportOnLeave_r16); err != nil {
				return err
			}
			if err := c.Sl_Hysteresis_r16.Encode(e); err != nil {
				return err
			}
			if err := c.Sl_TimeToTrigger_r16.Encode(e); err != nil {
				return err
			}
		case SL_EventTriggerConfig_r16_Sl_EventId_r16_EventS2_r16:
			c := (*ie.Sl_EventId_r16.EventS2_r16)
			sLEventTriggerConfigR16SlEventIdR16EventS2R16Seq := e.NewSequenceEncoder(sLEventTriggerConfigR16SlEventIdR16EventS2R16Constraints)
			if err := sLEventTriggerConfigR16SlEventIdR16EventS2R16Seq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := c.S2_Threshold_r16.Encode(e); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.Sl_ReportOnLeave_r16); err != nil {
				return err
			}
			if err := c.Sl_Hysteresis_r16.Encode(e); err != nil {
				return err
			}
			if err := c.Sl_TimeToTrigger_r16.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Sl_EventId_r16.Choice), Constraint: "index out of range"}
		}
	}

	// 3. sl-ReportInterval-r16: ref
	{
		if err := ie.Sl_ReportInterval_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-ReportAmount-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_ReportAmount_r16, sLEventTriggerConfigR16SlReportAmountR16Constraints); err != nil {
			return err
		}
	}

	// 5. sl-ReportQuantity-r16: ref
	{
		if err := ie.Sl_ReportQuantity_r16.Encode(e); err != nil {
			return err
		}
	}

	// 6. sl-RS-Type-r16: ref
	{
		if err := ie.Sl_RS_Type_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_EventTriggerConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLEventTriggerConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sl-EventId-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(sL_EventTriggerConfig_r16SlEventIdR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Sl_EventId_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SL_EventTriggerConfig_r16_Sl_EventId_r16_EventS1_r16:
			ie.Sl_EventId_r16.EventS1_r16 = &struct {
				S1_Threshold_r16     SL_MeasTriggerQuantity_r16
				Sl_ReportOnLeave_r16 bool
				Sl_Hysteresis_r16    Hysteresis
				Sl_TimeToTrigger_r16 TimeToTrigger
			}{}
			c := (*ie.Sl_EventId_r16.EventS1_r16)
			sLEventTriggerConfigR16SlEventIdR16EventS1R16Seq := d.NewSequenceDecoder(sLEventTriggerConfigR16SlEventIdR16EventS1R16Constraints)
			if err := sLEventTriggerConfigR16SlEventIdR16EventS1R16Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				if err := c.S1_Threshold_r16.Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.Sl_ReportOnLeave_r16 = v
			}
			{
				if err := c.Sl_Hysteresis_r16.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.Sl_TimeToTrigger_r16.Decode(d); err != nil {
					return err
				}
			}
		case SL_EventTriggerConfig_r16_Sl_EventId_r16_EventS2_r16:
			ie.Sl_EventId_r16.EventS2_r16 = &struct {
				S2_Threshold_r16     SL_MeasTriggerQuantity_r16
				Sl_ReportOnLeave_r16 bool
				Sl_Hysteresis_r16    Hysteresis
				Sl_TimeToTrigger_r16 TimeToTrigger
			}{}
			c := (*ie.Sl_EventId_r16.EventS2_r16)
			sLEventTriggerConfigR16SlEventIdR16EventS2R16Seq := d.NewSequenceDecoder(sLEventTriggerConfigR16SlEventIdR16EventS2R16Constraints)
			if err := sLEventTriggerConfigR16SlEventIdR16EventS2R16Seq.DecodeExtensionBit(); err != nil {
				return err
			}
			{
				if err := c.S2_Threshold_r16.Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.Sl_ReportOnLeave_r16 = v
			}
			{
				if err := c.Sl_Hysteresis_r16.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.Sl_TimeToTrigger_r16.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-ReportInterval-r16: ref
	{
		if err := ie.Sl_ReportInterval_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-ReportAmount-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(sLEventTriggerConfigR16SlReportAmountR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_ReportAmount_r16 = v2
	}

	// 5. sl-ReportQuantity-r16: ref
	{
		if err := ie.Sl_ReportQuantity_r16.Decode(d); err != nil {
			return err
		}
	}

	// 6. sl-RS-Type-r16: ref
	{
		if err := ie.Sl_RS_Type_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
