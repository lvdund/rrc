// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RLC-Config-v1700 (line 14146).

var rLCConfigV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-AM-RLC-v1700"},
		{Name: "dl-UM-RLC-v1700"},
	},
}

type RLC_Config_v1700 struct {
	Dl_AM_RLC_v1700 DL_AM_RLC_v1700
	Dl_UM_RLC_v1700 DL_UM_RLC_v1700
}

func (ie *RLC_Config_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rLCConfigV1700Constraints)
	_ = seq

	// 1. dl-AM-RLC-v1700: ref
	{
		if err := ie.Dl_AM_RLC_v1700.Encode(e); err != nil {
			return err
		}
	}

	// 2. dl-UM-RLC-v1700: ref
	{
		if err := ie.Dl_UM_RLC_v1700.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RLC_Config_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rLCConfigV1700Constraints)
	_ = seq

	// 1. dl-AM-RLC-v1700: ref
	{
		if err := ie.Dl_AM_RLC_v1700.Decode(d); err != nil {
			return err
		}
	}

	// 2. dl-UM-RLC-v1700: ref
	{
		if err := ie.Dl_UM_RLC_v1700.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
