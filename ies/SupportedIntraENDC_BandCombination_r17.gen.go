// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SupportedIntraENDC-BandCombination-r17 (line 17131).

var supportedIntraENDCBandCombinationR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandwidthCombinationSetIntraENDC-v1790", Optional: true},
		{Name: "mrdc-Parameters-v1790", Optional: true},
	},
}

var supportedIntraENDCBandCombinationR17SupportedBandwidthCombinationSetIntraENDCV1790Constraints = per.SizeRange(1, 32)

type SupportedIntraENDC_BandCombination_r17 struct {
	SupportedBandwidthCombinationSetIntraENDC_v1790 *per.BitString
	Mrdc_Parameters_v1790                           *MRDC_Parameters_v1790
}

func (ie *SupportedIntraENDC_BandCombination_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(supportedIntraENDCBandCombinationR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandwidthCombinationSetIntraENDC_v1790 != nil, ie.Mrdc_Parameters_v1790 != nil}); err != nil {
		return err
	}

	// 2. supportedBandwidthCombinationSetIntraENDC-v1790: bit-string
	{
		if ie.SupportedBandwidthCombinationSetIntraENDC_v1790 != nil {
			if err := e.EncodeBitString(*ie.SupportedBandwidthCombinationSetIntraENDC_v1790, supportedIntraENDCBandCombinationR17SupportedBandwidthCombinationSetIntraENDCV1790Constraints); err != nil {
				return err
			}
		}
	}

	// 3. mrdc-Parameters-v1790: ref
	{
		if ie.Mrdc_Parameters_v1790 != nil {
			if err := ie.Mrdc_Parameters_v1790.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SupportedIntraENDC_BandCombination_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(supportedIntraENDCBandCombinationR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedBandwidthCombinationSetIntraENDC-v1790: bit-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeBitString(supportedIntraENDCBandCombinationR17SupportedBandwidthCombinationSetIntraENDCV1790Constraints)
			if err != nil {
				return err
			}
			ie.SupportedBandwidthCombinationSetIntraENDC_v1790 = &v
		}
	}

	// 3. mrdc-Parameters-v1790: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Mrdc_Parameters_v1790 = new(MRDC_Parameters_v1790)
			if err := ie.Mrdc_Parameters_v1790.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
