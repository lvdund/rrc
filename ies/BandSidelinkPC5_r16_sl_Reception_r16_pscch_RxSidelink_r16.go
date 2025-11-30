package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandSidelinkPC5_r16_sl_Reception_r16_pscch_RxSidelink_r16_Enum_value1 aper.Enumerated = 0
	BandSidelinkPC5_r16_sl_Reception_r16_pscch_RxSidelink_r16_Enum_value2 aper.Enumerated = 1
)

type BandSidelinkPC5_r16_sl_Reception_r16_pscch_RxSidelink_r16 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *BandSidelinkPC5_r16_sl_Reception_r16_pscch_RxSidelink_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode BandSidelinkPC5_r16_sl_Reception_r16_pscch_RxSidelink_r16", err)
	}
	return nil
}

func (ie *BandSidelinkPC5_r16_sl_Reception_r16_pscch_RxSidelink_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode BandSidelinkPC5_r16_sl_Reception_r16_pscch_RxSidelink_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
