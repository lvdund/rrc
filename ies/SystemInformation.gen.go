// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SystemInformation (line 2341).

var systemInformationConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var systemInformationCriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "systemInformation"},
		{Name: "criticalExtensionsFuture-r16"},
	},
}

const (
	SystemInformation_CriticalExtensions_SystemInformation            = 0
	SystemInformation_CriticalExtensions_CriticalExtensionsFuture_r16 = 1
)

var systemInformationCriticalExtensionsCriticalExtensionsFutureR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "posSystemInformation-r16"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	SystemInformation_CriticalExtensions_CriticalExtensionsFuture_r16_PosSystemInformation_r16 = 0
	SystemInformation_CriticalExtensions_CriticalExtensionsFuture_r16_CriticalExtensionsFuture = 1
)

type SystemInformation struct {
	CriticalExtensions struct {
		Choice                       int
		SystemInformation            *SystemInformation_IEs
		CriticalExtensionsFuture_r16 *struct {
			Choice                   int
			PosSystemInformation_r16 *PosSystemInformation_r16_IEs
			CriticalExtensionsFuture *struct{}
		}
	}
}

func (ie *SystemInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(systemInformationConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(systemInformationCriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case SystemInformation_CriticalExtensions_SystemInformation:
			if err := ie.CriticalExtensions.SystemInformation.Encode(e); err != nil {
				return err
			}
		case SystemInformation_CriticalExtensions_CriticalExtensionsFuture_r16:
			choiceEnc := e.NewChoiceEncoder(systemInformationCriticalExtensionsCriticalExtensionsFutureR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.CriticalExtensions.CriticalExtensionsFuture_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.CriticalExtensions.CriticalExtensionsFuture_r16).Choice {
			case SystemInformation_CriticalExtensions_CriticalExtensionsFuture_r16_PosSystemInformation_r16:
				if err := (*ie.CriticalExtensions.CriticalExtensionsFuture_r16).PosSystemInformation_r16.Encode(e); err != nil {
					return err
				}
			case SystemInformation_CriticalExtensions_CriticalExtensionsFuture_r16_CriticalExtensionsFuture:
				// empty SEQUENCE alternative
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SystemInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(systemInformationConstraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(systemInformationCriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SystemInformation_CriticalExtensions_SystemInformation:
			ie.CriticalExtensions.SystemInformation = new(SystemInformation_IEs)
			if err := ie.CriticalExtensions.SystemInformation.Decode(d); err != nil {
				return err
			}
		case SystemInformation_CriticalExtensions_CriticalExtensionsFuture_r16:
			ie.CriticalExtensions.CriticalExtensionsFuture_r16 = &struct {
				Choice                   int
				PosSystemInformation_r16 *PosSystemInformation_r16_IEs
				CriticalExtensionsFuture *struct{}
			}{}
			choiceDec := d.NewChoiceDecoder(systemInformationCriticalExtensionsCriticalExtensionsFutureR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.CriticalExtensions.CriticalExtensionsFuture_r16).Choice = int(index)
			switch index {
			case SystemInformation_CriticalExtensions_CriticalExtensionsFuture_r16_PosSystemInformation_r16:
				(*ie.CriticalExtensions.CriticalExtensionsFuture_r16).PosSystemInformation_r16 = new(PosSystemInformation_r16_IEs)
				if err := (*ie.CriticalExtensions.CriticalExtensionsFuture_r16).PosSystemInformation_r16.Decode(d); err != nil {
					return err
				}
			case SystemInformation_CriticalExtensions_CriticalExtensionsFuture_r16_CriticalExtensionsFuture:
				(*ie.CriticalExtensions.CriticalExtensionsFuture_r16).CriticalExtensionsFuture = &struct{}{}
				// empty SEQUENCE alternative
			}
		}
	}

	return nil
}
