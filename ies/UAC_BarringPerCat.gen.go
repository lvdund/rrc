// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UAC-BarringPerCat (line 16199).

var uACBarringPerCatConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "accessCategory"},
		{Name: "uac-barringInfoSetIndex"},
	},
}

var uACBarringPerCatAccessCategoryConstraints = per.Constrained(1, common.MaxAccessCat_1)

type UAC_BarringPerCat struct {
	AccessCategory          int64
	Uac_BarringInfoSetIndex UAC_BarringInfoSetIndex
}

func (ie *UAC_BarringPerCat) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uACBarringPerCatConstraints)
	_ = seq

	// 1. accessCategory: integer
	{
		if err := e.EncodeInteger(ie.AccessCategory, uACBarringPerCatAccessCategoryConstraints); err != nil {
			return err
		}
	}

	// 2. uac-barringInfoSetIndex: ref
	{
		if err := ie.Uac_BarringInfoSetIndex.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UAC_BarringPerCat) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uACBarringPerCatConstraints)
	_ = seq

	// 1. accessCategory: integer
	{
		v0, err := d.DecodeInteger(uACBarringPerCatAccessCategoryConstraints)
		if err != nil {
			return err
		}
		ie.AccessCategory = v0
	}

	// 2. uac-barringInfoSetIndex: ref
	{
		if err := ie.Uac_BarringInfoSetIndex.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
