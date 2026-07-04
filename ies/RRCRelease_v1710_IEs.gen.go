// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCRelease-v1710-IEs (line 1257).

var rRCReleaseV1710IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "noLastCellUpdate-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCRelease_v1710_IEs_NoLastCellUpdate_r17_True = 0
)

var rRCReleaseV1710IEsNoLastCellUpdateR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCRelease_v1710_IEs_NoLastCellUpdate_r17_True},
}

type RRCRelease_v1710_IEs struct {
	NoLastCellUpdate_r17 *int64
	NonCriticalExtension *struct{}
}

func (ie *RRCRelease_v1710_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReleaseV1710IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.NoLastCellUpdate_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. noLastCellUpdate-r17: enumerated
	{
		if ie.NoLastCellUpdate_r17 != nil {
			if err := e.EncodeEnumerated(*ie.NoLastCellUpdate_r17, rRCReleaseV1710IEsNoLastCellUpdateR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *RRCRelease_v1710_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReleaseV1710IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. noLastCellUpdate-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rRCReleaseV1710IEsNoLastCellUpdateR17Constraints)
			if err != nil {
				return err
			}
			ie.NoLastCellUpdate_r17 = &idx
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
