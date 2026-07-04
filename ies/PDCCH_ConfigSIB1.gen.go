// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-ConfigSIB1 (line 11151).

var pDCCHConfigSIB1Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "controlResourceSetZero"},
		{Name: "searchSpaceZero"},
	},
}

type PDCCH_ConfigSIB1 struct {
	ControlResourceSetZero ControlResourceSetZero
	SearchSpaceZero        SearchSpaceZero
}

func (ie *PDCCH_ConfigSIB1) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHConfigSIB1Constraints)
	_ = seq

	// 1. controlResourceSetZero: ref
	{
		if err := ie.ControlResourceSetZero.Encode(e); err != nil {
			return err
		}
	}

	// 2. searchSpaceZero: ref
	{
		if err := ie.SearchSpaceZero.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PDCCH_ConfigSIB1) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHConfigSIB1Constraints)
	_ = seq

	// 1. controlResourceSetZero: ref
	{
		if err := ie.ControlResourceSetZero.Decode(d); err != nil {
			return err
		}
	}

	// 2. searchSpaceZero: ref
	{
		if err := ie.SearchSpaceZero.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
