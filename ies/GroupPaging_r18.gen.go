// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: GroupPaging-r18 (line 884).

var groupPagingR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "inactiveReceptionAllowed-r18", Optional: true},
	},
}

const (
	GroupPaging_r18_InactiveReceptionAllowed_r18_True = 0
)

var groupPagingR18InactiveReceptionAllowedR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GroupPaging_r18_InactiveReceptionAllowed_r18_True},
}

type GroupPaging_r18 struct {
	InactiveReceptionAllowed_r18 *int64
}

func (ie *GroupPaging_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(groupPagingR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InactiveReceptionAllowed_r18 != nil}); err != nil {
		return err
	}

	// 2. inactiveReceptionAllowed-r18: enumerated
	{
		if ie.InactiveReceptionAllowed_r18 != nil {
			if err := e.EncodeEnumerated(*ie.InactiveReceptionAllowed_r18, groupPagingR18InactiveReceptionAllowedR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *GroupPaging_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(groupPagingR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. inactiveReceptionAllowed-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(groupPagingR18InactiveReceptionAllowedR18Constraints)
			if err != nil {
				return err
			}
			ie.InactiveReceptionAllowed_r18 = &idx
		}
	}

	return nil
}
