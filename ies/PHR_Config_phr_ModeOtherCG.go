package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PHR_Config_phr_ModeOtherCG_Enum_real    aper.Enumerated = 0
	PHR_Config_phr_ModeOtherCG_Enum_virtual aper.Enumerated = 1
)

type PHR_Config_phr_ModeOtherCG struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PHR_Config_phr_ModeOtherCG) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PHR_Config_phr_ModeOtherCG", err)
	}
	return nil
}

func (ie *PHR_Config_phr_ModeOtherCG) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PHR_Config_phr_ModeOtherCG", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
