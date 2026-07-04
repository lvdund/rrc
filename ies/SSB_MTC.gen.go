// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-MTC (line 15851).

var sSBMTCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicityAndOffset"},
		{Name: "duration"},
	},
}

var sSB_MTCPeriodicityAndOffsetConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sf5"},
		{Name: "sf10"},
		{Name: "sf20"},
		{Name: "sf40"},
		{Name: "sf80"},
		{Name: "sf160"},
	},
}

const (
	SSB_MTC_PeriodicityAndOffset_Sf5   = 0
	SSB_MTC_PeriodicityAndOffset_Sf10  = 1
	SSB_MTC_PeriodicityAndOffset_Sf20  = 2
	SSB_MTC_PeriodicityAndOffset_Sf40  = 3
	SSB_MTC_PeriodicityAndOffset_Sf80  = 4
	SSB_MTC_PeriodicityAndOffset_Sf160 = 5
)

const (
	SSB_MTC_Duration_Sf1 = 0
	SSB_MTC_Duration_Sf2 = 1
	SSB_MTC_Duration_Sf3 = 2
	SSB_MTC_Duration_Sf4 = 3
	SSB_MTC_Duration_Sf5 = 4
)

var sSBMTCDurationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SSB_MTC_Duration_Sf1, SSB_MTC_Duration_Sf2, SSB_MTC_Duration_Sf3, SSB_MTC_Duration_Sf4, SSB_MTC_Duration_Sf5},
}

type SSB_MTC struct {
	PeriodicityAndOffset struct {
		Choice int
		Sf5    *int64
		Sf10   *int64
		Sf20   *int64
		Sf40   *int64
		Sf80   *int64
		Sf160  *int64
	}
	Duration int64
}

func (ie *SSB_MTC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBMTCConstraints)
	_ = seq

	// 1. periodicityAndOffset: choice
	{
		choiceEnc := e.NewChoiceEncoder(sSB_MTCPeriodicityAndOffsetConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.PeriodicityAndOffset.Choice), false, nil); err != nil {
			return err
		}
		switch ie.PeriodicityAndOffset.Choice {
		case SSB_MTC_PeriodicityAndOffset_Sf5:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset.Sf5), per.Constrained(0, 4)); err != nil {
				return err
			}
		case SSB_MTC_PeriodicityAndOffset_Sf10:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset.Sf10), per.Constrained(0, 9)); err != nil {
				return err
			}
		case SSB_MTC_PeriodicityAndOffset_Sf20:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset.Sf20), per.Constrained(0, 19)); err != nil {
				return err
			}
		case SSB_MTC_PeriodicityAndOffset_Sf40:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset.Sf40), per.Constrained(0, 39)); err != nil {
				return err
			}
		case SSB_MTC_PeriodicityAndOffset_Sf80:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset.Sf80), per.Constrained(0, 79)); err != nil {
				return err
			}
		case SSB_MTC_PeriodicityAndOffset_Sf160:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset.Sf160), per.Constrained(0, 159)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.PeriodicityAndOffset.Choice), Constraint: "index out of range"}
		}
	}

	// 2. duration: enumerated
	{
		if err := e.EncodeEnumerated(ie.Duration, sSBMTCDurationConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SSB_MTC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBMTCConstraints)
	_ = seq

	// 1. periodicityAndOffset: choice
	{
		choiceDec := d.NewChoiceDecoder(sSB_MTCPeriodicityAndOffsetConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.PeriodicityAndOffset.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SSB_MTC_PeriodicityAndOffset_Sf5:
			ie.PeriodicityAndOffset.Sf5 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 4))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset.Sf5) = v
		case SSB_MTC_PeriodicityAndOffset_Sf10:
			ie.PeriodicityAndOffset.Sf10 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset.Sf10) = v
		case SSB_MTC_PeriodicityAndOffset_Sf20:
			ie.PeriodicityAndOffset.Sf20 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset.Sf20) = v
		case SSB_MTC_PeriodicityAndOffset_Sf40:
			ie.PeriodicityAndOffset.Sf40 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset.Sf40) = v
		case SSB_MTC_PeriodicityAndOffset_Sf80:
			ie.PeriodicityAndOffset.Sf80 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset.Sf80) = v
		case SSB_MTC_PeriodicityAndOffset_Sf160:
			ie.PeriodicityAndOffset.Sf160 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset.Sf160) = v
		}
	}

	// 2. duration: enumerated
	{
		v1, err := d.DecodeEnumerated(sSBMTCDurationConstraints)
		if err != nil {
			return err
		}
		ie.Duration = v1
	}

	return nil
}
