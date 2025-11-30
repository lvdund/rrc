package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SidelinkPreconfigNR_r16_t400_r16_Enum_ms100  aper.Enumerated = 0
	SidelinkPreconfigNR_r16_t400_r16_Enum_ms200  aper.Enumerated = 1
	SidelinkPreconfigNR_r16_t400_r16_Enum_ms300  aper.Enumerated = 2
	SidelinkPreconfigNR_r16_t400_r16_Enum_ms400  aper.Enumerated = 3
	SidelinkPreconfigNR_r16_t400_r16_Enum_ms600  aper.Enumerated = 4
	SidelinkPreconfigNR_r16_t400_r16_Enum_ms1000 aper.Enumerated = 5
	SidelinkPreconfigNR_r16_t400_r16_Enum_ms1500 aper.Enumerated = 6
	SidelinkPreconfigNR_r16_t400_r16_Enum_ms2000 aper.Enumerated = 7
)

type SidelinkPreconfigNR_r16_t400_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SidelinkPreconfigNR_r16_t400_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SidelinkPreconfigNR_r16_t400_r16", err)
	}
	return nil
}

func (ie *SidelinkPreconfigNR_r16_t400_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SidelinkPreconfigNR_r16_t400_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
