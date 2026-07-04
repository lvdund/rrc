// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-UplinkTxSwitch-v1640 (line 16851).

var bandCombinationUplinkTxSwitchV1640Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-v1640", Optional: true},
	},
}

type BandCombination_UplinkTxSwitch_v1640 struct {
	BandCombination_v1640 *BandCombination_v1640
}

func (ie *BandCombination_UplinkTxSwitch_v1640) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchV1640Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_v1640 != nil}); err != nil {
		return err
	}

	// 2. bandCombination-v1640: ref
	{
		if ie.BandCombination_v1640 != nil {
			if err := ie.BandCombination_v1640.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1640) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchV1640Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandCombination-v1640: ref
	{
		if seq.IsComponentPresent(0) {
			ie.BandCombination_v1640 = new(BandCombination_v1640)
			if err := ie.BandCombination_v1640.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
