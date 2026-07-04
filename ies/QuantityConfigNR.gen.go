// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: QuantityConfigNR (line 12787).

var quantityConfigNRConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "quantityConfigCell"},
		{Name: "quantityConfigRS-Index", Optional: true},
	},
}

type QuantityConfigNR struct {
	QuantityConfigCell     QuantityConfigRS
	QuantityConfigRS_Index *QuantityConfigRS
}

func (ie *QuantityConfigNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(quantityConfigNRConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.QuantityConfigRS_Index != nil}); err != nil {
		return err
	}

	// 2. quantityConfigCell: ref
	{
		if err := ie.QuantityConfigCell.Encode(e); err != nil {
			return err
		}
	}

	// 3. quantityConfigRS-Index: ref
	{
		if ie.QuantityConfigRS_Index != nil {
			if err := ie.QuantityConfigRS_Index.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *QuantityConfigNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(quantityConfigNRConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. quantityConfigCell: ref
	{
		if err := ie.QuantityConfigCell.Decode(d); err != nil {
			return err
		}
	}

	// 3. quantityConfigRS-Index: ref
	{
		if seq.IsComponentPresent(1) {
			ie.QuantityConfigRS_Index = new(QuantityConfigRS)
			if err := ie.QuantityConfigRS_Index.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
