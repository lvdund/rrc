// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasWindowConfig-r18 (line 10083).

var measWindowConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "windowOffsetPeriodicity-r18"},
		{Name: "windowDuration-r18"},
	},
}

var measWindowConfig_r18WindowOffsetPeriodicityR18Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "periodicityMs40-r18"},
		{Name: "periodicityMs80-r18"},
	},
}

const (
	MeasWindowConfig_r18_WindowOffsetPeriodicity_r18_PeriodicityMs40_r18 = 0
	MeasWindowConfig_r18_WindowOffsetPeriodicity_r18_PeriodicityMs80_r18 = 1
)

const (
	MeasWindowConfig_r18_WindowDuration_r18_Ms2     = 0
	MeasWindowConfig_r18_WindowDuration_r18_Ms5     = 1
	MeasWindowConfig_r18_WindowDuration_r18_Ms5dot5 = 2
	MeasWindowConfig_r18_WindowDuration_r18_Spare1  = 3
)

var measWindowConfigR18WindowDurationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasWindowConfig_r18_WindowDuration_r18_Ms2, MeasWindowConfig_r18_WindowDuration_r18_Ms5, MeasWindowConfig_r18_WindowDuration_r18_Ms5dot5, MeasWindowConfig_r18_WindowDuration_r18_Spare1},
}

type MeasWindowConfig_r18 struct {
	WindowOffsetPeriodicity_r18 struct {
		Choice              int
		PeriodicityMs40_r18 *int64
		PeriodicityMs80_r18 *int64
	}
	WindowDuration_r18 int64
}

func (ie *MeasWindowConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measWindowConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. windowOffsetPeriodicity-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(measWindowConfig_r18WindowOffsetPeriodicityR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.WindowOffsetPeriodicity_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.WindowOffsetPeriodicity_r18.Choice {
		case MeasWindowConfig_r18_WindowOffsetPeriodicity_r18_PeriodicityMs40_r18:
			if err := e.EncodeInteger((*ie.WindowOffsetPeriodicity_r18.PeriodicityMs40_r18), per.Constrained(0, 39)); err != nil {
				return err
			}
		case MeasWindowConfig_r18_WindowOffsetPeriodicity_r18_PeriodicityMs80_r18:
			if err := e.EncodeInteger((*ie.WindowOffsetPeriodicity_r18.PeriodicityMs80_r18), per.Constrained(0, 79)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.WindowOffsetPeriodicity_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 3. windowDuration-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.WindowDuration_r18, measWindowConfigR18WindowDurationR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasWindowConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measWindowConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. windowOffsetPeriodicity-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(measWindowConfig_r18WindowOffsetPeriodicityR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.WindowOffsetPeriodicity_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MeasWindowConfig_r18_WindowOffsetPeriodicity_r18_PeriodicityMs40_r18:
			ie.WindowOffsetPeriodicity_r18.PeriodicityMs40_r18 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*ie.WindowOffsetPeriodicity_r18.PeriodicityMs40_r18) = v
		case MeasWindowConfig_r18_WindowOffsetPeriodicity_r18_PeriodicityMs80_r18:
			ie.WindowOffsetPeriodicity_r18.PeriodicityMs80_r18 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*ie.WindowOffsetPeriodicity_r18.PeriodicityMs80_r18) = v
		}
	}

	// 3. windowDuration-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(measWindowConfigR18WindowDurationR18Constraints)
		if err != nil {
			return err
		}
		ie.WindowDuration_r18 = v1
	}

	return nil
}
