package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer_Enum_s0     aper.Enumerated = 0
	LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer_Enum_s0dot4 aper.Enumerated = 1
	LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer_Enum_s0dot8 aper.Enumerated = 2
	LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer_Enum_s1dot6 aper.Enumerated = 3
	LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer_Enum_s3     aper.Enumerated = 4
	LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer_Enum_s6     aper.Enumerated = 5
	LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer_Enum_s12    aper.Enumerated = 6
	LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer_Enum_s30    aper.Enumerated = 7
)

type LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer", err)
	}
	return nil
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
