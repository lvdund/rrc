// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SwitchingPeriodUnaffectedBandDualUL-r18 (line 17019).

var switchingPeriodUnaffectedBandDualULR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandIndexUnaffected-r18"},
		{Name: "periodUnaffectedBandDualUL-r18"},
	},
}

var switchingPeriodUnaffectedBandDualULR18BandIndexUnaffectedR18Constraints = per.Constrained(1, common.MaxSimultaneousBands)

var switchingPeriodUnaffectedBandDualUL_r18PeriodUnaffectedBandDualULR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "maintainedUL-Trans-r18"},
		{Name: "periodOnULBands-r18"},
	},
}

const (
	SwitchingPeriodUnaffectedBandDualUL_r18_PeriodUnaffectedBandDualUL_r18_MaintainedUL_Trans_r18 = 0
	SwitchingPeriodUnaffectedBandDualUL_r18_PeriodUnaffectedBandDualUL_r18_PeriodOnULBands_r18    = 1
)

const (
	SwitchingPeriodUnaffectedBandDualUL_r18_PeriodUnaffectedBandDualUL_r18_PeriodOnULBands_r18_N35us  = 0
	SwitchingPeriodUnaffectedBandDualUL_r18_PeriodUnaffectedBandDualUL_r18_PeriodOnULBands_r18_N140us = 1
	SwitchingPeriodUnaffectedBandDualUL_r18_PeriodUnaffectedBandDualUL_r18_PeriodOnULBands_r18_N210us = 2
)

var switchingPeriodUnaffectedBandDualULR18PeriodUnaffectedBandDualULR18PeriodOnULBandsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SwitchingPeriodUnaffectedBandDualUL_r18_PeriodUnaffectedBandDualUL_r18_PeriodOnULBands_r18_N35us, SwitchingPeriodUnaffectedBandDualUL_r18_PeriodUnaffectedBandDualUL_r18_PeriodOnULBands_r18_N140us, SwitchingPeriodUnaffectedBandDualUL_r18_PeriodUnaffectedBandDualUL_r18_PeriodOnULBands_r18_N210us},
}

type SwitchingPeriodUnaffectedBandDualUL_r18 struct {
	BandIndexUnaffected_r18        int64
	PeriodUnaffectedBandDualUL_r18 struct {
		Choice                 int
		MaintainedUL_Trans_r18 *struct{}
		PeriodOnULBands_r18    *int64
	}
}

func (ie *SwitchingPeriodUnaffectedBandDualUL_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(switchingPeriodUnaffectedBandDualULR18Constraints)
	_ = seq

	// 1. bandIndexUnaffected-r18: integer
	{
		if err := e.EncodeInteger(ie.BandIndexUnaffected_r18, switchingPeriodUnaffectedBandDualULR18BandIndexUnaffectedR18Constraints); err != nil {
			return err
		}
	}

	// 2. periodUnaffectedBandDualUL-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(switchingPeriodUnaffectedBandDualUL_r18PeriodUnaffectedBandDualULR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.PeriodUnaffectedBandDualUL_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.PeriodUnaffectedBandDualUL_r18.Choice {
		case SwitchingPeriodUnaffectedBandDualUL_r18_PeriodUnaffectedBandDualUL_r18_MaintainedUL_Trans_r18:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case SwitchingPeriodUnaffectedBandDualUL_r18_PeriodUnaffectedBandDualUL_r18_PeriodOnULBands_r18:
			if err := e.EncodeEnumerated((*ie.PeriodUnaffectedBandDualUL_r18.PeriodOnULBands_r18), switchingPeriodUnaffectedBandDualULR18PeriodUnaffectedBandDualULR18PeriodOnULBandsR18Constraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.PeriodUnaffectedBandDualUL_r18.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SwitchingPeriodUnaffectedBandDualUL_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(switchingPeriodUnaffectedBandDualULR18Constraints)
	_ = seq

	// 1. bandIndexUnaffected-r18: integer
	{
		v0, err := d.DecodeInteger(switchingPeriodUnaffectedBandDualULR18BandIndexUnaffectedR18Constraints)
		if err != nil {
			return err
		}
		ie.BandIndexUnaffected_r18 = v0
	}

	// 2. periodUnaffectedBandDualUL-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(switchingPeriodUnaffectedBandDualUL_r18PeriodUnaffectedBandDualULR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.PeriodUnaffectedBandDualUL_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SwitchingPeriodUnaffectedBandDualUL_r18_PeriodUnaffectedBandDualUL_r18_MaintainedUL_Trans_r18:
			ie.PeriodUnaffectedBandDualUL_r18.MaintainedUL_Trans_r18 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case SwitchingPeriodUnaffectedBandDualUL_r18_PeriodUnaffectedBandDualUL_r18_PeriodOnULBands_r18:
			ie.PeriodUnaffectedBandDualUL_r18.PeriodOnULBands_r18 = new(int64)
			v, err := d.DecodeEnumerated(switchingPeriodUnaffectedBandDualULR18PeriodUnaffectedBandDualULR18PeriodOnULBandsR18Constraints)
			if err != nil {
				return err
			}
			(*ie.PeriodUnaffectedBandDualUL_r18.PeriodOnULBands_r18) = v
		}
	}

	return nil
}
