package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_ms0p02       aper.Enumerated = 0
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_ms0p04       aper.Enumerated = 1
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_ms0p0625     aper.Enumerated = 2
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_ms0p125      aper.Enumerated = 3
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_ms0p25       aper.Enumerated = 4
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_ms0p5        aper.Enumerated = 5
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_ms0p01_v1700 aper.Enumerated = 6
	LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration_Enum_spare1       aper.Enumerated = 7
)

type LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration", err)
	}
	return nil
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
