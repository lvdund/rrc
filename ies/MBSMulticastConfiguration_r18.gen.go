// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MBSMulticastConfiguration-r18 (line 662).

var mBSMulticastConfigurationR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var mBSMulticastConfiguration_r18CriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "mbsMulticastConfiguration-r18"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	MBSMulticastConfiguration_r18_CriticalExtensions_MbsMulticastConfiguration_r18 = 0
	MBSMulticastConfiguration_r18_CriticalExtensions_CriticalExtensionsFuture      = 1
)

type MBSMulticastConfiguration_r18 struct {
	CriticalExtensions struct {
		Choice                        int
		MbsMulticastConfiguration_r18 *MBSMulticastConfiguration_r18_IEs
		CriticalExtensionsFuture      *struct{}
	}
}

func (ie *MBSMulticastConfiguration_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSMulticastConfigurationR18Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(mBSMulticastConfiguration_r18CriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case MBSMulticastConfiguration_r18_CriticalExtensions_MbsMulticastConfiguration_r18:
			if err := ie.CriticalExtensions.MbsMulticastConfiguration_r18.Encode(e); err != nil {
				return err
			}
		case MBSMulticastConfiguration_r18_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *MBSMulticastConfiguration_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSMulticastConfigurationR18Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(mBSMulticastConfiguration_r18CriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MBSMulticastConfiguration_r18_CriticalExtensions_MbsMulticastConfiguration_r18:
			ie.CriticalExtensions.MbsMulticastConfiguration_r18 = new(MBSMulticastConfiguration_r18_IEs)
			if err := ie.CriticalExtensions.MbsMulticastConfiguration_r18.Decode(d); err != nil {
				return err
			}
		case MBSMulticastConfiguration_r18_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
