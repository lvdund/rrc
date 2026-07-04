// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReject (line 1205).

var rRCRejectConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var rRCRejectCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rrcReject"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	RRCReject_CriticalExtensions_RrcReject                = 0
	RRCReject_CriticalExtensions_CriticalExtensionsFuture = 1
)

type RRCReject struct {
	CriticalExtensions struct {
		Choice                   int
		RrcReject                *RRCReject_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *RRCReject) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCRejectConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(rRCRejectCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case RRCReject_CriticalExtensions_RrcReject:
			if err := ie.CriticalExtensions.RrcReject.Encode(e); err != nil {
				return err
			}
		case RRCReject_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *RRCReject) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCRejectConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(rRCRejectCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RRCReject_CriticalExtensions_RrcReject:
			ie.CriticalExtensions.RrcReject = new(RRCReject_IEs)
			if err := ie.CriticalExtensions.RrcReject.Decode(d); err != nil {
				return err
			}
		case RRCReject_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
