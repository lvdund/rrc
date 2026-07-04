// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MulticastRLC-BearerConfig-r17 (line 14038).

var multicastRLCBearerConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "servedMBS-RadioBearer-r17"},
		{Name: "isPTM-Entity-r17", Optional: true},
	},
}

const (
	MulticastRLC_BearerConfig_r17_IsPTM_Entity_r17_True = 0
)

var multicastRLCBearerConfigR17IsPTMEntityR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MulticastRLC_BearerConfig_r17_IsPTM_Entity_r17_True},
}

type MulticastRLC_BearerConfig_r17 struct {
	ServedMBS_RadioBearer_r17 MRB_Identity_r17
	IsPTM_Entity_r17          *int64
}

func (ie *MulticastRLC_BearerConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(multicastRLCBearerConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IsPTM_Entity_r17 != nil}); err != nil {
		return err
	}

	// 2. servedMBS-RadioBearer-r17: ref
	{
		if err := ie.ServedMBS_RadioBearer_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. isPTM-Entity-r17: enumerated
	{
		if ie.IsPTM_Entity_r17 != nil {
			if err := e.EncodeEnumerated(*ie.IsPTM_Entity_r17, multicastRLCBearerConfigR17IsPTMEntityR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MulticastRLC_BearerConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(multicastRLCBearerConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. servedMBS-RadioBearer-r17: ref
	{
		if err := ie.ServedMBS_RadioBearer_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. isPTM-Entity-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(multicastRLCBearerConfigR17IsPTMEntityR17Constraints)
			if err != nil {
				return err
			}
			ie.IsPTM_Entity_r17 = &idx
		}
	}

	return nil
}
