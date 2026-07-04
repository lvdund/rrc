// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CFR-ConfigMulticast-r17 (line 5897).

var cFRConfigMulticastR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "locationAndBandwidthMulticast-r17", Optional: true},
		{Name: "pdcch-ConfigMulticast-r17", Optional: true},
		{Name: "pdsch-ConfigMulticast-r17", Optional: true},
		{Name: "sps-ConfigMulticastToAddModList-r17", Optional: true},
		{Name: "sps-ConfigMulticastToReleaseList-r17", Optional: true},
	},
}

var cFRConfigMulticastR17LocationAndBandwidthMulticastR17Constraints = per.Constrained(0, 37949)

type CFR_ConfigMulticast_r17 struct {
	LocationAndBandwidthMulticast_r17    *int64
	Pdcch_ConfigMulticast_r17            *PDCCH_Config
	Pdsch_ConfigMulticast_r17            *PDSCH_Config
	Sps_ConfigMulticastToAddModList_r17  *SPS_ConfigMulticastToAddModList_r17
	Sps_ConfigMulticastToReleaseList_r17 *SPS_ConfigMulticastToReleaseList_r17
}

func (ie *CFR_ConfigMulticast_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cFRConfigMulticastR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LocationAndBandwidthMulticast_r17 != nil, ie.Pdcch_ConfigMulticast_r17 != nil, ie.Pdsch_ConfigMulticast_r17 != nil, ie.Sps_ConfigMulticastToAddModList_r17 != nil, ie.Sps_ConfigMulticastToReleaseList_r17 != nil}); err != nil {
		return err
	}

	// 2. locationAndBandwidthMulticast-r17: integer
	{
		if ie.LocationAndBandwidthMulticast_r17 != nil {
			if err := e.EncodeInteger(*ie.LocationAndBandwidthMulticast_r17, cFRConfigMulticastR17LocationAndBandwidthMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. pdcch-ConfigMulticast-r17: ref
	{
		if ie.Pdcch_ConfigMulticast_r17 != nil {
			if err := ie.Pdcch_ConfigMulticast_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. pdsch-ConfigMulticast-r17: ref
	{
		if ie.Pdsch_ConfigMulticast_r17 != nil {
			if err := ie.Pdsch_ConfigMulticast_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sps-ConfigMulticastToAddModList-r17: ref
	{
		if ie.Sps_ConfigMulticastToAddModList_r17 != nil {
			if err := ie.Sps_ConfigMulticastToAddModList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sps-ConfigMulticastToReleaseList-r17: ref
	{
		if ie.Sps_ConfigMulticastToReleaseList_r17 != nil {
			if err := ie.Sps_ConfigMulticastToReleaseList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CFR_ConfigMulticast_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cFRConfigMulticastR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. locationAndBandwidthMulticast-r17: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(cFRConfigMulticastR17LocationAndBandwidthMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.LocationAndBandwidthMulticast_r17 = &v
		}
	}

	// 3. pdcch-ConfigMulticast-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Pdcch_ConfigMulticast_r17 = new(PDCCH_Config)
			if err := ie.Pdcch_ConfigMulticast_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. pdsch-ConfigMulticast-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Pdsch_ConfigMulticast_r17 = new(PDSCH_Config)
			if err := ie.Pdsch_ConfigMulticast_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sps-ConfigMulticastToAddModList-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sps_ConfigMulticastToAddModList_r17 = new(SPS_ConfigMulticastToAddModList_r17)
			if err := ie.Sps_ConfigMulticastToAddModList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sps-ConfigMulticastToReleaseList-r17: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Sps_ConfigMulticastToReleaseList_r17 = new(SPS_ConfigMulticastToReleaseList_r17)
			if err := ie.Sps_ConfigMulticastToReleaseList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
