// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-v1580 (line 16679).

var bandCombinationV1580Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mrdc-Parameters-v1580"},
	},
}

type BandCombination_v1580 struct {
	Mrdc_Parameters_v1580 MRDC_Parameters_v1580
}

func (ie *BandCombination_v1580) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1580Constraints)
	_ = seq

	// 1. mrdc-Parameters-v1580: ref
	{
		if err := ie.Mrdc_Parameters_v1580.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BandCombination_v1580) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1580Constraints)
	_ = seq

	// 1. mrdc-Parameters-v1580: ref
	{
		if err := ie.Mrdc_Parameters_v1580.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
