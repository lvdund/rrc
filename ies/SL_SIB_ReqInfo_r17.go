package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_SIB_ReqInfo_r17_Enum_sib1    aper.Enumerated = 0
	SL_SIB_ReqInfo_r17_Enum_sib2    aper.Enumerated = 1
	SL_SIB_ReqInfo_r17_Enum_sib3    aper.Enumerated = 2
	SL_SIB_ReqInfo_r17_Enum_sib4    aper.Enumerated = 3
	SL_SIB_ReqInfo_r17_Enum_sib5    aper.Enumerated = 4
	SL_SIB_ReqInfo_r17_Enum_sib6    aper.Enumerated = 5
	SL_SIB_ReqInfo_r17_Enum_sib7    aper.Enumerated = 6
	SL_SIB_ReqInfo_r17_Enum_sib8    aper.Enumerated = 7
	SL_SIB_ReqInfo_r17_Enum_sib9    aper.Enumerated = 8
	SL_SIB_ReqInfo_r17_Enum_sib10   aper.Enumerated = 9
	SL_SIB_ReqInfo_r17_Enum_sib11   aper.Enumerated = 10
	SL_SIB_ReqInfo_r17_Enum_sib12   aper.Enumerated = 11
	SL_SIB_ReqInfo_r17_Enum_sib13   aper.Enumerated = 12
	SL_SIB_ReqInfo_r17_Enum_sib14   aper.Enumerated = 13
	SL_SIB_ReqInfo_r17_Enum_sib15   aper.Enumerated = 14
	SL_SIB_ReqInfo_r17_Enum_sib16   aper.Enumerated = 15
	SL_SIB_ReqInfo_r17_Enum_sib17   aper.Enumerated = 16
	SL_SIB_ReqInfo_r17_Enum_sib18   aper.Enumerated = 17
	SL_SIB_ReqInfo_r17_Enum_sib19   aper.Enumerated = 18
	SL_SIB_ReqInfo_r17_Enum_sib20   aper.Enumerated = 19
	SL_SIB_ReqInfo_r17_Enum_sib21   aper.Enumerated = 20
	SL_SIB_ReqInfo_r17_Enum_spare11 aper.Enumerated = 21
	SL_SIB_ReqInfo_r17_Enum_spare10 aper.Enumerated = 22
	SL_SIB_ReqInfo_r17_Enum_spare9  aper.Enumerated = 23
	SL_SIB_ReqInfo_r17_Enum_spare8  aper.Enumerated = 24
	SL_SIB_ReqInfo_r17_Enum_spare7  aper.Enumerated = 25
	SL_SIB_ReqInfo_r17_Enum_spare6  aper.Enumerated = 26
	SL_SIB_ReqInfo_r17_Enum_spare5  aper.Enumerated = 27
	SL_SIB_ReqInfo_r17_Enum_spare4  aper.Enumerated = 28
	SL_SIB_ReqInfo_r17_Enum_spare3  aper.Enumerated = 29
	SL_SIB_ReqInfo_r17_Enum_spare2  aper.Enumerated = 30
	SL_SIB_ReqInfo_r17_Enum_spare1  aper.Enumerated = 31
)

type SL_SIB_ReqInfo_r17 struct {
	Value aper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *SL_SIB_ReqInfo_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode SL_SIB_ReqInfo_r17", err)
	}
	return nil
}

func (ie *SL_SIB_ReqInfo_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode SL_SIB_ReqInfo_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
