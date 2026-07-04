// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR-v1550 (line 17292).

var cAParametersNRV1550Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dummy", Optional: true},
	},
}

const (
	CA_ParametersNR_v1550_Dummy_Supported = 0
)

var cAParametersNRV1550DummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1550_Dummy_Supported},
}

type CA_ParametersNR_v1550 struct {
	Dummy *int64
}

func (ie *CA_ParametersNR_v1550) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1550Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dummy != nil}); err != nil {
		return err
	}

	// 2. dummy: enumerated
	{
		if ie.Dummy != nil {
			if err := e.EncodeEnumerated(*ie.Dummy, cAParametersNRV1550DummyConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1550) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1550Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dummy: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1550DummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &idx
		}
	}

	return nil
}
