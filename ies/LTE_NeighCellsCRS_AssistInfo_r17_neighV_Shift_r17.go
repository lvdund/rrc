package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17_Enum_n0 aper.Enumerated = 0
	LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17_Enum_n1 aper.Enumerated = 1
	LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17_Enum_n2 aper.Enumerated = 2
	LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17_Enum_n3 aper.Enumerated = 3
	LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17_Enum_n4 aper.Enumerated = 4
	LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17_Enum_n5 aper.Enumerated = 5
)

type LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17 struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17", err)
	}
	return nil
}

func (ie *LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
