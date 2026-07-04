// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MBSBroadcastConfiguration-r17 (line 614).

var mBSBroadcastConfigurationR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "criticalExtensions"},
	},
}

var mBSBroadcastConfiguration_r17CriticalExtensionsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "mbsBroadcastConfiguration-r17"},
		{Name: "criticalExtensionsFuture"},
	},
}

const (
	MBSBroadcastConfiguration_r17_CriticalExtensions_MbsBroadcastConfiguration_r17 = 0
	MBSBroadcastConfiguration_r17_CriticalExtensions_CriticalExtensionsFuture      = 1
)

type MBSBroadcastConfiguration_r17 struct {
	CriticalExtensions struct {
		Choice                        int
		MbsBroadcastConfiguration_r17 *MBSBroadcastConfiguration_r17_IEs
		CriticalExtensionsFuture      *struct{}
	}
}

func (ie *MBSBroadcastConfiguration_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSBroadcastConfigurationR17Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceEnc := e.NewChoiceEncoder(mBSBroadcastConfiguration_r17CriticalExtensionsConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CriticalExtensions.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CriticalExtensions.Choice {
		case MBSBroadcastConfiguration_r17_CriticalExtensions_MbsBroadcastConfiguration_r17:
			if err := ie.CriticalExtensions.MbsBroadcastConfiguration_r17.Encode(e); err != nil {
				return err
			}
		case MBSBroadcastConfiguration_r17_CriticalExtensions_CriticalExtensionsFuture:
			// empty SEQUENCE alternative
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CriticalExtensions.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *MBSBroadcastConfiguration_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSBroadcastConfigurationR17Constraints)
	_ = seq

	// 1. criticalExtensions: choice
	{
		choiceDec := d.NewChoiceDecoder(mBSBroadcastConfiguration_r17CriticalExtensionsConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CriticalExtensions.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MBSBroadcastConfiguration_r17_CriticalExtensions_MbsBroadcastConfiguration_r17:
			ie.CriticalExtensions.MbsBroadcastConfiguration_r17 = new(MBSBroadcastConfiguration_r17_IEs)
			if err := ie.CriticalExtensions.MbsBroadcastConfiguration_r17.Decode(d); err != nil {
				return err
			}
		case MBSBroadcastConfiguration_r17_CriticalExtensions_CriticalExtensionsFuture:
			ie.CriticalExtensions.CriticalExtensionsFuture = &struct{}{}
			// empty SEQUENCE alternative
		}
	}

	return nil
}
