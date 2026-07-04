// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-DiscConfigCommon-r17 (line 4339).

var sLDiscConfigCommonR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RelayUE-ConfigCommon-r17"},
		{Name: "sl-RemoteUE-ConfigCommon-r17"},
	},
}

type SL_DiscConfigCommon_r17 struct {
	Sl_RelayUE_ConfigCommon_r17  SL_RelayUE_Config_r17
	Sl_RemoteUE_ConfigCommon_r17 SL_RemoteUE_Config_r17
}

func (ie *SL_DiscConfigCommon_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLDiscConfigCommonR17Constraints)
	_ = seq

	// 1. sl-RelayUE-ConfigCommon-r17: ref
	{
		if err := ie.Sl_RelayUE_ConfigCommon_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. sl-RemoteUE-ConfigCommon-r17: ref
	{
		if err := ie.Sl_RemoteUE_ConfigCommon_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_DiscConfigCommon_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLDiscConfigCommonR17Constraints)
	_ = seq

	// 1. sl-RelayUE-ConfigCommon-r17: ref
	{
		if err := ie.Sl_RelayUE_ConfigCommon_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. sl-RemoteUE-ConfigCommon-r17: ref
	{
		if err := ie.Sl_RemoteUE_ConfigCommon_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
