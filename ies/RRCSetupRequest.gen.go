// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCSetupRequest (line 1785).

var rRCSetupRequestConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrcSetupRequest"},
	},
}

type RRCSetupRequest struct {
	RrcSetupRequest RRCSetupRequest_IEs
}

func (ie *RRCSetupRequest) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSetupRequestConstraints)
	_ = seq

	// 1. rrcSetupRequest: ref
	{
		if err := ie.RrcSetupRequest.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RRCSetupRequest) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSetupRequestConstraints)
	_ = seq

	// 1. rrcSetupRequest: ref
	{
		if err := ie.RrcSetupRequest.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
