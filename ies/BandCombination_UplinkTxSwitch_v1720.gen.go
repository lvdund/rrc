// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-UplinkTxSwitch-v1720 (line 16886).

var bandCombinationUplinkTxSwitchV1720Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-v1720", Optional: true},
		{Name: "uplinkTxSwitching-OptionSupport2T2T-r17", Optional: true},
	},
}

const (
	BandCombination_UplinkTxSwitch_v1720_UplinkTxSwitching_OptionSupport2T2T_r17_SwitchedUL = 0
	BandCombination_UplinkTxSwitch_v1720_UplinkTxSwitching_OptionSupport2T2T_r17_DualUL     = 1
	BandCombination_UplinkTxSwitch_v1720_UplinkTxSwitching_OptionSupport2T2T_r17_Both       = 2
)

var bandCombinationUplinkTxSwitchV1720UplinkTxSwitchingOptionSupport2T2TR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandCombination_UplinkTxSwitch_v1720_UplinkTxSwitching_OptionSupport2T2T_r17_SwitchedUL, BandCombination_UplinkTxSwitch_v1720_UplinkTxSwitching_OptionSupport2T2T_r17_DualUL, BandCombination_UplinkTxSwitch_v1720_UplinkTxSwitching_OptionSupport2T2T_r17_Both},
}

type BandCombination_UplinkTxSwitch_v1720 struct {
	BandCombination_v1720                   *BandCombination_v1720
	UplinkTxSwitching_OptionSupport2T2T_r17 *int64
}

func (ie *BandCombination_UplinkTxSwitch_v1720) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchV1720Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_v1720 != nil, ie.UplinkTxSwitching_OptionSupport2T2T_r17 != nil}); err != nil {
		return err
	}

	// 2. bandCombination-v1720: ref
	{
		if ie.BandCombination_v1720 != nil {
			if err := ie.BandCombination_v1720.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. uplinkTxSwitching-OptionSupport2T2T-r17: enumerated
	{
		if ie.UplinkTxSwitching_OptionSupport2T2T_r17 != nil {
			if err := e.EncodeEnumerated(*ie.UplinkTxSwitching_OptionSupport2T2T_r17, bandCombinationUplinkTxSwitchV1720UplinkTxSwitchingOptionSupport2T2TR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1720) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchV1720Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandCombination-v1720: ref
	{
		if seq.IsComponentPresent(0) {
			ie.BandCombination_v1720 = new(BandCombination_v1720)
			if err := ie.BandCombination_v1720.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. uplinkTxSwitching-OptionSupport2T2T-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(bandCombinationUplinkTxSwitchV1720UplinkTxSwitchingOptionSupport2T2TR17Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitching_OptionSupport2T2T_r17 = &idx
		}
	}

	return nil
}
