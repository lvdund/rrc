// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PQI-r16 (line 27757).

var sLPQIR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-StandardizedPQI-r16"},
		{Name: "sl-Non-StandardizedPQI-r16"},
	},
}

const (
	SL_PQI_r16_Sl_StandardizedPQI_r16     = 0
	SL_PQI_r16_Sl_Non_StandardizedPQI_r16 = 1
)

var sLPQIR16SlNonStandardizedPQIR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ResourceType-r16", Optional: true},
		{Name: "sl-PriorityLevel-r16", Optional: true},
		{Name: "sl-PacketDelayBudget-r16", Optional: true},
		{Name: "sl-PacketErrorRate-r16", Optional: true},
		{Name: "sl-AveragingWindow-r16", Optional: true},
		{Name: "sl-MaxDataBurstVolume-r16", Optional: true},
	},
}

const (
	SL_PQI_r16_Sl_Non_StandardizedPQI_r16_Sl_ResourceType_r16_Gbr              = 0
	SL_PQI_r16_Sl_Non_StandardizedPQI_r16_Sl_ResourceType_r16_Non_GBR          = 1
	SL_PQI_r16_Sl_Non_StandardizedPQI_r16_Sl_ResourceType_r16_DelayCriticalGBR = 2
	SL_PQI_r16_Sl_Non_StandardizedPQI_r16_Sl_ResourceType_r16_Spare1           = 3
)

var sLPQIR16SlNonStandardizedPQIR16SlResourceTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PQI_r16_Sl_Non_StandardizedPQI_r16_Sl_ResourceType_r16_Gbr, SL_PQI_r16_Sl_Non_StandardizedPQI_r16_Sl_ResourceType_r16_Non_GBR, SL_PQI_r16_Sl_Non_StandardizedPQI_r16_Sl_ResourceType_r16_DelayCriticalGBR, SL_PQI_r16_Sl_Non_StandardizedPQI_r16_Sl_ResourceType_r16_Spare1},
}

type SL_PQI_r16 struct {
	Choice                     int
	Sl_StandardizedPQI_r16     *int64
	Sl_Non_StandardizedPQI_r16 *struct {
		Sl_ResourceType_r16       *int64
		Sl_PriorityLevel_r16      *int64
		Sl_PacketDelayBudget_r16  *int64
		Sl_PacketErrorRate_r16    *int64
		Sl_AveragingWindow_r16    *int64
		Sl_MaxDataBurstVolume_r16 *int64
	}
}

func (ie *SL_PQI_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sLPQIR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_PQI_r16_Sl_StandardizedPQI_r16:
		if err := e.EncodeInteger((*ie.Sl_StandardizedPQI_r16), per.Constrained(0, 255)); err != nil {
			return err
		}
	case SL_PQI_r16_Sl_Non_StandardizedPQI_r16:
		c := (*ie.Sl_Non_StandardizedPQI_r16)
		sLPQIR16SlNonStandardizedPQIR16Seq := e.NewSequenceEncoder(sLPQIR16SlNonStandardizedPQIR16Constraints)
		if err := sLPQIR16SlNonStandardizedPQIR16Seq.EncodeExtensionBit(false); err != nil {
			return err
		}
		if err := sLPQIR16SlNonStandardizedPQIR16Seq.EncodePreamble([]bool{c.Sl_ResourceType_r16 != nil, c.Sl_PriorityLevel_r16 != nil, c.Sl_PacketDelayBudget_r16 != nil, c.Sl_PacketErrorRate_r16 != nil, c.Sl_AveragingWindow_r16 != nil, c.Sl_MaxDataBurstVolume_r16 != nil}); err != nil {
			return err
		}
		if c.Sl_ResourceType_r16 != nil {
			if err := e.EncodeEnumerated((*c.Sl_ResourceType_r16), sLPQIR16SlNonStandardizedPQIR16SlResourceTypeR16Constraints); err != nil {
				return err
			}
		}
		if c.Sl_PriorityLevel_r16 != nil {
			if err := e.EncodeInteger((*c.Sl_PriorityLevel_r16), per.Constrained(1, 8)); err != nil {
				return err
			}
		}
		if c.Sl_PacketDelayBudget_r16 != nil {
			if err := e.EncodeInteger((*c.Sl_PacketDelayBudget_r16), per.Constrained(0, 1023)); err != nil {
				return err
			}
		}
		if c.Sl_PacketErrorRate_r16 != nil {
			if err := e.EncodeInteger((*c.Sl_PacketErrorRate_r16), per.Constrained(0, 9)); err != nil {
				return err
			}
		}
		if c.Sl_AveragingWindow_r16 != nil {
			if err := e.EncodeInteger((*c.Sl_AveragingWindow_r16), per.Constrained(0, 4095)); err != nil {
				return err
			}
		}
		if c.Sl_MaxDataBurstVolume_r16 != nil {
			if err := e.EncodeInteger((*c.Sl_MaxDataBurstVolume_r16), per.Constrained(0, 4095)); err != nil {
				return err
			}
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SL-PQI-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SL_PQI_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sLPQIR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SL-PQI-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SL_PQI_r16_Sl_StandardizedPQI_r16:
		ie.Sl_StandardizedPQI_r16 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 255))
		if err != nil {
			return err
		}
		(*ie.Sl_StandardizedPQI_r16) = v
	case SL_PQI_r16_Sl_Non_StandardizedPQI_r16:
		ie.Sl_Non_StandardizedPQI_r16 = &struct {
			Sl_ResourceType_r16       *int64
			Sl_PriorityLevel_r16      *int64
			Sl_PacketDelayBudget_r16  *int64
			Sl_PacketErrorRate_r16    *int64
			Sl_AveragingWindow_r16    *int64
			Sl_MaxDataBurstVolume_r16 *int64
		}{}
		c := (*ie.Sl_Non_StandardizedPQI_r16)
		sLPQIR16SlNonStandardizedPQIR16Seq := d.NewSequenceDecoder(sLPQIR16SlNonStandardizedPQIR16Constraints)
		if err := sLPQIR16SlNonStandardizedPQIR16Seq.DecodeExtensionBit(); err != nil {
			return err
		}
		if err := sLPQIR16SlNonStandardizedPQIR16Seq.DecodePreamble(); err != nil {
			return err
		}
		if sLPQIR16SlNonStandardizedPQIR16Seq.IsComponentPresent(0) {
			c.Sl_ResourceType_r16 = new(int64)
			v, err := d.DecodeEnumerated(sLPQIR16SlNonStandardizedPQIR16SlResourceTypeR16Constraints)
			if err != nil {
				return err
			}
			(*c.Sl_ResourceType_r16) = v
		}
		if sLPQIR16SlNonStandardizedPQIR16Seq.IsComponentPresent(1) {
			c.Sl_PriorityLevel_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(1, 8))
			if err != nil {
				return err
			}
			(*c.Sl_PriorityLevel_r16) = v
		}
		if sLPQIR16SlNonStandardizedPQIR16Seq.IsComponentPresent(2) {
			c.Sl_PacketDelayBudget_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1023))
			if err != nil {
				return err
			}
			(*c.Sl_PacketDelayBudget_r16) = v
		}
		if sLPQIR16SlNonStandardizedPQIR16Seq.IsComponentPresent(3) {
			c.Sl_PacketErrorRate_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*c.Sl_PacketErrorRate_r16) = v
		}
		if sLPQIR16SlNonStandardizedPQIR16Seq.IsComponentPresent(4) {
			c.Sl_AveragingWindow_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 4095))
			if err != nil {
				return err
			}
			(*c.Sl_AveragingWindow_r16) = v
		}
		if sLPQIR16SlNonStandardizedPQIR16Seq.IsComponentPresent(5) {
			c.Sl_MaxDataBurstVolume_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 4095))
			if err != nil {
				return err
			}
			(*c.Sl_MaxDataBurstVolume_r16) = v
		}
	default:
		return &per.DecodeError{TypeName: "SL-PQI-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
