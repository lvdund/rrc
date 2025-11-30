package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_TxResourceReqDisc_r17_sl_CastTypeDisc_r17_Enum_broadcast aper.Enumerated = 0
	SL_TxResourceReqDisc_r17_sl_CastTypeDisc_r17_Enum_groupcast aper.Enumerated = 1
	SL_TxResourceReqDisc_r17_sl_CastTypeDisc_r17_Enum_unicast   aper.Enumerated = 2
	SL_TxResourceReqDisc_r17_sl_CastTypeDisc_r17_Enum_spare1    aper.Enumerated = 3
)

type SL_TxResourceReqDisc_r17_sl_CastTypeDisc_r17 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SL_TxResourceReqDisc_r17_sl_CastTypeDisc_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SL_TxResourceReqDisc_r17_sl_CastTypeDisc_r17", err)
	}
	return nil
}

func (ie *SL_TxResourceReqDisc_r17_sl_CastTypeDisc_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SL_TxResourceReqDisc_r17_sl_CastTypeDisc_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
