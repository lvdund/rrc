// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-ReportPeriodicityAndOffset (line 7437).

var cSIReportPeriodicityAndOffsetConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "slots4"},
		{Name: "slots5"},
		{Name: "slots8"},
		{Name: "slots10"},
		{Name: "slots16"},
		{Name: "slots20"},
		{Name: "slots40"},
		{Name: "slots80"},
		{Name: "slots160"},
		{Name: "slots320"},
	},
}

const (
	CSI_ReportPeriodicityAndOffset_Slots4   = 0
	CSI_ReportPeriodicityAndOffset_Slots5   = 1
	CSI_ReportPeriodicityAndOffset_Slots8   = 2
	CSI_ReportPeriodicityAndOffset_Slots10  = 3
	CSI_ReportPeriodicityAndOffset_Slots16  = 4
	CSI_ReportPeriodicityAndOffset_Slots20  = 5
	CSI_ReportPeriodicityAndOffset_Slots40  = 6
	CSI_ReportPeriodicityAndOffset_Slots80  = 7
	CSI_ReportPeriodicityAndOffset_Slots160 = 8
	CSI_ReportPeriodicityAndOffset_Slots320 = 9
)

type CSI_ReportPeriodicityAndOffset struct {
	Choice   int
	Slots4   *int64
	Slots5   *int64
	Slots8   *int64
	Slots10  *int64
	Slots16  *int64
	Slots20  *int64
	Slots40  *int64
	Slots80  *int64
	Slots160 *int64
	Slots320 *int64
}

func (ie *CSI_ReportPeriodicityAndOffset) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(cSIReportPeriodicityAndOffsetConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportPeriodicityAndOffset_Slots4:
		if err := e.EncodeInteger((*ie.Slots4), per.Constrained(0, 3)); err != nil {
			return err
		}
	case CSI_ReportPeriodicityAndOffset_Slots5:
		if err := e.EncodeInteger((*ie.Slots5), per.Constrained(0, 4)); err != nil {
			return err
		}
	case CSI_ReportPeriodicityAndOffset_Slots8:
		if err := e.EncodeInteger((*ie.Slots8), per.Constrained(0, 7)); err != nil {
			return err
		}
	case CSI_ReportPeriodicityAndOffset_Slots10:
		if err := e.EncodeInteger((*ie.Slots10), per.Constrained(0, 9)); err != nil {
			return err
		}
	case CSI_ReportPeriodicityAndOffset_Slots16:
		if err := e.EncodeInteger((*ie.Slots16), per.Constrained(0, 15)); err != nil {
			return err
		}
	case CSI_ReportPeriodicityAndOffset_Slots20:
		if err := e.EncodeInteger((*ie.Slots20), per.Constrained(0, 19)); err != nil {
			return err
		}
	case CSI_ReportPeriodicityAndOffset_Slots40:
		if err := e.EncodeInteger((*ie.Slots40), per.Constrained(0, 39)); err != nil {
			return err
		}
	case CSI_ReportPeriodicityAndOffset_Slots80:
		if err := e.EncodeInteger((*ie.Slots80), per.Constrained(0, 79)); err != nil {
			return err
		}
	case CSI_ReportPeriodicityAndOffset_Slots160:
		if err := e.EncodeInteger((*ie.Slots160), per.Constrained(0, 159)); err != nil {
			return err
		}
	case CSI_ReportPeriodicityAndOffset_Slots320:
		if err := e.EncodeInteger((*ie.Slots320), per.Constrained(0, 319)); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "CSI-ReportPeriodicityAndOffset",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *CSI_ReportPeriodicityAndOffset) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(cSIReportPeriodicityAndOffsetConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "CSI-ReportPeriodicityAndOffset", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case CSI_ReportPeriodicityAndOffset_Slots4:
		ie.Slots4 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 3))
		if err != nil {
			return err
		}
		(*ie.Slots4) = v
	case CSI_ReportPeriodicityAndOffset_Slots5:
		ie.Slots5 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 4))
		if err != nil {
			return err
		}
		(*ie.Slots5) = v
	case CSI_ReportPeriodicityAndOffset_Slots8:
		ie.Slots8 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 7))
		if err != nil {
			return err
		}
		(*ie.Slots8) = v
	case CSI_ReportPeriodicityAndOffset_Slots10:
		ie.Slots10 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 9))
		if err != nil {
			return err
		}
		(*ie.Slots10) = v
	case CSI_ReportPeriodicityAndOffset_Slots16:
		ie.Slots16 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 15))
		if err != nil {
			return err
		}
		(*ie.Slots16) = v
	case CSI_ReportPeriodicityAndOffset_Slots20:
		ie.Slots20 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 19))
		if err != nil {
			return err
		}
		(*ie.Slots20) = v
	case CSI_ReportPeriodicityAndOffset_Slots40:
		ie.Slots40 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 39))
		if err != nil {
			return err
		}
		(*ie.Slots40) = v
	case CSI_ReportPeriodicityAndOffset_Slots80:
		ie.Slots80 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 79))
		if err != nil {
			return err
		}
		(*ie.Slots80) = v
	case CSI_ReportPeriodicityAndOffset_Slots160:
		ie.Slots160 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 159))
		if err != nil {
			return err
		}
		(*ie.Slots160) = v
	case CSI_ReportPeriodicityAndOffset_Slots320:
		ie.Slots320 = new(int64)
		v, err := d.DecodeInteger(per.Constrained(0, 319))
		if err != nil {
			return err
		}
		(*ie.Slots320) = v
	default:
		return &per.DecodeError{TypeName: "CSI-ReportPeriodicityAndOffset", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
