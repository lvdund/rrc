// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ULTxSwitchingBandPair-v1840 (line 16982).

var uLTxSwitchingBandPairV1840Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "configured1T1T-OnTwoBands-r18", Optional: true},
	},
}

const (
	ULTxSwitchingBandPair_v1840_Configured1T1T_OnTwoBands_r18_Supported = 0
)

var uLTxSwitchingBandPairV1840Configured1T1TOnTwoBandsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ULTxSwitchingBandPair_v1840_Configured1T1T_OnTwoBands_r18_Supported},
}

type ULTxSwitchingBandPair_v1840 struct {
	Configured1T1T_OnTwoBands_r18 *int64
}

func (ie *ULTxSwitchingBandPair_v1840) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLTxSwitchingBandPairV1840Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Configured1T1T_OnTwoBands_r18 != nil}); err != nil {
		return err
	}

	// 2. configured1T1T-OnTwoBands-r18: enumerated
	{
		if ie.Configured1T1T_OnTwoBands_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Configured1T1T_OnTwoBands_r18, uLTxSwitchingBandPairV1840Configured1T1TOnTwoBandsR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ULTxSwitchingBandPair_v1840) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLTxSwitchingBandPairV1840Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. configured1T1T-OnTwoBands-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uLTxSwitchingBandPairV1840Configured1T1TOnTwoBandsR18Constraints)
			if err != nil {
				return err
			}
			ie.Configured1T1T_OnTwoBands_r18 = &idx
		}
	}

	return nil
}
