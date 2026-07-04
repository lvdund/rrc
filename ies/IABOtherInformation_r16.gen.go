// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IABOtherInformation-r16 (line 428).

var iABOtherInformationR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dummy"},
		{Name: "criticalExtensions"},
	},
}

var iABOtherInformation_r16CriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "iabOtherInformation-r16"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	IABOtherInformation_r16_CriticalExtensions_IabOtherInformation_r16  = 0
	IABOtherInformation_r16_CriticalExtensions_CriticalExtensionsFuture = 1
)

type IABOtherInformation_r16 struct {
	Dummy              RRC_TransactionIdentifier
	CriticalExtensions struct {
		Choice                   int
		IabOtherInformation_r16  *IABOtherInformation_r16_IEs
		CriticalExtensionsFuture *struct{}
	}
}

func (ie *IABOtherInformation_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABOtherInformationR16Constraints)
	_ = seq

	// 1. dummy: ref
	{
		if err := ie.Dummy.Encode(e); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(iABOtherInformation_r16CriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case IABOtherInformation_r16_CriticalExtensions_IabOtherInformation_r16:
			if err := ie.CriticalExtensions.IabOtherInformation_r16.Encode(e); err != nil {
				return err
			}
		case IABOtherInformation_r16_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *IABOtherInformation_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABOtherInformationR16Constraints)
	_ = seq

	// 1. dummy: ref
	{
		if err := ie.Dummy.Decode(d); err != nil {
			return err
		}
	}

	// 2. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(iABOtherInformation_r16CriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case IABOtherInformation_r16_CriticalExtensions_IabOtherInformation_r16:
			ie.CriticalExtensions.IabOtherInformation_r16 = new(IABOtherInformation_r16_IEs)
			if err := ie.CriticalExtensions.IabOtherInformation_r16.Decode(d); err != nil {
				return err
			}
		case IABOtherInformation_r16_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
