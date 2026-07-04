// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: T-StatusProhibit (line 14127).

const (
	T_StatusProhibit_Ms0    = 0
	T_StatusProhibit_Ms5    = 1
	T_StatusProhibit_Ms10   = 2
	T_StatusProhibit_Ms15   = 3
	T_StatusProhibit_Ms20   = 4
	T_StatusProhibit_Ms25   = 5
	T_StatusProhibit_Ms30   = 6
	T_StatusProhibit_Ms35   = 7
	T_StatusProhibit_Ms40   = 8
	T_StatusProhibit_Ms45   = 9
	T_StatusProhibit_Ms50   = 10
	T_StatusProhibit_Ms55   = 11
	T_StatusProhibit_Ms60   = 12
	T_StatusProhibit_Ms65   = 13
	T_StatusProhibit_Ms70   = 14
	T_StatusProhibit_Ms75   = 15
	T_StatusProhibit_Ms80   = 16
	T_StatusProhibit_Ms85   = 17
	T_StatusProhibit_Ms90   = 18
	T_StatusProhibit_Ms95   = 19
	T_StatusProhibit_Ms100  = 20
	T_StatusProhibit_Ms105  = 21
	T_StatusProhibit_Ms110  = 22
	T_StatusProhibit_Ms115  = 23
	T_StatusProhibit_Ms120  = 24
	T_StatusProhibit_Ms125  = 25
	T_StatusProhibit_Ms130  = 26
	T_StatusProhibit_Ms135  = 27
	T_StatusProhibit_Ms140  = 28
	T_StatusProhibit_Ms145  = 29
	T_StatusProhibit_Ms150  = 30
	T_StatusProhibit_Ms155  = 31
	T_StatusProhibit_Ms160  = 32
	T_StatusProhibit_Ms165  = 33
	T_StatusProhibit_Ms170  = 34
	T_StatusProhibit_Ms175  = 35
	T_StatusProhibit_Ms180  = 36
	T_StatusProhibit_Ms185  = 37
	T_StatusProhibit_Ms190  = 38
	T_StatusProhibit_Ms195  = 39
	T_StatusProhibit_Ms200  = 40
	T_StatusProhibit_Ms205  = 41
	T_StatusProhibit_Ms210  = 42
	T_StatusProhibit_Ms215  = 43
	T_StatusProhibit_Ms220  = 44
	T_StatusProhibit_Ms225  = 45
	T_StatusProhibit_Ms230  = 46
	T_StatusProhibit_Ms235  = 47
	T_StatusProhibit_Ms240  = 48
	T_StatusProhibit_Ms245  = 49
	T_StatusProhibit_Ms250  = 50
	T_StatusProhibit_Ms300  = 51
	T_StatusProhibit_Ms350  = 52
	T_StatusProhibit_Ms400  = 53
	T_StatusProhibit_Ms450  = 54
	T_StatusProhibit_Ms500  = 55
	T_StatusProhibit_Ms800  = 56
	T_StatusProhibit_Ms1000 = 57
	T_StatusProhibit_Ms1200 = 58
	T_StatusProhibit_Ms1600 = 59
	T_StatusProhibit_Ms2000 = 60
	T_StatusProhibit_Ms2400 = 61
	T_StatusProhibit_Spare2 = 62
	T_StatusProhibit_Spare1 = 63
)

var tStatusProhibitConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{T_StatusProhibit_Ms0, T_StatusProhibit_Ms5, T_StatusProhibit_Ms10, T_StatusProhibit_Ms15, T_StatusProhibit_Ms20, T_StatusProhibit_Ms25, T_StatusProhibit_Ms30, T_StatusProhibit_Ms35, T_StatusProhibit_Ms40, T_StatusProhibit_Ms45, T_StatusProhibit_Ms50, T_StatusProhibit_Ms55, T_StatusProhibit_Ms60, T_StatusProhibit_Ms65, T_StatusProhibit_Ms70, T_StatusProhibit_Ms75, T_StatusProhibit_Ms80, T_StatusProhibit_Ms85, T_StatusProhibit_Ms90, T_StatusProhibit_Ms95, T_StatusProhibit_Ms100, T_StatusProhibit_Ms105, T_StatusProhibit_Ms110, T_StatusProhibit_Ms115, T_StatusProhibit_Ms120, T_StatusProhibit_Ms125, T_StatusProhibit_Ms130, T_StatusProhibit_Ms135, T_StatusProhibit_Ms140, T_StatusProhibit_Ms145, T_StatusProhibit_Ms150, T_StatusProhibit_Ms155, T_StatusProhibit_Ms160, T_StatusProhibit_Ms165, T_StatusProhibit_Ms170, T_StatusProhibit_Ms175, T_StatusProhibit_Ms180, T_StatusProhibit_Ms185, T_StatusProhibit_Ms190, T_StatusProhibit_Ms195, T_StatusProhibit_Ms200, T_StatusProhibit_Ms205, T_StatusProhibit_Ms210, T_StatusProhibit_Ms215, T_StatusProhibit_Ms220, T_StatusProhibit_Ms225, T_StatusProhibit_Ms230, T_StatusProhibit_Ms235, T_StatusProhibit_Ms240, T_StatusProhibit_Ms245, T_StatusProhibit_Ms250, T_StatusProhibit_Ms300, T_StatusProhibit_Ms350, T_StatusProhibit_Ms400, T_StatusProhibit_Ms450, T_StatusProhibit_Ms500, T_StatusProhibit_Ms800, T_StatusProhibit_Ms1000, T_StatusProhibit_Ms1200, T_StatusProhibit_Ms1600, T_StatusProhibit_Ms2000, T_StatusProhibit_Ms2400, T_StatusProhibit_Spare2, T_StatusProhibit_Spare1},
}

type T_StatusProhibit struct {
	Value int64
}

func (v *T_StatusProhibit) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, tStatusProhibitConstraints)
}

func (v *T_StatusProhibit) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(tStatusProhibitConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
