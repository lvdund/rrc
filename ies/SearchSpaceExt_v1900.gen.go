// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SearchSpaceExt-v1900 (line 14537).

var searchSpaceExtV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "searchSpaceLinkingId-CE-r19", Optional: true},
	},
}

var searchSpaceExtV1900SearchSpaceLinkingIdCER19Constraints = per.Constrained(0, common.MaxNrofSearchSpacesLinks_1_r17)

type SearchSpaceExt_v1900 struct {
	SearchSpaceLinkingId_CE_r19 *int64
}

func (ie *SearchSpaceExt_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(searchSpaceExtV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SearchSpaceLinkingId_CE_r19 != nil}); err != nil {
		return err
	}

	// 2. searchSpaceLinkingId-CE-r19: integer
	{
		if ie.SearchSpaceLinkingId_CE_r19 != nil {
			if err := e.EncodeInteger(*ie.SearchSpaceLinkingId_CE_r19, searchSpaceExtV1900SearchSpaceLinkingIdCER19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SearchSpaceExt_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(searchSpaceExtV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. searchSpaceLinkingId-CE-r19: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(searchSpaceExtV1900SearchSpaceLinkingIdCER19Constraints)
			if err != nil {
				return err
			}
			ie.SearchSpaceLinkingId_CE_r19 = &v
		}
	}

	return nil
}
