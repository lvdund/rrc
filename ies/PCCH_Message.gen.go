// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PCCH-Message (line 111).

var pCCHMessageConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "message"},
	},
}

type PCCH_Message struct {
	Message PCCH_MessageType
}

func (ie *PCCH_Message) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pCCHMessageConstraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PCCH_Message) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pCCHMessageConstraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
