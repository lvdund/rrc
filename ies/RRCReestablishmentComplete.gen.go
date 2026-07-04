// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReestablishmentComplete (line 913).

var rRCReestablishmentCompleteConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var rRCReestablishmentCompleteCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcReestablishmentComplete"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	RRCReestablishmentComplete_CriticalExtensions_RrcReestablishmentComplete = 0
	RRCReestablishmentComplete_CriticalExtensions_CriticalExtensionsFuture   = 1
)

type RRCReestablishmentComplete struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                     int
		RrcReestablishmentComplete *RRCReestablishmentComplete_IEs
		CriticalExtensionsFuture   *struct{}
	}
}

func (ie *RRCReestablishmentComplete) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReestablishmentCompleteConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(rRCReestablishmentCompleteCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case RRCReestablishmentComplete_CriticalExtensions_RrcReestablishmentComplete:
			if err := ie.CriticalExtensions.RrcReestablishmentComplete.Encode(e); err != nil {
				return err
			}
		case RRCReestablishmentComplete_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *RRCReestablishmentComplete) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReestablishmentCompleteConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(rRCReestablishmentCompleteCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RRCReestablishmentComplete_CriticalExtensions_RrcReestablishmentComplete:
			ie.CriticalExtensions.RrcReestablishmentComplete = new(RRCReestablishmentComplete_IEs)
			if err := ie.CriticalExtensions.RrcReestablishmentComplete.Decode(d); err != nil {
				return err
			}
		case RRCReestablishmentComplete_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
