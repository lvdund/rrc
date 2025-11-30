package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_PDCP_ConfigPC5_r16_sl_PDCP_SN_Size_r16_Enum_len12bits aper.Enumerated = 0
	SL_PDCP_ConfigPC5_r16_sl_PDCP_SN_Size_r16_Enum_len18bits aper.Enumerated = 1
)

type SL_PDCP_ConfigPC5_r16_sl_PDCP_SN_Size_r16 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SL_PDCP_ConfigPC5_r16_sl_PDCP_SN_Size_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SL_PDCP_ConfigPC5_r16_sl_PDCP_SN_Size_r16", err)
	}
	return nil
}

func (ie *SL_PDCP_ConfigPC5_r16_sl_PDCP_SN_Size_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SL_PDCP_ConfigPC5_r16_sl_PDCP_SN_Size_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
