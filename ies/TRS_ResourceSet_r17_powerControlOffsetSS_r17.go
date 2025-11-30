package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	TRS_ResourceSet_r17_powerControlOffsetSS_r17_Enum_db_3 aper.Enumerated = 0
	TRS_ResourceSet_r17_powerControlOffsetSS_r17_Enum_db0  aper.Enumerated = 1
	TRS_ResourceSet_r17_powerControlOffsetSS_r17_Enum_db3  aper.Enumerated = 2
	TRS_ResourceSet_r17_powerControlOffsetSS_r17_Enum_db6  aper.Enumerated = 3
)

type TRS_ResourceSet_r17_powerControlOffsetSS_r17 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *TRS_ResourceSet_r17_powerControlOffsetSS_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode TRS_ResourceSet_r17_powerControlOffsetSS_r17", err)
	}
	return nil
}

func (ie *TRS_ResourceSet_r17_powerControlOffsetSS_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode TRS_ResourceSet_r17_powerControlOffsetSS_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
