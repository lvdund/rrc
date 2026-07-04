// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MPE-Config-FR2-r16 (line 11576).

var mPEConfigFR2R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mpe-ProhibitTimer-r16"},
		{Name: "mpe-Threshold-r16"},
	},
}

const (
	MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf0    = 0
	MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf10   = 1
	MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf20   = 2
	MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf50   = 3
	MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf100  = 4
	MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf200  = 5
	MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf500  = 6
	MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf1000 = 7
)

var mPEConfigFR2R16MpeProhibitTimerR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf0, MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf10, MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf20, MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf50, MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf100, MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf200, MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf500, MPE_Config_FR2_r16_Mpe_ProhibitTimer_r16_Sf1000},
}

const (
	MPE_Config_FR2_r16_Mpe_Threshold_r16_DB3  = 0
	MPE_Config_FR2_r16_Mpe_Threshold_r16_DB6  = 1
	MPE_Config_FR2_r16_Mpe_Threshold_r16_DB9  = 2
	MPE_Config_FR2_r16_Mpe_Threshold_r16_DB12 = 3
)

var mPEConfigFR2R16MpeThresholdR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MPE_Config_FR2_r16_Mpe_Threshold_r16_DB3, MPE_Config_FR2_r16_Mpe_Threshold_r16_DB6, MPE_Config_FR2_r16_Mpe_Threshold_r16_DB9, MPE_Config_FR2_r16_Mpe_Threshold_r16_DB12},
}

type MPE_Config_FR2_r16 struct {
	Mpe_ProhibitTimer_r16 int64
	Mpe_Threshold_r16     int64
}

func (ie *MPE_Config_FR2_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mPEConfigFR2R16Constraints)
	_ = seq

	// 1. mpe-ProhibitTimer-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mpe_ProhibitTimer_r16, mPEConfigFR2R16MpeProhibitTimerR16Constraints); err != nil {
			return err
		}
	}

	// 2. mpe-Threshold-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mpe_Threshold_r16, mPEConfigFR2R16MpeThresholdR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MPE_Config_FR2_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mPEConfigFR2R16Constraints)
	_ = seq

	// 1. mpe-ProhibitTimer-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(mPEConfigFR2R16MpeProhibitTimerR16Constraints)
		if err != nil {
			return err
		}
		ie.Mpe_ProhibitTimer_r16 = v0
	}

	// 2. mpe-Threshold-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(mPEConfigFR2R16MpeThresholdR16Constraints)
		if err != nil {
			return err
		}
		ie.Mpe_Threshold_r16 = v1
	}

	return nil
}
