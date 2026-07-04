// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-Gap-r17 (line 10204).

var mUSIMGapR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "musim-GapId-r17"},
		{Name: "musim-GapInfo-r17"},
	},
}

type MUSIM_Gap_r17 struct {
	Musim_GapId_r17   MUSIM_GapId_r17
	Musim_GapInfo_r17 MUSIM_GapInfo_r17
}

func (ie *MUSIM_Gap_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMGapR17Constraints)
	_ = seq

	// 1. musim-GapId-r17: ref
	{
		if err := ie.Musim_GapId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. musim-GapInfo-r17: ref
	{
		if err := ie.Musim_GapInfo_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MUSIM_Gap_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMGapR17Constraints)
	_ = seq

	// 1. musim-GapId-r17: ref
	{
		if err := ie.Musim_GapId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. musim-GapInfo-r17: ref
	{
		if err := ie.Musim_GapInfo_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
