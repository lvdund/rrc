// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-MultiBandInfo (line 26160).

var eUTRAMultiBandInfoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "eutra-FreqBandIndicator"},
		{Name: "eutra-NS-PmaxList", Optional: true},
	},
}

type EUTRA_MultiBandInfo struct {
	Eutra_FreqBandIndicator FreqBandIndicatorEUTRA
	Eutra_NS_PmaxList       *EUTRA_NS_PmaxList
}

func (ie *EUTRA_MultiBandInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRAMultiBandInfoConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Eutra_NS_PmaxList != nil}); err != nil {
		return err
	}

	// 2. eutra-FreqBandIndicator: ref
	{
		if err := ie.Eutra_FreqBandIndicator.Encode(e); err != nil {
			return err
		}
	}

	// 3. eutra-NS-PmaxList: ref
	{
		if ie.Eutra_NS_PmaxList != nil {
			if err := ie.Eutra_NS_PmaxList.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *EUTRA_MultiBandInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRAMultiBandInfoConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. eutra-FreqBandIndicator: ref
	{
		if err := ie.Eutra_FreqBandIndicator.Decode(d); err != nil {
			return err
		}
	}

	// 3. eutra-NS-PmaxList: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Eutra_NS_PmaxList = new(EUTRA_NS_PmaxList)
			if err := ie.Eutra_NS_PmaxList.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
