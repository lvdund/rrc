// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandParameters-v1780 (line 17086).

var bandParametersV1780Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-BandwidthClassDL-NR-r17", Optional: true},
		{Name: "ca-BandwidthClassUL-NR-r17", Optional: true},
		{Name: "supportedAggBW-FR2-r17", Optional: true},
	},
}

var bandParametersV1780SupportedAggBWFR2R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedAggBW-DL-r17", Optional: true},
		{Name: "supportedAggBW-UL-r17", Optional: true},
	},
}

type BandParameters_v1780 struct {
	Ca_BandwidthClassDL_NR_r17 *CA_BandwidthClassNR_r17
	Ca_BandwidthClassUL_NR_r17 *CA_BandwidthClassNR_r17
	SupportedAggBW_FR2_r17     *struct {
		SupportedAggBW_DL_r17 *SupportedAggBandwidth_r17
		SupportedAggBW_UL_r17 *SupportedAggBandwidth_r17
	}
}

func (ie *BandParameters_v1780) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandParametersV1780Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_BandwidthClassDL_NR_r17 != nil, ie.Ca_BandwidthClassUL_NR_r17 != nil, ie.SupportedAggBW_FR2_r17 != nil}); err != nil {
		return err
	}

	// 2. ca-BandwidthClassDL-NR-r17: ref
	{
		if ie.Ca_BandwidthClassDL_NR_r17 != nil {
			if err := ie.Ca_BandwidthClassDL_NR_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ca-BandwidthClassUL-NR-r17: ref
	{
		if ie.Ca_BandwidthClassUL_NR_r17 != nil {
			if err := ie.Ca_BandwidthClassUL_NR_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. supportedAggBW-FR2-r17: sequence
	{
		if ie.SupportedAggBW_FR2_r17 != nil {
			c := ie.SupportedAggBW_FR2_r17
			bandParametersV1780SupportedAggBWFR2R17Seq := e.NewSequenceEncoder(bandParametersV1780SupportedAggBWFR2R17Constraints)
			if err := bandParametersV1780SupportedAggBWFR2R17Seq.EncodePreamble([]bool{c.SupportedAggBW_DL_r17 != nil, c.SupportedAggBW_UL_r17 != nil}); err != nil {
				return err
			}
			if c.SupportedAggBW_DL_r17 != nil {
				if err := c.SupportedAggBW_DL_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.SupportedAggBW_UL_r17 != nil {
				if err := c.SupportedAggBW_UL_r17.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *BandParameters_v1780) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandParametersV1780Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-BandwidthClassDL-NR-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_BandwidthClassDL_NR_r17 = new(CA_BandwidthClassNR_r17)
			if err := ie.Ca_BandwidthClassDL_NR_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ca-BandwidthClassUL-NR-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_BandwidthClassUL_NR_r17 = new(CA_BandwidthClassNR_r17)
			if err := ie.Ca_BandwidthClassUL_NR_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. supportedAggBW-FR2-r17: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.SupportedAggBW_FR2_r17 = &struct {
				SupportedAggBW_DL_r17 *SupportedAggBandwidth_r17
				SupportedAggBW_UL_r17 *SupportedAggBandwidth_r17
			}{}
			c := ie.SupportedAggBW_FR2_r17
			bandParametersV1780SupportedAggBWFR2R17Seq := d.NewSequenceDecoder(bandParametersV1780SupportedAggBWFR2R17Constraints)
			if err := bandParametersV1780SupportedAggBWFR2R17Seq.DecodePreamble(); err != nil {
				return err
			}
			if bandParametersV1780SupportedAggBWFR2R17Seq.IsComponentPresent(0) {
				c.SupportedAggBW_DL_r17 = new(SupportedAggBandwidth_r17)
				if err := (*c.SupportedAggBW_DL_r17).Decode(d); err != nil {
					return err
				}
			}
			if bandParametersV1780SupportedAggBWFR2R17Seq.IsComponentPresent(1) {
				c.SupportedAggBW_UL_r17 = new(SupportedAggBandwidth_r17)
				if err := (*c.SupportedAggBW_UL_r17).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
