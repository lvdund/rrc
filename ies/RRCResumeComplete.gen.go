// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCResumeComplete (line 1589).

var rRCResumeCompleteConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var rRCResumeCompleteCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcResumeComplete"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	RRCResumeComplete_CriticalExtensions_RrcResumeComplete        = 0
	RRCResumeComplete_CriticalExtensions_CriticalExtensionsFuture = 1
)

type RRCResumeComplete struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		RrcResumeComplete        *RRCResumeComplete_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *RRCResumeComplete) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeCompleteConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(rRCResumeCompleteCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case RRCResumeComplete_CriticalExtensions_RrcResumeComplete:
			if err := ie.CriticalExtensions.RrcResumeComplete.Encode(e); err != nil {
				return err
			}
		case RRCResumeComplete_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *RRCResumeComplete) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeCompleteConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(rRCResumeCompleteCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RRCResumeComplete_CriticalExtensions_RrcResumeComplete:
			ie.CriticalExtensions.RrcResumeComplete = new(RRCResumeComplete_IEs)
			if err := ie.CriticalExtensions.RrcResumeComplete.Decode(d); err != nil {
				return err
			}
		case RRCResumeComplete_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
