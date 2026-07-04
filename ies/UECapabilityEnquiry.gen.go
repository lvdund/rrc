// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UECapabilityEnquiry (line 2848).

var uECapabilityEnquiryConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var uECapabilityEnquiryCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ueCapabilityEnquiry"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	UECapabilityEnquiry_CriticalExtensions_UeCapabilityEnquiry      = 0
	UECapabilityEnquiry_CriticalExtensions_CriticalExtensionsFuture = 1
)

type UECapabilityEnquiry struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		UeCapabilityEnquiry      *UECapabilityEnquiry_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *UECapabilityEnquiry) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uECapabilityEnquiryConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(uECapabilityEnquiryCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case UECapabilityEnquiry_CriticalExtensions_UeCapabilityEnquiry:
			if err := ie.CriticalExtensions.UeCapabilityEnquiry.Encode(e); err != nil {
				return err
			}
		case UECapabilityEnquiry_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *UECapabilityEnquiry) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uECapabilityEnquiryConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(uECapabilityEnquiryCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case UECapabilityEnquiry_CriticalExtensions_UeCapabilityEnquiry:
			ie.CriticalExtensions.UeCapabilityEnquiry = new(UECapabilityEnquiry_IEs)
			if err := ie.CriticalExtensions.UeCapabilityEnquiry.Decode(d); err != nil {
				return err
			}
		case UECapabilityEnquiry_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
