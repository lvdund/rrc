// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EpochTime-r17 (line 8292).

var epochTimeR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sfn-r17"},
		{Name: "subFrameNR-r17"},
	},
}

var epochTimeR17SfnR17Constraints = per.Constrained(0, 1023)

var epochTimeR17SubFrameNRR17Constraints = per.Constrained(0, 9)

type EpochTime_r17 struct {
	Sfn_r17        int64
	SubFrameNR_r17 int64
}

func (ie *EpochTime_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(epochTimeR17Constraints)
	_ = seq

	// 1. sfn-r17: integer
	{
		if err := e.EncodeInteger(ie.Sfn_r17, epochTimeR17SfnR17Constraints); err != nil {
			return err
		}
	}

	// 2. subFrameNR-r17: integer
	{
		if err := e.EncodeInteger(ie.SubFrameNR_r17, epochTimeR17SubFrameNRR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *EpochTime_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(epochTimeR17Constraints)
	_ = seq

	// 1. sfn-r17: integer
	{
		v0, err := d.DecodeInteger(epochTimeR17SfnR17Constraints)
		if err != nil {
			return err
		}
		ie.Sfn_r17 = v0
	}

	// 2. subFrameNR-r17: integer
	{
		v1, err := d.DecodeInteger(epochTimeR17SubFrameNRR17Constraints)
		if err != nil {
			return err
		}
		ie.SubFrameNR_r17 = v1
	}

	return nil
}
