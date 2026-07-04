// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCRelease-v1650-IEs (line 1252).

var rRCReleaseV1650IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mpsPriorityIndication-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCRelease_v1650_IEs_MpsPriorityIndication_r16_True = 0
)

var rRCReleaseV1650IEsMpsPriorityIndicationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCRelease_v1650_IEs_MpsPriorityIndication_r16_True},
}

type RRCRelease_v1650_IEs struct {
	MpsPriorityIndication_r16 *int64
	NonCriticalExtension      *RRCRelease_v1710_IEs
}

func (ie *RRCRelease_v1650_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReleaseV1650IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MpsPriorityIndication_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. mpsPriorityIndication-r16: enumerated
	{
		if ie.MpsPriorityIndication_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MpsPriorityIndication_r16, rRCReleaseV1650IEsMpsPriorityIndicationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCRelease_v1650_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReleaseV1650IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mpsPriorityIndication-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rRCReleaseV1650IEsMpsPriorityIndicationR16Constraints)
			if err != nil {
				return err
			}
			ie.MpsPriorityIndication_r16 = &idx
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(RRCRelease_v1710_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
