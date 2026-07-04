// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NCR-SlotPeriodicityAndSlotOffset-r18 (line 10325).

var nCRSlotPeriodicityAndSlotOffsetR18Constraints = per.ChoiceConstraints{
	Extensible: true,
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
		{Name: "sl128"},
		{Name: "sl160"},
		{Name: "sl256"},
		{Name: "sl320"},
		{Name: "sl512"},
		{Name: "sl640"},
		{Name: "sl1024"},
		{Name: "sl1280"},
		{Name: "sl2560"},
		{Name: "sl5120"},
		{Name: "sl10240"},
	},
}

const (
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl1     = 0
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl2     = 1
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl4     = 2
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl5     = 3
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl8     = 4
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl10    = 5
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl16    = 6
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl20    = 7
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl32    = 8
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl40    = 9
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl64    = 10
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl80    = 11
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl128   = 12
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl160   = 13
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl256   = 14
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl320   = 15
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl512   = 16
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl640   = 17
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl1024  = 18
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl1280  = 19
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl2560  = 20
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl5120  = 21
	NCR_SlotPeriodicityAndSlotOffset_r18_Sl10240 = 22
)

type NCR_SlotPeriodicityAndSlotOffset_r18 struct {
	Choice  int
	Sl1     *struct{}
	Sl2     *int64
	Sl4     *int64
	Sl5     *int64
	Sl8     *int64
	Sl10    *int64
	Sl16    *int64
	Sl20    *int64
	Sl32    *int64
	Sl40    *int64
	Sl64    *int64
	Sl80    *int64
	Sl128   *int64
	Sl160   *int64
	Sl256   *int64
	Sl320   *int64
	Sl512   *int64
	Sl640   *int64
	Sl1024  *int64
	Sl1280  *int64
	Sl2560  *int64
	Sl5120  *int64
	Sl10240 *int64
}

func (ie *NCR_SlotPeriodicityAndSlotOffset_r18) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(nCRSlotPeriodicityAndSlotOffsetR18Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl1:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl2:
		if err := e.EncodeInteger((*ie.Sl2), per.Constrained(0, 1)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl4:
		if err := e.EncodeInteger((*ie.Sl4), per.Constrained(0, 3)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl5:
		if err := e.EncodeInteger((*ie.Sl5), per.Constrained(0, 4)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl8:
		if err := e.EncodeInteger((*ie.Sl8), per.Constrained(0, 7)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl10:
		if err := e.EncodeInteger((*ie.Sl10), per.Constrained(0, 9)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl16:
		if err := e.EncodeInteger((*ie.Sl16), per.Constrained(0, 15)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl20:
		if err := e.EncodeInteger((*ie.Sl20), per.Constrained(0, 19)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl32:
		if err := e.EncodeInteger((*ie.Sl32), per.Constrained(0, 31)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl40:
		if err := e.EncodeInteger((*ie.Sl40), per.Constrained(0, 39)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl64:
		if err := e.EncodeInteger((*ie.Sl64), per.Constrained(0, 63)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl80:
		if err := e.EncodeInteger((*ie.Sl80), per.Constrained(0, 79)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl128:
		if err := e.EncodeInteger((*ie.Sl128), per.Constrained(0, 127)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl160:
		if err := e.EncodeInteger((*ie.Sl160), per.Constrained(0, 159)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl256:
		if err := e.EncodeInteger((*ie.Sl256), per.Constrained(0, 255)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl320:
		if err := e.EncodeInteger((*ie.Sl320), per.Constrained(0, 319)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl512:
		if err := e.EncodeInteger((*ie.Sl512), per.Constrained(0, 511)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl640:
		if err := e.EncodeInteger((*ie.Sl640), per.Constrained(0, 639)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl1024:
		if err := e.EncodeInteger((*ie.Sl1024), per.Constrained(0, 1023)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl1280:
		if err := e.EncodeInteger((*ie.Sl1280), per.Constrained(0, 1279)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl2560:
		if err := e.EncodeInteger((*ie.Sl2560), per.Constrained(0, 2559)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl5120:
		if err := e.EncodeInteger((*ie.Sl5120), per.Constrained(0, 5119)); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl10240:
		if err := e.EncodeInteger((*ie.Sl10240), per.Constrained(0, 10239)); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "NCR-SlotPeriodicityAndSlotOffset-r18",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *NCR_SlotPeriodicityAndSlotOffset_r18) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(nCRSlotPeriodicityAndSlotOffsetR18Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "NCR-SlotPeriodicityAndSlotOffset-r18", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl1:
		ie.Sl1 = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl2:
		ie.Sl2 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 1))
		if err != nil {
			return err
		}
		(*ie.Sl2) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl4:
		ie.Sl4 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 3))
		if err != nil {
			return err
		}
		(*ie.Sl4) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl5:
		ie.Sl5 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 4))
		if err != nil {
			return err
		}
		(*ie.Sl5) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl8:
		ie.Sl8 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 7))
		if err != nil {
			return err
		}
		(*ie.Sl8) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl10:
		ie.Sl10 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 9))
		if err != nil {
			return err
		}
		(*ie.Sl10) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl16:
		ie.Sl16 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 15))
		if err != nil {
			return err
		}
		(*ie.Sl16) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl20:
		ie.Sl20 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 19))
		if err != nil {
			return err
		}
		(*ie.Sl20) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl32:
		ie.Sl32 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 31))
		if err != nil {
			return err
		}
		(*ie.Sl32) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl40:
		ie.Sl40 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 39))
		if err != nil {
			return err
		}
		(*ie.Sl40) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl64:
		ie.Sl64 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 63))
		if err != nil {
			return err
		}
		(*ie.Sl64) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl80:
		ie.Sl80 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 79))
		if err != nil {
			return err
		}
		(*ie.Sl80) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl128:
		ie.Sl128 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 127))
		if err != nil {
			return err
		}
		(*ie.Sl128) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl160:
		ie.Sl160 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 159))
		if err != nil {
			return err
		}
		(*ie.Sl160) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl256:
		ie.Sl256 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 255))
		if err != nil {
			return err
		}
		(*ie.Sl256) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl320:
		ie.Sl320 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 319))
		if err != nil {
			return err
		}
		(*ie.Sl320) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl512:
		ie.Sl512 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 511))
		if err != nil {
			return err
		}
		(*ie.Sl512) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl640:
		ie.Sl640 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 639))
		if err != nil {
			return err
		}
		(*ie.Sl640) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl1024:
		ie.Sl1024 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 1023))
		if err != nil {
			return err
		}
		(*ie.Sl1024) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl1280:
		ie.Sl1280 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 1279))
		if err != nil {
			return err
		}
		(*ie.Sl1280) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl2560:
		ie.Sl2560 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 2559))
		if err != nil {
			return err
		}
		(*ie.Sl2560) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl5120:
		ie.Sl5120 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 5119))
		if err != nil {
			return err
		}
		(*ie.Sl5120) = v
	case NCR_SlotPeriodicityAndSlotOffset_r18_Sl10240:
		ie.Sl10240 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 10239))
		if err != nil {
			return err
		}
		(*ie.Sl10240) = v
	default:
		return &per.DecodeError{TypeName: "NCR-SlotPeriodicityAndSlotOffset-r18", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
