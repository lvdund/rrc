// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfigurationComplete (line 1136).

var rRCReconfigurationCompleteConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var rRCReconfigurationCompleteCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcReconfigurationComplete"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	RRCReconfigurationComplete_CriticalExtensions_RrcReconfigurationComplete = 0
	RRCReconfigurationComplete_CriticalExtensions_CriticalExtensionsFuture   = 1
)

type RRCReconfigurationComplete struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                     int
		RrcReconfigurationComplete *RRCReconfigurationComplete_IEs
		CriticalExtensionsFuture   *struct{}
	}
}

func (ie *RRCReconfigurationComplete) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationCompleteConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(rRCReconfigurationCompleteCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case RRCReconfigurationComplete_CriticalExtensions_RrcReconfigurationComplete:
			if err := ie.CriticalExtensions.RrcReconfigurationComplete.Encode(e); err != nil {
				return err
			}
		case RRCReconfigurationComplete_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *RRCReconfigurationComplete) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationCompleteConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(rRCReconfigurationCompleteCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RRCReconfigurationComplete_CriticalExtensions_RrcReconfigurationComplete:
			ie.CriticalExtensions.RrcReconfigurationComplete = new(RRCReconfigurationComplete_IEs)
			if err := ie.CriticalExtensions.RrcReconfigurationComplete.Decode(d); err != nil {
				return err
			}
		case RRCReconfigurationComplete_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
