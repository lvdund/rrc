// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PagingRecord-v1800 (line 874).

var pagingRecordV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mt-SDT", Optional: true},
	},
}

const (
	PagingRecord_v1800_Mt_SDT_True = 0
)

var pagingRecordV1800MtSDTConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PagingRecord_v1800_Mt_SDT_True},
}

type PagingRecord_v1800 struct {
	Mt_SDT *int64
}

func (ie *PagingRecord_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pagingRecordV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mt_SDT != nil}); err != nil {
		return err
	}

	// 2. mt-SDT: enumerated
	{
		if ie.Mt_SDT != nil {
			if err := e.EncodeEnumerated(*ie.Mt_SDT, pagingRecordV1800MtSDTConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PagingRecord_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pagingRecordV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mt-SDT: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pagingRecordV1800MtSDTConstraints)
			if err != nil {
				return err
			}
			ie.Mt_SDT = &idx
		}
	}

	return nil
}
