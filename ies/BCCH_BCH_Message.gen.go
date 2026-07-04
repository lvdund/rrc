// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BCCH-BCH-Message (line 10).

var bCCHBCHMessageConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "message"},
	},
}

type BCCH_BCH_Message struct {
	Message BCCH_BCH_MessageType
}

func (ie *BCCH_BCH_Message) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bCCHBCHMessageConstraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BCCH_BCH_Message) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bCCHBCHMessageConstraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
