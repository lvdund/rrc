package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	T_StatusProhibit_v1610_Enum_ms1    aper.Enumerated = 0
	T_StatusProhibit_v1610_Enum_ms2    aper.Enumerated = 1
	T_StatusProhibit_v1610_Enum_ms3    aper.Enumerated = 2
	T_StatusProhibit_v1610_Enum_ms4    aper.Enumerated = 3
	T_StatusProhibit_v1610_Enum_spare4 aper.Enumerated = 4
	T_StatusProhibit_v1610_Enum_spare3 aper.Enumerated = 5
	T_StatusProhibit_v1610_Enum_spare2 aper.Enumerated = 6
	T_StatusProhibit_v1610_Enum_spare1 aper.Enumerated = 7
)

type T_StatusProhibit_v1610 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *T_StatusProhibit_v1610) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode T_StatusProhibit_v1610", err)
	}
	return nil
}

func (ie *T_StatusProhibit_v1610) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode T_StatusProhibit_v1610", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
