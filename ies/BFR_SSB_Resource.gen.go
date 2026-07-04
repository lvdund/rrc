// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BFR-SSB-Resource (line 5140).

var bFRSSBResourceConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb"},
		{Name: "ra-PreambleIndex"},
	},
}

var bFRSSBResourceRaPreambleIndexConstraints = per.Constrained(0, 63)

type BFR_SSB_Resource struct {
	Ssb              SSB_Index
	Ra_PreambleIndex int64
}

func (ie *BFR_SSB_Resource) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bFRSSBResourceConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. ssb: ref
	{
		if err := ie.Ssb.Encode(e); err != nil {
			return err
		}
	}

	// 3. ra-PreambleIndex: integer
	{
		if err := e.EncodeInteger(ie.Ra_PreambleIndex, bFRSSBResourceRaPreambleIndexConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BFR_SSB_Resource) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bFRSSBResourceConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. ssb: ref
	{
		if err := ie.Ssb.Decode(d); err != nil {
			return err
		}
	}

	// 3. ra-PreambleIndex: integer
	{
		v1, err := d.DecodeInteger(bFRSSBResourceRaPreambleIndexConstraints)
		if err != nil {
			return err
		}
		ie.Ra_PreambleIndex = v1
	}

	return nil
}
