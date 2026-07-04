// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-MappingToAddMod-r17 (line 28269).

var sLMappingToAddModR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RemoteUE-RB-Identity-r17"},
		{Name: "sl-EgressRLC-ChannelUu-r17", Optional: true},
		{Name: "sl-EgressRLC-ChannelPC5-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

type SL_MappingToAddMod_r17 struct {
	Sl_RemoteUE_RB_Identity_r17 SL_RemoteUE_RB_Identity_r17
	Sl_EgressRLC_ChannelUu_r17  *Uu_RelayRLC_ChannelID_r17
	Sl_EgressRLC_ChannelPC5_r17 *SL_RLC_ChannelID_r17
	Sl_EgressRLC_Channel_UL_r19 *SL_RLC_ChannelID_r17
	Sl_EgressRLC_Channel_DL_r19 *SL_RLC_ChannelID_r17
}

func (ie *SL_MappingToAddMod_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLMappingToAddModR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_EgressRLC_Channel_UL_r19 != nil || ie.Sl_EgressRLC_Channel_DL_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_EgressRLC_ChannelUu_r17 != nil, ie.Sl_EgressRLC_ChannelPC5_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-RemoteUE-RB-Identity-r17: ref
	{
		if err := ie.Sl_RemoteUE_RB_Identity_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-EgressRLC-ChannelUu-r17: ref
	{
		if ie.Sl_EgressRLC_ChannelUu_r17 != nil {
			if err := ie.Sl_EgressRLC_ChannelUu_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-EgressRLC-ChannelPC5-r17: ref
	{
		if ie.Sl_EgressRLC_ChannelPC5_r17 != nil {
			if err := ie.Sl_EgressRLC_ChannelPC5_r17.Encode(e); err != nil {
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
					{Name: "sl-EgressRLC-Channel-UL-r19", Optional: true},
					{Name: "sl-EgressRLC-Channel-DL-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_EgressRLC_Channel_UL_r19 != nil, ie.Sl_EgressRLC_Channel_DL_r19 != nil}); err != nil {
				return err
			}

			if ie.Sl_EgressRLC_Channel_UL_r19 != nil {
				if err := ie.Sl_EgressRLC_Channel_UL_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_EgressRLC_Channel_DL_r19 != nil {
				if err := ie.Sl_EgressRLC_Channel_DL_r19.Encode(ex); err != nil {
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

func (ie *SL_MappingToAddMod_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLMappingToAddModR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-RemoteUE-RB-Identity-r17: ref
	{
		if err := ie.Sl_RemoteUE_RB_Identity_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-EgressRLC-ChannelUu-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_EgressRLC_ChannelUu_r17 = new(Uu_RelayRLC_ChannelID_r17)
			if err := ie.Sl_EgressRLC_ChannelUu_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-EgressRLC-ChannelPC5-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_EgressRLC_ChannelPC5_r17 = new(SL_RLC_ChannelID_r17)
			if err := ie.Sl_EgressRLC_ChannelPC5_r17.Decode(d); err != nil {
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
				{Name: "sl-EgressRLC-Channel-UL-r19", Optional: true},
				{Name: "sl-EgressRLC-Channel-DL-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_EgressRLC_Channel_UL_r19 = new(SL_RLC_ChannelID_r17)
			if err := ie.Sl_EgressRLC_Channel_UL_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sl_EgressRLC_Channel_DL_r19 = new(SL_RLC_ChannelID_r17)
			if err := ie.Sl_EgressRLC_Channel_DL_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
