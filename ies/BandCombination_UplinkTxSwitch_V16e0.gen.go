// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-UplinkTxSwitch-v16e0 (line 16871).

var bandCombinationUplinkTxSwitchV16e0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-v15n0", Optional: true},
	},
}

type BandCombination_UplinkTxSwitch_V16e0 struct {
	BandCombination_V15n0 *BandCombination_V15n0
}

func (ie *BandCombination_UplinkTxSwitch_V16e0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchV16e0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_V15n0 != nil}); err != nil {
		return err
	}

	// 2. bandCombination-v15n0: ref
	{
		if ie.BandCombination_V15n0 != nil {
			if err := ie.BandCombination_V15n0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_V16e0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchV16e0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandCombination-v15n0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.BandCombination_V15n0 = new(BandCombination_V15n0)
			if err := ie.BandCombination_V15n0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
