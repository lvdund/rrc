// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-UplinkTxSwitch-v1700 (line 16879).

var bandCombinationUplinkTxSwitchV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-v1700", Optional: true},
		{Name: "supportedBandPairListNR-v1700", Optional: true},
		{Name: "uplinkTxSwitchingBandParametersList-v1700", Optional: true},
	},
}

var bandCombinationUplinkTxSwitchV1700SupportedBandPairListNRV1700Constraints = per.SizeRange(1, common.MaxULTxSwitchingBandPairs)

var bandCombinationUplinkTxSwitchV1700UplinkTxSwitchingBandParametersListV1700Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

type BandCombination_UplinkTxSwitch_v1700 struct {
	BandCombination_v1700                     *BandCombination_v1700
	SupportedBandPairListNR_v1700             []ULTxSwitchingBandPair_v1700
	UplinkTxSwitchingBandParametersList_v1700 []UplinkTxSwitchingBandParameters_v1700
}

func (ie *BandCombination_UplinkTxSwitch_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_v1700 != nil, ie.SupportedBandPairListNR_v1700 != nil, ie.UplinkTxSwitchingBandParametersList_v1700 != nil}); err != nil {
		return err
	}

	// 2. bandCombination-v1700: ref
	{
		if ie.BandCombination_v1700 != nil {
			if err := ie.BandCombination_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. supportedBandPairListNR-v1700: sequence-of
	{
		if ie.SupportedBandPairListNR_v1700 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationUplinkTxSwitchV1700SupportedBandPairListNRV1700Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SupportedBandPairListNR_v1700))); err != nil {
				return err
			}
			for i := range ie.SupportedBandPairListNR_v1700 {
				if err := ie.SupportedBandPairListNR_v1700[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. uplinkTxSwitchingBandParametersList-v1700: sequence-of
	{
		if ie.UplinkTxSwitchingBandParametersList_v1700 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationUplinkTxSwitchV1700UplinkTxSwitchingBandParametersListV1700Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.UplinkTxSwitchingBandParametersList_v1700))); err != nil {
				return err
			}
			for i := range ie.UplinkTxSwitchingBandParametersList_v1700 {
				if err := ie.UplinkTxSwitchingBandParametersList_v1700[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandCombination-v1700: ref
	{
		if seq.IsComponentPresent(0) {
			ie.BandCombination_v1700 = new(BandCombination_v1700)
			if err := ie.BandCombination_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. supportedBandPairListNR-v1700: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationUplinkTxSwitchV1700SupportedBandPairListNRV1700Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SupportedBandPairListNR_v1700 = make([]ULTxSwitchingBandPair_v1700, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SupportedBandPairListNR_v1700[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. uplinkTxSwitchingBandParametersList-v1700: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationUplinkTxSwitchV1700UplinkTxSwitchingBandParametersListV1700Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.UplinkTxSwitchingBandParametersList_v1700 = make([]UplinkTxSwitchingBandParameters_v1700, n)
			for i := int64(0); i < n; i++ {
				if err := ie.UplinkTxSwitchingBandParametersList_v1700[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
