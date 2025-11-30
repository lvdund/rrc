package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PLMN_IdentityInfo_cellReservedForOperatorUse_Enum_reserved    aper.Enumerated = 0
	PLMN_IdentityInfo_cellReservedForOperatorUse_Enum_notReserved aper.Enumerated = 1
)

type PLMN_IdentityInfo_cellReservedForOperatorUse struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PLMN_IdentityInfo_cellReservedForOperatorUse) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PLMN_IdentityInfo_cellReservedForOperatorUse", err)
	}
	return nil
}

func (ie *PLMN_IdentityInfo_cellReservedForOperatorUse) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PLMN_IdentityInfo_cellReservedForOperatorUse", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
