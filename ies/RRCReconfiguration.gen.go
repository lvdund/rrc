// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfiguration (line 963).

var rRCReconfigurationConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var rRCReconfigurationCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcReconfiguration"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	RRCReconfiguration_CriticalExtensions_RrcReconfiguration       = 0
	RRCReconfiguration_CriticalExtensions_CriticalExtensionsFuture = 1
)

type RRCReconfiguration struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		RrcReconfiguration       *RRCReconfiguration_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *RRCReconfiguration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(rRCReconfigurationCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case RRCReconfiguration_CriticalExtensions_RrcReconfiguration:
			if err := ie.CriticalExtensions.RrcReconfiguration.Encode(e); err != nil {
				return err
			}
		case RRCReconfiguration_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *RRCReconfiguration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(rRCReconfigurationCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RRCReconfiguration_CriticalExtensions_RrcReconfiguration:
			ie.CriticalExtensions.RrcReconfiguration = new(RRCReconfiguration_IEs)
			if err := ie.CriticalExtensions.RrcReconfiguration.Decode(d); err != nil {
				return err
			}
		case RRCReconfiguration_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
