// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-DiscConfigCommon-v1800 (line 4344).

var sLDiscConfigCommonV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RelayUE-ConfigCommonU2U-r18"},
		{Name: "sl-RemoteUE-ConfigCommonU2U-r18"},
	},
}

type SL_DiscConfigCommon_v1800 struct {
	Sl_RelayUE_ConfigCommonU2U_r18  SL_RelayUE_ConfigU2U_r18
	Sl_RemoteUE_ConfigCommonU2U_r18 SL_RemoteUE_ConfigU2U_r18
}

func (ie *SL_DiscConfigCommon_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLDiscConfigCommonV1800Constraints)
	_ = seq

	// 1. sl-RelayUE-ConfigCommonU2U-r18: ref
	{
		if err := ie.Sl_RelayUE_ConfigCommonU2U_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. sl-RemoteUE-ConfigCommonU2U-r18: ref
	{
		if err := ie.Sl_RemoteUE_ConfigCommonU2U_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_DiscConfigCommon_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLDiscConfigCommonV1800Constraints)
	_ = seq

	// 1. sl-RelayUE-ConfigCommonU2U-r18: ref
	{
		if err := ie.Sl_RelayUE_ConfigCommonU2U_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. sl-RemoteUE-ConfigCommonU2U-r18: ref
	{
		if err := ie.Sl_RemoteUE_ConfigCommonU2U_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
