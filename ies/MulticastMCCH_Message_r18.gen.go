// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MulticastMCCH-Message-r18 (line 96).

var multicastMCCHMessageR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "message"},
	},
}

type MulticastMCCH_Message_r18 struct {
	Message MulticastMCCH_MessageType_r18
}

func (ie *MulticastMCCH_Message_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(multicastMCCHMessageR18Constraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MulticastMCCH_Message_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(multicastMCCHMessageR18Constraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
