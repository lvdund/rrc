// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AreaValidityTA-Config-r18 (line 1498).

var areaValidityTAConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "inactivePosSRS-ValidityAreaTAT-r18"},
		{Name: "inactivePosSRS-ValidityAreaRSRP-r18", Optional: true},
		{Name: "autonomousTA-AdjustmentEnabled-r18", Optional: true},
	},
}

const (
	AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Ms1280   = 0
	AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Ms1920   = 1
	AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Ms2560   = 2
	AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Ms5120   = 3
	AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Ms10240  = 4
	AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Ms20480  = 5
	AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Ms40960  = 6
	AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Infinity = 7
)

var areaValidityTAConfigR18InactivePosSRSValidityAreaTATR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Ms1280, AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Ms1920, AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Ms2560, AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Ms5120, AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Ms10240, AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Ms20480, AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Ms40960, AreaValidityTA_Config_r18_InactivePosSRS_ValidityAreaTAT_r18_Infinity},
}

const (
	AreaValidityTA_Config_r18_AutonomousTA_AdjustmentEnabled_r18_True = 0
)

var areaValidityTAConfigR18AutonomousTAAdjustmentEnabledR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AreaValidityTA_Config_r18_AutonomousTA_AdjustmentEnabled_r18_True},
}

type AreaValidityTA_Config_r18 struct {
	InactivePosSRS_ValidityAreaTAT_r18  int64
	InactivePosSRS_ValidityAreaRSRP_r18 *RSRP_ChangeThreshold_r17
	AutonomousTA_AdjustmentEnabled_r18  *int64
}

func (ie *AreaValidityTA_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(areaValidityTAConfigR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InactivePosSRS_ValidityAreaRSRP_r18 != nil, ie.AutonomousTA_AdjustmentEnabled_r18 != nil}); err != nil {
		return err
	}

	// 2. inactivePosSRS-ValidityAreaTAT-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.InactivePosSRS_ValidityAreaTAT_r18, areaValidityTAConfigR18InactivePosSRSValidityAreaTATR18Constraints); err != nil {
			return err
		}
	}

	// 3. inactivePosSRS-ValidityAreaRSRP-r18: ref
	{
		if ie.InactivePosSRS_ValidityAreaRSRP_r18 != nil {
			if err := ie.InactivePosSRS_ValidityAreaRSRP_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. autonomousTA-AdjustmentEnabled-r18: enumerated
	{
		if ie.AutonomousTA_AdjustmentEnabled_r18 != nil {
			if err := e.EncodeEnumerated(*ie.AutonomousTA_AdjustmentEnabled_r18, areaValidityTAConfigR18AutonomousTAAdjustmentEnabledR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *AreaValidityTA_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(areaValidityTAConfigR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. inactivePosSRS-ValidityAreaTAT-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(areaValidityTAConfigR18InactivePosSRSValidityAreaTATR18Constraints)
		if err != nil {
			return err
		}
		ie.InactivePosSRS_ValidityAreaTAT_r18 = v0
	}

	// 3. inactivePosSRS-ValidityAreaRSRP-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.InactivePosSRS_ValidityAreaRSRP_r18 = new(RSRP_ChangeThreshold_r17)
			if err := ie.InactivePosSRS_ValidityAreaRSRP_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. autonomousTA-AdjustmentEnabled-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(areaValidityTAConfigR18AutonomousTAAdjustmentEnabledR18Constraints)
			if err != nil {
				return err
			}
			ie.AutonomousTA_AdjustmentEnabled_r18 = &idx
		}
	}

	return nil
}
