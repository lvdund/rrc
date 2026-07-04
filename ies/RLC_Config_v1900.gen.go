// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RLC-Config-v1900 (line 14151).

var rLCConfigV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-AM-RLC-v1900"},
		{Name: "ul-AM-RLC-v1900"},
	},
}

type RLC_Config_v1900 struct {
	Dl_AM_RLC_v1900 DL_AM_RLC_v1900
	Ul_AM_RLC_v1900 UL_AM_RLC_v1900
}

func (ie *RLC_Config_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rLCConfigV1900Constraints)
	_ = seq

	// 1. dl-AM-RLC-v1900: ref
	{
		if err := ie.Dl_AM_RLC_v1900.Encode(e); err != nil {
			return err
		}
	}

	// 2. ul-AM-RLC-v1900: ref
	{
		if err := ie.Ul_AM_RLC_v1900.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RLC_Config_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rLCConfigV1900Constraints)
	_ = seq

	// 1. dl-AM-RLC-v1900: ref
	{
		if err := ie.Dl_AM_RLC_v1900.Decode(d); err != nil {
			return err
		}
	}

	// 2. ul-AM-RLC-v1900: ref
	{
		if err := ie.Ul_AM_RLC_v1900.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
