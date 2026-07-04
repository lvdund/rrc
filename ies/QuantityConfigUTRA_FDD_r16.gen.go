// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: QuantityConfigUTRA-FDD-r16 (line 12808).

var quantityConfigUTRAFDDR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "filterCoefficientRSCP-r16", Optional: true},
		{Name: "filterCoefficientEcNO-r16", Optional: true},
	},
}

type QuantityConfigUTRA_FDD_r16 struct {
	FilterCoefficientRSCP_r16 *FilterCoefficient
	FilterCoefficientEcNO_r16 *FilterCoefficient
}

func (ie *QuantityConfigUTRA_FDD_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(quantityConfigUTRAFDDR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FilterCoefficientRSCP_r16 != nil && ie.FilterCoefficientRSCP_r16.Value != FilterCoefficient_Fc4, ie.FilterCoefficientEcNO_r16 != nil && ie.FilterCoefficientEcNO_r16.Value != FilterCoefficient_Fc4}); err != nil {
		return err
	}

	// 2. filterCoefficientRSCP-r16: ref
	{
		if ie.FilterCoefficientRSCP_r16 != nil && ie.FilterCoefficientRSCP_r16.Value != FilterCoefficient_Fc4 {
			if err := ie.FilterCoefficientRSCP_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. filterCoefficientEcNO-r16: ref
	{
		if ie.FilterCoefficientEcNO_r16 != nil && ie.FilterCoefficientEcNO_r16.Value != FilterCoefficient_Fc4 {
			if err := ie.FilterCoefficientEcNO_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *QuantityConfigUTRA_FDD_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(quantityConfigUTRAFDDR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. filterCoefficientRSCP-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FilterCoefficientRSCP_r16 = new(FilterCoefficient)
			if err := ie.FilterCoefficientRSCP_r16.Decode(d); err != nil {
				return err
			}
		} else {
			ie.FilterCoefficientRSCP_r16 = &FilterCoefficient{Value: FilterCoefficient_Fc4}
		}
	}

	// 3. filterCoefficientEcNO-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.FilterCoefficientEcNO_r16 = new(FilterCoefficient)
			if err := ie.FilterCoefficientEcNO_r16.Decode(d); err != nil {
				return err
			}
		} else {
			ie.FilterCoefficientEcNO_r16 = &FilterCoefficient{Value: FilterCoefficient_Fc4}
		}
	}

	return nil
}
