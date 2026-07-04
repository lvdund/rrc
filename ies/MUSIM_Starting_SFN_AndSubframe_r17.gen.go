// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-Starting-SFN-AndSubframe-r17 (line 10238).

var mUSIMStartingSFNAndSubframeR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "starting-SFN-r17"},
		{Name: "startingSubframe-r17"},
	},
}

var mUSIMStartingSFNAndSubframeR17StartingSFNR17Constraints = per.Constrained(0, 1023)

var mUSIMStartingSFNAndSubframeR17StartingSubframeR17Constraints = per.Constrained(0, 9)

type MUSIM_Starting_SFN_AndSubframe_r17 struct {
	Starting_SFN_r17     int64
	StartingSubframe_r17 int64
}

func (ie *MUSIM_Starting_SFN_AndSubframe_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMStartingSFNAndSubframeR17Constraints)
	_ = seq

	// 1. starting-SFN-r17: integer
	{
		if err := e.EncodeInteger(ie.Starting_SFN_r17, mUSIMStartingSFNAndSubframeR17StartingSFNR17Constraints); err != nil {
			return err
		}
	}

	// 2. startingSubframe-r17: integer
	{
		if err := e.EncodeInteger(ie.StartingSubframe_r17, mUSIMStartingSFNAndSubframeR17StartingSubframeR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MUSIM_Starting_SFN_AndSubframe_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMStartingSFNAndSubframeR17Constraints)
	_ = seq

	// 1. starting-SFN-r17: integer
	{
		v0, err := d.DecodeInteger(mUSIMStartingSFNAndSubframeR17StartingSFNR17Constraints)
		if err != nil {
			return err
		}
		ie.Starting_SFN_r17 = v0
	}

	// 2. startingSubframe-r17: integer
	{
		v1, err := d.DecodeInteger(mUSIMStartingSFNAndSubframeR17StartingSubframeR17Constraints)
		if err != nil {
			return err
		}
		ie.StartingSubframe_r17 = v1
	}

	return nil
}
