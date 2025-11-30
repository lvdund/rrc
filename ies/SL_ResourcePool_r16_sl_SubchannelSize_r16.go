package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n10  aper.Enumerated = 0
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n12  aper.Enumerated = 1
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n15  aper.Enumerated = 2
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n20  aper.Enumerated = 3
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n25  aper.Enumerated = 4
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n50  aper.Enumerated = 5
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n75  aper.Enumerated = 6
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n100 aper.Enumerated = 7
)

type SL_ResourcePool_r16_sl_SubchannelSize_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SL_ResourcePool_r16_sl_SubchannelSize_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SL_ResourcePool_r16_sl_SubchannelSize_r16", err)
	}
	return nil
}

func (ie *SL_ResourcePool_r16_sl_SubchannelSize_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SL_ResourcePool_r16_sl_SubchannelSize_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
