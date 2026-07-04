// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-UplinkTxSwitch-v1860 (line 16940).

var bandCombinationUplinkTxSwitchV1860Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-v1860", Optional: true},
	},
}

type BandCombination_UplinkTxSwitch_v1860 struct {
	BandCombination_v1860 *BandCombination_v1860
}

func (ie *BandCombination_UplinkTxSwitch_v1860) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchV1860Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_v1860 != nil}); err != nil {
		return err
	}

	// 2. bandCombination-v1860: ref
	{
		if ie.BandCombination_v1860 != nil {
			if err := ie.BandCombination_v1860.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1860) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchV1860Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandCombination-v1860: ref
	{
		if seq.IsComponentPresent(0) {
			ie.BandCombination_v1860 = new(BandCombination_v1860)
			if err := ie.BandCombination_v1860.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
