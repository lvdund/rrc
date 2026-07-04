// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEInformationRequest-r16 (line 2897).

var uEInformationRequestR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var uEInformationRequest_r16CriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ueInformationRequest-r16"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	UEInformationRequest_r16_CriticalExtensions_UeInformationRequest_r16 = 0
	UEInformationRequest_r16_CriticalExtensions_CriticalExtensionsFuture = 1
)

type UEInformationRequest_r16 struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		UeInformationRequest_r16 *UEInformationRequest_r16_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *UEInformationRequest_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEInformationRequestR16Constraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(uEInformationRequest_r16CriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case UEInformationRequest_r16_CriticalExtensions_UeInformationRequest_r16:
			if err := ie.CriticalExtensions.UeInformationRequest_r16.Encode(e); err != nil {
				return err
			}
		case UEInformationRequest_r16_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *UEInformationRequest_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEInformationRequestR16Constraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(uEInformationRequest_r16CriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case UEInformationRequest_r16_CriticalExtensions_UeInformationRequest_r16:
			ie.CriticalExtensions.UeInformationRequest_r16 = new(UEInformationRequest_r16_IEs)
			if err := ie.CriticalExtensions.UeInformationRequest_r16.Decode(d); err != nil {
				return err
			}
		case UEInformationRequest_r16_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
