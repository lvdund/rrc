// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-QoS-Info-v1800 (line 2282).

var sLQoSInfoV1800Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-TxInterestedFreqList-r18"},
		{Name: "sl-TxProfile-r18", Optional: true},
	},
}

type SL_QoS_Info_v1800 struct {
	Sl_TxInterestedFreqList_r18 SL_TxInterestedFreqList_r16
	Sl_TxProfile_r18            *SL_TxProfile_r18
}

func (ie *SL_QoS_Info_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLQoSInfoV1800Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_TxProfile_r18 != nil}); err != nil {
		return err
	}

	// 3. sl-TxInterestedFreqList-r18: ref
	{
		if err := ie.Sl_TxInterestedFreqList_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-TxProfile-r18: ref
	{
		if ie.Sl_TxProfile_r18 != nil {
			if err := ie.Sl_TxProfile_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_QoS_Info_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLQoSInfoV1800Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-TxInterestedFreqList-r18: ref
	{
		if err := ie.Sl_TxInterestedFreqList_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-TxProfile-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_TxProfile_r18 = new(SL_TxProfile_r18)
			if err := ie.Sl_TxProfile_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
