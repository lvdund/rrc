// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-UplinkTxSwitch-v1670 (line 16859).

var bandCombinationUplinkTxSwitchV1670Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-v15g0", Optional: true},
	},
}

type BandCombination_UplinkTxSwitch_v1670 struct {
	BandCombination_V15g0 *BandCombination_V15g0
}

func (ie *BandCombination_UplinkTxSwitch_v1670) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchV1670Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_V15g0 != nil}); err != nil {
		return err
	}

	// 2. bandCombination-v15g0: ref
	{
		if ie.BandCombination_V15g0 != nil {
			if err := ie.BandCombination_V15g0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1670) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchV1670Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandCombination-v15g0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.BandCombination_V15g0 = new(BandCombination_V15g0)
			if err := ie.BandCombination_V15g0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
