// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultsSL-r16 (line 10004).

var measResultsSLR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultsListSL-r16"},
	},
}

var measResultsSL_r16MeasResultsListSLR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "measResultNR-SL-r16"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "measResultNR-SL-r18"},
	},
}

const (
	MeasResultsSL_r16_MeasResultsListSL_r16_MeasResultNR_SL_r16 = 0
)

type MeasResultsSL_r16 struct {
	MeasResultsListSL_r16 struct {
		Choice              int
		MeasResultNR_SL_r16 *MeasResultNR_SL_r16
	}
}

func (ie *MeasResultsSL_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultsSLR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. measResultsListSL-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(measResultsSL_r16MeasResultsListSLR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.MeasResultsListSL_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.MeasResultsListSL_r16.Choice {
		case MeasResultsSL_r16_MeasResultsListSL_r16_MeasResultNR_SL_r16:
			if err := ie.MeasResultsListSL_r16.MeasResultNR_SL_r16.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.MeasResultsListSL_r16.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *MeasResultsSL_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultsSLR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. measResultsListSL-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(measResultsSL_r16MeasResultsListSLR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.MeasResultsListSL_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MeasResultsSL_r16_MeasResultsListSL_r16_MeasResultNR_SL_r16:
			ie.MeasResultsListSL_r16.MeasResultNR_SL_r16 = new(MeasResultNR_SL_r16)
			if err := ie.MeasResultsListSL_r16.MeasResultNR_SL_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
