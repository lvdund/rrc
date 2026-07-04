// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-GapExt-v1820 (line 10209).

var mUSIMGapExtV1820Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "gapPriority-r18"},
	},
}

type MUSIM_GapExt_v1820 struct {
	GapPriority_r18 GapPriority_r17
}

func (ie *MUSIM_GapExt_v1820) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMGapExtV1820Constraints)
	_ = seq

	// 1. gapPriority-r18: ref
	{
		if err := ie.GapPriority_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MUSIM_GapExt_v1820) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMGapExtV1820Constraints)
	_ = seq

	// 1. gapPriority-r18: ref
	{
		if err := ie.GapPriority_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
