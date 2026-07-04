// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MPE-Config-FR2-r17 (line 11581).

var mPEConfigFR2R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mpe-ProhibitTimer-r17"},
		{Name: "mpe-Threshold-r17"},
		{Name: "numberOfN-r17"},
	},
}

const (
	MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf0    = 0
	MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf10   = 1
	MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf20   = 2
	MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf50   = 3
	MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf100  = 4
	MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf200  = 5
	MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf500  = 6
	MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf1000 = 7
)

var mPEConfigFR2R17MpeProhibitTimerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf0, MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf10, MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf20, MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf50, MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf100, MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf200, MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf500, MPE_Config_FR2_r17_Mpe_ProhibitTimer_r17_Sf1000},
}

const (
	MPE_Config_FR2_r17_Mpe_Threshold_r17_DB3  = 0
	MPE_Config_FR2_r17_Mpe_Threshold_r17_DB6  = 1
	MPE_Config_FR2_r17_Mpe_Threshold_r17_DB9  = 2
	MPE_Config_FR2_r17_Mpe_Threshold_r17_DB12 = 3
)

var mPEConfigFR2R17MpeThresholdR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MPE_Config_FR2_r17_Mpe_Threshold_r17_DB3, MPE_Config_FR2_r17_Mpe_Threshold_r17_DB6, MPE_Config_FR2_r17_Mpe_Threshold_r17_DB9, MPE_Config_FR2_r17_Mpe_Threshold_r17_DB12},
}

var mPEConfigFR2R17NumberOfNR17Constraints = per.Constrained(1, 4)

type MPE_Config_FR2_r17 struct {
	Mpe_ProhibitTimer_r17 int64
	Mpe_Threshold_r17     int64
	NumberOfN_r17         int64
}

func (ie *MPE_Config_FR2_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mPEConfigFR2R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. mpe-ProhibitTimer-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mpe_ProhibitTimer_r17, mPEConfigFR2R17MpeProhibitTimerR17Constraints); err != nil {
			return err
		}
	}

	// 3. mpe-Threshold-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mpe_Threshold_r17, mPEConfigFR2R17MpeThresholdR17Constraints); err != nil {
			return err
		}
	}

	// 4. numberOfN-r17: integer
	{
		if err := e.EncodeInteger(ie.NumberOfN_r17, mPEConfigFR2R17NumberOfNR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MPE_Config_FR2_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mPEConfigFR2R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. mpe-ProhibitTimer-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(mPEConfigFR2R17MpeProhibitTimerR17Constraints)
		if err != nil {
			return err
		}
		ie.Mpe_ProhibitTimer_r17 = v0
	}

	// 3. mpe-Threshold-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(mPEConfigFR2R17MpeThresholdR17Constraints)
		if err != nil {
			return err
		}
		ie.Mpe_Threshold_r17 = v1
	}

	// 4. numberOfN-r17: integer
	{
		v2, err := d.DecodeInteger(mPEConfigFR2R17NumberOfNR17Constraints)
		if err != nil {
			return err
		}
		ie.NumberOfN_r17 = v2
	}

	return nil
}
