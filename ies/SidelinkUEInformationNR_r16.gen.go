// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SidelinkUEInformationNR-r16 (line 2120).

var sidelinkUEInformationNRR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var sidelinkUEInformationNR_r16CriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sidelinkUEInformationNR-r16"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	SidelinkUEInformationNR_r16_CriticalExtensions_SidelinkUEInformationNR_r16 = 0
	SidelinkUEInformationNR_r16_CriticalExtensions_CriticalExtensionsFuture    = 1
)

type SidelinkUEInformationNR_r16 struct {
	CriticalExtensions struct {
		Choice                      int
		SidelinkUEInformationNR_r16 *SidelinkUEInformationNR_r16_IEs
		CriticalExtensionsFuture    *struct{}
	}
}

func (ie *SidelinkUEInformationNR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sidelinkUEInformationNRR16Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(sidelinkUEInformationNR_r16CriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case SidelinkUEInformationNR_r16_CriticalExtensions_SidelinkUEInformationNR_r16:
			if err := ie.CriticalExtensions.SidelinkUEInformationNR_r16.Encode(e); err != nil {
				return err
			}
		case SidelinkUEInformationNR_r16_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SidelinkUEInformationNR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sidelinkUEInformationNRR16Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(sidelinkUEInformationNR_r16CriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SidelinkUEInformationNR_r16_CriticalExtensions_SidelinkUEInformationNR_r16:
			ie.CriticalExtensions.SidelinkUEInformationNR_r16 = new(SidelinkUEInformationNR_r16_IEs)
			if err := ie.CriticalExtensions.SidelinkUEInformationNR_r16.Decode(d); err != nil {
				return err
			}
		case SidelinkUEInformationNR_r16_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
