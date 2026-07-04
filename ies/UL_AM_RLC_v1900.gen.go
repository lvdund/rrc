// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-AM-RLC-v1900 (line 14169).

var uLAMRLCV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "stopReTxDiscardedSDU-r19", Optional: true},
	},
}

const (
	UL_AM_RLC_v1900_StopReTxDiscardedSDU_r19_Enabled = 0
)

var uLAMRLCV1900StopReTxDiscardedSDUR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UL_AM_RLC_v1900_StopReTxDiscardedSDU_r19_Enabled},
}

type UL_AM_RLC_v1900 struct {
	StopReTxDiscardedSDU_r19 *int64
}

func (ie *UL_AM_RLC_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLAMRLCV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.StopReTxDiscardedSDU_r19 != nil}); err != nil {
		return err
	}

	// 2. stopReTxDiscardedSDU-r19: enumerated
	{
		if ie.StopReTxDiscardedSDU_r19 != nil {
			if err := e.EncodeEnumerated(*ie.StopReTxDiscardedSDU_r19, uLAMRLCV1900StopReTxDiscardedSDUR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UL_AM_RLC_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLAMRLCV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. stopReTxDiscardedSDU-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uLAMRLCV1900StopReTxDiscardedSDUR19Constraints)
			if err != nil {
				return err
			}
			ie.StopReTxDiscardedSDU_r19 = &idx
		}
	}

	return nil
}
