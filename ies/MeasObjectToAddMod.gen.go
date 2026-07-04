// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasObjectToAddMod (line 9647).

var measObjectToAddModConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measObjectId"},
		{Name: "measObject"},
	},
}

var measObjectToAddModMeasObjectConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "measObjectNR"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "measObjectEUTRA"},
		{Name: "measObjectUTRA-FDD-r16"},
		{Name: "measObjectNR-SL-r16"},
		{Name: "measObjectCLI-r16"},
		{Name: "measObjectRxTxDiff-r17"},
		{Name: "measObjectRelay-r17"},
		{Name: "measObjectNR-SL-r18"},
	},
}

const (
	MeasObjectToAddMod_MeasObject_MeasObjectNR = 0
)

type MeasObjectToAddMod struct {
	MeasObjectId MeasObjectId
	MeasObject   struct {
		Choice       int
		MeasObjectNR *MeasObjectNR
	}
}

func (ie *MeasObjectToAddMod) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measObjectToAddModConstraints)
	_ = seq

	// 1. measObjectId: ref
	{
		if err := ie.MeasObjectId.Encode(e); err != nil {
			return err
		}
	}

	// 2. measObject: choice
	{
		choiceEnc := e.NewChoiceEncoder(measObjectToAddModMeasObjectConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.MeasObject.Choice), false, nil); err != nil {
			return err
		}
		switch ie.MeasObject.Choice {
		case MeasObjectToAddMod_MeasObject_MeasObjectNR:
			if err := ie.MeasObject.MeasObjectNR.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.MeasObject.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *MeasObjectToAddMod) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measObjectToAddModConstraints)
	_ = seq

	// 1. measObjectId: ref
	{
		if err := ie.MeasObjectId.Decode(d); err != nil {
			return err
		}
	}

	// 2. measObject: choice
	{
		choiceDec := d.NewChoiceDecoder(measObjectToAddModMeasObjectConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.MeasObject.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MeasObjectToAddMod_MeasObject_MeasObjectNR:
			ie.MeasObject.MeasObjectNR = new(MeasObjectNR)
			if err := ie.MeasObject.MeasObjectNR.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
