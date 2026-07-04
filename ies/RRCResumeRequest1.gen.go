// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCResumeRequest1 (line 1669).

var rRCResumeRequest1Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrcResumeRequest1"},
	},
}

type RRCResumeRequest1 struct {
	RrcResumeRequest1 RRCResumeRequest1_IEs
}

func (ie *RRCResumeRequest1) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeRequest1Constraints)
	_ = seq

	// 1. rrcResumeRequest1: ref
	{
		if err := ie.RrcResumeRequest1.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RRCResumeRequest1) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeRequest1Constraints)
	_ = seq

	// 1. rrcResumeRequest1: ref
	{
		if err := ie.RrcResumeRequest1.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
