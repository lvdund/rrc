// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RefLocOffset-r19 (line 4099).

var refLocOffsetR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "offsetLon-r19", Optional: true},
		{Name: "offsetLat-r19", Optional: true},
	},
}

var refLocOffsetR19OffsetLonR19Constraints = per.Constrained(-32766, 32766)

var refLocOffsetR19OffsetLatR19Constraints = per.Constrained(-32766, 32766)

type RefLocOffset_r19 struct {
	OffsetLon_r19 *int64
	OffsetLat_r19 *int64
}

func (ie *RefLocOffset_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(refLocOffsetR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OffsetLon_r19 != nil && *ie.OffsetLon_r19 != 0, ie.OffsetLat_r19 != nil && *ie.OffsetLat_r19 != 0}); err != nil {
		return err
	}

	// 2. offsetLon-r19: integer
	{
		if ie.OffsetLon_r19 != nil && *ie.OffsetLon_r19 != 0 {
			if err := e.EncodeInteger(*ie.OffsetLon_r19, refLocOffsetR19OffsetLonR19Constraints); err != nil {
				return err
			}
		}
	}

	// 3. offsetLat-r19: integer
	{
		if ie.OffsetLat_r19 != nil && *ie.OffsetLat_r19 != 0 {
			if err := e.EncodeInteger(*ie.OffsetLat_r19, refLocOffsetR19OffsetLatR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RefLocOffset_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(refLocOffsetR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. offsetLon-r19: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(refLocOffsetR19OffsetLonR19Constraints)
			if err != nil {
				return err
			}
			ie.OffsetLon_r19 = &v
		} else {
			v := int64(0)
			ie.OffsetLon_r19 = &v
		}
	}

	// 3. offsetLat-r19: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(refLocOffsetR19OffsetLatR19Constraints)
			if err != nil {
				return err
			}
			ie.OffsetLat_r19 = &v
		} else {
			v := int64(0)
			ie.OffsetLat_r19 = &v
		}
	}

	return nil
}
