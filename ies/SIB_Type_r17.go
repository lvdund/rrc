package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB_Type_r17_Enum_sibType2        aper.Enumerated = 0
	SIB_Type_r17_Enum_sibType3        aper.Enumerated = 1
	SIB_Type_r17_Enum_sibType4        aper.Enumerated = 2
	SIB_Type_r17_Enum_sibType5        aper.Enumerated = 3
	SIB_Type_r17_Enum_sibType9        aper.Enumerated = 4
	SIB_Type_r17_Enum_sibType10_v1610 aper.Enumerated = 5
	SIB_Type_r17_Enum_sibType11_v1610 aper.Enumerated = 6
	SIB_Type_r17_Enum_sibType12_v1610 aper.Enumerated = 7
	SIB_Type_r17_Enum_sibType13_v1610 aper.Enumerated = 8
	SIB_Type_r17_Enum_sibType14_v1610 aper.Enumerated = 9
	SIB_Type_r17_Enum_spare6          aper.Enumerated = 10
	SIB_Type_r17_Enum_spare5          aper.Enumerated = 11
	SIB_Type_r17_Enum_spare4          aper.Enumerated = 12
	SIB_Type_r17_Enum_spare3          aper.Enumerated = 13
	SIB_Type_r17_Enum_spare2          aper.Enumerated = 14
	SIB_Type_r17_Enum_spare1          aper.Enumerated = 15
)

type SIB_Type_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SIB_Type_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SIB_Type_r17", err)
	}
	return nil
}

func (ie *SIB_Type_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SIB_Type_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
