package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates_Enum_twoStates aper.Enumerated = 0
)

type PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates", err)
	}
	return nil
}

func (ie *PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
