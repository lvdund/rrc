// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNRDC-v1760 (line 18436).

var cAParametersNRDCV1760Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-ForDC-v1760"},
	},
}

type CA_ParametersNRDC_v1760 struct {
	Ca_ParametersNR_ForDC_v1760 CA_ParametersNR_v1760
}

func (ie *CA_ParametersNRDC_v1760) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRDCV1760Constraints)
	_ = seq

	// 1. ca-ParametersNR-ForDC-v1760: ref
	{
		if err := ie.Ca_ParametersNR_ForDC_v1760.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CA_ParametersNRDC_v1760) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRDCV1760Constraints)
	_ = seq

	// 1. ca-ParametersNR-ForDC-v1760: ref
	{
		if err := ie.Ca_ParametersNR_ForDC_v1760.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
