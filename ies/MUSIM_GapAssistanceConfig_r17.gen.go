// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-GapAssistanceConfig-r17 (line 26389).

var mUSIMGapAssistanceConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "musim-GapProhibitTimer-r17"},
	},
}

const (
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S0     = 0
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S0dot1 = 1
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S0dot2 = 2
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S0dot3 = 3
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S0dot4 = 4
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S0dot5 = 5
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S1     = 6
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S2     = 7
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S3     = 8
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S4     = 9
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S5     = 10
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S6     = 11
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S7     = 12
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S8     = 13
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S9     = 14
	MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S10    = 15
)

var mUSIMGapAssistanceConfigR17MusimGapProhibitTimerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S0, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S0dot1, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S0dot2, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S0dot3, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S0dot4, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S0dot5, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S1, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S2, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S3, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S4, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S5, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S6, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S7, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S8, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S9, MUSIM_GapAssistanceConfig_r17_Musim_GapProhibitTimer_r17_S10},
}

type MUSIM_GapAssistanceConfig_r17 struct {
	Musim_GapProhibitTimer_r17 int64
}

func (ie *MUSIM_GapAssistanceConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMGapAssistanceConfigR17Constraints)
	_ = seq

	// 1. musim-GapProhibitTimer-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Musim_GapProhibitTimer_r17, mUSIMGapAssistanceConfigR17MusimGapProhibitTimerR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MUSIM_GapAssistanceConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMGapAssistanceConfigR17Constraints)
	_ = seq

	// 1. musim-GapProhibitTimer-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(mUSIMGapAssistanceConfigR17MusimGapProhibitTimerR17Constraints)
		if err != nil {
			return err
		}
		ie.Musim_GapProhibitTimer_r17 = v0
	}

	return nil
}
