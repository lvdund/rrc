// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNRDC-v17b0 (line 18444).

var cAParametersNRDCV17b0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-ForDC-v17b0", Optional: true},
	},
}

type CA_ParametersNRDC_V17b0 struct {
	Ca_ParametersNR_ForDC_V17b0 *CA_ParametersNR_v1740
}

func (ie *CA_ParametersNRDC_V17b0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRDCV17b0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_ForDC_V17b0 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC-v17b0: ref
	{
		if ie.Ca_ParametersNR_ForDC_V17b0 != nil {
			if err := ie.Ca_ParametersNR_ForDC_V17b0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNRDC_V17b0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRDCV17b0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC-v17b0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_ForDC_V17b0 = new(CA_ParametersNR_v1740)
			if err := ie.Ca_ParametersNR_ForDC_V17b0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
