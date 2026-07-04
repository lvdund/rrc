// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ULInformationTransfer (line 3633).

var uLInformationTransferConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var uLInformationTransferCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ulInformationTransfer"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	ULInformationTransfer_CriticalExtensions_UlInformationTransfer    = 0
	ULInformationTransfer_CriticalExtensions_CriticalExtensionsFuture = 1
)

type ULInformationTransfer struct {
	CriticalExtensions struct {
		Choice                   int
		UlInformationTransfer    *ULInformationTransfer_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *ULInformationTransfer) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLInformationTransferConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(uLInformationTransferCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case ULInformationTransfer_CriticalExtensions_UlInformationTransfer:
			if err := ie.CriticalExtensions.UlInformationTransfer.Encode(e); err != nil {
				return err
			}
		case ULInformationTransfer_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *ULInformationTransfer) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLInformationTransferConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(uLInformationTransferCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case ULInformationTransfer_CriticalExtensions_UlInformationTransfer:
			ie.CriticalExtensions.UlInformationTransfer = new(ULInformationTransfer_IEs)
			if err := ie.CriticalExtensions.UlInformationTransfer.Decode(d); err != nil {
				return err
			}
		case ULInformationTransfer_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
