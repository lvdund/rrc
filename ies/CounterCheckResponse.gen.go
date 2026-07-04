// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CounterCheckResponse (line 232).

var counterCheckResponseConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var counterCheckResponseCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "counterCheckResponse"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	CounterCheckResponse_CriticalExtensions_CounterCheckResponse     = 0
	CounterCheckResponse_CriticalExtensions_CriticalExtensionsFuture = 1
)

type CounterCheckResponse struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		CounterCheckResponse     *CounterCheckResponse_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *CounterCheckResponse) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(counterCheckResponseConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(counterCheckResponseCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case CounterCheckResponse_CriticalExtensions_CounterCheckResponse:
			if err := ie.CriticalExtensions.CounterCheckResponse.Encode(e); err != nil {
				return err
			}
		case CounterCheckResponse_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *CounterCheckResponse) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(counterCheckResponseConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(counterCheckResponseCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CounterCheckResponse_CriticalExtensions_CounterCheckResponse:
			ie.CriticalExtensions.CounterCheckResponse = new(CounterCheckResponse_IEs)
			if err := ie.CriticalExtensions.CounterCheckResponse.Decode(d); err != nil {
				return err
			}
		case CounterCheckResponse_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
