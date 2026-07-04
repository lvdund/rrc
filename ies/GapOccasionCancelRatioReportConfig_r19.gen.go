// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: GapOccasionCancelRatioReportConfig-r19 (line 26527).

var gapOccasionCancelRatioReportConfigR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "gapOccasionCancelRatioProhibitTimer-r19"},
	},
}

const (
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S0     = 0
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S0dot5 = 1
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S1     = 2
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S2     = 3
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S5     = 4
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S10    = 5
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S20    = 6
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S30    = 7
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S60    = 8
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S90    = 9
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S120   = 10
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S300   = 11
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S600   = 12
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_Spare3 = 13
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_Spare2 = 14
	GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_Spare1 = 15
)

var gapOccasionCancelRatioReportConfigR19GapOccasionCancelRatioProhibitTimerR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S0, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S0dot5, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S1, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S2, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S5, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S10, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S20, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S30, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S60, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S90, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S120, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S300, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_S600, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_Spare3, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_Spare2, GapOccasionCancelRatioReportConfig_r19_GapOccasionCancelRatioProhibitTimer_r19_Spare1},
}

type GapOccasionCancelRatioReportConfig_r19 struct {
	GapOccasionCancelRatioProhibitTimer_r19 int64
}

func (ie *GapOccasionCancelRatioReportConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(gapOccasionCancelRatioReportConfigR19Constraints)
	_ = seq

	// 1. gapOccasionCancelRatioProhibitTimer-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.GapOccasionCancelRatioProhibitTimer_r19, gapOccasionCancelRatioReportConfigR19GapOccasionCancelRatioProhibitTimerR19Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *GapOccasionCancelRatioReportConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(gapOccasionCancelRatioReportConfigR19Constraints)
	_ = seq

	// 1. gapOccasionCancelRatioProhibitTimer-r19: enumerated
	{
		v0, err := d.DecodeEnumerated(gapOccasionCancelRatioReportConfigR19GapOccasionCancelRatioProhibitTimerR19Constraints)
		if err != nil {
			return err
		}
		ie.GapOccasionCancelRatioProhibitTimer_r19 = v0
	}

	return nil
}
