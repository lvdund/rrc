// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PagingRecord-v1700 (line 870).

var pagingRecordV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pagingCause-r17", Optional: true},
	},
}

const (
	PagingRecord_v1700_PagingCause_r17_Voice = 0
)

var pagingRecordV1700PagingCauseR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PagingRecord_v1700_PagingCause_r17_Voice},
}

type PagingRecord_v1700 struct {
	PagingCause_r17 *int64
}

func (ie *PagingRecord_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pagingRecordV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PagingCause_r17 != nil}); err != nil {
		return err
	}

	// 2. pagingCause-r17: enumerated
	{
		if ie.PagingCause_r17 != nil {
			if err := e.EncodeEnumerated(*ie.PagingCause_r17, pagingRecordV1700PagingCauseR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PagingRecord_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pagingRecordV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pagingCause-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pagingRecordV1700PagingCauseR17Constraints)
			if err != nil {
				return err
			}
			ie.PagingCause_r17 = &idx
		}
	}

	return nil
}
