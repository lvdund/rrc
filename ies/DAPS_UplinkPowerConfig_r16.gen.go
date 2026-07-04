// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DAPS-UplinkPowerConfig-r16 (line 5724).

var dAPSUplinkPowerConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "p-DAPS-Source-r16"},
		{Name: "p-DAPS-Target-r16"},
		{Name: "uplinkPowerSharingDAPS-Mode-r16"},
	},
}

const (
	DAPS_UplinkPowerConfig_r16_UplinkPowerSharingDAPS_Mode_r16_Semi_Static_Mode1 = 0
	DAPS_UplinkPowerConfig_r16_UplinkPowerSharingDAPS_Mode_r16_Semi_Static_Mode2 = 1
	DAPS_UplinkPowerConfig_r16_UplinkPowerSharingDAPS_Mode_r16_Dynamic           = 2
)

var dAPSUplinkPowerConfigR16UplinkPowerSharingDAPSModeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DAPS_UplinkPowerConfig_r16_UplinkPowerSharingDAPS_Mode_r16_Semi_Static_Mode1, DAPS_UplinkPowerConfig_r16_UplinkPowerSharingDAPS_Mode_r16_Semi_Static_Mode2, DAPS_UplinkPowerConfig_r16_UplinkPowerSharingDAPS_Mode_r16_Dynamic},
}

type DAPS_UplinkPowerConfig_r16 struct {
	P_DAPS_Source_r16               P_Max
	P_DAPS_Target_r16               P_Max
	UplinkPowerSharingDAPS_Mode_r16 int64
}

func (ie *DAPS_UplinkPowerConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dAPSUplinkPowerConfigR16Constraints)
	_ = seq

	// 1. p-DAPS-Source-r16: ref
	{
		if err := ie.P_DAPS_Source_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. p-DAPS-Target-r16: ref
	{
		if err := ie.P_DAPS_Target_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. uplinkPowerSharingDAPS-Mode-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.UplinkPowerSharingDAPS_Mode_r16, dAPSUplinkPowerConfigR16UplinkPowerSharingDAPSModeR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DAPS_UplinkPowerConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dAPSUplinkPowerConfigR16Constraints)
	_ = seq

	// 1. p-DAPS-Source-r16: ref
	{
		if err := ie.P_DAPS_Source_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. p-DAPS-Target-r16: ref
	{
		if err := ie.P_DAPS_Target_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. uplinkPowerSharingDAPS-Mode-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(dAPSUplinkPowerConfigR16UplinkPowerSharingDAPSModeR16Constraints)
		if err != nil {
			return err
		}
		ie.UplinkPowerSharingDAPS_Mode_r16 = v2
	}

	return nil
}
