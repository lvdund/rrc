// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReleasePreference-r16 (line 2686).

var releasePreferenceR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "preferredRRC-State-r16"},
	},
}

const (
	ReleasePreference_r16_PreferredRRC_State_r16_Idle           = 0
	ReleasePreference_r16_PreferredRRC_State_r16_Inactive       = 1
	ReleasePreference_r16_PreferredRRC_State_r16_Connected      = 2
	ReleasePreference_r16_PreferredRRC_State_r16_OutOfConnected = 3
)

var releasePreferenceR16PreferredRRCStateR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReleasePreference_r16_PreferredRRC_State_r16_Idle, ReleasePreference_r16_PreferredRRC_State_r16_Inactive, ReleasePreference_r16_PreferredRRC_State_r16_Connected, ReleasePreference_r16_PreferredRRC_State_r16_OutOfConnected},
}

type ReleasePreference_r16 struct {
	PreferredRRC_State_r16 int64
}

func (ie *ReleasePreference_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(releasePreferenceR16Constraints)
	_ = seq

	// 1. preferredRRC-State-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.PreferredRRC_State_r16, releasePreferenceR16PreferredRRCStateR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *ReleasePreference_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(releasePreferenceR16Constraints)
	_ = seq

	// 1. preferredRRC-State-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(releasePreferenceR16PreferredRRCStateR16Constraints)
		if err != nil {
			return err
		}
		ie.PreferredRRC_State_r16 = v0
	}

	return nil
}
