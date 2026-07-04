// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SecurityModeComplete (line 1948).

var securityModeCompleteConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var securityModeCompleteCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "securityModeComplete"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	SecurityModeComplete_CriticalExtensions_SecurityModeComplete     = 0
	SecurityModeComplete_CriticalExtensions_CriticalExtensionsFuture = 1
)

type SecurityModeComplete struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		SecurityModeComplete     *SecurityModeComplete_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *SecurityModeComplete) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(securityModeCompleteConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(securityModeCompleteCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case SecurityModeComplete_CriticalExtensions_SecurityModeComplete:
			if err := ie.CriticalExtensions.SecurityModeComplete.Encode(e); err != nil {
				return err
			}
		case SecurityModeComplete_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SecurityModeComplete) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(securityModeCompleteConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(securityModeCompleteCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SecurityModeComplete_CriticalExtensions_SecurityModeComplete:
			ie.CriticalExtensions.SecurityModeComplete = new(SecurityModeComplete_IEs)
			if err := ie.CriticalExtensions.SecurityModeComplete.Decode(d); err != nil {
				return err
			}
		case SecurityModeComplete_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
