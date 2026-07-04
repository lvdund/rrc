// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MBS-ServiceInfo-r17 (line 28491).

var mBSServiceInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tmgi-r17"},
	},
}

type MBS_ServiceInfo_r17 struct {
	Tmgi_r17 TMGI_r17
}

func (ie *MBS_ServiceInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mBSServiceInfoR17Constraints)
	_ = seq

	// 1. tmgi-r17: ref
	{
		if err := ie.Tmgi_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MBS_ServiceInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mBSServiceInfoR17Constraints)
	_ = seq

	// 1. tmgi-r17: ref
	{
		if err := ie.Tmgi_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
