// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-UplinkTxSwitch-v1760 (line 16899).

var bandCombinationUplinkTxSwitchV1760Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandCombination-v1760", Optional: true},
	},
}

type BandCombination_UplinkTxSwitch_v1760 struct {
	BandCombination_v1760 *BandCombination_v1760
}

func (ie *BandCombination_UplinkTxSwitch_v1760) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationUplinkTxSwitchV1760Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandCombination_v1760 != nil}); err != nil {
		return err
	}

	// 2. bandCombination-v1760: ref
	{
		if ie.BandCombination_v1760 != nil {
			if err := ie.BandCombination_v1760.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1760) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationUplinkTxSwitchV1760Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandCombination-v1760: ref
	{
		if seq.IsComponentPresent(0) {
			ie.BandCombination_v1760 = new(BandCombination_v1760)
			if err := ie.BandCombination_v1760.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
