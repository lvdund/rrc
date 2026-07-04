// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNRDC-v16a0 (line 18410).

var cAParametersNRDCV16a0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-ForDC-v16a0", Optional: true},
	},
}

type CA_ParametersNRDC_V16a0 struct {
	Ca_ParametersNR_ForDC_V16a0 *CA_ParametersNR_V16a0
}

func (ie *CA_ParametersNRDC_V16a0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRDCV16a0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_ForDC_V16a0 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC-v16a0: ref
	{
		if ie.Ca_ParametersNR_ForDC_V16a0 != nil {
			if err := ie.Ca_ParametersNR_ForDC_V16a0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNRDC_V16a0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRDCV16a0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC-v16a0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_ForDC_V16a0 = new(CA_ParametersNR_V16a0)
			if err := ie.Ca_ParametersNR_ForDC_V16a0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
