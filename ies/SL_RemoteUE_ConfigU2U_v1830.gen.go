// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RemoteUE-ConfigU2U-v1830 (line 27864).

var sLRemoteUEConfigU2UV1830Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-FilterCoefficientU2U-r18", Optional: true},
	},
}

type SL_RemoteUE_ConfigU2U_v1830 struct {
	Sl_FilterCoefficientU2U_r18 *FilterCoefficient
}

func (ie *SL_RemoteUE_ConfigU2U_v1830) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLRemoteUEConfigU2UV1830Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_FilterCoefficientU2U_r18 != nil}); err != nil {
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

	return nil
}

func (ie *SL_RemoteUE_ConfigU2U_v1830) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLRemoteUEConfigU2UV1830Constraints)

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

	return nil
}
