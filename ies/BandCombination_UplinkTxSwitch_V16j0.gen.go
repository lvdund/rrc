// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-UplinkTxSwitch-v16j0 (line 16875).

var bandCombinationUplinkTxSwitchV16j0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-v16j0", Optional: true},
	},
}

type BandCombination_UplinkTxSwitch_V16j0 struct {
	BandCombination_V16j0 *BandCombination_V16j0
}

func (ie *BandCombination_UplinkTxSwitch_V16j0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchV16j0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_V16j0 != nil}); err != nil {
		return err
	}

	// 2. bandCombination-v16j0: ref
	{
		if ie.BandCombination_V16j0 != nil {
			if err := ie.BandCombination_V16j0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_V16j0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchV16j0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandCombination-v16j0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.BandCombination_V16j0 = new(BandCombination_V16j0)
			if err := ie.BandCombination_V16j0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
