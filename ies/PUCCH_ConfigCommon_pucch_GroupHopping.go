package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_ConfigCommon_pucch_GroupHopping_Enum_neither aper.Enumerated = 0
	PUCCH_ConfigCommon_pucch_GroupHopping_Enum_enable  aper.Enumerated = 1
	PUCCH_ConfigCommon_pucch_GroupHopping_Enum_disable aper.Enumerated = 2
)

type PUCCH_ConfigCommon_pucch_GroupHopping struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PUCCH_ConfigCommon_pucch_GroupHopping) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PUCCH_ConfigCommon_pucch_GroupHopping", err)
	}
	return nil
}

func (ie *PUCCH_ConfigCommon_pucch_GroupHopping) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PUCCH_ConfigCommon_pucch_GroupHopping", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
