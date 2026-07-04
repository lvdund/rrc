// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-PPW-PreConfig-r17 (line 7643).

var dLPPWPreConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-PPW-ID-r17"},
		{Name: "dl-PPW-PeriodicityAndStartSlot-r17"},
		{Name: "length-r17"},
		{Name: "type-r17", Optional: true},
		{Name: "priority-r17", Optional: true},
	},
}

var dLPPWPreConfigR17LengthR17Constraints = per.Constrained(1, 160)

const (
	DL_PPW_PreConfig_r17_Type_r17_Type1A = 0
	DL_PPW_PreConfig_r17_Type_r17_Type1B = 1
	DL_PPW_PreConfig_r17_Type_r17_Type2  = 2
)

var dLPPWPreConfigR17TypeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DL_PPW_PreConfig_r17_Type_r17_Type1A, DL_PPW_PreConfig_r17_Type_r17_Type1B, DL_PPW_PreConfig_r17_Type_r17_Type2},
}

const (
	DL_PPW_PreConfig_r17_Priority_r17_St1 = 0
	DL_PPW_PreConfig_r17_Priority_r17_St2 = 1
	DL_PPW_PreConfig_r17_Priority_r17_St3 = 2
)

var dLPPWPreConfigR17PriorityR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DL_PPW_PreConfig_r17_Priority_r17_St1, DL_PPW_PreConfig_r17_Priority_r17_St2, DL_PPW_PreConfig_r17_Priority_r17_St3},
}

type DL_PPW_PreConfig_r17 struct {
	Dl_PPW_ID_r17                      DL_PPW_ID_r17
	Dl_PPW_PeriodicityAndStartSlot_r17 DL_PPW_PeriodicityAndStartSlot_r17
	Length_r17                         int64
	Type_r17                           *int64
	Priority_r17                       *int64
}

func (ie *DL_PPW_PreConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLPPWPreConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Type_r17 != nil, ie.Priority_r17 != nil}); err != nil {
		return err
	}

	// 2. dl-PPW-ID-r17: ref
	{
		if err := ie.Dl_PPW_ID_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. dl-PPW-PeriodicityAndStartSlot-r17: ref
	{
		if err := ie.Dl_PPW_PeriodicityAndStartSlot_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. length-r17: integer
	{
		if err := e.EncodeInteger(ie.Length_r17, dLPPWPreConfigR17LengthR17Constraints); err != nil {
			return err
		}
	}

	// 5. type-r17: enumerated
	{
		if ie.Type_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Type_r17, dLPPWPreConfigR17TypeR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. priority-r17: enumerated
	{
		if ie.Priority_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Priority_r17, dLPPWPreConfigR17PriorityR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DL_PPW_PreConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLPPWPreConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dl-PPW-ID-r17: ref
	{
		if err := ie.Dl_PPW_ID_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. dl-PPW-PeriodicityAndStartSlot-r17: ref
	{
		if err := ie.Dl_PPW_PeriodicityAndStartSlot_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. length-r17: integer
	{
		v2, err := d.DecodeInteger(dLPPWPreConfigR17LengthR17Constraints)
		if err != nil {
			return err
		}
		ie.Length_r17 = v2
	}

	// 5. type-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(dLPPWPreConfigR17TypeR17Constraints)
			if err != nil {
				return err
			}
			ie.Type_r17 = &idx
		}
	}

	// 6. priority-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(dLPPWPreConfigR17PriorityR17Constraints)
			if err != nil {
				return err
			}
			ie.Priority_r17 = &idx
		}
	}

	return nil
}
