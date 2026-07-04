// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: T-Reassembly (line 14120).

const (
	T_Reassembly_Ms0    = 0
	T_Reassembly_Ms5    = 1
	T_Reassembly_Ms10   = 2
	T_Reassembly_Ms15   = 3
	T_Reassembly_Ms20   = 4
	T_Reassembly_Ms25   = 5
	T_Reassembly_Ms30   = 6
	T_Reassembly_Ms35   = 7
	T_Reassembly_Ms40   = 8
	T_Reassembly_Ms45   = 9
	T_Reassembly_Ms50   = 10
	T_Reassembly_Ms55   = 11
	T_Reassembly_Ms60   = 12
	T_Reassembly_Ms65   = 13
	T_Reassembly_Ms70   = 14
	T_Reassembly_Ms75   = 15
	T_Reassembly_Ms80   = 16
	T_Reassembly_Ms85   = 17
	T_Reassembly_Ms90   = 18
	T_Reassembly_Ms95   = 19
	T_Reassembly_Ms100  = 20
	T_Reassembly_Ms110  = 21
	T_Reassembly_Ms120  = 22
	T_Reassembly_Ms130  = 23
	T_Reassembly_Ms140  = 24
	T_Reassembly_Ms150  = 25
	T_Reassembly_Ms160  = 26
	T_Reassembly_Ms170  = 27
	T_Reassembly_Ms180  = 28
	T_Reassembly_Ms190  = 29
	T_Reassembly_Ms200  = 30
	T_Reassembly_Spare1 = 31
)

var tReassemblyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{T_Reassembly_Ms0, T_Reassembly_Ms5, T_Reassembly_Ms10, T_Reassembly_Ms15, T_Reassembly_Ms20, T_Reassembly_Ms25, T_Reassembly_Ms30, T_Reassembly_Ms35, T_Reassembly_Ms40, T_Reassembly_Ms45, T_Reassembly_Ms50, T_Reassembly_Ms55, T_Reassembly_Ms60, T_Reassembly_Ms65, T_Reassembly_Ms70, T_Reassembly_Ms75, T_Reassembly_Ms80, T_Reassembly_Ms85, T_Reassembly_Ms90, T_Reassembly_Ms95, T_Reassembly_Ms100, T_Reassembly_Ms110, T_Reassembly_Ms120, T_Reassembly_Ms130, T_Reassembly_Ms140, T_Reassembly_Ms150, T_Reassembly_Ms160, T_Reassembly_Ms170, T_Reassembly_Ms180, T_Reassembly_Ms190, T_Reassembly_Ms200, T_Reassembly_Spare1},
}

type T_Reassembly struct {
	Value int64
}

func (v *T_Reassembly) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, tReassemblyConstraints)
}

func (v *T_Reassembly) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(tReassemblyConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
