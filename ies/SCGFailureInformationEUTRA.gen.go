// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SCGFailureInformationEUTRA (line 1890).

var sCGFailureInformationEUTRAConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var sCGFailureInformationEUTRACriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "scgFailureInformationEUTRA"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	SCGFailureInformationEUTRA_CriticalExtensions_ScgFailureInformationEUTRA = 0
	SCGFailureInformationEUTRA_CriticalExtensions_CriticalExtensionsFuture   = 1
)

type SCGFailureInformationEUTRA struct {
	CriticalExtensions struct {
		Choice                     int
		ScgFailureInformationEUTRA *SCGFailureInformationEUTRA_IEs
		CriticalExtensionsFuture   *struct{}
	}
}

func (ie *SCGFailureInformationEUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sCGFailureInformationEUTRAConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(sCGFailureInformationEUTRACriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case SCGFailureInformationEUTRA_CriticalExtensions_ScgFailureInformationEUTRA:
			if err := ie.CriticalExtensions.ScgFailureInformationEUTRA.Encode(e); err != nil {
				return err
			}
		case SCGFailureInformationEUTRA_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SCGFailureInformationEUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sCGFailureInformationEUTRAConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(sCGFailureInformationEUTRACriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SCGFailureInformationEUTRA_CriticalExtensions_ScgFailureInformationEUTRA:
			ie.CriticalExtensions.ScgFailureInformationEUTRA = new(SCGFailureInformationEUTRA_IEs)
			if err := ie.CriticalExtensions.ScgFailureInformationEUTRA.Decode(d); err != nil {
				return err
			}
		case SCGFailureInformationEUTRA_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
