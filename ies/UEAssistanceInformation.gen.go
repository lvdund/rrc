// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEAssistanceInformation (line 2390).

var uEAssistanceInformationConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var uEAssistanceInformationCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ueAssistanceInformation"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	UEAssistanceInformation_CriticalExtensions_UeAssistanceInformation  = 0
	UEAssistanceInformation_CriticalExtensions_CriticalExtensionsFuture = 1
)

type UEAssistanceInformation struct {
	CriticalExtensions struct {
		Choice                   int
		UeAssistanceInformation  *UEAssistanceInformation_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *UEAssistanceInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEAssistanceInformationConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(uEAssistanceInformationCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case UEAssistanceInformation_CriticalExtensions_UeAssistanceInformation:
			if err := ie.CriticalExtensions.UeAssistanceInformation.Encode(e); err != nil {
				return err
			}
		case UEAssistanceInformation_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *UEAssistanceInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEAssistanceInformationConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(uEAssistanceInformationCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case UEAssistanceInformation_CriticalExtensions_UeAssistanceInformation:
			ie.CriticalExtensions.UeAssistanceInformation = new(UEAssistanceInformation_IEs)
			if err := ie.CriticalExtensions.UeAssistanceInformation.Decode(d); err != nil {
				return err
			}
		case UEAssistanceInformation_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
