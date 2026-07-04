// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SI-RequestResourcesRepetition-r18 (line 15044).

var sIRequestResourcesRepetitionR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ra-PreambleStartIndex-r18"},
	},
}

var sIRequestResourcesRepetitionR18RaPreambleStartIndexR18Constraints = per.Constrained(0, 63)

type SI_RequestResourcesRepetition_r18 struct {
	Ra_PreambleStartIndex_r18 int64
}

func (ie *SI_RequestResourcesRepetition_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIRequestResourcesRepetitionR18Constraints)
	_ = seq

	// 1. ra-PreambleStartIndex-r18: integer
	{
		if err := e.EncodeInteger(ie.Ra_PreambleStartIndex_r18, sIRequestResourcesRepetitionR18RaPreambleStartIndexR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SI_RequestResourcesRepetition_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIRequestResourcesRepetitionR18Constraints)
	_ = seq

	// 1. ra-PreambleStartIndex-r18: integer
	{
		v0, err := d.DecodeInteger(sIRequestResourcesRepetitionR18RaPreambleStartIndexR18Constraints)
		if err != nil {
			return err
		}
		ie.Ra_PreambleStartIndex_r18 = v0
	}

	return nil
}
