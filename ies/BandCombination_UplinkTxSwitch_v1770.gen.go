// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-UplinkTxSwitch-v1770 (line 16903).

var bandCombinationUplinkTxSwitchV1770Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-v1770", Optional: true},
	},
}

type BandCombination_UplinkTxSwitch_v1770 struct {
	BandCombination_v1770 *BandCombination_v1770
}

func (ie *BandCombination_UplinkTxSwitch_v1770) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchV1770Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_v1770 != nil}); err != nil {
		return err
	}

	// 2. bandCombination-v1770: ref
	{
		if ie.BandCombination_v1770 != nil {
			if err := ie.BandCombination_v1770.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1770) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchV1770Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandCombination-v1770: ref
	{
		if seq.IsComponentPresent(0) {
			ie.BandCombination_v1770 = new(BandCombination_v1770)
			if err := ie.BandCombination_v1770.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
