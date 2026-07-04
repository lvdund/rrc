// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RelayUE-ConfigU2U-v1840 (line 27832).

var sLRelayUEConfigU2UV1840Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-FilterCoefficientU2U-r18", Optional: true},
		{Name: "sd-FilterCoefficientU2U-r18", Optional: true},
	},
}

type SL_RelayUE_ConfigU2U_v1840 struct {
	Sl_FilterCoefficientU2U_r18 *FilterCoefficient
	Sd_FilterCoefficientU2U_r18 *FilterCoefficient
}

func (ie *SL_RelayUE_ConfigU2U_v1840) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLRelayUEConfigU2UV1840Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_FilterCoefficientU2U_r18 != nil, ie.Sd_FilterCoefficientU2U_r18 != nil}); err != nil {
		return err
	}

	// 3. sl-FilterCoefficientU2U-r18: ref
	{
		if ie.Sl_FilterCoefficientU2U_r18 != nil {
			if err := ie.Sl_FilterCoefficientU2U_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sd-FilterCoefficientU2U-r18: ref
	{
		if ie.Sd_FilterCoefficientU2U_r18 != nil {
			if err := ie.Sd_FilterCoefficientU2U_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_RelayUE_ConfigU2U_v1840) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLRelayUEConfigU2UV1840Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-FilterCoefficientU2U-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_FilterCoefficientU2U_r18 = new(FilterCoefficient)
			if err := ie.Sl_FilterCoefficientU2U_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sd-FilterCoefficientU2U-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sd_FilterCoefficientU2U_r18 = new(FilterCoefficient)
			if err := ie.Sd_FilterCoefficientU2U_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
