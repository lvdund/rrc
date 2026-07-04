// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-PeriodicityAndOffset (line 15655).

var sRSPeriodicityAndOffsetConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl1"},
		{Name: "sl2"},
		{Name: "sl4"},
		{Name: "sl5"},
		{Name: "sl8"},
		{Name: "sl10"},
		{Name: "sl16"},
		{Name: "sl20"},
		{Name: "sl32"},
		{Name: "sl40"},
		{Name: "sl64"},
		{Name: "sl80"},
		{Name: "sl160"},
		{Name: "sl320"},
		{Name: "sl640"},
		{Name: "sl1280"},
		{Name: "sl2560"},
	},
}

const (
	SRS_PeriodicityAndOffset_Sl1    = 0
	SRS_PeriodicityAndOffset_Sl2    = 1
	SRS_PeriodicityAndOffset_Sl4    = 2
	SRS_PeriodicityAndOffset_Sl5    = 3
	SRS_PeriodicityAndOffset_Sl8    = 4
	SRS_PeriodicityAndOffset_Sl10   = 5
	SRS_PeriodicityAndOffset_Sl16   = 6
	SRS_PeriodicityAndOffset_Sl20   = 7
	SRS_PeriodicityAndOffset_Sl32   = 8
	SRS_PeriodicityAndOffset_Sl40   = 9
	SRS_PeriodicityAndOffset_Sl64   = 10
	SRS_PeriodicityAndOffset_Sl80   = 11
	SRS_PeriodicityAndOffset_Sl160  = 12
	SRS_PeriodicityAndOffset_Sl320  = 13
	SRS_PeriodicityAndOffset_Sl640  = 14
	SRS_PeriodicityAndOffset_Sl1280 = 15
	SRS_PeriodicityAndOffset_Sl2560 = 16
)

type SRS_PeriodicityAndOffset struct {
	Choice int
	Sl1    *struct{}
	Sl2    *int64
	Sl4    *int64
	Sl5    *int64
	Sl8    *int64
	Sl10   *int64
	Sl16   *int64
	Sl20   *int64
	Sl32   *int64
	Sl40   *int64
	Sl64   *int64
	Sl80   *int64
	Sl160  *int64
	Sl320  *int64
	Sl640  *int64
	Sl1280 *int64
	Sl2560 *int64
}

func (ie *SRS_PeriodicityAndOffset) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sRSPeriodicityAndOffsetConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PeriodicityAndOffset_Sl1:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl2:
		if err := e.EncodeInteger((*ie.Sl2), per.Constrained(0, 1)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl4:
		if err := e.EncodeInteger((*ie.Sl4), per.Constrained(0, 3)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl5:
		if err := e.EncodeInteger((*ie.Sl5), per.Constrained(0, 4)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl8:
		if err := e.EncodeInteger((*ie.Sl8), per.Constrained(0, 7)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl10:
		if err := e.EncodeInteger((*ie.Sl10), per.Constrained(0, 9)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl16:
		if err := e.EncodeInteger((*ie.Sl16), per.Constrained(0, 15)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl20:
		if err := e.EncodeInteger((*ie.Sl20), per.Constrained(0, 19)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl32:
		if err := e.EncodeInteger((*ie.Sl32), per.Constrained(0, 31)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl40:
		if err := e.EncodeInteger((*ie.Sl40), per.Constrained(0, 39)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl64:
		if err := e.EncodeInteger((*ie.Sl64), per.Constrained(0, 63)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl80:
		if err := e.EncodeInteger((*ie.Sl80), per.Constrained(0, 79)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl160:
		if err := e.EncodeInteger((*ie.Sl160), per.Constrained(0, 159)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl320:
		if err := e.EncodeInteger((*ie.Sl320), per.Constrained(0, 319)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl640:
		if err := e.EncodeInteger((*ie.Sl640), per.Constrained(0, 639)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl1280:
		if err := e.EncodeInteger((*ie.Sl1280), per.Constrained(0, 1279)); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl2560:
		if err := e.EncodeInteger((*ie.Sl2560), per.Constrained(0, 2559)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SRS-PeriodicityAndOffset",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SRS_PeriodicityAndOffset) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sRSPeriodicityAndOffsetConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SRS-PeriodicityAndOffset", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SRS_PeriodicityAndOffset_Sl1:
		ie.Sl1 = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case SRS_PeriodicityAndOffset_Sl2:
		ie.Sl2 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 1))
		if err != nil {
			return err
		}
		(*ie.Sl2) = v
	case SRS_PeriodicityAndOffset_Sl4:
		ie.Sl4 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 3))
		if err != nil {
			return err
		}
		(*ie.Sl4) = v
	case SRS_PeriodicityAndOffset_Sl5:
		ie.Sl5 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 4))
		if err != nil {
			return err
		}
		(*ie.Sl5) = v
	case SRS_PeriodicityAndOffset_Sl8:
		ie.Sl8 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 7))
		if err != nil {
			return err
		}
		(*ie.Sl8) = v
	case SRS_PeriodicityAndOffset_Sl10:
		ie.Sl10 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 9))
		if err != nil {
			return err
		}
		(*ie.Sl10) = v
	case SRS_PeriodicityAndOffset_Sl16:
		ie.Sl16 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 15))
		if err != nil {
			return err
		}
		(*ie.Sl16) = v
	case SRS_PeriodicityAndOffset_Sl20:
		ie.Sl20 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 19))
		if err != nil {
			return err
		}
		(*ie.Sl20) = v
	case SRS_PeriodicityAndOffset_Sl32:
		ie.Sl32 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 31))
		if err != nil {
			return err
		}
		(*ie.Sl32) = v
	case SRS_PeriodicityAndOffset_Sl40:
		ie.Sl40 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 39))
		if err != nil {
			return err
		}
		(*ie.Sl40) = v
	case SRS_PeriodicityAndOffset_Sl64:
		ie.Sl64 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 63))
		if err != nil {
			return err
		}
		(*ie.Sl64) = v
	case SRS_PeriodicityAndOffset_Sl80:
		ie.Sl80 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 79))
		if err != nil {
			return err
		}
		(*ie.Sl80) = v
	case SRS_PeriodicityAndOffset_Sl160:
		ie.Sl160 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 159))
		if err != nil {
			return err
		}
		(*ie.Sl160) = v
	case SRS_PeriodicityAndOffset_Sl320:
		ie.Sl320 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 319))
		if err != nil {
			return err
		}
		(*ie.Sl320) = v
	case SRS_PeriodicityAndOffset_Sl640:
		ie.Sl640 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 639))
		if err != nil {
			return err
		}
		(*ie.Sl640) = v
	case SRS_PeriodicityAndOffset_Sl1280:
		ie.Sl1280 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 1279))
		if err != nil {
			return err
		}
		(*ie.Sl1280) = v
	case SRS_PeriodicityAndOffset_Sl2560:
		ie.Sl2560 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 2559))
		if err != nil {
			return err
		}
		(*ie.Sl2560) = v
	default:
		return &per.DecodeError{TypeName: "SRS-PeriodicityAndOffset", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
