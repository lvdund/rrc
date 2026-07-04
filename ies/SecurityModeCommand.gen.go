// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SecurityModeCommand (line 1925).

var securityModeCommandConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var securityModeCommandCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "securityModeCommand"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	SecurityModeCommand_CriticalExtensions_SecurityModeCommand      = 0
	SecurityModeCommand_CriticalExtensions_CriticalExtensionsFuture = 1
)

type SecurityModeCommand struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		SecurityModeCommand      *SecurityModeCommand_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *SecurityModeCommand) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(securityModeCommandConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(securityModeCommandCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case SecurityModeCommand_CriticalExtensions_SecurityModeCommand:
			if err := ie.CriticalExtensions.SecurityModeCommand.Encode(e); err != nil {
				return err
			}
		case SecurityModeCommand_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SecurityModeCommand) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(securityModeCommandConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(securityModeCommandCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SecurityModeCommand_CriticalExtensions_SecurityModeCommand:
			ie.CriticalExtensions.SecurityModeCommand = new(SecurityModeCommand_IEs)
			if err := ie.CriticalExtensions.SecurityModeCommand.Decode(d); err != nil {
				return err
			}
		case SecurityModeCommand_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
