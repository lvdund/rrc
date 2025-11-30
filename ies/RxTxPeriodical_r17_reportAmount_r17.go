package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RxTxPeriodical_r17_reportAmount_r17_Enum_r1       aper.Enumerated = 0
	RxTxPeriodical_r17_reportAmount_r17_Enum_infinity aper.Enumerated = 1
	RxTxPeriodical_r17_reportAmount_r17_Enum_spare6   aper.Enumerated = 2
	RxTxPeriodical_r17_reportAmount_r17_Enum_spare5   aper.Enumerated = 3
	RxTxPeriodical_r17_reportAmount_r17_Enum_spare4   aper.Enumerated = 4
	RxTxPeriodical_r17_reportAmount_r17_Enum_spare3   aper.Enumerated = 5
	RxTxPeriodical_r17_reportAmount_r17_Enum_spare2   aper.Enumerated = 6
	RxTxPeriodical_r17_reportAmount_r17_Enum_spare1   aper.Enumerated = 7
)

type RxTxPeriodical_r17_reportAmount_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RxTxPeriodical_r17_reportAmount_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RxTxPeriodical_r17_reportAmount_r17", err)
	}
	return nil
}

func (ie *RxTxPeriodical_r17_reportAmount_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RxTxPeriodical_r17_reportAmount_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
