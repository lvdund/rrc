// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MsgA-PUSCH-Config-r16 (line 10121).

var msgAPUSCHConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "msgA-PUSCH-ResourceGroupA-r16", Optional: true},
		{Name: "msgA-PUSCH-ResourceGroupB-r16", Optional: true},
		{Name: "msgA-TransformPrecoder-r16", Optional: true},
		{Name: "msgA-DataScramblingIndex-r16", Optional: true},
		{Name: "msgA-DeltaPreamble-r16", Optional: true},
	},
}

const (
	MsgA_PUSCH_Config_r16_MsgA_TransformPrecoder_r16_Enabled  = 0
	MsgA_PUSCH_Config_r16_MsgA_TransformPrecoder_r16_Disabled = 1
)

var msgAPUSCHConfigR16MsgATransformPrecoderR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MsgA_PUSCH_Config_r16_MsgA_TransformPrecoder_r16_Enabled, MsgA_PUSCH_Config_r16_MsgA_TransformPrecoder_r16_Disabled},
}

var msgAPUSCHConfigR16MsgADataScramblingIndexR16Constraints = per.Constrained(0, 1023)

var msgAPUSCHConfigR16MsgADeltaPreambleR16Constraints = per.Constrained(-1, 6)

type MsgA_PUSCH_Config_r16 struct {
	MsgA_PUSCH_ResourceGroupA_r16 *MsgA_PUSCH_Resource_r16
	MsgA_PUSCH_ResourceGroupB_r16 *MsgA_PUSCH_Resource_r16
	MsgA_TransformPrecoder_r16    *int64
	MsgA_DataScramblingIndex_r16  *int64
	MsgA_DeltaPreamble_r16        *int64
}

func (ie *MsgA_PUSCH_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(msgAPUSCHConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MsgA_PUSCH_ResourceGroupA_r16 != nil, ie.MsgA_PUSCH_ResourceGroupB_r16 != nil, ie.MsgA_TransformPrecoder_r16 != nil, ie.MsgA_DataScramblingIndex_r16 != nil, ie.MsgA_DeltaPreamble_r16 != nil}); err != nil {
		return err
	}

	// 2. msgA-PUSCH-ResourceGroupA-r16: ref
	{
		if ie.MsgA_PUSCH_ResourceGroupA_r16 != nil {
			if err := ie.MsgA_PUSCH_ResourceGroupA_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. msgA-PUSCH-ResourceGroupB-r16: ref
	{
		if ie.MsgA_PUSCH_ResourceGroupB_r16 != nil {
			if err := ie.MsgA_PUSCH_ResourceGroupB_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. msgA-TransformPrecoder-r16: enumerated
	{
		if ie.MsgA_TransformPrecoder_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MsgA_TransformPrecoder_r16, msgAPUSCHConfigR16MsgATransformPrecoderR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. msgA-DataScramblingIndex-r16: integer
	{
		if ie.MsgA_DataScramblingIndex_r16 != nil {
			if err := e.EncodeInteger(*ie.MsgA_DataScramblingIndex_r16, msgAPUSCHConfigR16MsgADataScramblingIndexR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. msgA-DeltaPreamble-r16: integer
	{
		if ie.MsgA_DeltaPreamble_r16 != nil {
			if err := e.EncodeInteger(*ie.MsgA_DeltaPreamble_r16, msgAPUSCHConfigR16MsgADeltaPreambleR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MsgA_PUSCH_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(msgAPUSCHConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. msgA-PUSCH-ResourceGroupA-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.MsgA_PUSCH_ResourceGroupA_r16 = new(MsgA_PUSCH_Resource_r16)
			if err := ie.MsgA_PUSCH_ResourceGroupA_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. msgA-PUSCH-ResourceGroupB-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MsgA_PUSCH_ResourceGroupB_r16 = new(MsgA_PUSCH_Resource_r16)
			if err := ie.MsgA_PUSCH_ResourceGroupB_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. msgA-TransformPrecoder-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(msgAPUSCHConfigR16MsgATransformPrecoderR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_TransformPrecoder_r16 = &idx
		}
	}

	// 5. msgA-DataScramblingIndex-r16: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(msgAPUSCHConfigR16MsgADataScramblingIndexR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_DataScramblingIndex_r16 = &v
		}
	}

	// 6. msgA-DeltaPreamble-r16: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(msgAPUSCHConfigR16MsgADeltaPreambleR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_DeltaPreamble_r16 = &v
		}
	}

	return nil
}
