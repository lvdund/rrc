// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FBS-PreferenceReportingConfig-r19 (line 26631).

var fBSPreferenceReportingConfigR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fbs-PreferenceReportingConfigProhibitTimer-19"},
	},
}

const (
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S0     = 0
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S0dot5 = 1
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S1     = 2
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S2     = 3
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S3     = 4
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S4     = 5
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S5     = 6
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S6     = 7
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S7     = 8
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S8     = 9
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S9     = 10
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S10    = 11
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S20    = 12
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S30    = 13
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_Spare2 = 14
	FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_Spare1 = 15
)

var fBSPreferenceReportingConfigR19FbsPreferenceReportingConfigProhibitTimer19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S0, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S0dot5, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S1, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S2, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S3, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S4, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S5, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S6, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S7, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S8, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S9, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S10, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S20, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_S30, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_Spare2, FBS_PreferenceReportingConfig_r19_Fbs_PreferenceReportingConfigProhibitTimer_19_Spare1},
}

type FBS_PreferenceReportingConfig_r19 struct {
	Fbs_PreferenceReportingConfigProhibitTimer_19 int64
}

func (ie *FBS_PreferenceReportingConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(fBSPreferenceReportingConfigR19Constraints)
	_ = seq

	// 1. fbs-PreferenceReportingConfigProhibitTimer-19: enumerated
	{
		if err := e.EncodeEnumerated(ie.Fbs_PreferenceReportingConfigProhibitTimer_19, fBSPreferenceReportingConfigR19FbsPreferenceReportingConfigProhibitTimer19Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *FBS_PreferenceReportingConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(fBSPreferenceReportingConfigR19Constraints)
	_ = seq

	// 1. fbs-PreferenceReportingConfigProhibitTimer-19: enumerated
	{
		v0, err := d.DecodeEnumerated(fBSPreferenceReportingConfigR19FbsPreferenceReportingConfigProhibitTimer19Constraints)
		if err != nil {
			return err
		}
		ie.Fbs_PreferenceReportingConfigProhibitTimer_19 = v0
	}

	return nil
}
