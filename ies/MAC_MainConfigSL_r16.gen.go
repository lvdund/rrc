// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MAC-MainConfigSL-r16 (line 28208).

var mACMainConfigSLR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-BSR-Config-r16", Optional: true},
		{Name: "ul-PrioritizationThres-r16", Optional: true},
		{Name: "sl-PrioritizationThres-r16", Optional: true},
	},
}

var mACMainConfigSLR16UlPrioritizationThresR16Constraints = per.Constrained(1, 16)

var mACMainConfigSLR16SlPrioritizationThresR16Constraints = per.Constrained(1, 8)

type MAC_MainConfigSL_r16 struct {
	Sl_BSR_Config_r16          *BSR_Config
	Ul_PrioritizationThres_r16 *int64
	Sl_PrioritizationThres_r16 *int64
}

func (ie *MAC_MainConfigSL_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mACMainConfigSLR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_BSR_Config_r16 != nil, ie.Ul_PrioritizationThres_r16 != nil, ie.Sl_PrioritizationThres_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-BSR-Config-r16: ref
	{
		if ie.Sl_BSR_Config_r16 != nil {
			if err := ie.Sl_BSR_Config_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ul-PrioritizationThres-r16: integer
	{
		if ie.Ul_PrioritizationThres_r16 != nil {
			if err := e.EncodeInteger(*ie.Ul_PrioritizationThres_r16, mACMainConfigSLR16UlPrioritizationThresR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-PrioritizationThres-r16: integer
	{
		if ie.Sl_PrioritizationThres_r16 != nil {
			if err := e.EncodeInteger(*ie.Sl_PrioritizationThres_r16, mACMainConfigSLR16SlPrioritizationThresR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MAC_MainConfigSL_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mACMainConfigSLR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-BSR-Config-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_BSR_Config_r16 = new(BSR_Config)
			if err := ie.Sl_BSR_Config_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ul-PrioritizationThres-r16: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(mACMainConfigSLR16UlPrioritizationThresR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_PrioritizationThres_r16 = &v
		}
	}

	// 5. sl-PrioritizationThres-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(mACMainConfigSLR16SlPrioritizationThresR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PrioritizationThres_r16 = &v
		}
	}

	return nil
}
