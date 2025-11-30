package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_TxResourceReq_v1700_sl_DRX_Indication_r17_Enum_on  aper.Enumerated = 0
	SL_TxResourceReq_v1700_sl_DRX_Indication_r17_Enum_off aper.Enumerated = 1
)

type SL_TxResourceReq_v1700_sl_DRX_Indication_r17 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SL_TxResourceReq_v1700_sl_DRX_Indication_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SL_TxResourceReq_v1700_sl_DRX_Indication_r17", err)
	}
	return nil
}

func (ie *SL_TxResourceReq_v1700_sl_DRX_Indication_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SL_TxResourceReq_v1700_sl_DRX_Indication_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
