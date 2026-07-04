// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SPS-PUCCH-AN-r16 (line 15263).

var sPSPUCCHANR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sps-PUCCH-AN-ResourceID-r16"},
		{Name: "maxPayloadSize-r16", Optional: true},
	},
}

var sPSPUCCHANR16MaxPayloadSizeR16Constraints = per.Constrained(4, 256)

type SPS_PUCCH_AN_r16 struct {
	Sps_PUCCH_AN_ResourceID_r16 PUCCH_ResourceId
	MaxPayloadSize_r16          *int64
}

func (ie *SPS_PUCCH_AN_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sPSPUCCHANR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxPayloadSize_r16 != nil}); err != nil {
		return err
	}

	// 2. sps-PUCCH-AN-ResourceID-r16: ref
	{
		if err := ie.Sps_PUCCH_AN_ResourceID_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. maxPayloadSize-r16: integer
	{
		if ie.MaxPayloadSize_r16 != nil {
			if err := e.EncodeInteger(*ie.MaxPayloadSize_r16, sPSPUCCHANR16MaxPayloadSizeR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SPS_PUCCH_AN_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sPSPUCCHANR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sps-PUCCH-AN-ResourceID-r16: ref
	{
		if err := ie.Sps_PUCCH_AN_ResourceID_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. maxPayloadSize-r16: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sPSPUCCHANR16MaxPayloadSizeR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxPayloadSize_r16 = &v
		}
	}

	return nil
}
