// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReleasePreferenceConfig-r16 (line 26462).

var releasePreferenceConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "releasePreferenceProhibitTimer-r16"},
		{Name: "connectedReporting-r16", Optional: true},
	},
}

const (
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S0       = 0
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S0dot5   = 1
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S1       = 2
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S2       = 3
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S3       = 4
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S4       = 5
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S5       = 6
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S6       = 7
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S7       = 8
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S8       = 9
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S9       = 10
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S10      = 11
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S20      = 12
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S30      = 13
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_Infinity = 14
	ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_Spare1   = 15
)

var releasePreferenceConfigR16ReleasePreferenceProhibitTimerR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S0, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S0dot5, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S1, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S2, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S3, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S4, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S5, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S6, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S7, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S8, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S9, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S10, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S20, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_S30, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_Infinity, ReleasePreferenceConfig_r16_ReleasePreferenceProhibitTimer_r16_Spare1},
}

const (
	ReleasePreferenceConfig_r16_ConnectedReporting_r16_True = 0
)

var releasePreferenceConfigR16ConnectedReportingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReleasePreferenceConfig_r16_ConnectedReporting_r16_True},
}

type ReleasePreferenceConfig_r16 struct {
	ReleasePreferenceProhibitTimer_r16 int64
	ConnectedReporting_r16             *int64
}

func (ie *ReleasePreferenceConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(releasePreferenceConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ConnectedReporting_r16 != nil}); err != nil {
		return err
	}

	// 2. releasePreferenceProhibitTimer-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.ReleasePreferenceProhibitTimer_r16, releasePreferenceConfigR16ReleasePreferenceProhibitTimerR16Constraints); err != nil {
			return err
		}
	}

	// 3. connectedReporting-r16: enumerated
	{
		if ie.ConnectedReporting_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ConnectedReporting_r16, releasePreferenceConfigR16ConnectedReportingR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ReleasePreferenceConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(releasePreferenceConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. releasePreferenceProhibitTimer-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(releasePreferenceConfigR16ReleasePreferenceProhibitTimerR16Constraints)
		if err != nil {
			return err
		}
		ie.ReleasePreferenceProhibitTimer_r16 = v0
	}

	// 3. connectedReporting-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(releasePreferenceConfigR16ConnectedReportingR16Constraints)
			if err != nil {
				return err
			}
			ie.ConnectedReporting_r16 = &idx
		}
	}

	return nil
}
