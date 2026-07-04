// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-DCCH-Message (line 162).

var uLDCCHMessageConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "message"},
	},
}

type UL_DCCH_Message struct {
	Message UL_DCCH_MessageType
}

func (ie *UL_DCCH_Message) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLDCCHMessageConstraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UL_DCCH_Message) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLDCCHMessageConstraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
