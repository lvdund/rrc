package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_PDCP_Config_r16_sl_OutOfOrderDelivery_Enum_true aper.Enumerated = 0
)

type SL_PDCP_Config_r16_sl_OutOfOrderDelivery struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *SL_PDCP_Config_r16_sl_OutOfOrderDelivery) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode SL_PDCP_Config_r16_sl_OutOfOrderDelivery", err)
	}
	return nil
}

func (ie *SL_PDCP_Config_r16_sl_OutOfOrderDelivery) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode SL_PDCP_Config_r16_sl_OutOfOrderDelivery", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
