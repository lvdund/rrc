// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MsgA-ConfigCommon-r16 (line 10113).

var msgAConfigCommonR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rach-ConfigCommonTwoStepRA-r16"},
		{Name: "msgA-PUSCH-Config-r16", Optional: true},
	},
}

type MsgA_ConfigCommon_r16 struct {
	Rach_ConfigCommonTwoStepRA_r16 RACH_ConfigCommonTwoStepRA_r16
	MsgA_PUSCH_Config_r16          *MsgA_PUSCH_Config_r16
}

func (ie *MsgA_ConfigCommon_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(msgAConfigCommonR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MsgA_PUSCH_Config_r16 != nil}); err != nil {
		return err
	}

	// 2. rach-ConfigCommonTwoStepRA-r16: ref
	{
		if err := ie.Rach_ConfigCommonTwoStepRA_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. msgA-PUSCH-Config-r16: ref
	{
		if ie.MsgA_PUSCH_Config_r16 != nil {
			if err := ie.MsgA_PUSCH_Config_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MsgA_ConfigCommon_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(msgAConfigCommonR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. rach-ConfigCommonTwoStepRA-r16: ref
	{
		if err := ie.Rach_ConfigCommonTwoStepRA_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. msgA-PUSCH-Config-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MsgA_PUSCH_Config_r16 = new(MsgA_PUSCH_Config_r16)
			if err := ie.MsgA_PUSCH_Config_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
