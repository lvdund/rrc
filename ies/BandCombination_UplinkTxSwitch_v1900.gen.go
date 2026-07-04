// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-UplinkTxSwitch-v1900 (line 16944).

var bandCombinationUplinkTxSwitchV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-v1900", Optional: true},
		{Name: "simultaneousSRS-UplinkTxSwitch-r19", Optional: true},
		{Name: "supportedBandPairListNR-v1900", Optional: true},
		{Name: "uplinkTxSwitchingBandParametersList-v1900", Optional: true},
	},
}

const (
	BandCombination_UplinkTxSwitch_v1900_SimultaneousSRS_UplinkTxSwitch_r19_Max = 0
	BandCombination_UplinkTxSwitch_v1900_SimultaneousSRS_UplinkTxSwitch_r19_Sum = 1
)

var bandCombinationUplinkTxSwitchV1900SimultaneousSRSUplinkTxSwitchR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandCombination_UplinkTxSwitch_v1900_SimultaneousSRS_UplinkTxSwitch_r19_Max, BandCombination_UplinkTxSwitch_v1900_SimultaneousSRS_UplinkTxSwitch_r19_Sum},
}

var bandCombinationUplinkTxSwitchV1900SupportedBandPairListNRV1900Constraints = per.SizeRange(1, common.MaxULTxSwitchingBandPairs)

var bandCombinationUplinkTxSwitchV1900UplinkTxSwitchingBandParametersListV1900Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

type BandCombination_UplinkTxSwitch_v1900 struct {
	BandCombination_v1900                     *BandCombination_v1900
	SimultaneousSRS_UplinkTxSwitch_r19        *int64
	SupportedBandPairListNR_v1900             []ULTxSwitchingBandPair_v1900
	UplinkTxSwitchingBandParametersList_v1900 []UplinkTxSwitchingBandParameters_v1900
}

func (ie *BandCombination_UplinkTxSwitch_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_v1900 != nil, ie.SimultaneousSRS_UplinkTxSwitch_r19 != nil, ie.SupportedBandPairListNR_v1900 != nil, ie.UplinkTxSwitchingBandParametersList_v1900 != nil}); err != nil {
		return err
	}

	// 2. bandCombination-v1900: ref
	{
		if ie.BandCombination_v1900 != nil {
			if err := ie.BandCombination_v1900.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. simultaneousSRS-UplinkTxSwitch-r19: enumerated
	{
		if ie.SimultaneousSRS_UplinkTxSwitch_r19 != nil {
			if err := e.EncodeEnumerated(*ie.SimultaneousSRS_UplinkTxSwitch_r19, bandCombinationUplinkTxSwitchV1900SimultaneousSRSUplinkTxSwitchR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. supportedBandPairListNR-v1900: sequence-of
	{
		if ie.SupportedBandPairListNR_v1900 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationUplinkTxSwitchV1900SupportedBandPairListNRV1900Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SupportedBandPairListNR_v1900))); err != nil {
				return err
			}
			for i := range ie.SupportedBandPairListNR_v1900 {
				if err := ie.SupportedBandPairListNR_v1900[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. uplinkTxSwitchingBandParametersList-v1900: sequence-of
	{
		if ie.UplinkTxSwitchingBandParametersList_v1900 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationUplinkTxSwitchV1900UplinkTxSwitchingBandParametersListV1900Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.UplinkTxSwitchingBandParametersList_v1900))); err != nil {
				return err
			}
			for i := range ie.UplinkTxSwitchingBandParametersList_v1900 {
				if err := ie.UplinkTxSwitchingBandParametersList_v1900[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandCombination-v1900: ref
	{
		if seq.IsComponentPresent(0) {
			ie.BandCombination_v1900 = new(BandCombination_v1900)
			if err := ie.BandCombination_v1900.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. simultaneousSRS-UplinkTxSwitch-r19: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(bandCombinationUplinkTxSwitchV1900SimultaneousSRSUplinkTxSwitchR19Constraints)
			if err != nil {
				return err
			}
			ie.SimultaneousSRS_UplinkTxSwitch_r19 = &idx
		}
	}

	// 4. supportedBandPairListNR-v1900: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationUplinkTxSwitchV1900SupportedBandPairListNRV1900Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SupportedBandPairListNR_v1900 = make([]ULTxSwitchingBandPair_v1900, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SupportedBandPairListNR_v1900[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. uplinkTxSwitchingBandParametersList-v1900: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationUplinkTxSwitchV1900UplinkTxSwitchingBandParametersListV1900Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.UplinkTxSwitchingBandParametersList_v1900 = make([]UplinkTxSwitchingBandParameters_v1900, n)
			for i := int64(0); i < n; i++ {
				if err := ie.UplinkTxSwitchingBandParametersList_v1900[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
