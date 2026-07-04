// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-CCCH-Message (line 127).

var uLCCCHMessageConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "message"},
	},
}

type UL_CCCH_Message struct {
	Message UL_CCCH_MessageType
}

func (ie *UL_CCCH_Message) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLCCCHMessageConstraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UL_CCCH_Message) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLCCCHMessageConstraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
