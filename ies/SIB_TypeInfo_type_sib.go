package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB_TypeInfo_type_sib_Enum_sibType2        aper.Enumerated = 0
	SIB_TypeInfo_type_sib_Enum_sibType3        aper.Enumerated = 1
	SIB_TypeInfo_type_sib_Enum_sibType4        aper.Enumerated = 2
	SIB_TypeInfo_type_sib_Enum_sibType5        aper.Enumerated = 3
	SIB_TypeInfo_type_sib_Enum_sibType6        aper.Enumerated = 4
	SIB_TypeInfo_type_sib_Enum_sibType7        aper.Enumerated = 5
	SIB_TypeInfo_type_sib_Enum_sibType8        aper.Enumerated = 6
	SIB_TypeInfo_type_sib_Enum_sibType9        aper.Enumerated = 7
	SIB_TypeInfo_type_sib_Enum_sibType10_v1610 aper.Enumerated = 8
	SIB_TypeInfo_type_sib_Enum_sibType11_v1610 aper.Enumerated = 9
	SIB_TypeInfo_type_sib_Enum_sibType12_v1610 aper.Enumerated = 10
	SIB_TypeInfo_type_sib_Enum_sibType13_v1610 aper.Enumerated = 11
	SIB_TypeInfo_type_sib_Enum_sibType14_v1610 aper.Enumerated = 12
	SIB_TypeInfo_type_sib_Enum_spare3          aper.Enumerated = 13
	SIB_TypeInfo_type_sib_Enum_spare2          aper.Enumerated = 14
	SIB_TypeInfo_type_sib_Enum_spare1          aper.Enumerated = 15
)

type SIB_TypeInfo_type_sib struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SIB_TypeInfo_type_sib) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SIB_TypeInfo_type_sib", err)
	}
	return nil
}

func (ie *SIB_TypeInfo_type_sib) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SIB_TypeInfo_type_sib", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
