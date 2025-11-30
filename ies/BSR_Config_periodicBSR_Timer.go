package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BSR_Config_periodicBSR_Timer_Enum_sf1      aper.Enumerated = 0
	BSR_Config_periodicBSR_Timer_Enum_sf5      aper.Enumerated = 1
	BSR_Config_periodicBSR_Timer_Enum_sf10     aper.Enumerated = 2
	BSR_Config_periodicBSR_Timer_Enum_sf16     aper.Enumerated = 3
	BSR_Config_periodicBSR_Timer_Enum_sf20     aper.Enumerated = 4
	BSR_Config_periodicBSR_Timer_Enum_sf32     aper.Enumerated = 5
	BSR_Config_periodicBSR_Timer_Enum_sf40     aper.Enumerated = 6
	BSR_Config_periodicBSR_Timer_Enum_sf64     aper.Enumerated = 7
	BSR_Config_periodicBSR_Timer_Enum_sf80     aper.Enumerated = 8
	BSR_Config_periodicBSR_Timer_Enum_sf128    aper.Enumerated = 9
	BSR_Config_periodicBSR_Timer_Enum_sf160    aper.Enumerated = 10
	BSR_Config_periodicBSR_Timer_Enum_sf320    aper.Enumerated = 11
	BSR_Config_periodicBSR_Timer_Enum_sf640    aper.Enumerated = 12
	BSR_Config_periodicBSR_Timer_Enum_sf1280   aper.Enumerated = 13
	BSR_Config_periodicBSR_Timer_Enum_sf2560   aper.Enumerated = 14
	BSR_Config_periodicBSR_Timer_Enum_infinity aper.Enumerated = 15
)

type BSR_Config_periodicBSR_Timer struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *BSR_Config_periodicBSR_Timer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode BSR_Config_periodicBSR_Timer", err)
	}
	return nil
}

func (ie *BSR_Config_periodicBSR_Timer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode BSR_Config_periodicBSR_Timer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
