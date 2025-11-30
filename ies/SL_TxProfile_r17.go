package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_TxProfile_r17_Enum_drx_Compatible   aper.Enumerated = 0
	SL_TxProfile_r17_Enum_drx_Incompatible aper.Enumerated = 1
	SL_TxProfile_r17_Enum_spare6           aper.Enumerated = 2
	SL_TxProfile_r17_Enum_spare5           aper.Enumerated = 3
	SL_TxProfile_r17_Enum_spare4           aper.Enumerated = 4
	SL_TxProfile_r17_Enum_spare3           aper.Enumerated = 5
	SL_TxProfile_r17_Enum_spare2           aper.Enumerated = 6
	SL_TxProfile_r17_Enum_spare1           aper.Enumerated = 7
)

type SL_TxProfile_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SL_TxProfile_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SL_TxProfile_r17", err)
	}
	return nil
}

func (ie *SL_TxProfile_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SL_TxProfile_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
