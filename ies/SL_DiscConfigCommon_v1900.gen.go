// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-DiscConfigCommon-v1900 (line 4354).

var sLDiscConfigCommonV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RelayUE-ConfigCommonMH-r19"},
	},
}

type SL_DiscConfigCommon_v1900 struct {
	Sl_RelayUE_ConfigCommonMH_r19 SL_RelayUE_ConfigMH_r19
}

func (ie *SL_DiscConfigCommon_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLDiscConfigCommonV1900Constraints)
	_ = seq

	// 1. sl-RelayUE-ConfigCommonMH-r19: ref
	{
		if err := ie.Sl_RelayUE_ConfigCommonMH_r19.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_DiscConfigCommon_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLDiscConfigCommonV1900Constraints)
	_ = seq

	// 1. sl-RelayUE-ConfigCommonMH-r19: ref
	{
		if err := ie.Sl_RelayUE_ConfigCommonMH_r19.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
