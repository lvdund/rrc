// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCResumeRequest (line 1655).

var rRCResumeRequestConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrcResumeRequest"},
	},
}

type RRCResumeRequest struct {
	RrcResumeRequest RRCResumeRequest_IEs
}

func (ie *RRCResumeRequest) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeRequestConstraints)
	_ = seq

	// 1. rrcResumeRequest: ref
	{
		if err := ie.RrcResumeRequest.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RRCResumeRequest) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeRequestConstraints)
	_ = seq

	// 1. rrcResumeRequest: ref
	{
		if err := ie.RrcResumeRequest.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
