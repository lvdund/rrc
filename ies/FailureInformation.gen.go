// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FailureInformation (line 397).

var failureInformationConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var failureInformationCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "failureInformation"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	FailureInformation_CriticalExtensions_FailureInformation       = 0
	FailureInformation_CriticalExtensions_CriticalExtensionsFuture = 1
)

type FailureInformation struct {
	CriticalExtensions struct {
		Choice                   int
		FailureInformation       *FailureInformation_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *FailureInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(failureInformationConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(failureInformationCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case FailureInformation_CriticalExtensions_FailureInformation:
			if err := ie.CriticalExtensions.FailureInformation.Encode(e); err != nil {
				return err
			}
		case FailureInformation_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *FailureInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(failureInformationConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(failureInformationCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case FailureInformation_CriticalExtensions_FailureInformation:
			ie.CriticalExtensions.FailureInformation = new(FailureInformation_IEs)
			if err := ie.CriticalExtensions.FailureInformation.Decode(d); err != nil {
				return err
			}
		case FailureInformation_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
