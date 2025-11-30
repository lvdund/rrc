package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	LogicalChannelConfig_bitRateMultiplier_r16_Enum_x40  aper.Enumerated = 0
	LogicalChannelConfig_bitRateMultiplier_r16_Enum_x70  aper.Enumerated = 1
	LogicalChannelConfig_bitRateMultiplier_r16_Enum_x100 aper.Enumerated = 2
	LogicalChannelConfig_bitRateMultiplier_r16_Enum_x200 aper.Enumerated = 3
)

type LogicalChannelConfig_bitRateMultiplier_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *LogicalChannelConfig_bitRateMultiplier_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode LogicalChannelConfig_bitRateMultiplier_r16", err)
	}
	return nil
}

func (ie *LogicalChannelConfig_bitRateMultiplier_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode LogicalChannelConfig_bitRateMultiplier_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
