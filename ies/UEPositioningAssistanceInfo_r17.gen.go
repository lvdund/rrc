// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEPositioningAssistanceInfo-r17 (line 3573).

var uEPositioningAssistanceInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var uEPositioningAssistanceInfo_r17CriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "uePositioningAssistanceInfo-r17"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	UEPositioningAssistanceInfo_r17_CriticalExtensions_UePositioningAssistanceInfo_r17 = 0
	UEPositioningAssistanceInfo_r17_CriticalExtensions_CriticalExtensionsFuture        = 1
)

type UEPositioningAssistanceInfo_r17 struct {
	CriticalExtensions struct {
		Choice                          int
		UePositioningAssistanceInfo_r17 *UEPositioningAssistanceInfo_r17_IEs
		CriticalExtensionsFuture        *struct{}
	}
}

func (ie *UEPositioningAssistanceInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEPositioningAssistanceInfoR17Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(uEPositioningAssistanceInfo_r17CriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case UEPositioningAssistanceInfo_r17_CriticalExtensions_UePositioningAssistanceInfo_r17:
			if err := ie.CriticalExtensions.UePositioningAssistanceInfo_r17.Encode(e); err != nil {
				return err
			}
		case UEPositioningAssistanceInfo_r17_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *UEPositioningAssistanceInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEPositioningAssistanceInfoR17Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(uEPositioningAssistanceInfo_r17CriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case UEPositioningAssistanceInfo_r17_CriticalExtensions_UePositioningAssistanceInfo_r17:
			ie.CriticalExtensions.UePositioningAssistanceInfo_r17 = new(UEPositioningAssistanceInfo_r17_IEs)
			if err := ie.CriticalExtensions.UePositioningAssistanceInfo_r17.Decode(d); err != nil {
				return err
			}
		case UEPositioningAssistanceInfo_r17_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
