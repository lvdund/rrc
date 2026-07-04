// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SchedulingRequestResourceConfigExt-v1700 (line 14309).

var schedulingRequestResourceConfigExtV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicityAndOffset-r17", Optional: true},
	},
}

var schedulingRequestResourceConfigExt_v1700PeriodicityAndOffsetR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl1280"},
		{Name: "sl2560"},
		{Name: "sl5120"},
	},
}

const (
	SchedulingRequestResourceConfigExt_v1700_PeriodicityAndOffset_r17_Sl1280 = 0
	SchedulingRequestResourceConfigExt_v1700_PeriodicityAndOffset_r17_Sl2560 = 1
	SchedulingRequestResourceConfigExt_v1700_PeriodicityAndOffset_r17_Sl5120 = 2
)

type SchedulingRequestResourceConfigExt_v1700 struct {
	PeriodicityAndOffset_r17 *struct {
		Choice int
		Sl1280 *int64
		Sl2560 *int64
		Sl5120 *int64
	}
}

func (ie *SchedulingRequestResourceConfigExt_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(schedulingRequestResourceConfigExtV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PeriodicityAndOffset_r17 != nil}); err != nil {
		return err
	}

	// 2. periodicityAndOffset-r17: choice
	{
		if ie.PeriodicityAndOffset_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(schedulingRequestResourceConfigExt_v1700PeriodicityAndOffsetR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.PeriodicityAndOffset_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.PeriodicityAndOffset_r17).Choice {
			case SchedulingRequestResourceConfigExt_v1700_PeriodicityAndOffset_r17_Sl1280:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset_r17).Sl1280), per.Constrained(0, 1279)); err != nil {
					return err
				}
			case SchedulingRequestResourceConfigExt_v1700_PeriodicityAndOffset_r17_Sl2560:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset_r17).Sl2560), per.Constrained(0, 2559)); err != nil {
					return err
				}
			case SchedulingRequestResourceConfigExt_v1700_PeriodicityAndOffset_r17_Sl5120:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset_r17).Sl5120), per.Constrained(0, 5119)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.PeriodicityAndOffset_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *SchedulingRequestResourceConfigExt_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(schedulingRequestResourceConfigExtV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. periodicityAndOffset-r17: choice
	{
		if seq.IsComponentPresent(0) {
			ie.PeriodicityAndOffset_r17 = &struct {
				Choice int
				Sl1280 *int64
				Sl2560 *int64
				Sl5120 *int64
			}{}
			choiceDec := d.NewChoiceDecoder(schedulingRequestResourceConfigExt_v1700PeriodicityAndOffsetR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SchedulingRequestResourceConfigExt_v1700_PeriodicityAndOffset_r17_Sl1280:
				(*ie.PeriodicityAndOffset_r17).Sl1280 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 1279))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset_r17).Sl1280) = v
			case SchedulingRequestResourceConfigExt_v1700_PeriodicityAndOffset_r17_Sl2560:
				(*ie.PeriodicityAndOffset_r17).Sl2560 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 2559))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset_r17).Sl2560) = v
			case SchedulingRequestResourceConfigExt_v1700_PeriodicityAndOffset_r17_Sl5120:
				(*ie.PeriodicityAndOffset_r17).Sl5120 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 5119))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset_r17).Sl5120) = v
			}
		}
	}

	return nil
}
