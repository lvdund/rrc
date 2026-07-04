// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BCCH-DL-SCH-Message (line 22).

var bCCHDLSCHMessageConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "message"},
	},
}

type BCCH_DL_SCH_Message struct {
	Message BCCH_DL_SCH_MessageType
}

func (ie *BCCH_DL_SCH_Message) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bCCHDLSCHMessageConstraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BCCH_DL_SCH_Message) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bCCHDLSCHMessageConstraints)
	_ = seq

	// 1. message: ref
	{
		if err := ie.Message.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
