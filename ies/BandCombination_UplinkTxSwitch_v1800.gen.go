// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-UplinkTxSwitch-v1800 (line 16919).

var bandCombinationUplinkTxSwitchV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-v1800", Optional: true},
		{Name: "supportedBandPairListNR-r18", Optional: true},
		{Name: "uplinkTxSwitchingMinimumSeparationTime-r18", Optional: true},
		{Name: "uplinkTxSwitchingAdditionalPeriodDualUL-List-r18", Optional: true},
		{Name: "switchingPeriodRestriction-r18", Optional: true},
	},
}

var bandCombinationUplinkTxSwitchV1800SupportedBandPairListNRR18Constraints = per.SizeRange(1, common.MaxULTxSwitchingBandPairs)

const (
	BandCombination_UplinkTxSwitch_v1800_UplinkTxSwitchingMinimumSeparationTime_r18_N0us   = 0
	BandCombination_UplinkTxSwitch_v1800_UplinkTxSwitchingMinimumSeparationTime_r18_N500us = 1
)

var bandCombinationUplinkTxSwitchV1800UplinkTxSwitchingMinimumSeparationTimeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandCombination_UplinkTxSwitch_v1800_UplinkTxSwitchingMinimumSeparationTime_r18_N0us, BandCombination_UplinkTxSwitch_v1800_UplinkTxSwitchingMinimumSeparationTime_r18_N500us},
}

var bandCombinationUplinkTxSwitchV1800UplinkTxSwitchingAdditionalPeriodDualULListR18Constraints = per.SizeRange(1, common.MaxULTxSwitchingBetweenBandPairs_r18)

const (
	BandCombination_UplinkTxSwitch_v1800_SwitchingPeriodRestriction_r18_True = 0
)

var bandCombinationUplinkTxSwitchV1800SwitchingPeriodRestrictionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandCombination_UplinkTxSwitch_v1800_SwitchingPeriodRestriction_r18_True},
}

type BandCombination_UplinkTxSwitch_v1800 struct {
	BandCombination_v1800                            *BandCombination_v1800
	SupportedBandPairListNR_r18                      []ULTxSwitchingBandPair_r18
	UplinkTxSwitchingMinimumSeparationTime_r18       *int64
	UplinkTxSwitchingAdditionalPeriodDualUL_List_r18 []UplinkTxSwitchingAdditionalPeriodDualUL_r18
	SwitchingPeriodRestriction_r18                   *int64
}

func (ie *BandCombination_UplinkTxSwitch_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_v1800 != nil, ie.SupportedBandPairListNR_r18 != nil, ie.UplinkTxSwitchingMinimumSeparationTime_r18 != nil, ie.UplinkTxSwitchingAdditionalPeriodDualUL_List_r18 != nil, ie.SwitchingPeriodRestriction_r18 != nil}); err != nil {
		return err
	}

	// 2. bandCombination-v1800: ref
	{
		if ie.BandCombination_v1800 != nil {
			if err := ie.BandCombination_v1800.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. supportedBandPairListNR-r18: sequence-of
	{
		if ie.SupportedBandPairListNR_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationUplinkTxSwitchV1800SupportedBandPairListNRR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SupportedBandPairListNR_r18))); err != nil {
				return err
			}
			for i := range ie.SupportedBandPairListNR_r18 {
				if err := ie.SupportedBandPairListNR_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. uplinkTxSwitchingMinimumSeparationTime-r18: enumerated
	{
		if ie.UplinkTxSwitchingMinimumSeparationTime_r18 != nil {
			if err := e.EncodeEnumerated(*ie.UplinkTxSwitchingMinimumSeparationTime_r18, bandCombinationUplinkTxSwitchV1800UplinkTxSwitchingMinimumSeparationTimeR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. uplinkTxSwitchingAdditionalPeriodDualUL-List-r18: sequence-of
	{
		if ie.UplinkTxSwitchingAdditionalPeriodDualUL_List_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationUplinkTxSwitchV1800UplinkTxSwitchingAdditionalPeriodDualULListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.UplinkTxSwitchingAdditionalPeriodDualUL_List_r18))); err != nil {
				return err
			}
			for i := range ie.UplinkTxSwitchingAdditionalPeriodDualUL_List_r18 {
				if err := ie.UplinkTxSwitchingAdditionalPeriodDualUL_List_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. switchingPeriodRestriction-r18: enumerated
	{
		if ie.SwitchingPeriodRestriction_r18 != nil {
			if err := e.EncodeEnumerated(*ie.SwitchingPeriodRestriction_r18, bandCombinationUplinkTxSwitchV1800SwitchingPeriodRestrictionR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandCombination-v1800: ref
	{
		if seq.IsComponentPresent(0) {
			ie.BandCombination_v1800 = new(BandCombination_v1800)
			if err := ie.BandCombination_v1800.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. supportedBandPairListNR-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationUplinkTxSwitchV1800SupportedBandPairListNRR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SupportedBandPairListNR_r18 = make([]ULTxSwitchingBandPair_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SupportedBandPairListNR_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. uplinkTxSwitchingMinimumSeparationTime-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(bandCombinationUplinkTxSwitchV1800UplinkTxSwitchingMinimumSeparationTimeR18Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitchingMinimumSeparationTime_r18 = &idx
		}
	}

	// 5. uplinkTxSwitchingAdditionalPeriodDualUL-List-r18: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationUplinkTxSwitchV1800UplinkTxSwitchingAdditionalPeriodDualULListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.UplinkTxSwitchingAdditionalPeriodDualUL_List_r18 = make([]UplinkTxSwitchingAdditionalPeriodDualUL_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.UplinkTxSwitchingAdditionalPeriodDualUL_List_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. switchingPeriodRestriction-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(bandCombinationUplinkTxSwitchV1800SwitchingPeriodRestrictionR18Constraints)
			if err != nil {
				return err
			}
			ie.SwitchingPeriodRestriction_r18 = &idx
		}
	}

	return nil
}
