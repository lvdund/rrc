// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ULTxSwitchingBandPair-v1700 (line 16961).

var uLTxSwitchingBandPairV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uplinkTxSwitchingPeriod2T2T-r17", Optional: true},
	},
}

const (
	ULTxSwitchingBandPair_v1700_UplinkTxSwitchingPeriod2T2T_r17_N35us  = 0
	ULTxSwitchingBandPair_v1700_UplinkTxSwitchingPeriod2T2T_r17_N140us = 1
	ULTxSwitchingBandPair_v1700_UplinkTxSwitchingPeriod2T2T_r17_N210us = 2
)

var uLTxSwitchingBandPairV1700UplinkTxSwitchingPeriod2T2TR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ULTxSwitchingBandPair_v1700_UplinkTxSwitchingPeriod2T2T_r17_N35us, ULTxSwitchingBandPair_v1700_UplinkTxSwitchingPeriod2T2T_r17_N140us, ULTxSwitchingBandPair_v1700_UplinkTxSwitchingPeriod2T2T_r17_N210us},
}

type ULTxSwitchingBandPair_v1700 struct {
	UplinkTxSwitchingPeriod2T2T_r17 *int64
}

func (ie *ULTxSwitchingBandPair_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLTxSwitchingBandPairV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.UplinkTxSwitchingPeriod2T2T_r17 != nil}); err != nil {
		return err
	}

	// 2. uplinkTxSwitchingPeriod2T2T-r17: enumerated
	{
		if ie.UplinkTxSwitchingPeriod2T2T_r17 != nil {
			if err := e.EncodeEnumerated(*ie.UplinkTxSwitchingPeriod2T2T_r17, uLTxSwitchingBandPairV1700UplinkTxSwitchingPeriod2T2TR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ULTxSwitchingBandPair_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLTxSwitchingBandPairV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. uplinkTxSwitchingPeriod2T2T-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uLTxSwitchingBandPairV1700UplinkTxSwitchingPeriod2T2TR17Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitchingPeriod2T2T_r17 = &idx
		}
	}

	return nil
}
