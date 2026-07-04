// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SchedulingRequestResourceConfig (line 14281).

var schedulingRequestResourceConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "schedulingRequestResourceId"},
		{Name: "schedulingRequestID"},
		{Name: "periodicityAndOffset", Optional: true},
		{Name: "resource", Optional: true},
	},
}

var schedulingRequestResourceConfigPeriodicityAndOffsetConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sym2"},
		{Name: "sym6or7"},
		{Name: "sl1"},
		{Name: "sl2"},
		{Name: "sl4"},
		{Name: "sl5"},
		{Name: "sl8"},
		{Name: "sl10"},
		{Name: "sl16"},
		{Name: "sl20"},
		{Name: "sl40"},
		{Name: "sl80"},
		{Name: "sl160"},
		{Name: "sl320"},
		{Name: "sl640"},
	},
}

const (
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sym2    = 0
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sym6or7 = 1
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl1     = 2
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl2     = 3
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl4     = 4
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl5     = 5
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl8     = 6
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl10    = 7
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl16    = 8
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl20    = 9
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl40    = 10
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl80    = 11
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl160   = 12
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl320   = 13
	SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl640   = 14
)

type SchedulingRequestResourceConfig struct {
	SchedulingRequestResourceId SchedulingRequestResourceId
	SchedulingRequestID         SchedulingRequestId
	PeriodicityAndOffset        *struct {
		Choice  int
		Sym2    *struct{}
		Sym6or7 *struct{}
		Sl1     *struct{}
		Sl2     *int64
		Sl4     *int64
		Sl5     *int64
		Sl8     *int64
		Sl10    *int64
		Sl16    *int64
		Sl20    *int64
		Sl40    *int64
		Sl80    *int64
		Sl160   *int64
		Sl320   *int64
		Sl640   *int64
	}
	Resource *PUCCH_ResourceId
}

func (ie *SchedulingRequestResourceConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(schedulingRequestResourceConfigConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PeriodicityAndOffset != nil, ie.Resource != nil}); err != nil {
		return err
	}

	// 2. schedulingRequestResourceId: ref
	{
		if err := ie.SchedulingRequestResourceId.Encode(e); err != nil {
			return err
		}
	}

	// 3. schedulingRequestID: ref
	{
		if err := ie.SchedulingRequestID.Encode(e); err != nil {
			return err
		}
	}

	// 4. periodicityAndOffset: choice
	{
		if ie.PeriodicityAndOffset != nil {
			choiceEnc := e.NewChoiceEncoder(schedulingRequestResourceConfigPeriodicityAndOffsetConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.PeriodicityAndOffset).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.PeriodicityAndOffset).Choice {
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sym2:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sym6or7:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl1:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl2:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset).Sl2), per.Constrained(0, 1)); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl4:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset).Sl4), per.Constrained(0, 3)); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl5:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset).Sl5), per.Constrained(0, 4)); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl8:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset).Sl8), per.Constrained(0, 7)); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl10:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset).Sl10), per.Constrained(0, 9)); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl16:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset).Sl16), per.Constrained(0, 15)); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl20:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset).Sl20), per.Constrained(0, 19)); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl40:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset).Sl40), per.Constrained(0, 39)); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl80:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset).Sl80), per.Constrained(0, 79)); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl160:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset).Sl160), per.Constrained(0, 159)); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl320:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset).Sl320), per.Constrained(0, 319)); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl640:
				if err := e.EncodeInteger((*(*ie.PeriodicityAndOffset).Sl640), per.Constrained(0, 639)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.PeriodicityAndOffset).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. resource: ref
	{
		if ie.Resource != nil {
			if err := ie.Resource.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SchedulingRequestResourceConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(schedulingRequestResourceConfigConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. schedulingRequestResourceId: ref
	{
		if err := ie.SchedulingRequestResourceId.Decode(d); err != nil {
			return err
		}
	}

	// 3. schedulingRequestID: ref
	{
		if err := ie.SchedulingRequestID.Decode(d); err != nil {
			return err
		}
	}

	// 4. periodicityAndOffset: choice
	{
		if seq.IsComponentPresent(2) {
			ie.PeriodicityAndOffset = &struct {
				Choice  int
				Sym2    *struct{}
				Sym6or7 *struct{}
				Sl1     *struct{}
				Sl2     *int64
				Sl4     *int64
				Sl5     *int64
				Sl8     *int64
				Sl10    *int64
				Sl16    *int64
				Sl20    *int64
				Sl40    *int64
				Sl80    *int64
				Sl160   *int64
				Sl320   *int64
				Sl640   *int64
			}{}
			choiceDec := d.NewChoiceDecoder(schedulingRequestResourceConfigPeriodicityAndOffsetConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sym2:
				(*ie.PeriodicityAndOffset).Sym2 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sym6or7:
				(*ie.PeriodicityAndOffset).Sym6or7 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl1:
				(*ie.PeriodicityAndOffset).Sl1 = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl2:
				(*ie.PeriodicityAndOffset).Sl2 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 1))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset).Sl2) = v
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl4:
				(*ie.PeriodicityAndOffset).Sl4 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 3))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset).Sl4) = v
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl5:
				(*ie.PeriodicityAndOffset).Sl5 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 4))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset).Sl5) = v
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl8:
				(*ie.PeriodicityAndOffset).Sl8 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 7))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset).Sl8) = v
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl10:
				(*ie.PeriodicityAndOffset).Sl10 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 9))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset).Sl10) = v
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl16:
				(*ie.PeriodicityAndOffset).Sl16 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset).Sl16) = v
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl20:
				(*ie.PeriodicityAndOffset).Sl20 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 19))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset).Sl20) = v
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl40:
				(*ie.PeriodicityAndOffset).Sl40 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 39))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset).Sl40) = v
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl80:
				(*ie.PeriodicityAndOffset).Sl80 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 79))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset).Sl80) = v
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl160:
				(*ie.PeriodicityAndOffset).Sl160 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 159))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset).Sl160) = v
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl320:
				(*ie.PeriodicityAndOffset).Sl320 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 319))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset).Sl320) = v
			case SchedulingRequestResourceConfig_PeriodicityAndOffset_Sl640:
				(*ie.PeriodicityAndOffset).Sl640 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 639))
				if err != nil {
					return err
				}
				(*(*ie.PeriodicityAndOffset).Sl640) = v
			}
		}
	}

	// 5. resource: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Resource = new(PUCCH_ResourceId)
			if err := ie.Resource.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
