// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DedicatedSIBRequest-r16 (line 258).

var dedicatedSIBRequestR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var dedicatedSIBRequest_r16CriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "dedicatedSIBRequest-r16"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	DedicatedSIBRequest_r16_CriticalExtensions_DedicatedSIBRequest_r16  = 0
	DedicatedSIBRequest_r16_CriticalExtensions_CriticalExtensionsFuture = 1
)

type DedicatedSIBRequest_r16 struct {
	CriticalExtensions struct {
		Choice                   int
		DedicatedSIBRequest_r16  *DedicatedSIBRequest_r16_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *DedicatedSIBRequest_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dedicatedSIBRequestR16Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(dedicatedSIBRequest_r16CriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case DedicatedSIBRequest_r16_CriticalExtensions_DedicatedSIBRequest_r16:
			if err := ie.CriticalExtensions.DedicatedSIBRequest_r16.Encode(e); err != nil {
				return err
			}
		case DedicatedSIBRequest_r16_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *DedicatedSIBRequest_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dedicatedSIBRequestR16Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(dedicatedSIBRequest_r16CriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case DedicatedSIBRequest_r16_CriticalExtensions_DedicatedSIBRequest_r16:
			ie.CriticalExtensions.DedicatedSIBRequest_r16 = new(DedicatedSIBRequest_r16_IEs)
			if err := ie.CriticalExtensions.DedicatedSIBRequest_r16.Decode(d); err != nil {
				return err
			}
		case DedicatedSIBRequest_r16_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
