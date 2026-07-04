// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-RLC-BearerConfig-r16 (line 28136).

var sLRLCBearerConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RLC-BearerConfigIndex-r16"},
		{Name: "sl-ServedRadioBearer-r16", Optional: true},
		{Name: "sl-RLC-Config-r16", Optional: true},
		{Name: "sl-MAC-LogicalChannelConfig-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

type SL_RLC_BearerConfig_r16 struct {
	Sl_RLC_BearerConfigIndex_r16    SL_RLC_BearerConfigIndex_r16
	Sl_ServedRadioBearer_r16        *SLRB_Uu_ConfigIndex_r16
	Sl_RLC_Config_r16               *SL_RLC_Config_r16
	Sl_MAC_LogicalChannelConfig_r16 *SL_LogicalChannelConfig_r16
	Sl_RLC_BearerConfigIndex_v1800  *SL_RLC_BearerConfigIndex_v1800
}

func (ie *SL_RLC_BearerConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLRLCBearerConfigR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_RLC_BearerConfigIndex_v1800 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_ServedRadioBearer_r16 != nil, ie.Sl_RLC_Config_r16 != nil, ie.Sl_MAC_LogicalChannelConfig_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-RLC-BearerConfigIndex-r16: ref
	{
		if err := ie.Sl_RLC_BearerConfigIndex_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-ServedRadioBearer-r16: ref
	{
		if ie.Sl_ServedRadioBearer_r16 != nil {
			if err := ie.Sl_ServedRadioBearer_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-RLC-Config-r16: ref
	{
		if ie.Sl_RLC_Config_r16 != nil {
			if err := ie.Sl_RLC_Config_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sl-MAC-LogicalChannelConfig-r16: ref
	{
		if ie.Sl_MAC_LogicalChannelConfig_r16 != nil {
			if err := ie.Sl_MAC_LogicalChannelConfig_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-RLC-BearerConfigIndex-v1800", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_RLC_BearerConfigIndex_v1800 != nil}); err != nil {
				return err
			}

			if ie.Sl_RLC_BearerConfigIndex_v1800 != nil {
				if err := ie.Sl_RLC_BearerConfigIndex_v1800.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_RLC_BearerConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLRLCBearerConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-RLC-BearerConfigIndex-r16: ref
	{
		if err := ie.Sl_RLC_BearerConfigIndex_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-ServedRadioBearer-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_ServedRadioBearer_r16 = new(SLRB_Uu_ConfigIndex_r16)
			if err := ie.Sl_ServedRadioBearer_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-RLC-Config-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_RLC_Config_r16 = new(SL_RLC_Config_r16)
			if err := ie.Sl_RLC_Config_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sl-MAC-LogicalChannelConfig-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sl_MAC_LogicalChannelConfig_r16 = new(SL_LogicalChannelConfig_r16)
			if err := ie.Sl_MAC_LogicalChannelConfig_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sl-RLC-BearerConfigIndex-v1800", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_RLC_BearerConfigIndex_v1800 = new(SL_RLC_BearerConfigIndex_v1800)
			if err := ie.Sl_RLC_BearerConfigIndex_v1800.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
