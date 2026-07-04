// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CondTriggerConfig-r16 (line 13562).

var condTriggerConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "condEventId"},
		{Name: "rsType-r16"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var condTriggerConfig_r16CondEventIdConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "condEventA3"},
		{Name: "condEventA5"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "condEventA4-r17"},
		{Name: "condEventD1-r17"},
		{Name: "condEventT1-r17"},
		{Name: "condEventD2-r18"},
		{Name: "condEventA3H1-r19"},
		{Name: "condEventA3H2-r19"},
		{Name: "condEventA5H1-r19"},
		{Name: "condEventA5H2-r19"},
	},
}

const (
	CondTriggerConfig_r16_CondEventId_CondEventA3 = 0
	CondTriggerConfig_r16_CondEventId_CondEventA5 = 1
)

const (
	CondTriggerConfig_r16_Ext_NesEvent_r18_True = 0
)

var condTriggerConfigR16ExtNesEventR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CondTriggerConfig_r16_Ext_NesEvent_r18_True},
}

type CondTriggerConfig_r16 struct {
	CondEventId struct {
		Choice      int
		CondEventA3 *struct {
			A3_Offset     MeasTriggerQuantityOffset
			Hysteresis    Hysteresis
			TimeToTrigger TimeToTrigger
		}
		CondEventA5 *struct {
			A5_Threshold1 MeasTriggerQuantity
			A5_Threshold2 MeasTriggerQuantity
			Hysteresis    Hysteresis
			TimeToTrigger TimeToTrigger
		}
	}
	RsType_r16   NR_RS_Type
	NesEvent_r18 *int64
}

func (ie *CondTriggerConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(condTriggerConfigR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.NesEvent_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. condEventId: choice
	{
		choiceEnc := e.NewChoiceEncoder(condTriggerConfig_r16CondEventIdConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CondEventId.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CondEventId.Choice {
		case CondTriggerConfig_r16_CondEventId_CondEventA3:
			c := (*ie.CondEventId.CondEventA3)
			if err := c.A3_Offset.Encode(e); err != nil {
				return err
			}
			if err := c.Hysteresis.Encode(e); err != nil {
				return err
			}
			if err := c.TimeToTrigger.Encode(e); err != nil {
				return err
			}
		case CondTriggerConfig_r16_CondEventId_CondEventA5:
			c := (*ie.CondEventId.CondEventA5)
			if err := c.A5_Threshold1.Encode(e); err != nil {
				return err
			}
			if err := c.A5_Threshold2.Encode(e); err != nil {
				return err
			}
			if err := c.Hysteresis.Encode(e); err != nil {
				return err
			}
			if err := c.TimeToTrigger.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CondEventId.Choice), Constraint: "index out of range"}
		}
	}

	// 3. rsType-r16: ref
	{
		if err := ie.RsType_r16.Encode(e); err != nil {
			return err
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "nesEvent-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.NesEvent_r18 != nil}); err != nil {
				return err
			}

			if ie.NesEvent_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.NesEvent_r18, condTriggerConfigR16ExtNesEventR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CondTriggerConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(condTriggerConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. condEventId: choice
	{
		choiceDec := d.NewChoiceDecoder(condTriggerConfig_r16CondEventIdConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CondEventId.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CondTriggerConfig_r16_CondEventId_CondEventA3:
			ie.CondEventId.CondEventA3 = &struct {
				A3_Offset     MeasTriggerQuantityOffset
				Hysteresis    Hysteresis
				TimeToTrigger TimeToTrigger
			}{}
			c := (*ie.CondEventId.CondEventA3)
			{
				if err := c.A3_Offset.Decode(d); err != nil {
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
		case CondTriggerConfig_r16_CondEventId_CondEventA5:
			ie.CondEventId.CondEventA5 = &struct {
				A5_Threshold1 MeasTriggerQuantity
				A5_Threshold2 MeasTriggerQuantity
				Hysteresis    Hysteresis
				TimeToTrigger TimeToTrigger
			}{}
			c := (*ie.CondEventId.CondEventA5)
			{
				if err := c.A5_Threshold1.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.A5_Threshold2.Decode(d); err != nil {
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
		}
	}

	// 3. rsType-r16: ref
	{
		if err := ie.RsType_r16.Decode(d); err != nil {
			return err
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "nesEvent-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(condTriggerConfigR16ExtNesEventR18Constraints)
			if err != nil {
				return err
			}
			ie.NesEvent_r18 = &v
		}
	}

	return nil
}
