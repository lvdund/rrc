package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	T_Reassembly_Enum_ms0    aper.Enumerated = 0
	T_Reassembly_Enum_ms5    aper.Enumerated = 1
	T_Reassembly_Enum_ms10   aper.Enumerated = 2
	T_Reassembly_Enum_ms15   aper.Enumerated = 3
	T_Reassembly_Enum_ms20   aper.Enumerated = 4
	T_Reassembly_Enum_ms25   aper.Enumerated = 5
	T_Reassembly_Enum_ms30   aper.Enumerated = 6
	T_Reassembly_Enum_ms35   aper.Enumerated = 7
	T_Reassembly_Enum_ms40   aper.Enumerated = 8
	T_Reassembly_Enum_ms45   aper.Enumerated = 9
	T_Reassembly_Enum_ms50   aper.Enumerated = 10
	T_Reassembly_Enum_ms55   aper.Enumerated = 11
	T_Reassembly_Enum_ms60   aper.Enumerated = 12
	T_Reassembly_Enum_ms65   aper.Enumerated = 13
	T_Reassembly_Enum_ms70   aper.Enumerated = 14
	T_Reassembly_Enum_ms75   aper.Enumerated = 15
	T_Reassembly_Enum_ms80   aper.Enumerated = 16
	T_Reassembly_Enum_ms85   aper.Enumerated = 17
	T_Reassembly_Enum_ms90   aper.Enumerated = 18
	T_Reassembly_Enum_ms95   aper.Enumerated = 19
	T_Reassembly_Enum_ms100  aper.Enumerated = 20
	T_Reassembly_Enum_ms110  aper.Enumerated = 21
	T_Reassembly_Enum_ms120  aper.Enumerated = 22
	T_Reassembly_Enum_ms130  aper.Enumerated = 23
	T_Reassembly_Enum_ms140  aper.Enumerated = 24
	T_Reassembly_Enum_ms150  aper.Enumerated = 25
	T_Reassembly_Enum_ms160  aper.Enumerated = 26
	T_Reassembly_Enum_ms170  aper.Enumerated = 27
	T_Reassembly_Enum_ms180  aper.Enumerated = 28
	T_Reassembly_Enum_ms190  aper.Enumerated = 29
	T_Reassembly_Enum_ms200  aper.Enumerated = 30
	T_Reassembly_Enum_spare1 aper.Enumerated = 31
)

type T_Reassembly struct {
	Value aper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *T_Reassembly) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode T_Reassembly", err)
	}
	return nil
}

func (ie *T_Reassembly) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode T_Reassembly", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
