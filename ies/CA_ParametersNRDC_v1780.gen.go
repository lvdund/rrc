// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNRDC-v1780 (line 18440).

var cAParametersNRDCV1780Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-ForDC-v1780", Optional: true},
	},
}

type CA_ParametersNRDC_v1780 struct {
	Ca_ParametersNR_ForDC_v1780 *CA_ParametersNR_v1780
}

func (ie *CA_ParametersNRDC_v1780) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRDCV1780Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_ForDC_v1780 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC-v1780: ref
	{
		if ie.Ca_ParametersNR_ForDC_v1780 != nil {
			if err := ie.Ca_ParametersNR_ForDC_v1780.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNRDC_v1780) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRDCV1780Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC-v1780: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_ForDC_v1780 = new(CA_ParametersNR_v1780)
			if err := ie.Ca_ParametersNR_ForDC_v1780.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
