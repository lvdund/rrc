package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetUplinkPerCC_channelBW_90mhz_Enum_supported aper.Enumerated = 0
)

type FeatureSetUplinkPerCC_channelBW_90mhz struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *FeatureSetUplinkPerCC_channelBW_90mhz) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode FeatureSetUplinkPerCC_channelBW_90mhz", err)
	}
	return nil
}

func (ie *FeatureSetUplinkPerCC_channelBW_90mhz) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode FeatureSetUplinkPerCC_channelBW_90mhz", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
