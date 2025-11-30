package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_ms1     aper.Enumerated = 0
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_ms2     aper.Enumerated = 1
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_ms2dot5 aper.Enumerated = 2
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_ms4     aper.Enumerated = 3
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_ms5     aper.Enumerated = 4
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_ms10    aper.Enumerated = 5
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_spare2  aper.Enumerated = 6
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_spare1  aper.Enumerated = 7
)

type SemiStaticChannelAccessConfigUE_r17_periodUE_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SemiStaticChannelAccessConfigUE_r17_periodUE_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SemiStaticChannelAccessConfigUE_r17_periodUE_r17", err)
	}
	return nil
}

func (ie *SemiStaticChannelAccessConfigUE_r17_periodUE_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SemiStaticChannelAccessConfigUE_r17_periodUE_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
