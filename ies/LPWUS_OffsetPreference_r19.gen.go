// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LPWUS-OffsetPreference-r19 (line 2682).

var lPWUSOffsetPreferenceR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "timeOffset-r19", Optional: true},
	},
}

const (
	LPWUS_OffsetPreference_r19_TimeOffset_r19_Ms5  = 0
	LPWUS_OffsetPreference_r19_TimeOffset_r19_Ms13 = 1
	LPWUS_OffsetPreference_r19_TimeOffset_r19_Ms37 = 2
)

var lPWUSOffsetPreferenceR19TimeOffsetR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LPWUS_OffsetPreference_r19_TimeOffset_r19_Ms5, LPWUS_OffsetPreference_r19_TimeOffset_r19_Ms13, LPWUS_OffsetPreference_r19_TimeOffset_r19_Ms37},
}

type LPWUS_OffsetPreference_r19 struct {
	TimeOffset_r19 *int64
}

func (ie *LPWUS_OffsetPreference_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lPWUSOffsetPreferenceR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.TimeOffset_r19 != nil}); err != nil {
		return err
	}

	// 2. timeOffset-r19: enumerated
	{
		if ie.TimeOffset_r19 != nil {
			if err := e.EncodeEnumerated(*ie.TimeOffset_r19, lPWUSOffsetPreferenceR19TimeOffsetR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LPWUS_OffsetPreference_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lPWUSOffsetPreferenceR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. timeOffset-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(lPWUSOffsetPreferenceR19TimeOffsetR19Constraints)
			if err != nil {
				return err
			}
			ie.TimeOffset_r19 = &idx
		}
	}

	return nil
}
