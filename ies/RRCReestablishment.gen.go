// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReestablishment (line 891).

var rRCReestablishmentConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var rRCReestablishmentCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcReestablishment"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	RRCReestablishment_CriticalExtensions_RrcReestablishment       = 0
	RRCReestablishment_CriticalExtensions_CriticalExtensionsFuture = 1
)

type RRCReestablishment struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		RrcReestablishment       *RRCReestablishment_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *RRCReestablishment) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReestablishmentConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(rRCReestablishmentCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case RRCReestablishment_CriticalExtensions_RrcReestablishment:
			if err := ie.CriticalExtensions.RrcReestablishment.Encode(e); err != nil {
				return err
			}
		case RRCReestablishment_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *RRCReestablishment) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReestablishmentConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(rRCReestablishmentCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RRCReestablishment_CriticalExtensions_RrcReestablishment:
			ie.CriticalExtensions.RrcReestablishment = new(RRCReestablishment_IEs)
			if err := ie.CriticalExtensions.RrcReestablishment.Decode(d); err != nil {
				return err
			}
		case RRCReestablishment_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
