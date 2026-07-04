// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandCombination-v15g0 (line 16688).

var bandCombinationV15g0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-v15g0", Optional: true},
		{Name: "ca-ParametersNRDC-v15g0", Optional: true},
		{Name: "mrdc-Parameters-v15g0", Optional: true},
	},
}

type BandCombination_V15g0 struct {
	Ca_ParametersNR_V15g0   *CA_ParametersNR_V15g0
	Ca_ParametersNRDC_V15g0 *CA_ParametersNRDC_V15g0
	Mrdc_Parameters_V15g0   *MRDC_Parameters_V15g0
}

func (ie *BandCombination_V15g0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV15g0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_V15g0 != nil, ie.Ca_ParametersNRDC_V15g0 != nil, ie.Mrdc_Parameters_V15g0 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v15g0: ref
	{
		if ie.Ca_ParametersNR_V15g0 != nil {
			if err := ie.Ca_ParametersNR_V15g0.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v15g0: ref
	{
		if ie.Ca_ParametersNRDC_V15g0 != nil {
			if err := ie.Ca_ParametersNRDC_V15g0.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. mrdc-Parameters-v15g0: ref
	{
		if ie.Mrdc_Parameters_V15g0 != nil {
			if err := ie.Mrdc_Parameters_V15g0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_V15g0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV15g0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v15g0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_V15g0 = new(CA_ParametersNR_V15g0)
			if err := ie.Ca_ParametersNR_V15g0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v15g0: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_ParametersNRDC_V15g0 = new(CA_ParametersNRDC_V15g0)
			if err := ie.Ca_ParametersNRDC_V15g0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. mrdc-Parameters-v15g0: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Mrdc_Parameters_V15g0 = new(MRDC_Parameters_V15g0)
			if err := ie.Mrdc_Parameters_V15g0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
