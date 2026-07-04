// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-CapabilityRestrictedBandParameters-r18 (line 2655).

var mUSIMCapabilityRestrictedBandParametersR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "musim-bandEntryIndex-r18"},
		{Name: "musim-CapabilityRestricted-r18"},
	},
}

var mUSIMCapabilityRestrictedBandParametersR18MusimCapabilityRestrictedR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "musim-MIMO-Layers-DL-r18", Optional: true},
		{Name: "musim-MIMO-Layers-UL-r18", Optional: true},
		{Name: "musim-SupportedBandwidth-DL-r18", Optional: true},
		{Name: "musim-SupportedBandwidth-UL-r18", Optional: true},
	},
}

type MUSIM_CapabilityRestrictedBandParameters_r18 struct {
	Musim_BandEntryIndex_r18       MUSIM_BandEntryIndex_r18
	Musim_CapabilityRestricted_r18 struct {
		Musim_MIMO_Layers_DL_r18        *int64
		Musim_MIMO_Layers_UL_r18        *int64
		Musim_SupportedBandwidth_DL_r18 *SupportedBandwidth_v1700
		Musim_SupportedBandwidth_UL_r18 *SupportedBandwidth_v1700
	}
}

func (ie *MUSIM_CapabilityRestrictedBandParameters_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMCapabilityRestrictedBandParametersR18Constraints)
	_ = seq

	// 1. musim-bandEntryIndex-r18: ref
	{
		if err := ie.Musim_BandEntryIndex_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. musim-CapabilityRestricted-r18: sequence
	{
		{
			c := &ie.Musim_CapabilityRestricted_r18
			mUSIMCapabilityRestrictedBandParametersR18MusimCapabilityRestrictedR18Seq := e.NewSequenceEncoder(mUSIMCapabilityRestrictedBandParametersR18MusimCapabilityRestrictedR18Constraints)
			if err := mUSIMCapabilityRestrictedBandParametersR18MusimCapabilityRestrictedR18Seq.EncodePreamble([]bool{c.Musim_MIMO_Layers_DL_r18 != nil, c.Musim_MIMO_Layers_UL_r18 != nil, c.Musim_SupportedBandwidth_DL_r18 != nil, c.Musim_SupportedBandwidth_UL_r18 != nil}); err != nil {
				return err
			}
			if c.Musim_MIMO_Layers_DL_r18 != nil {
				if err := e.EncodeInteger((*c.Musim_MIMO_Layers_DL_r18), per.Constrained(1, 8)); err != nil {
					return err
				}
			}
			if c.Musim_MIMO_Layers_UL_r18 != nil {
				if err := e.EncodeInteger((*c.Musim_MIMO_Layers_UL_r18), per.Constrained(1, 4)); err != nil {
					return err
				}
			}
			if c.Musim_SupportedBandwidth_DL_r18 != nil {
				if err := c.Musim_SupportedBandwidth_DL_r18.Encode(e); err != nil {
					return err
				}
			}
			if c.Musim_SupportedBandwidth_UL_r18 != nil {
				if err := c.Musim_SupportedBandwidth_UL_r18.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MUSIM_CapabilityRestrictedBandParameters_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMCapabilityRestrictedBandParametersR18Constraints)
	_ = seq

	// 1. musim-bandEntryIndex-r18: ref
	{
		if err := ie.Musim_BandEntryIndex_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. musim-CapabilityRestricted-r18: sequence
	{
		{
			c := &ie.Musim_CapabilityRestricted_r18
			mUSIMCapabilityRestrictedBandParametersR18MusimCapabilityRestrictedR18Seq := d.NewSequenceDecoder(mUSIMCapabilityRestrictedBandParametersR18MusimCapabilityRestrictedR18Constraints)
			if err := mUSIMCapabilityRestrictedBandParametersR18MusimCapabilityRestrictedR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if mUSIMCapabilityRestrictedBandParametersR18MusimCapabilityRestrictedR18Seq.IsComponentPresent(0) {
				c.Musim_MIMO_Layers_DL_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*c.Musim_MIMO_Layers_DL_r18) = v
			}
			if mUSIMCapabilityRestrictedBandParametersR18MusimCapabilityRestrictedR18Seq.IsComponentPresent(1) {
				c.Musim_MIMO_Layers_UL_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				(*c.Musim_MIMO_Layers_UL_r18) = v
			}
			if mUSIMCapabilityRestrictedBandParametersR18MusimCapabilityRestrictedR18Seq.IsComponentPresent(2) {
				c.Musim_SupportedBandwidth_DL_r18 = new(SupportedBandwidth_v1700)
				if err := (*c.Musim_SupportedBandwidth_DL_r18).Decode(d); err != nil {
					return err
				}
			}
			if mUSIMCapabilityRestrictedBandParametersR18MusimCapabilityRestrictedR18Seq.IsComponentPresent(3) {
				c.Musim_SupportedBandwidth_UL_r18 = new(SupportedBandwidth_v1700)
				if err := (*c.Musim_SupportedBandwidth_UL_r18).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
