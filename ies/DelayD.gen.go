// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DelayD (line 7302).
// DelayD ::=                          ENUMERATED { symb4, slot1, slot2, slot3, slot4, slot5, slot6, slot10 }

const (
	DelayD_Symb4  = 0
	DelayD_Slot1  = 1
	DelayD_Slot2  = 2
	DelayD_Slot3  = 3
	DelayD_Slot4  = 4
	DelayD_Slot5  = 5
	DelayD_Slot6  = 6
	DelayD_Slot10 = 7
)

var delayDConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DelayD_Symb4, DelayD_Slot1, DelayD_Slot2, DelayD_Slot3, DelayD_Slot4, DelayD_Slot5, DelayD_Slot6, DelayD_Slot10},
}

type DelayD struct {
	Value int64
}

func (v *DelayD) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, delayDConstraints)
}

func (v *DelayD) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(delayDConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
