// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-InterUE-CoordinationConfig-r17 (line 27348).

var sLInterUECoordinationConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-InterUE-CoordinationScheme1-r17", Optional: true},
		{Name: "sl-InterUE-CoordinationScheme2-r17", Optional: true},
	},
}

type SL_InterUE_CoordinationConfig_r17 struct {
	Sl_InterUE_CoordinationScheme1_r17 *SL_InterUE_CoordinationScheme1_r17
	Sl_InterUE_CoordinationScheme2_r17 *SL_InterUE_CoordinationScheme2_r17
}

func (ie *SL_InterUE_CoordinationConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLInterUECoordinationConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_InterUE_CoordinationScheme1_r17 != nil, ie.Sl_InterUE_CoordinationScheme2_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-InterUE-CoordinationScheme1-r17: ref
	{
		if ie.Sl_InterUE_CoordinationScheme1_r17 != nil {
			if err := ie.Sl_InterUE_CoordinationScheme1_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-InterUE-CoordinationScheme2-r17: ref
	{
		if ie.Sl_InterUE_CoordinationScheme2_r17 != nil {
			if err := ie.Sl_InterUE_CoordinationScheme2_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_InterUE_CoordinationConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLInterUECoordinationConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-InterUE-CoordinationScheme1-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_InterUE_CoordinationScheme1_r17 = new(SL_InterUE_CoordinationScheme1_r17)
			if err := ie.Sl_InterUE_CoordinationScheme1_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-InterUE-CoordinationScheme2-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_InterUE_CoordinationScheme2_r17 = new(SL_InterUE_CoordinationScheme2_r17)
			if err := ie.Sl_InterUE_CoordinationScheme2_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
