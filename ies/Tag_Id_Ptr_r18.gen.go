// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Tag-Id-ptr-r18 (line 15991).
// Tag-Id-ptr-r18 ::= ENUMERATED {n0,n1}

const (
	Tag_Id_Ptr_r18_N0 = 0
	Tag_Id_Ptr_r18_N1 = 1
)

var tagIdPtrR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Tag_Id_Ptr_r18_N0, Tag_Id_Ptr_r18_N1},
}

type Tag_Id_Ptr_r18 struct {
	Value int64
}

func (v *Tag_Id_Ptr_r18) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, tagIdPtrR18Constraints)
}

func (v *Tag_Id_Ptr_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(tagIdPtrR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
