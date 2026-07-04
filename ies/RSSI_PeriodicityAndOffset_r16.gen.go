// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RSSI-PeriodicityAndOffset-r16 (line 9373).

var rSSIPeriodicityAndOffsetR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl10"},
		{Name: "sl20"},
		{Name: "sl40"},
		{Name: "sl80"},
		{Name: "sl160"},
		{Name: "sl320"},
		{Name: "s1640"},
	},
}

const (
	RSSI_PeriodicityAndOffset_r16_Sl10  = 0
	RSSI_PeriodicityAndOffset_r16_Sl20  = 1
	RSSI_PeriodicityAndOffset_r16_Sl40  = 2
	RSSI_PeriodicityAndOffset_r16_Sl80  = 3
	RSSI_PeriodicityAndOffset_r16_Sl160 = 4
	RSSI_PeriodicityAndOffset_r16_Sl320 = 5
	RSSI_PeriodicityAndOffset_r16_S1640 = 6
)

type RSSI_PeriodicityAndOffset_r16 struct {
	Choice int
	Sl10   *int64
	Sl20   *int64
	Sl40   *int64
	Sl80   *int64
	Sl160  *int64
	Sl320  *int64
	S1640  *int64
}

func (ie *RSSI_PeriodicityAndOffset_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(rSSIPeriodicityAndOffsetR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case RSSI_PeriodicityAndOffset_r16_Sl10:
		if err := e.EncodeInteger((*ie.Sl10), per.Constrained(0, 9)); err != nil {
			return err
		}
	case RSSI_PeriodicityAndOffset_r16_Sl20:
		if err := e.EncodeInteger((*ie.Sl20), per.Constrained(0, 19)); err != nil {
			return err
		}
	case RSSI_PeriodicityAndOffset_r16_Sl40:
		if err := e.EncodeInteger((*ie.Sl40), per.Constrained(0, 39)); err != nil {
			return err
		}
	case RSSI_PeriodicityAndOffset_r16_Sl80:
		if err := e.EncodeInteger((*ie.Sl80), per.Constrained(0, 79)); err != nil {
			return err
		}
	case RSSI_PeriodicityAndOffset_r16_Sl160:
		if err := e.EncodeInteger((*ie.Sl160), per.Constrained(0, 159)); err != nil {
			return err
		}
	case RSSI_PeriodicityAndOffset_r16_Sl320:
		if err := e.EncodeInteger((*ie.Sl320), per.Constrained(0, 319)); err != nil {
			return err
		}
	case RSSI_PeriodicityAndOffset_r16_S1640:
		if err := e.EncodeInteger((*ie.S1640), per.Constrained(0, 639)); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "RSSI-PeriodicityAndOffset-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *RSSI_PeriodicityAndOffset_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(rSSIPeriodicityAndOffsetR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "RSSI-PeriodicityAndOffset-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case RSSI_PeriodicityAndOffset_r16_Sl10:
		ie.Sl10 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 9))
		if err != nil {
			return err
		}
		(*ie.Sl10) = v
	case RSSI_PeriodicityAndOffset_r16_Sl20:
		ie.Sl20 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 19))
		if err != nil {
			return err
		}
		(*ie.Sl20) = v
	case RSSI_PeriodicityAndOffset_r16_Sl40:
		ie.Sl40 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 39))
		if err != nil {
			return err
		}
		(*ie.Sl40) = v
	case RSSI_PeriodicityAndOffset_r16_Sl80:
		ie.Sl80 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 79))
		if err != nil {
			return err
		}
		(*ie.Sl80) = v
	case RSSI_PeriodicityAndOffset_r16_Sl160:
		ie.Sl160 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 159))
		if err != nil {
			return err
		}
		(*ie.Sl160) = v
	case RSSI_PeriodicityAndOffset_r16_Sl320:
		ie.Sl320 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 319))
		if err != nil {
			return err
		}
		(*ie.Sl320) = v
	case RSSI_PeriodicityAndOffset_r16_S1640:
		ie.S1640 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 639))
		if err != nil {
			return err
		}
		(*ie.S1640) = v
	default:
		return &per.DecodeError{TypeName: "RSSI-PeriodicityAndOffset-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
