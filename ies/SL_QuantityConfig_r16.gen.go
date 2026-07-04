// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-QuantityConfig-r16 (line 27773).

var sLQuantityConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-FilterCoefficientDMRS-r16", Optional: true},
	},
}

type SL_QuantityConfig_r16 struct {
	Sl_FilterCoefficientDMRS_r16 *FilterCoefficient
}

func (ie *SL_QuantityConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLQuantityConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_FilterCoefficientDMRS_r16 != nil && ie.Sl_FilterCoefficientDMRS_r16.Value != FilterCoefficient_Fc4}); err != nil {
		return err
	}

	// 3. sl-FilterCoefficientDMRS-r16: ref
	{
		if ie.Sl_FilterCoefficientDMRS_r16 != nil && ie.Sl_FilterCoefficientDMRS_r16.Value != FilterCoefficient_Fc4 {
			if err := ie.Sl_FilterCoefficientDMRS_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_QuantityConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLQuantityConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-FilterCoefficientDMRS-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_FilterCoefficientDMRS_r16 = new(FilterCoefficient)
			if err := ie.Sl_FilterCoefficientDMRS_r16.Decode(d); err != nil {
				return err
			}
		} else {
			ie.Sl_FilterCoefficientDMRS_r16 = &FilterCoefficient{Value: FilterCoefficient_Fc4}
		}
	}

	return nil
}
