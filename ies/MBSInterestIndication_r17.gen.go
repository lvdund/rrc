// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MBSInterestIndication-r17 (line 639).

var mBSInterestIndicationR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var mBSInterestIndication_r17CriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "mbsInterestIndication-r17"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	MBSInterestIndication_r17_CriticalExtensions_MbsInterestIndication_r17 = 0
	MBSInterestIndication_r17_CriticalExtensions_CriticalExtensionsFuture  = 1
)

type MBSInterestIndication_r17 struct {
	CriticalExtensions struct {
		Choice                    int
		MbsInterestIndication_r17 *MBSInterestIndication_r17_IEs
		CriticalExtensionsFuture  *struct{}
	}
}

func (ie *MBSInterestIndication_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSInterestIndicationR17Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(mBSInterestIndication_r17CriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case MBSInterestIndication_r17_CriticalExtensions_MbsInterestIndication_r17:
			if err := ie.CriticalExtensions.MbsInterestIndication_r17.Encode(e); err != nil {
				return err
			}
		case MBSInterestIndication_r17_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *MBSInterestIndication_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSInterestIndicationR17Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(mBSInterestIndication_r17CriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MBSInterestIndication_r17_CriticalExtensions_MbsInterestIndication_r17:
			ie.CriticalExtensions.MbsInterestIndication_r17 = new(MBSInterestIndication_r17_IEs)
			if err := ie.CriticalExtensions.MbsInterestIndication_r17.Decode(d); err != nil {
				return err
			}
		case MBSInterestIndication_r17_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
