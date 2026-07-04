// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-CCCH-Message (line 37).

var dLCCCHMessageConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "message"},
	},
}

type DL_CCCH_Message struct {
	Message DL_CCCH_MessageType
}

func (ie *DL_CCCH_Message) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLCCCHMessageConstraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DL_CCCH_Message) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLCCCHMessageConstraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
