package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SI_RequestConfig_si_RequestPeriod_Enum_one     aper.Enumerated = 0
	SI_RequestConfig_si_RequestPeriod_Enum_two     aper.Enumerated = 1
	SI_RequestConfig_si_RequestPeriod_Enum_four    aper.Enumerated = 2
	SI_RequestConfig_si_RequestPeriod_Enum_six     aper.Enumerated = 3
	SI_RequestConfig_si_RequestPeriod_Enum_eight   aper.Enumerated = 4
	SI_RequestConfig_si_RequestPeriod_Enum_ten     aper.Enumerated = 5
	SI_RequestConfig_si_RequestPeriod_Enum_twelve  aper.Enumerated = 6
	SI_RequestConfig_si_RequestPeriod_Enum_sixteen aper.Enumerated = 7
)

type SI_RequestConfig_si_RequestPeriod struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SI_RequestConfig_si_RequestPeriod) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SI_RequestConfig_si_RequestPeriod", err)
	}
	return nil
}

func (ie *SI_RequestConfig_si_RequestPeriod) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SI_RequestConfig_si_RequestPeriod", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
