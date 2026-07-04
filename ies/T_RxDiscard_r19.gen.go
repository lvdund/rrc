// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: T-RxDiscard-r19 (line 14181).

const (
	T_RxDiscard_r19_Ms10    = 0
	T_RxDiscard_r19_Ms20    = 1
	T_RxDiscard_r19_Ms30    = 2
	T_RxDiscard_r19_Ms40    = 3
	T_RxDiscard_r19_Ms50    = 4
	T_RxDiscard_r19_Ms60    = 5
	T_RxDiscard_r19_Ms75    = 6
	T_RxDiscard_r19_Ms100   = 7
	T_RxDiscard_r19_Ms150   = 8
	T_RxDiscard_r19_Ms200   = 9
	T_RxDiscard_r19_Ms250   = 10
	T_RxDiscard_r19_Ms300   = 11
	T_RxDiscard_r19_Ms400   = 12
	T_RxDiscard_r19_Ms500   = 13
	T_RxDiscard_r19_Ms750   = 14
	T_RxDiscard_r19_Ms1000  = 15
	T_RxDiscard_r19_Ms1500  = 16
	T_RxDiscard_r19_Ms2000  = 17
	T_RxDiscard_r19_Ms3000  = 18
	T_RxDiscard_r19_Spare13 = 19
	T_RxDiscard_r19_Spare12 = 20
	T_RxDiscard_r19_Spare11 = 21
	T_RxDiscard_r19_Spare10 = 22
	T_RxDiscard_r19_Spare9  = 23
	T_RxDiscard_r19_Spare8  = 24
	T_RxDiscard_r19_Spare7  = 25
	T_RxDiscard_r19_Spare6  = 26
	T_RxDiscard_r19_Spare5  = 27
	T_RxDiscard_r19_Spare4  = 28
	T_RxDiscard_r19_Spare3  = 29
	T_RxDiscard_r19_Spare2  = 30
	T_RxDiscard_r19_Spare1  = 31
)

var tRxDiscardR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{T_RxDiscard_r19_Ms10, T_RxDiscard_r19_Ms20, T_RxDiscard_r19_Ms30, T_RxDiscard_r19_Ms40, T_RxDiscard_r19_Ms50, T_RxDiscard_r19_Ms60, T_RxDiscard_r19_Ms75, T_RxDiscard_r19_Ms100, T_RxDiscard_r19_Ms150, T_RxDiscard_r19_Ms200, T_RxDiscard_r19_Ms250, T_RxDiscard_r19_Ms300, T_RxDiscard_r19_Ms400, T_RxDiscard_r19_Ms500, T_RxDiscard_r19_Ms750, T_RxDiscard_r19_Ms1000, T_RxDiscard_r19_Ms1500, T_RxDiscard_r19_Ms2000, T_RxDiscard_r19_Ms3000, T_RxDiscard_r19_Spare13, T_RxDiscard_r19_Spare12, T_RxDiscard_r19_Spare11, T_RxDiscard_r19_Spare10, T_RxDiscard_r19_Spare9, T_RxDiscard_r19_Spare8, T_RxDiscard_r19_Spare7, T_RxDiscard_r19_Spare6, T_RxDiscard_r19_Spare5, T_RxDiscard_r19_Spare4, T_RxDiscard_r19_Spare3, T_RxDiscard_r19_Spare2, T_RxDiscard_r19_Spare1},
}

type T_RxDiscard_r19 struct {
	Value int64
}

func (v *T_RxDiscard_r19) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, tRxDiscardR19Constraints)
}

func (v *T_RxDiscard_r19) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(tRxDiscardR19Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
