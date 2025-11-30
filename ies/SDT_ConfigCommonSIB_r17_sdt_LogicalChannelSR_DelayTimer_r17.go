package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_sf20   aper.Enumerated = 0
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_sf40   aper.Enumerated = 1
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_sf64   aper.Enumerated = 2
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_sf128  aper.Enumerated = 3
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_sf512  aper.Enumerated = 4
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_sf1024 aper.Enumerated = 5
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_sf2560 aper.Enumerated = 6
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_spare1 aper.Enumerated = 7
)

type SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17", err)
	}
	return nil
}

func (ie *SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
