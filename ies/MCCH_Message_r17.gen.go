// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MCCH-Message-r17 (line 81).

var mCCHMessageR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "message"},
	},
}

type MCCH_Message_r17 struct {
	Message MCCH_MessageType_r17
}

func (ie *MCCH_Message_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mCCHMessageR17Constraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MCCH_Message_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mCCHMessageR17Constraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
