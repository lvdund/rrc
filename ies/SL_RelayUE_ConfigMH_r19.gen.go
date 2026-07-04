// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RelayUE-ConfigMH-r19 (line 27818).

var sLRelayUEConfigMHR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sd-RSRP-ThreshDiscConfigMH-r19"},
		{Name: "sd-hystMaxRelayMH-r19"},
	},
}

type SL_RelayUE_ConfigMH_r19 struct {
	Sd_RSRP_ThreshDiscConfigMH_r19 SL_RSRP_Range_r16
	Sd_HystMaxRelayMH_r19          Hysteresis
}

func (ie *SL_RelayUE_ConfigMH_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLRelayUEConfigMHR19Constraints)
	_ = seq

	// 1. sd-RSRP-ThreshDiscConfigMH-r19: ref
	{
		if err := ie.Sd_RSRP_ThreshDiscConfigMH_r19.Encode(e); err != nil {
			return err
		}
	}

	// 2. sd-hystMaxRelayMH-r19: ref
	{
		if err := ie.Sd_HystMaxRelayMH_r19.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_RelayUE_ConfigMH_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLRelayUEConfigMHR19Constraints)
	_ = seq

	// 1. sd-RSRP-ThreshDiscConfigMH-r19: ref
	{
		if err := ie.Sd_RSRP_ThreshDiscConfigMH_r19.Decode(d); err != nil {
			return err
		}
	}

	// 2. sd-hystMaxRelayMH-r19: ref
	{
		if err := ie.Sd_HystMaxRelayMH_r19.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
