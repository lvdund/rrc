// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BAP-Config-r16 (line 1079).

var bAPConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "bap-Address-r16", Optional: true},
		{Name: "defaultUL-BAP-RoutingID-r16", Optional: true},
		{Name: "defaultUL-BH-RLC-Channel-r16", Optional: true},
		{Name: "flowControlFeedbackType-r16", Optional: true},
	},
}

var bAPConfigR16BapAddressR16Constraints = per.FixedSize(10)

const (
	BAP_Config_r16_FlowControlFeedbackType_r16_PerBH_RLC_Channel = 0
	BAP_Config_r16_FlowControlFeedbackType_r16_PerRoutingID      = 1
	BAP_Config_r16_FlowControlFeedbackType_r16_Both              = 2
)

var bAPConfigR16FlowControlFeedbackTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BAP_Config_r16_FlowControlFeedbackType_r16_PerBH_RLC_Channel, BAP_Config_r16_FlowControlFeedbackType_r16_PerRoutingID, BAP_Config_r16_FlowControlFeedbackType_r16_Both},
}

type BAP_Config_r16 struct {
	Bap_Address_r16              *per.BitString
	DefaultUL_BAP_RoutingID_r16  *BAP_RoutingID_r16
	DefaultUL_BH_RLC_Channel_r16 *BH_RLC_ChannelID_r16
	FlowControlFeedbackType_r16  *int64
}

func (ie *BAP_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bAPConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Bap_Address_r16 != nil, ie.DefaultUL_BAP_RoutingID_r16 != nil, ie.DefaultUL_BH_RLC_Channel_r16 != nil, ie.FlowControlFeedbackType_r16 != nil}); err != nil {
		return err
	}

	// 3. bap-Address-r16: bit-string
	{
		if ie.Bap_Address_r16 != nil {
			if err := e.EncodeBitString(*ie.Bap_Address_r16, bAPConfigR16BapAddressR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. defaultUL-BAP-RoutingID-r16: ref
	{
		if ie.DefaultUL_BAP_RoutingID_r16 != nil {
			if err := ie.DefaultUL_BAP_RoutingID_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. defaultUL-BH-RLC-Channel-r16: ref
	{
		if ie.DefaultUL_BH_RLC_Channel_r16 != nil {
			if err := ie.DefaultUL_BH_RLC_Channel_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. flowControlFeedbackType-r16: enumerated
	{
		if ie.FlowControlFeedbackType_r16 != nil {
			if err := e.EncodeEnumerated(*ie.FlowControlFeedbackType_r16, bAPConfigR16FlowControlFeedbackTypeR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BAP_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bAPConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. bap-Address-r16: bit-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeBitString(bAPConfigR16BapAddressR16Constraints)
			if err != nil {
				return err
			}
			ie.Bap_Address_r16 = &v
		}
	}

	// 4. defaultUL-BAP-RoutingID-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.DefaultUL_BAP_RoutingID_r16 = new(BAP_RoutingID_r16)
			if err := ie.DefaultUL_BAP_RoutingID_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. defaultUL-BH-RLC-Channel-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.DefaultUL_BH_RLC_Channel_r16 = new(BH_RLC_ChannelID_r16)
			if err := ie.DefaultUL_BH_RLC_Channel_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. flowControlFeedbackType-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(bAPConfigR16FlowControlFeedbackTypeR16Constraints)
			if err != nil {
				return err
			}
			ie.FlowControlFeedbackType_r16 = &idx
		}
	}

	return nil
}
