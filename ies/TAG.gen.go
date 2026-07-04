// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TAG (line 15955).

var tAGConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "tag-Id"},
		{Name: "timeAlignmentTimer"},
	},
}

type TAG struct {
	Tag_Id             TAG_Id
	TimeAlignmentTimer TimeAlignmentTimer
}

func (ie *TAG) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tAGConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. tag-Id: ref
	{
		if err := ie.Tag_Id.Encode(e); err != nil {
			return err
		}
	}

	// 3. timeAlignmentTimer: ref
	{
		if err := ie.TimeAlignmentTimer.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *TAG) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tAGConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. tag-Id: ref
	{
		if err := ie.Tag_Id.Decode(d); err != nil {
			return err
		}
	}

	// 3. timeAlignmentTimer: ref
	{
		if err := ie.TimeAlignmentTimer.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
