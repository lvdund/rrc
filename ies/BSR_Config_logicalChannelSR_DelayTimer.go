package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BSR_Config_logicalChannelSR_DelayTimer_Enum_sf20   aper.Enumerated = 0
	BSR_Config_logicalChannelSR_DelayTimer_Enum_sf40   aper.Enumerated = 1
	BSR_Config_logicalChannelSR_DelayTimer_Enum_sf64   aper.Enumerated = 2
	BSR_Config_logicalChannelSR_DelayTimer_Enum_sf128  aper.Enumerated = 3
	BSR_Config_logicalChannelSR_DelayTimer_Enum_sf512  aper.Enumerated = 4
	BSR_Config_logicalChannelSR_DelayTimer_Enum_sf1024 aper.Enumerated = 5
	BSR_Config_logicalChannelSR_DelayTimer_Enum_sf2560 aper.Enumerated = 6
	BSR_Config_logicalChannelSR_DelayTimer_Enum_spare1 aper.Enumerated = 7
)

type BSR_Config_logicalChannelSR_DelayTimer struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *BSR_Config_logicalChannelSR_DelayTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode BSR_Config_logicalChannelSR_DelayTimer", err)
	}
	return nil
}

func (ie *BSR_Config_logicalChannelSR_DelayTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode BSR_Config_logicalChannelSR_DelayTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
