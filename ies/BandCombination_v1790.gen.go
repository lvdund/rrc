// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-v1790 (line 16787).

var bandCombinationV1790Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedIntraENDC-BandCombinationList-r17", Optional: true},
	},
}

var bandCombinationV1790SupportedIntraENDCBandCombinationListR17Constraints = per.SizeRange(1, common.MaxNrofIntraEndc_Components_r17)

type BandCombination_v1790 struct {
	SupportedIntraENDC_BandCombinationList_r17 []SupportedIntraENDC_BandCombination_r17
}

func (ie *BandCombination_v1790) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1790Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedIntraENDC_BandCombinationList_r17 != nil}); err != nil {
		return err
	}

	// 2. supportedIntraENDC-BandCombinationList-r17: sequence-of
	{
		if ie.SupportedIntraENDC_BandCombinationList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationV1790SupportedIntraENDCBandCombinationListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SupportedIntraENDC_BandCombinationList_r17))); err != nil {
				return err
			}
			for i := range ie.SupportedIntraENDC_BandCombinationList_r17 {
				if err := ie.SupportedIntraENDC_BandCombinationList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1790) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1790Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedIntraENDC-BandCombinationList-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationV1790SupportedIntraENDCBandCombinationListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SupportedIntraENDC_BandCombinationList_r17 = make([]SupportedIntraENDC_BandCombination_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SupportedIntraENDC_BandCombinationList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
