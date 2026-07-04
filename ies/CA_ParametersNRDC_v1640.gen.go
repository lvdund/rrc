// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNRDC-v1640 (line 18402).

var cAParametersNRDCV1640Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-ForDC-v1640", Optional: true},
	},
}

type CA_ParametersNRDC_v1640 struct {
	Ca_ParametersNR_ForDC_v1640 *CA_ParametersNR_v1640
}

func (ie *CA_ParametersNRDC_v1640) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRDCV1640Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_ForDC_v1640 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC-v1640: ref
	{
		if ie.Ca_ParametersNR_ForDC_v1640 != nil {
			if err := ie.Ca_ParametersNR_ForDC_v1640.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNRDC_v1640) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRDCV1640Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC-v1640: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_ForDC_v1640 = new(CA_ParametersNR_v1640)
			if err := ie.Ca_ParametersNR_ForDC_v1640.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
