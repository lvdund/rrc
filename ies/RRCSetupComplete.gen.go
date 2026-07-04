// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCSetupComplete (line 1707).

var rRCSetupCompleteConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var rRCSetupCompleteCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcSetupComplete"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	RRCSetupComplete_CriticalExtensions_RrcSetupComplete         = 0
	RRCSetupComplete_CriticalExtensions_CriticalExtensionsFuture = 1
)

type RRCSetupComplete struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		RrcSetupComplete         *RRCSetupComplete_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *RRCSetupComplete) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSetupCompleteConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(rRCSetupCompleteCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case RRCSetupComplete_CriticalExtensions_RrcSetupComplete:
			if err := ie.CriticalExtensions.RrcSetupComplete.Encode(e); err != nil {
				return err
			}
		case RRCSetupComplete_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *RRCSetupComplete) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSetupCompleteConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(rRCSetupCompleteCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RRCSetupComplete_CriticalExtensions_RrcSetupComplete:
			ie.CriticalExtensions.RrcSetupComplete = new(RRCSetupComplete_IEs)
			if err := ie.CriticalExtensions.RrcSetupComplete.Decode(d); err != nil {
				return err
			}
		case RRCSetupComplete_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
