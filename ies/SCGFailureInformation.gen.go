// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SCGFailureInformation (line 1831).

var sCGFailureInformationConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var sCGFailureInformationCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "scgFailureInformation"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	SCGFailureInformation_CriticalExtensions_ScgFailureInformation    = 0
	SCGFailureInformation_CriticalExtensions_CriticalExtensionsFuture = 1
)

type SCGFailureInformation struct {
	CriticalExtensions struct {
		Choice                   int
		ScgFailureInformation    *SCGFailureInformation_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *SCGFailureInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sCGFailureInformationConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(sCGFailureInformationCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case SCGFailureInformation_CriticalExtensions_ScgFailureInformation:
			if err := ie.CriticalExtensions.ScgFailureInformation.Encode(e); err != nil {
				return err
			}
		case SCGFailureInformation_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SCGFailureInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sCGFailureInformationConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(sCGFailureInformationCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SCGFailureInformation_CriticalExtensions_ScgFailureInformation:
			ie.CriticalExtensions.ScgFailureInformation = new(SCGFailureInformation_IEs)
			if err := ie.CriticalExtensions.ScgFailureInformation.Decode(d); err != nil {
				return err
			}
		case SCGFailureInformation_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
