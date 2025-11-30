package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed_Enum_true aper.Enumerated = 0
)

type LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed", err)
	}
	return nil
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
