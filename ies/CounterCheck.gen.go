// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CounterCheck (line 207).

var counterCheckConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var counterCheckCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "counterCheck"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	CounterCheck_CriticalExtensions_CounterCheck             = 0
	CounterCheck_CriticalExtensions_CriticalExtensionsFuture = 1
)

type CounterCheck struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		CounterCheck             *CounterCheck_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *CounterCheck) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(counterCheckConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(counterCheckCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case CounterCheck_CriticalExtensions_CounterCheck:
			if err := ie.CriticalExtensions.CounterCheck.Encode(e); err != nil {
				return err
			}
		case CounterCheck_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *CounterCheck) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(counterCheckConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(counterCheckCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CounterCheck_CriticalExtensions_CounterCheck:
			ie.CriticalExtensions.CounterCheck = new(CounterCheck_IEs)
			if err := ie.CriticalExtensions.CounterCheck.Decode(d); err != nil {
				return err
			}
		case CounterCheck_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
