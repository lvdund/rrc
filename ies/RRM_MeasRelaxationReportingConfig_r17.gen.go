// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRM-MeasRelaxationReportingConfig-r17 (line 26485).

var rRMMeasRelaxationReportingConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "s-SearchDeltaP-Stationary-r17"},
		{Name: "t-SearchDeltaP-Stationary-r17"},
	},
}

const (
	RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_DB2    = 0
	RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_DB3    = 1
	RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_DB6    = 2
	RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_DB9    = 3
	RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_DB12   = 4
	RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_DB15   = 5
	RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_Spare2 = 6
	RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_Spare1 = 7
)

var rRMMeasRelaxationReportingConfigR17SSearchDeltaPStationaryR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_DB2, RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_DB3, RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_DB6, RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_DB9, RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_DB12, RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_DB15, RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_Spare2, RRM_MeasRelaxationReportingConfig_r17_S_SearchDeltaP_Stationary_r17_Spare1},
}

const (
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S5     = 0
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S10    = 1
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S20    = 2
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S30    = 3
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S60    = 4
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S120   = 5
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S180   = 6
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S240   = 7
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S300   = 8
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_Spare7 = 9
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_Spare6 = 10
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_Spare5 = 11
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_Spare4 = 12
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_Spare3 = 13
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_Spare2 = 14
	RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_Spare1 = 15
)

var rRMMeasRelaxationReportingConfigR17TSearchDeltaPStationaryR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S5, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S10, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S20, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S30, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S60, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S120, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S180, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S240, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_S300, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_Spare7, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_Spare6, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_Spare5, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_Spare4, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_Spare3, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_Spare2, RRM_MeasRelaxationReportingConfig_r17_T_SearchDeltaP_Stationary_r17_Spare1},
}

type RRM_MeasRelaxationReportingConfig_r17 struct {
	S_SearchDeltaP_Stationary_r17 int64
	T_SearchDeltaP_Stationary_r17 int64
}

func (ie *RRM_MeasRelaxationReportingConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRMMeasRelaxationReportingConfigR17Constraints)
	_ = seq

	// 1. s-SearchDeltaP-Stationary-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.S_SearchDeltaP_Stationary_r17, rRMMeasRelaxationReportingConfigR17SSearchDeltaPStationaryR17Constraints); err != nil {
			return err
		}
	}

	// 2. t-SearchDeltaP-Stationary-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.T_SearchDeltaP_Stationary_r17, rRMMeasRelaxationReportingConfigR17TSearchDeltaPStationaryR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RRM_MeasRelaxationReportingConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRMMeasRelaxationReportingConfigR17Constraints)
	_ = seq

	// 1. s-SearchDeltaP-Stationary-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(rRMMeasRelaxationReportingConfigR17SSearchDeltaPStationaryR17Constraints)
		if err != nil {
			return err
		}
		ie.S_SearchDeltaP_Stationary_r17 = v0
	}

	// 2. t-SearchDeltaP-Stationary-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(rRMMeasRelaxationReportingConfigR17TSearchDeltaPStationaryR17Constraints)
		if err != nil {
			return err
		}
		ie.T_SearchDeltaP_Stationary_r17 = v1
	}

	return nil
}
