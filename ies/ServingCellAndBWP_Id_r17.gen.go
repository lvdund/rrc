// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ServingCellAndBWP-Id-r17 (line 14603).

var servingCellAndBWPIdR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "servingcell-r17"},
		{Name: "bwp-r17"},
	},
}

type ServingCellAndBWP_Id_r17 struct {
	Servingcell_r17 ServCellIndex
	Bwp_r17         BWP_Id
}

func (ie *ServingCellAndBWP_Id_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(servingCellAndBWPIdR17Constraints)
	_ = seq

	// 1. servingcell-r17: ref
	{
		if err := ie.Servingcell_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. bwp-r17: ref
	{
		if err := ie.Bwp_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *ServingCellAndBWP_Id_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(servingCellAndBWPIdR17Constraints)
	_ = seq

	// 1. servingcell-r17: ref
	{
		if err := ie.Servingcell_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. bwp-r17: ref
	{
		if err := ie.Bwp_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
