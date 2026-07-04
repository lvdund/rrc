// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SubSlot-Config-r16 (line 20449).

var subSlotConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sub-SlotConfig-NCP-r16", Optional: true},
		{Name: "sub-SlotConfig-ECP-r16", Optional: true},
	},
}

const (
	SubSlot_Config_r16_Sub_SlotConfig_NCP_r16_N4 = 0
	SubSlot_Config_r16_Sub_SlotConfig_NCP_r16_N5 = 1
	SubSlot_Config_r16_Sub_SlotConfig_NCP_r16_N6 = 2
	SubSlot_Config_r16_Sub_SlotConfig_NCP_r16_N7 = 3
)

var subSlotConfigR16SubSlotConfigNCPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SubSlot_Config_r16_Sub_SlotConfig_NCP_r16_N4, SubSlot_Config_r16_Sub_SlotConfig_NCP_r16_N5, SubSlot_Config_r16_Sub_SlotConfig_NCP_r16_N6, SubSlot_Config_r16_Sub_SlotConfig_NCP_r16_N7},
}

const (
	SubSlot_Config_r16_Sub_SlotConfig_ECP_r16_N4 = 0
	SubSlot_Config_r16_Sub_SlotConfig_ECP_r16_N5 = 1
	SubSlot_Config_r16_Sub_SlotConfig_ECP_r16_N6 = 2
)

var subSlotConfigR16SubSlotConfigECPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SubSlot_Config_r16_Sub_SlotConfig_ECP_r16_N4, SubSlot_Config_r16_Sub_SlotConfig_ECP_r16_N5, SubSlot_Config_r16_Sub_SlotConfig_ECP_r16_N6},
}

type SubSlot_Config_r16 struct {
	Sub_SlotConfig_NCP_r16 *int64
	Sub_SlotConfig_ECP_r16 *int64
}

func (ie *SubSlot_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(subSlotConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sub_SlotConfig_NCP_r16 != nil, ie.Sub_SlotConfig_ECP_r16 != nil}); err != nil {
		return err
	}

	// 2. sub-SlotConfig-NCP-r16: enumerated
	{
		if ie.Sub_SlotConfig_NCP_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sub_SlotConfig_NCP_r16, subSlotConfigR16SubSlotConfigNCPR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. sub-SlotConfig-ECP-r16: enumerated
	{
		if ie.Sub_SlotConfig_ECP_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sub_SlotConfig_ECP_r16, subSlotConfigR16SubSlotConfigECPR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SubSlot_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(subSlotConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sub-SlotConfig-NCP-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(subSlotConfigR16SubSlotConfigNCPR16Constraints)
			if err != nil {
				return err
			}
			ie.Sub_SlotConfig_NCP_r16 = &idx
		}
	}

	// 3. sub-SlotConfig-ECP-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(subSlotConfigR16SubSlotConfigECPR16Constraints)
			if err != nil {
				return err
			}
			ie.Sub_SlotConfig_ECP_r16 = &idx
		}
	}

	return nil
}
