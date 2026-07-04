// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-DCCH-Message (line 54).

var dLDCCHMessageConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "message"},
	},
}

type DL_DCCH_Message struct {
	Message DL_DCCH_MessageType
}

func (ie *DL_DCCH_Message) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLDCCHMessageConstraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DL_DCCH_Message) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLDCCHMessageConstraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
