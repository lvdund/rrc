// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DLInformationTransfer (line 338).

var dLInformationTransferConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-TransactionIdentifier"},
		{Name: "criticalExtensions"},
	},
}

var dLInformationTransferCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "dlInformationTransfer"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	DLInformationTransfer_CriticalExtensions_DlInformationTransfer    = 0
	DLInformationTransfer_CriticalExtensions_CriticalExtensionsFuture = 1
)

type DLInformationTransfer struct {
	Rrc_TransactionIdentifier RRC_TransactionIdentifier
	CriticalExtensions        struct {
		Choice                   int
		DlInformationTransfer    *DLInformationTransfer_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *DLInformationTransfer) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLInformationTransferConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(dLInformationTransferCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case DLInformationTransfer_CriticalExtensions_DlInformationTransfer:
			if err := ie.CriticalExtensions.DlInformationTransfer.Encode(e); err != nil {
				return err
			}
		case DLInformationTransfer_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *DLInformationTransfer) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLInformationTransferConstraints)
	_ = seq

	// 1. rrc-TransactionIdentifier: ref
	{
		if err := ie.Rrc_TransactionIdentifier.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(dLInformationTransferCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case DLInformationTransfer_CriticalExtensions_DlInformationTransfer:
			ie.CriticalExtensions.DlInformationTransfer = new(DLInformationTransfer_IEs)
			if err := ie.CriticalExtensions.DlInformationTransfer.Decode(d); err != nil {
				return err
			}
		case DLInformationTransfer_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
