// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-UplinkTxSwitch-v1790 (line 16911).

var bandCombinationUplinkTxSwitchV1790Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-v1790", Optional: true},
	},
}

type BandCombination_UplinkTxSwitch_v1790 struct {
	BandCombination_v1790 *BandCombination_v1790
}

func (ie *BandCombination_UplinkTxSwitch_v1790) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchV1790Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_v1790 != nil}); err != nil {
		return err
	}

	// 2. bandCombination-v1790: ref
	{
		if ie.BandCombination_v1790 != nil {
			if err := ie.BandCombination_v1790.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1790) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchV1790Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandCombination-v1790: ref
	{
		if seq.IsComponentPresent(0) {
			ie.BandCombination_v1790 = new(BandCombination_v1790)
			if err := ie.BandCombination_v1790.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
