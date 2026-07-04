// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-v1590 (line 16683).

var bandCombinationV1590Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandwidthCombinationSetIntraENDC", Optional: true},
		{Name: "mrdc-Parameters-v1590"},
	},
}

var bandCombinationV1590SupportedBandwidthCombinationSetIntraENDCConstraints = per.SizeRange(1, 32)

type BandCombination_v1590 struct {
	SupportedBandwidthCombinationSetIntraENDC *per.BitString
	Mrdc_Parameters_v1590                     MRDC_Parameters_v1590
}

func (ie *BandCombination_v1590) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1590Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandwidthCombinationSetIntraENDC != nil}); err != nil {
		return err
	}

	// 2. supportedBandwidthCombinationSetIntraENDC: bit-string
	{
		if ie.SupportedBandwidthCombinationSetIntraENDC != nil {
			if err := e.EncodeBitString(*ie.SupportedBandwidthCombinationSetIntraENDC, bandCombinationV1590SupportedBandwidthCombinationSetIntraENDCConstraints); err != nil {
				return err
			}
		}
	}

	// 3. mrdc-Parameters-v1590: ref
	{
		if err := ie.Mrdc_Parameters_v1590.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BandCombination_v1590) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1590Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedBandwidthCombinationSetIntraENDC: bit-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeBitString(bandCombinationV1590SupportedBandwidthCombinationSetIntraENDCConstraints)
			if err != nil {
				return err
			}
			ie.SupportedBandwidthCombinationSetIntraENDC = &v
		}
	}

	// 3. mrdc-Parameters-v1590: ref
	{
		if err := ie.Mrdc_Parameters_v1590.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
