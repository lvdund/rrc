// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-CapabilityRestrictionConfig-r18 (line 26397).

var mUSIMCapabilityRestrictionConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "musim-CandidateBandList-r18", Optional: true},
		{Name: "musim-WaitTimer-r18"},
		{Name: "musim-ProhibitTimer-r18"},
	},
}

const (
	MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Ms10   = 0
	MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Ms20   = 1
	MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Ms40   = 2
	MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Ms60   = 3
	MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Ms80   = 4
	MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Ms100  = 5
	MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Spare2 = 6
	MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Spare1 = 7
)

var mUSIMCapabilityRestrictionConfigR18MusimWaitTimerR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Ms10, MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Ms20, MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Ms40, MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Ms60, MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Ms80, MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Ms100, MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Spare2, MUSIM_CapabilityRestrictionConfig_r18_Musim_WaitTimer_r18_Spare1},
}

const (
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S0     = 0
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S0dot1 = 1
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S0dot2 = 2
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S0dot3 = 3
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S0dot4 = 4
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S0dot5 = 5
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S1     = 6
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S2     = 7
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S3     = 8
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S4     = 9
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S5     = 10
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S6     = 11
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S7     = 12
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S8     = 13
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S9     = 14
	MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S10    = 15
)

var mUSIMCapabilityRestrictionConfigR18MusimProhibitTimerR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S0, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S0dot1, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S0dot2, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S0dot3, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S0dot4, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S0dot5, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S1, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S2, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S3, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S4, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S5, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S6, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S7, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S8, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S9, MUSIM_CapabilityRestrictionConfig_r18_Musim_ProhibitTimer_r18_S10},
}

type MUSIM_CapabilityRestrictionConfig_r18 struct {
	Musim_CandidateBandList_r18 *MUSIM_CandidateBandList_r18
	Musim_WaitTimer_r18         int64
	Musim_ProhibitTimer_r18     int64
}

func (ie *MUSIM_CapabilityRestrictionConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMCapabilityRestrictionConfigR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Musim_CandidateBandList_r18 != nil}); err != nil {
		return err
	}

	// 2. musim-CandidateBandList-r18: ref
	{
		if ie.Musim_CandidateBandList_r18 != nil {
			if err := ie.Musim_CandidateBandList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. musim-WaitTimer-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.Musim_WaitTimer_r18, mUSIMCapabilityRestrictionConfigR18MusimWaitTimerR18Constraints); err != nil {
			return err
		}
	}

	// 4. musim-ProhibitTimer-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.Musim_ProhibitTimer_r18, mUSIMCapabilityRestrictionConfigR18MusimProhibitTimerR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MUSIM_CapabilityRestrictionConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMCapabilityRestrictionConfigR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. musim-CandidateBandList-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Musim_CandidateBandList_r18 = new(MUSIM_CandidateBandList_r18)
			if err := ie.Musim_CandidateBandList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. musim-WaitTimer-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(mUSIMCapabilityRestrictionConfigR18MusimWaitTimerR18Constraints)
		if err != nil {
			return err
		}
		ie.Musim_WaitTimer_r18 = v1
	}

	// 4. musim-ProhibitTimer-r18: enumerated
	{
		v2, err := d.DecodeEnumerated(mUSIMCapabilityRestrictionConfigR18MusimProhibitTimerR18Constraints)
		if err != nil {
			return err
		}
		ie.Musim_ProhibitTimer_r18 = v2
	}

	return nil
}
