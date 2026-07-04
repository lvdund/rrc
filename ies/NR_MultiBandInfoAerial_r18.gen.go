// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NR-MultiBandInfoAerial-r18 (line 10185).

var nRMultiBandInfoAerialR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "freqBandIndicatorNR-Aerial-r18", Optional: true},
		{Name: "nr-NS-PmaxListAerial-r18", Optional: true},
	},
}

type NR_MultiBandInfoAerial_r18 struct {
	FreqBandIndicatorNR_Aerial_r18 *FreqBandIndicatorNR
	Nr_NS_PmaxListAerial_r18       *NR_NS_PmaxListAerial_r18
}

func (ie *NR_MultiBandInfoAerial_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRMultiBandInfoAerialR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FreqBandIndicatorNR_Aerial_r18 != nil, ie.Nr_NS_PmaxListAerial_r18 != nil}); err != nil {
		return err
	}

	// 2. freqBandIndicatorNR-Aerial-r18: ref
	{
		if ie.FreqBandIndicatorNR_Aerial_r18 != nil {
			if err := ie.FreqBandIndicatorNR_Aerial_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. nr-NS-PmaxListAerial-r18: ref
	{
		if ie.Nr_NS_PmaxListAerial_r18 != nil {
			if err := ie.Nr_NS_PmaxListAerial_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NR_MultiBandInfoAerial_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRMultiBandInfoAerialR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. freqBandIndicatorNR-Aerial-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FreqBandIndicatorNR_Aerial_r18 = new(FreqBandIndicatorNR)
			if err := ie.FreqBandIndicatorNR_Aerial_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nr-NS-PmaxListAerial-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Nr_NS_PmaxListAerial_r18 = new(NR_NS_PmaxListAerial_r18)
			if err := ie.Nr_NS_PmaxListAerial_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
