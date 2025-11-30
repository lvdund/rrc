package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB_ReqInfo_r16_Enum_sib12       aper.Enumerated = 0
	SIB_ReqInfo_r16_Enum_sib13       aper.Enumerated = 1
	SIB_ReqInfo_r16_Enum_sib14       aper.Enumerated = 2
	SIB_ReqInfo_r16_Enum_sib20_v1700 aper.Enumerated = 3
	SIB_ReqInfo_r16_Enum_sib21_v1700 aper.Enumerated = 4
	SIB_ReqInfo_r16_Enum_spare3      aper.Enumerated = 5
	SIB_ReqInfo_r16_Enum_spare2      aper.Enumerated = 6
	SIB_ReqInfo_r16_Enum_spare1      aper.Enumerated = 7
)

type SIB_ReqInfo_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SIB_ReqInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SIB_ReqInfo_r16", err)
	}
	return nil
}

func (ie *SIB_ReqInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SIB_ReqInfo_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
