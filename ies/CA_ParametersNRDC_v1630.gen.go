// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNRDC-v1630 (line 18397).

var cAParametersNRDCV1630Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-ForDC-v1610", Optional: true},
		{Name: "ca-ParametersNR-ForDC-v1630", Optional: true},
	},
}

type CA_ParametersNRDC_v1630 struct {
	Ca_ParametersNR_ForDC_v1610 *CA_ParametersNR_v1610
	Ca_ParametersNR_ForDC_v1630 *CA_ParametersNR_v1630
}

func (ie *CA_ParametersNRDC_v1630) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRDCV1630Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_ForDC_v1610 != nil, ie.Ca_ParametersNR_ForDC_v1630 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC-v1610: ref
	{
		if ie.Ca_ParametersNR_ForDC_v1610 != nil {
			if err := ie.Ca_ParametersNR_ForDC_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNR-ForDC-v1630: ref
	{
		if ie.Ca_ParametersNR_ForDC_v1630 != nil {
			if err := ie.Ca_ParametersNR_ForDC_v1630.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNRDC_v1630) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRDCV1630Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC-v1610: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_ForDC_v1610 = new(CA_ParametersNR_v1610)
			if err := ie.Ca_ParametersNR_ForDC_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNR-ForDC-v1630: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_ParametersNR_ForDC_v1630 = new(CA_ParametersNR_v1630)
			if err := ie.Ca_ParametersNR_ForDC_v1630.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
