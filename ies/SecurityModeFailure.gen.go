// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SecurityModeFailure (line 1964).

var securityModeFailureConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var securityModeFailureCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "securityModeFailure"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	SecurityModeFailure_CriticalExtensions_SecurityModeFailure      = 0
	SecurityModeFailure_CriticalExtensions_CriticalExtensionsFuture = 1
)

type SecurityModeFailure struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		SecurityModeFailure      *SecurityModeFailure_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *SecurityModeFailure) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(securityModeFailureConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(securityModeFailureCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case SecurityModeFailure_CriticalExtensions_SecurityModeFailure:
			if err := ie.CriticalExtensions.SecurityModeFailure.Encode(e); err != nil {
				return err
			}
		case SecurityModeFailure_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SecurityModeFailure) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(securityModeFailureConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(securityModeFailureCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SecurityModeFailure_CriticalExtensions_SecurityModeFailure:
			ie.CriticalExtensions.SecurityModeFailure = new(SecurityModeFailure_IEs)
			if err := ie.CriticalExtensions.SecurityModeFailure.Decode(d); err != nil {
				return err
			}
		case SecurityModeFailure_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
