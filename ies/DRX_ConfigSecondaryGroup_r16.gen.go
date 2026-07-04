// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DRX-ConfigSecondaryGroup-r16 (line 8210).

var dRXConfigSecondaryGroupR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "drx-onDurationTimer-r16"},
		{Name: "drx-InactivityTimer-r16"},
	},
}

var dRX_ConfigSecondaryGroup_r16DrxOnDurationTimerR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "subMilliSeconds"},
		{Name: "milliSeconds"},
	},
}

const (
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_SubMilliSeconds = 0
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds    = 1
)

const (
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms0    = 0
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms1    = 1
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms2    = 2
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms3    = 3
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms4    = 4
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms5    = 5
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms6    = 6
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms8    = 7
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms10   = 8
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms20   = 9
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms30   = 10
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms40   = 11
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms50   = 12
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms60   = 13
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms80   = 14
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms100  = 15
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms200  = 16
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms300  = 17
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms500  = 18
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms750  = 19
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms1280 = 20
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms1920 = 21
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms2560 = 22
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare9 = 23
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare8 = 24
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare7 = 25
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare6 = 26
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare5 = 27
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare4 = 28
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare3 = 29
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare2 = 30
	DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare1 = 31
)

var dRXConfigSecondaryGroupR16DrxInactivityTimerR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms0, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms1, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms2, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms3, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms4, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms5, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms6, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms8, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms10, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms20, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms30, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms40, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms50, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms60, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms80, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms100, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms200, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms300, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms500, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms750, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms1280, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms1920, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Ms2560, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare9, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare8, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare7, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare6, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare5, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare4, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare3, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare2, DRX_ConfigSecondaryGroup_r16_Drx_InactivityTimer_r16_Spare1},
}

const (
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms1    = 0
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms2    = 1
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms3    = 2
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms4    = 3
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms5    = 4
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms6    = 5
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms8    = 6
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms10   = 7
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms20   = 8
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms30   = 9
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms40   = 10
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms50   = 11
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms60   = 12
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms80   = 13
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms100  = 14
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms200  = 15
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms300  = 16
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms400  = 17
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms500  = 18
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms600  = 19
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms800  = 20
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms1000 = 21
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms1200 = 22
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms1600 = 23
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare8 = 24
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare7 = 25
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare6 = 26
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare5 = 27
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare4 = 28
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare3 = 29
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare2 = 30
	DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare1 = 31
)

var dRXConfigSecondaryGroupR16DrxOnDurationTimerR16MilliSecondsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms1, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms2, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms3, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms4, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms5, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms6, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms8, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms10, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms20, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms30, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms40, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms50, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms60, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms80, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms100, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms200, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms300, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms400, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms500, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms600, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms800, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms1000, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms1200, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Ms1600, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare8, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare7, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare6, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare5, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare4, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare3, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare2, DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds_Spare1},
}

type DRX_ConfigSecondaryGroup_r16 struct {
	Drx_OnDurationTimer_r16 struct {
		Choice          int
		SubMilliSeconds *int64
		MilliSeconds    *int64
	}
	Drx_InactivityTimer_r16 int64
}

func (ie *DRX_ConfigSecondaryGroup_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRXConfigSecondaryGroupR16Constraints)
	_ = seq

	// 1. drx-onDurationTimer-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(dRX_ConfigSecondaryGroup_r16DrxOnDurationTimerR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Drx_OnDurationTimer_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Drx_OnDurationTimer_r16.Choice {
		case DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_SubMilliSeconds:
			if err := e.EncodeInteger((*ie.Drx_OnDurationTimer_r16.SubMilliSeconds), per.Constrained(1, 31)); err != nil {
				return err
			}
		case DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds:
			if err := e.EncodeEnumerated((*ie.Drx_OnDurationTimer_r16.MilliSeconds), dRXConfigSecondaryGroupR16DrxOnDurationTimerR16MilliSecondsConstraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Drx_OnDurationTimer_r16.Choice), Constraint: "index out of range"}
		}
	}

	// 2. drx-InactivityTimer-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Drx_InactivityTimer_r16, dRXConfigSecondaryGroupR16DrxInactivityTimerR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DRX_ConfigSecondaryGroup_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRXConfigSecondaryGroupR16Constraints)
	_ = seq

	// 1. drx-onDurationTimer-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(dRX_ConfigSecondaryGroup_r16DrxOnDurationTimerR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Drx_OnDurationTimer_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_SubMilliSeconds:
			ie.Drx_OnDurationTimer_r16.SubMilliSeconds = new(int64)
			v, err := d.DecodeInteger(per.Constrained(1, 31))
			if err != nil {
				return err
			}
			(*ie.Drx_OnDurationTimer_r16.SubMilliSeconds) = v
		case DRX_ConfigSecondaryGroup_r16_Drx_OnDurationTimer_r16_MilliSeconds:
			ie.Drx_OnDurationTimer_r16.MilliSeconds = new(int64)
			v, err := d.DecodeEnumerated(dRXConfigSecondaryGroupR16DrxOnDurationTimerR16MilliSecondsConstraints)
			if err != nil {
				return err
			}
			(*ie.Drx_OnDurationTimer_r16.MilliSeconds) = v
		}
	}

	// 2. drx-InactivityTimer-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(dRXConfigSecondaryGroupR16DrxInactivityTimerR16Constraints)
		if err != nil {
			return err
		}
		ie.Drx_InactivityTimer_r16 = v1
	}

	return nil
}
