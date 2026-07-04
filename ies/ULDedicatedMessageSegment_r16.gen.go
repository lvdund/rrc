// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ULDedicatedMessageSegment-r16 (line 3615).

var uLDedicatedMessageSegmentR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var uLDedicatedMessageSegment_r16CriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ulDedicatedMessageSegment-r16"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	ULDedicatedMessageSegment_r16_CriticalExtensions_UlDedicatedMessageSegment_r16 = 0
	ULDedicatedMessageSegment_r16_CriticalExtensions_CriticalExtensionsFuture      = 1
)

type ULDedicatedMessageSegment_r16 struct {
	CriticalExtensions struct {
		Choice                        int
		UlDedicatedMessageSegment_r16 *ULDedicatedMessageSegment_r16_IEs
		CriticalExtensionsFuture      *struct{}
	}
}

func (ie *ULDedicatedMessageSegment_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLDedicatedMessageSegmentR16Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(uLDedicatedMessageSegment_r16CriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case ULDedicatedMessageSegment_r16_CriticalExtensions_UlDedicatedMessageSegment_r16:
			if err := ie.CriticalExtensions.UlDedicatedMessageSegment_r16.Encode(e); err != nil {
				return err
			}
		case ULDedicatedMessageSegment_r16_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *ULDedicatedMessageSegment_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLDedicatedMessageSegmentR16Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(uLDedicatedMessageSegment_r16CriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case ULDedicatedMessageSegment_r16_CriticalExtensions_UlDedicatedMessageSegment_r16:
			ie.CriticalExtensions.UlDedicatedMessageSegment_r16 = new(ULDedicatedMessageSegment_r16_IEs)
			if err := ie.CriticalExtensions.UlDedicatedMessageSegment_r16.Decode(d); err != nil {
				return err
			}
		case ULDedicatedMessageSegment_r16_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
