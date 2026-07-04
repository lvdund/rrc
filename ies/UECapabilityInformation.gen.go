// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UECapabilityInformation (line 2880).

var uECapabilityInformationConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var uECapabilityInformationCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ueCapabilityInformation"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	UECapabilityInformation_CriticalExtensions_UeCapabilityInformation  = 0
	UECapabilityInformation_CriticalExtensions_CriticalExtensionsFuture = 1
)

type UECapabilityInformation struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		UeCapabilityInformation  *UECapabilityInformation_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *UECapabilityInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uECapabilityInformationConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(uECapabilityInformationCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case UECapabilityInformation_CriticalExtensions_UeCapabilityInformation:
			if err := ie.CriticalExtensions.UeCapabilityInformation.Encode(e); err != nil {
				return err
			}
		case UECapabilityInformation_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *UECapabilityInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uECapabilityInformationConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(uECapabilityInformationCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case UECapabilityInformation_CriticalExtensions_UeCapabilityInformation:
			ie.CriticalExtensions.UeCapabilityInformation = new(UECapabilityInformation_IEs)
			if err := ie.CriticalExtensions.UeCapabilityInformation.Decode(d); err != nil {
				return err
			}
		case UECapabilityInformation_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
