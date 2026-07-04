// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CFR-ConfigMCCH-MTCH-r17 (line 28392).

var cFRConfigMCCHMTCHR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "locationAndBandwidthBroadcast-r17", Optional: true},
		{Name: "pdsch-ConfigMCCH-r17", Optional: true},
		{Name: "commonControlResourceSetExt-r17", Optional: true},
	},
}

type CFR_ConfigMCCH_MTCH_r17 struct {
	LocationAndBandwidthBroadcast_r17 *LocationAndBandwidthBroadcast_r17
	Pdsch_ConfigMCCH_r17              *PDSCH_ConfigBroadcast_r17
	CommonControlResourceSetExt_r17   *ControlResourceSet
}

func (ie *CFR_ConfigMCCH_MTCH_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cFRConfigMCCHMTCHR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LocationAndBandwidthBroadcast_r17 != nil, ie.Pdsch_ConfigMCCH_r17 != nil, ie.CommonControlResourceSetExt_r17 != nil}); err != nil {
		return err
	}

	// 2. locationAndBandwidthBroadcast-r17: ref
	{
		if ie.LocationAndBandwidthBroadcast_r17 != nil {
			if err := ie.LocationAndBandwidthBroadcast_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. pdsch-ConfigMCCH-r17: ref
	{
		if ie.Pdsch_ConfigMCCH_r17 != nil {
			if err := ie.Pdsch_ConfigMCCH_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. commonControlResourceSetExt-r17: ref
	{
		if ie.CommonControlResourceSetExt_r17 != nil {
			if err := ie.CommonControlResourceSetExt_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CFR_ConfigMCCH_MTCH_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cFRConfigMCCHMTCHR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. locationAndBandwidthBroadcast-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.LocationAndBandwidthBroadcast_r17 = new(LocationAndBandwidthBroadcast_r17)
			if err := ie.LocationAndBandwidthBroadcast_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. pdsch-ConfigMCCH-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Pdsch_ConfigMCCH_r17 = new(PDSCH_ConfigBroadcast_r17)
			if err := ie.Pdsch_ConfigMCCH_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. commonControlResourceSetExt-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.CommonControlResourceSetExt_r17 = new(ControlResourceSet)
			if err := ie.CommonControlResourceSetExt_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
