// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-DiscConfigCommon-v1840 (line 4349).

var sLDiscConfigCommonV1840Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RelayUE-ConfigCommonU2U-v1840"},
		{Name: "sl-RemoteUE-ConfigCommonU2U-v1840"},
	},
}

type SL_DiscConfigCommon_v1840 struct {
	Sl_RelayUE_ConfigCommonU2U_v1840  SL_RelayUE_ConfigU2U_v1840
	Sl_RemoteUE_ConfigCommonU2U_v1840 SL_RemoteUE_ConfigU2U_v1830
}

func (ie *SL_DiscConfigCommon_v1840) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLDiscConfigCommonV1840Constraints)
	_ = seq

	// 1. sl-RelayUE-ConfigCommonU2U-v1840: ref
	{
		if err := ie.Sl_RelayUE_ConfigCommonU2U_v1840.Encode(e); err != nil {
			return err
		}
	}

	// 2. sl-RemoteUE-ConfigCommonU2U-v1840: ref
	{
		if err := ie.Sl_RemoteUE_ConfigCommonU2U_v1840.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_DiscConfigCommon_v1840) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLDiscConfigCommonV1840Constraints)
	_ = seq

	// 1. sl-RelayUE-ConfigCommonU2U-v1840: ref
	{
		if err := ie.Sl_RelayUE_ConfigCommonU2U_v1840.Decode(d); err != nil {
			return err
		}
	}

	// 2. sl-RemoteUE-ConfigCommonU2U-v1840: ref
	{
		if err := ie.Sl_RemoteUE_ConfigCommonU2U_v1840.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
