// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCSetup (line 1683).

var rRCSetupConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var rRCSetupCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcSetup"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	RRCSetup_CriticalExtensions_RrcSetup                 = 0
	RRCSetup_CriticalExtensions_CriticalExtensionsFuture = 1
)

type RRCSetup struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		RrcSetup                 *RRCSetup_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *RRCSetup) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSetupConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(rRCSetupCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case RRCSetup_CriticalExtensions_RrcSetup:
			if err := ie.CriticalExtensions.RrcSetup.Encode(e); err != nil {
				return err
			}
		case RRCSetup_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *RRCSetup) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSetupConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(rRCSetupCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RRCSetup_CriticalExtensions_RrcSetup:
			ie.CriticalExtensions.RrcSetup = new(RRCSetup_IEs)
			if err := ie.CriticalExtensions.RrcSetup.Decode(d); err != nil {
				return err
			}
		case RRCSetup_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
