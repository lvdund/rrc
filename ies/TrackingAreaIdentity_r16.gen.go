// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TrackingAreaIdentity-r16 (line 26081).

var trackingAreaIdentityR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-Identity-r16"},
		{Name: "trackingAreaCode-r16"},
	},
}

type TrackingAreaIdentity_r16 struct {
	Plmn_Identity_r16    PLMN_Identity
	TrackingAreaCode_r16 TrackingAreaCode
}

func (ie *TrackingAreaIdentity_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(trackingAreaIdentityR16Constraints)
	_ = seq

	// 1. plmn-Identity-r16: ref
	{
		if err := ie.Plmn_Identity_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. trackingAreaCode-r16: ref
	{
		if err := ie.TrackingAreaCode_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *TrackingAreaIdentity_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(trackingAreaIdentityR16Constraints)
	_ = seq

	// 1. plmn-Identity-r16: ref
	{
		if err := ie.Plmn_Identity_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. trackingAreaCode-r16: ref
	{
		if err := ie.TrackingAreaCode_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
