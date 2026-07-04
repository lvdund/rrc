// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RLC-Config-v1610 (line 14142).

var rLCConfigV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-AM-RLC-v1610"},
	},
}

type RLC_Config_v1610 struct {
	Dl_AM_RLC_v1610 DL_AM_RLC_v1610
}

func (ie *RLC_Config_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rLCConfigV1610Constraints)
	_ = seq

	// 1. dl-AM-RLC-v1610: ref
	{
		if err := ie.Dl_AM_RLC_v1610.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RLC_Config_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rLCConfigV1610Constraints)
	_ = seq

	// 1. dl-AM-RLC-v1610: ref
	{
		if err := ie.Dl_AM_RLC_v1610.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
