package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PHR_Config_phr_ProhibitTimer_Enum_sf0    aper.Enumerated = 0
	PHR_Config_phr_ProhibitTimer_Enum_sf10   aper.Enumerated = 1
	PHR_Config_phr_ProhibitTimer_Enum_sf20   aper.Enumerated = 2
	PHR_Config_phr_ProhibitTimer_Enum_sf50   aper.Enumerated = 3
	PHR_Config_phr_ProhibitTimer_Enum_sf100  aper.Enumerated = 4
	PHR_Config_phr_ProhibitTimer_Enum_sf200  aper.Enumerated = 5
	PHR_Config_phr_ProhibitTimer_Enum_sf500  aper.Enumerated = 6
	PHR_Config_phr_ProhibitTimer_Enum_sf1000 aper.Enumerated = 7
)

type PHR_Config_phr_ProhibitTimer struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *PHR_Config_phr_ProhibitTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode PHR_Config_phr_ProhibitTimer", err)
	}
	return nil
}

func (ie *PHR_Config_phr_ProhibitTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode PHR_Config_phr_ProhibitTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
