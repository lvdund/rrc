// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MobilityFromNRCommand (line 812).

var mobilityFromNRCommandConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var mobilityFromNRCommandCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "mobilityFromNRCommand"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	MobilityFromNRCommand_CriticalExtensions_MobilityFromNRCommand    = 0
	MobilityFromNRCommand_CriticalExtensions_CriticalExtensionsFuture = 1
)

type MobilityFromNRCommand struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		MobilityFromNRCommand    *MobilityFromNRCommand_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *MobilityFromNRCommand) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mobilityFromNRCommandConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(mobilityFromNRCommandCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case MobilityFromNRCommand_CriticalExtensions_MobilityFromNRCommand:
			if err := ie.CriticalExtensions.MobilityFromNRCommand.Encode(e); err != nil {
				return err
			}
		case MobilityFromNRCommand_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *MobilityFromNRCommand) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mobilityFromNRCommandConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(mobilityFromNRCommandCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MobilityFromNRCommand_CriticalExtensions_MobilityFromNRCommand:
			ie.CriticalExtensions.MobilityFromNRCommand = new(MobilityFromNRCommand_IEs)
			if err := ie.CriticalExtensions.MobilityFromNRCommand.Decode(d); err != nil {
				return err
			}
		case MobilityFromNRCommand_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
