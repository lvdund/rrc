// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-UplinkTxSwitch-v1840 (line 16935).

var bandCombinationUplinkTxSwitchV1840Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-v1840", Optional: true},
		{Name: "supportedBandPairListNR-v1840", Optional: true},
	},
}

var bandCombinationUplinkTxSwitchV1840SupportedBandPairListNRV1840Constraints = per.SizeRange(1, common.MaxULTxSwitchingBandPairs)

type BandCombination_UplinkTxSwitch_v1840 struct {
	BandCombination_v1840         *BandCombination_v1840
	SupportedBandPairListNR_v1840 []ULTxSwitchingBandPair_v1840
}

func (ie *BandCombination_UplinkTxSwitch_v1840) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchV1840Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_v1840 != nil, ie.SupportedBandPairListNR_v1840 != nil}); err != nil {
		return err
	}

	// 2. bandCombination-v1840: ref
	{
		if ie.BandCombination_v1840 != nil {
			if err := ie.BandCombination_v1840.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. supportedBandPairListNR-v1840: sequence-of
	{
		if ie.SupportedBandPairListNR_v1840 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationUplinkTxSwitchV1840SupportedBandPairListNRV1840Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SupportedBandPairListNR_v1840))); err != nil {
				return err
			}
			for i := range ie.SupportedBandPairListNR_v1840 {
				if err := ie.SupportedBandPairListNR_v1840[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1840) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchV1840Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandCombination-v1840: ref
	{
		if seq.IsComponentPresent(0) {
			ie.BandCombination_v1840 = new(BandCombination_v1840)
			if err := ie.BandCombination_v1840.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. supportedBandPairListNR-v1840: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationUplinkTxSwitchV1840SupportedBandPairListNRV1840Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SupportedBandPairListNR_v1840 = make([]ULTxSwitchingBandPair_v1840, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SupportedBandPairListNR_v1840[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
