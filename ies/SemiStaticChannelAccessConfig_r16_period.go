package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SemiStaticChannelAccessConfig_r16_period_Enum_ms1     aper.Enumerated = 0
	SemiStaticChannelAccessConfig_r16_period_Enum_ms2     aper.Enumerated = 1
	SemiStaticChannelAccessConfig_r16_period_Enum_ms2dot5 aper.Enumerated = 2
	SemiStaticChannelAccessConfig_r16_period_Enum_ms4     aper.Enumerated = 3
	SemiStaticChannelAccessConfig_r16_period_Enum_ms5     aper.Enumerated = 4
	SemiStaticChannelAccessConfig_r16_period_Enum_ms10    aper.Enumerated = 5
)

type SemiStaticChannelAccessConfig_r16_period struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *SemiStaticChannelAccessConfig_r16_period) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode SemiStaticChannelAccessConfig_r16_period", err)
	}
	return nil
}

func (ie *SemiStaticChannelAccessConfig_r16_period) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode SemiStaticChannelAccessConfig_r16_period", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
