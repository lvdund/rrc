// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCRelease (line 1221).

var rRCReleaseConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var rRCReleaseCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcRelease"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	RRCRelease_CriticalExtensions_RrcRelease               = 0
	RRCRelease_CriticalExtensions_CriticalExtensionsFuture = 1
)

type RRCRelease struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		RrcRelease               *RRCRelease_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *RRCRelease) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReleaseConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(rRCReleaseCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case RRCRelease_CriticalExtensions_RrcRelease:
			if err := ie.CriticalExtensions.RrcRelease.Encode(e); err != nil {
				return err
			}
		case RRCRelease_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *RRCRelease) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReleaseConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(rRCReleaseCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RRCRelease_CriticalExtensions_RrcRelease:
			ie.CriticalExtensions.RrcRelease = new(RRCRelease_IEs)
			if err := ie.CriticalExtensions.RrcRelease.Decode(d); err != nil {
				return err
			}
		case RRCRelease_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
