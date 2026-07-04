// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ConfigDedicatedNR-v16k0 (line 26971).

var sLConfigDedicatedNRV16k0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PHY-MAC-RLC-Config-v16k0", Optional: true},
	},
}

type SL_ConfigDedicatedNR_V16k0 struct {
	Sl_PHY_MAC_RLC_Config_V16k0 *SL_PHY_MAC_RLC_Config_V16k0
}

func (ie *SL_ConfigDedicatedNR_V16k0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLConfigDedicatedNRV16k0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PHY_MAC_RLC_Config_V16k0 != nil}); err != nil {
		return err
	}

	// 2. sl-PHY-MAC-RLC-Config-v16k0: ref
	{
		if ie.Sl_PHY_MAC_RLC_Config_V16k0 != nil {
			if err := ie.Sl_PHY_MAC_RLC_Config_V16k0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_ConfigDedicatedNR_V16k0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLConfigDedicatedNRV16k0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-PHY-MAC-RLC-Config-v16k0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_PHY_MAC_RLC_Config_V16k0 = new(SL_PHY_MAC_RLC_Config_V16k0)
			if err := ie.Sl_PHY_MAC_RLC_Config_V16k0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
