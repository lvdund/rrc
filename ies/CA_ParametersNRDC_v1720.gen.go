// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNRDC-v1720 (line 18427).

var cAParametersNRDCV1720Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-ForDC-v1700", Optional: true},
		{Name: "ca-ParametersNR-ForDC-v1720", Optional: true},
	},
}

type CA_ParametersNRDC_v1720 struct {
	Ca_ParametersNR_ForDC_v1700 *CA_ParametersNR_v1700
	Ca_ParametersNR_ForDC_v1720 *CA_ParametersNR_v1720
}

func (ie *CA_ParametersNRDC_v1720) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRDCV1720Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_ForDC_v1700 != nil, ie.Ca_ParametersNR_ForDC_v1720 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC-v1700: ref
	{
		if ie.Ca_ParametersNR_ForDC_v1700 != nil {
			if err := ie.Ca_ParametersNR_ForDC_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNR-ForDC-v1720: ref
	{
		if ie.Ca_ParametersNR_ForDC_v1720 != nil {
			if err := ie.Ca_ParametersNR_ForDC_v1720.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNRDC_v1720) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRDCV1720Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC-v1700: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_ForDC_v1700 = new(CA_ParametersNR_v1700)
			if err := ie.Ca_ParametersNR_ForDC_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNR-ForDC-v1720: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_ParametersNR_ForDC_v1720 = new(CA_ParametersNR_v1720)
			if err := ie.Ca_ParametersNR_ForDC_v1720.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
