// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: GuardBand-r16 (line 14777).

var guardBandR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "startCRB-r16"},
		{Name: "nrofCRBs-r16"},
	},
}

var guardBandR16StartCRBR16Constraints = per.Constrained(0, 274)

var guardBandR16NrofCRBsR16Constraints = per.Constrained(0, 15)

type GuardBand_r16 struct {
	StartCRB_r16 int64
	NrofCRBs_r16 int64
}

func (ie *GuardBand_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(guardBandR16Constraints)
	_ = seq

	// 1. startCRB-r16: integer
	{
		if err := e.EncodeInteger(ie.StartCRB_r16, guardBandR16StartCRBR16Constraints); err != nil {
			return err
		}
	}

	// 2. nrofCRBs-r16: integer
	{
		if err := e.EncodeInteger(ie.NrofCRBs_r16, guardBandR16NrofCRBsR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *GuardBand_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(guardBandR16Constraints)
	_ = seq

	// 1. startCRB-r16: integer
	{
		v0, err := d.DecodeInteger(guardBandR16StartCRBR16Constraints)
		if err != nil {
			return err
		}
		ie.StartCRB_r16 = v0
	}

	// 2. nrofCRBs-r16: integer
	{
		v1, err := d.DecodeInteger(guardBandR16NrofCRBsR16Constraints)
		if err != nil {
			return err
		}
		ie.NrofCRBs_r16 = v1
	}

	return nil
}
