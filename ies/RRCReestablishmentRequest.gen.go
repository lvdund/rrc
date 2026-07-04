// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReestablishmentRequest (line 942).

var rRCReestablishmentRequestConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrcReestablishmentRequest"},
	},
}

type RRCReestablishmentRequest struct {
	RrcReestablishmentRequest RRCReestablishmentRequest_IEs
}

func (ie *RRCReestablishmentRequest) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReestablishmentRequestConstraints)
	_ = seq

	// 1. rrcReestablishmentRequest: ref
	{
		if err := ie.RrcReestablishmentRequest.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RRCReestablishmentRequest) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReestablishmentRequestConstraints)
	_ = seq

	// 1. rrcReestablishmentRequest: ref
	{
		if err := ie.RrcReestablishmentRequest.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
