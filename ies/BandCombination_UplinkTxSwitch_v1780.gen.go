// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-UplinkTxSwitch-v1780 (line 16907).

var bandCombinationUplinkTxSwitchV1780Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-v1780", Optional: true},
	},
}

type BandCombination_UplinkTxSwitch_v1780 struct {
	BandCombination_v1780 *BandCombination_v1780
}

func (ie *BandCombination_UplinkTxSwitch_v1780) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchV1780Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_v1780 != nil}); err != nil {
		return err
	}

	// 2. bandCombination-v1780: ref
	{
		if ie.BandCombination_v1780 != nil {
			if err := ie.BandCombination_v1780.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1780) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchV1780Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandCombination-v1780: ref
	{
		if seq.IsComponentPresent(0) {
			ie.BandCombination_v1780 = new(BandCombination_v1780)
			if err := ie.BandCombination_v1780.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
