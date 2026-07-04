// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NR-MultiBandInfo (line 10172).

var nRMultiBandInfoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "freqBandIndicatorNR", Optional: true},
		{Name: "nr-NS-PmaxList", Optional: true},
	},
}

type NR_MultiBandInfo struct {
	FreqBandIndicatorNR *FreqBandIndicatorNR
	Nr_NS_PmaxList      *NR_NS_PmaxList
}

func (ie *NR_MultiBandInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRMultiBandInfoConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FreqBandIndicatorNR != nil, ie.Nr_NS_PmaxList != nil}); err != nil {
		return err
	}

	// 2. freqBandIndicatorNR: ref
	{
		if ie.FreqBandIndicatorNR != nil {
			if err := ie.FreqBandIndicatorNR.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. nr-NS-PmaxList: ref
	{
		if ie.Nr_NS_PmaxList != nil {
			if err := ie.Nr_NS_PmaxList.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NR_MultiBandInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRMultiBandInfoConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. freqBandIndicatorNR: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FreqBandIndicatorNR = new(FreqBandIndicatorNR)
			if err := ie.FreqBandIndicatorNR.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nr-NS-PmaxList: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Nr_NS_PmaxList = new(NR_NS_PmaxList)
			if err := ie.Nr_NS_PmaxList.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
