// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNRDC (line 18375).

var cAParametersNRDCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-ForDC", Optional: true},
		{Name: "ca-ParametersNR-ForDC-v1540", Optional: true},
		{Name: "ca-ParametersNR-ForDC-v1550", Optional: true},
		{Name: "ca-ParametersNR-ForDC-v1560", Optional: true},
		{Name: "featureSetCombinationDC", Optional: true},
	},
}

type CA_ParametersNRDC struct {
	Ca_ParametersNR_ForDC       *CA_ParametersNR
	Ca_ParametersNR_ForDC_v1540 *CA_ParametersNR_v1540
	Ca_ParametersNR_ForDC_v1550 *CA_ParametersNR_v1550
	Ca_ParametersNR_ForDC_v1560 *CA_ParametersNR_v1560
	FeatureSetCombinationDC     *FeatureSetCombinationId
}

func (ie *CA_ParametersNRDC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRDCConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_ForDC != nil, ie.Ca_ParametersNR_ForDC_v1540 != nil, ie.Ca_ParametersNR_ForDC_v1550 != nil, ie.Ca_ParametersNR_ForDC_v1560 != nil, ie.FeatureSetCombinationDC != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC: ref
	{
		if ie.Ca_ParametersNR_ForDC != nil {
			if err := ie.Ca_ParametersNR_ForDC.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNR-ForDC-v1540: ref
	{
		if ie.Ca_ParametersNR_ForDC_v1540 != nil {
			if err := ie.Ca_ParametersNR_ForDC_v1540.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ca-ParametersNR-ForDC-v1550: ref
	{
		if ie.Ca_ParametersNR_ForDC_v1550 != nil {
			if err := ie.Ca_ParametersNR_ForDC_v1550.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. ca-ParametersNR-ForDC-v1560: ref
	{
		if ie.Ca_ParametersNR_ForDC_v1560 != nil {
			if err := ie.Ca_ParametersNR_ForDC_v1560.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. featureSetCombinationDC: ref
	{
		if ie.FeatureSetCombinationDC != nil {
			if err := ie.FeatureSetCombinationDC.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNRDC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRDCConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-ForDC: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_ForDC = new(CA_ParametersNR)
			if err := ie.Ca_ParametersNR_ForDC.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNR-ForDC-v1540: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_ParametersNR_ForDC_v1540 = new(CA_ParametersNR_v1540)
			if err := ie.Ca_ParametersNR_ForDC_v1540.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ca-ParametersNR-ForDC-v1550: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ca_ParametersNR_ForDC_v1550 = new(CA_ParametersNR_v1550)
			if err := ie.Ca_ParametersNR_ForDC_v1550.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. ca-ParametersNR-ForDC-v1560: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Ca_ParametersNR_ForDC_v1560 = new(CA_ParametersNR_v1560)
			if err := ie.Ca_ParametersNR_ForDC_v1560.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. featureSetCombinationDC: ref
	{
		if seq.IsComponentPresent(4) {
			ie.FeatureSetCombinationDC = new(FeatureSetCombinationId)
			if err := ie.FeatureSetCombinationDC.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
